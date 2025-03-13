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
            "command_on": "docker start home-assistant",
            "command_off": "docker stop home-assistant",
            "tag": "container home-assistant",
            "type": "Terminal"
        },
        {
            "entity_type": "sensor",
            "command_state": "docker ps --filter \"name=home-assistant\" --format \"{{.Status}}\"",
            "tag": "container home-assistant",
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