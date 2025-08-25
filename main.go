package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	configFile := flag.String("config", "./my-cloud-box.yaml", "configuration file")
	loadConfig(configFile)

}

func loadConfig(configFile string) {
	_, err := os.Stat(configFile)
	if err != nil {
		fmt.Printf("Configuration file with name %s does not exits.\n", configFile)
		log.Fatal(err)
	}

	configFile, error := os.ReadFile(configFile)

	if error != nil {
		log.Printf("There was an error reading, config file: %v", error)
	}
}
