package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Entity struct {
	EntityType      string `json:"entity_type,omitempty"`
	CommandOn       string `json:"command_on,omitempty"`
	CommandOff      string `json:"command_off,omitempty"`
	CommandState    string `json:"command_state,omitempty"`
	Tag             string `json:"tag,omitempty"`
	Type            string `json:"type"`
	EnableHibernate string `json:"enable_hibernate,omitempty"`
}

type Warehouse struct {
	Address  string `json:"address"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	AddName  string `json:"add_name"`
	UseTag   string `json:"use_tag"`
	Type     string `json:"type"`
}

type Setting struct {
	UpdateInterval int    `json:"update_interval"`
	RetryInterval  int    `json:"retry_interval"`
	Type           string `json:"type"`
}

type Config struct {
	ActiveEntities   []Entity    `json:"active_entities"`
	ActiveWarehouses []Warehouse `json:"active_warehouses"`
	Settings         []Setting   `json:"settings"`
}

func getRunningContainers() ([]string, error) {
	// Run "docker ps --format '{{.Names}}'" to get names of all running containers
	cmd := exec.Command("docker", "ps", "--format", "{{.Names}}")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	// Split the output into individual container names
	containerNames := strings.Split(out.String(), "\n")
	// Remove empty string from the split (if there's a trailing newline)
	if len(containerNames) > 0 && containerNames[len(containerNames)-1] == "" {
		containerNames = containerNames[:len(containerNames)-1]
	}

	return containerNames, nil
}

func generateConfigurations(containers []string) ([]Entity, error) {
	var configurations []Entity

	// Add the static entities to the active entities list
	staticEntities := []Entity{
		{Type: "AppInfo"},
		{Type: "BootTime"},
		{Type: "Hostname"},
		{EnableHibernate: "N", Type: "Power"},
		{Type: "Ram"},
		{Type: "Temperature"},
		{Type: "OperatingSystem"},
		{Type: "UpTime"},
		{Type: "Time"},
	}

	// Append the static entities
	configurations = append(configurations, staticEntities...)

	// Loop through all containers and create the switch and sensor configurations
	for _, container := range containers {
		// Create the switch configuration
		switchConfig := Entity{
			EntityType: "switch",
			CommandOn:  fmt.Sprintf("docker start %s", container),
			CommandOff: fmt.Sprintf("docker stop %s", container),
			Tag:        fmt.Sprintf("container %s", container),
			Type:       "Terminal",
		}

		// Create the sensor configuration
		sensorConfig := Entity{
			EntityType:   "sensor",
			CommandState: fmt.Sprintf("docker ps --filter \"name=%s\" --format \"{{.Status}}\"", container),
			Tag:          fmt.Sprintf("container %s", container),
			Type:         "Terminal",
		}

		// Append both configurations to the result list
		configurations = append(configurations, switchConfig, sensorConfig)
	}

	return configurations, nil
}

func generateFullConfig(mqttHost, mqttUser, mqttPassword string) (Config, error) {
	// Get the list of running containers
	containers, err := getRunningContainers()
	if err != nil {
		return Config{}, err
	}

	// Generate the switch and sensor configurations for containers
	activeEntities, err := generateConfigurations(containers)
	if err != nil {
		return Config{}, err
	}

	// Generate the full configuration
	config := Config{
		ActiveEntities: activeEntities,
		ActiveWarehouses: []Warehouse{
			{
				Address:  mqttHost,
				Port:     "1883",
				Name:     "thor",
				Username: mqttUser,
				Password: mqttPassword,
				AddName:  "Y",
				UseTag:   "Y",
				Type:     "HomeAssistant",
			},
		},
		Settings: []Setting{
			{
				UpdateInterval: 5,
				RetryInterval:  1,
				Type:           "App",
			},
		},
	}

	return config, nil
}

func writeConfigToFile(config Config) error {
	// Create or open the configurations.json file
	file, err := os.Create("configurations.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// Convert the configuration struct to JSON
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print
	return encoder.Encode(config)
}

func main() {
	// Input MQTT credentials
	var mqttHost, mqttUser, mqttPassword string
	fmt.Print("Enter MQTT Host: ")
	fmt.Scanln(&mqttHost)
	fmt.Print("Enter MQTT Username: ")
	fmt.Scanln(&mqttUser)
	fmt.Print("Enter MQTT Password: ")
	fmt.Scanln(&mqttPassword)

	// Generate the full configuration JSON
	config, err := generateFullConfig(mqttHost, mqttUser, mqttPassword)
	if err != nil {
		fmt.Println("Error generating configuration:", err)
		return
	}

	// Write the configuration to a file
	err = writeConfigToFile(config)
	if err != nil {
		fmt.Println("Error writing configuration to file:", err)
		return
	}

	fmt.Println("Configuration saved to configurations.json")
}
