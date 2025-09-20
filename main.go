package main

import (
	"flag"
	"github.com/laertida/mycloudbox/endpoints"
	"log"
)

func main() {
	configFile := flag.String("config", "./my-cloud-box.yaml", "configuration file")
	flag.Parse()
	log.Println("started process with configFile", *configFile)

	fileEndpoint := endpoints.File{Path: "./", Protocol: "file", Properties: "?test=true"}
	fileEndpoint.Log()
}
