# Go Configuration Generator for IoTuring

Link to the repository: [https://github.com/richibrics/IoTuring](https://github.com/richibrics/IoTuring)

This Go script automatically generates a configuration JSON file for managing Docker containers through MQTT using IoTuring. It creates switches and sensors for each running container, allowing you to control and monitor containers easily via MQTT in Home Assistant.

## Features

- **Generates IoTuring Configuration**: Creates a configuration JSON file with entities for Docker containers.
- **Supports Switches and Sensors**: Allows turning containers on/off and monitors their state.
- **Customizable MQTT Settings**: Takes MQTT host, username, and password as input.
- **Static Entities**: Includes common system information like hostname, RAM, uptime, etc., statically provided by IoTuring.

## Prerequisites

- **Go**: The Go programming language must be installed. You can download it from [here](https://golang.org/dl/).
- **Docker**: The script interacts with Docker to retrieve the list of running containers.
- **IoTuring**: You must have IoTuring installed to integrate the generated configuration.

## Installation

1. **Clone the repository** (if you haven't already):
   ```bash
   git clone https://github.com/richibrics/IoTuring.git
   cd IoTuring
2. **Build the Go binary** : After cloning the repository, run the following commands to build the Go binary:
   ```bash
   go mod tidy
   go build -o generate_configurations

3. **Ensure Docker is running**: Make sure Docker is installed and running on your system:
   ```bash
   docker ps

4. **Run the Script**: Run the Go binary that generates the configuration:
   ```bash
   ./docker-mqtt-config-generator

   You will be prompted to enter the MQTT host, username, and password for the configuration.

5. **Check the Generated File**: After running the script, the configurations.json file will be created in the same directory.

## Example Output
Here is an example of the configurations.json file that will be generated:
  ```json
  {
    "active_entities": [
        {
            "type": "AppInfo"
        },
        {
            "type": "BootTime"
        },
        {
            "type": "Hostname"
        },
        {
            "enable_hibernate": "N",
            "type": "Power"
        },
        {
            "type": "Ram"
        },
        {
            "type": "Temperature"
        },
        {
            "type": "OperatingSystem"
        },
        {
            "type": "UpTime"
        },
        {
            "type": "Time"
        },
        {
            "entity_type": "switch",
            "command_on": "docker start beszel",
            "command_off": "docker stop beszel",
            "tag": "container beszel",
            "type": "Terminal"
        },
        {
            "entity_type": "sensor",
            "command_state": "docker ps --filter \"name=beszel\" --format \"{{.Status}}\"",
            "tag": "container beszel",
            "type": "Terminal"
        },
        {
            "entity_type": "switch",
            "command_on": "docker start teslamate",
            "command_off": "docker stop teslamate",
            "tag": "container teslamate",
            "type": "Terminal"
        },
        {
            "entity_type": "sensor",
            "command_state": "docker ps --filter \"name=teslamate\" --format \"{{.Status}}\"",
            "tag": "container teslamate",
            "type": "Terminal"
        },
        {
            "entity_type": "switch",
            "command_on": "docker start teslamate-grafana",
            "command_off": "docker stop teslamate-grafana",
            "tag": "container teslamate-grafana",
            "type": "Terminal"
        },
        {
            "entity_type": "sensor",
            "command_state": "docker ps --filter \"name=teslamate-grafana\" --format \"{{.Status}}\"",
            "tag": "container teslamate-grafana",
            "type": "Terminal"
        },
        {
            "entity_type": "switch",
            "command_on": "docker start homepage",
            "command_off": "docker stop homepage",
            "tag": "container homepage",
            "type": "Terminal"
        },
        {
            "entity_type": "sensor",
            "command_state": "docker ps --filter \"name=homepage\" --format \"{{.Status}}\"",
            "tag": "container homepage",
            "type": "Terminal"
        },
        {
            "entity_type": "switch",
            "command_on": "docker start homepage_tailscale",
            "command_off": "docker stop homepage_tailscale",
            "tag": "container homepage_tailscale",
            "type": "Terminal"
        },
        {
            "entity_type": "sensor",
            "command_state": "docker ps --filter \"name=homepage_tailscale\" --format \"{{.Status}}\"",
            "tag": "container homepage_tailscale",
            "type": "Terminal"
        },
        {
            "entity_type": "switch",
            "command_on": "docker start mealie",
            "command_off": "docker stop mealie",
            "tag": "container mealie",
            "type": "Terminal"
        },
        {
            "entity_type": "sensor",
            "command_state": "docker ps --filter \"name=mealie\" --format \"{{.Status}}\"",
            "tag": "container mealie",
            "type": "Terminal"
        },
        {
            "entity_type": "switch",
            "command_on": "docker start pdf",
            "command_off": "docker stop pdf",
            "tag": "container pdf",
            "type": "Terminal"
        },
        {
            "entity_type": "sensor",
            "command_state": "docker ps --filter \"name=pdf\" --format \"{{.Status}}\"",
            "tag": "container pdf",
            "type": "Terminal"
        },
        {
            "entity_type": "switch",
            "command_on": "docker start drawio",
            "command_off": "docker stop drawio",
            "tag": "container drawio",
            "type": "Terminal"
        },
        {
            "entity_type": "sensor",
            "command_state": "docker ps --filter \"name=drawio\" --format \"{{.Status}}\"",
            "tag": "container drawio",
            "type": "Terminal"
        },
        {
            "entity_type": "switch",
            "command_on": "docker start firefly_iii_core",
            "command_off": "docker stop firefly_iii_core",
            "tag": "container firefly_iii_core",
            "type": "Terminal"
        },
        {
            "entity_type": "sensor",
            "command_state": "docker ps --filter \"name=firefly_iii_core\" --format \"{{.Status}}\"",
            "tag": "container firefly_iii_core",
            "type": "Terminal"
        }
    ],
    "active_warehouses": [
        {
            "address": "{MQTT_HOST}",
            "port": "1883",
            "name": "thor",
            "username": "{MQTT_USER}",
            "password": "{MQTT_PASSWORD}",
            "add_name": "Y",
            "use_tag": "Y",
            "type": "HomeAssistant"
        }
    ],
    "settings": [
        {
            "update_interval": 5,
            "retry_interval": 1,
            "type": "App"
        }
    ]
}

## How It Works
### Entities: The script generates switch entities to start/stop containers and sensor entities to monitor their state.
### Static Entities: Predefined system metrics (such as Ram, Temperature, etc.) are also added to the configuration.
### MQTT Integration: The script allows you to customize the MQTT connection settings, ensuring integration with your MQTT broker.