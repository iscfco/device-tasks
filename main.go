package main

import (
	"device-tasks/model"
	"device-tasks/modules/configurator"
	"device-tasks/modules/scenariosdecoder"
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

	// Getting optimal configuration:
	optimalConfigs := configurator.GetOptimalConfig(scenarios)

	// Writing result:
	WriteResult(optimalConfigs)
}

// ReadScenariosFromFile will read the possible scenarios from a file, we assume the file is
// located in the same level of this file
func ReadScenariosFromFile() (*string, error) {
	file, err := ioutil.ReadFile("challenge.in")
	if err != nil {
		return nil, err
	}
	data := string(file)
	return &data, nil
}

// WriteResult will write the result in the output file
func WriteResult(configs []model.DeviceConfigList) error {
	var result string
	for index, config := range configs {
		result += fmt.Sprint(config)
		if index+1 < len(configs) {
			result += fmt.Sprint("\n\n")
		}
	}
	return ioutil.WriteFile("challenge.out", []byte(result), 0777)
}
