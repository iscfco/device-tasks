package main

import (
	"device-tasks/modules/scenariosdecoder"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Read scenarios from file
	scenariosStr, err := ReadScenariosFromFile()
	if err != nil {
		log.Fatal("Error at moment of read file", err)
	}

	// Scenarios decoding:
	scenarios, err := scenariosdecoder.DecodeScenarios(*scenariosStr)
	if err != nil {
		log.Fatal("Error at moment of decode scenarios", err)
	}
	jsonS, _ := json.Marshal(scenarios)
	fmt.Println(string(jsonS))

	// Getting optimal configuration:

	// Writing result:
}

// ReadScenarios will read the possible scenarios from a file, we assume the file is
// located in the same level of this file
func ReadScenariosFromFile() (*string, error) {
	file, err := ioutil.ReadFile("challenge.in")
	if err != nil {
		return nil, err
	}
	data := string(file)
	return &data, nil
}
