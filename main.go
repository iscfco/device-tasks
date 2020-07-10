package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Read scenarios:
	scenarios, err := ReadScenarios()
	if err != nil {
		log.Fatal("Error at moment of read file", err)
	}
	fmt.Println(*scenarios)

	// Scenarios Parser:
	// - Will read data and determine "Resource Capacity", "Background Tasks" and "Foreground Tasks"

	// Getting optimal configuration:

	// Writing result:
}

// ReadScenarios will read the possible scenarios from a file, we assume the file is
// located in the same level of this file
func ReadScenarios() (*string, error) {
	file, err := ioutil.ReadFile("challenge.in")
	if err != nil {
		return nil, err
	}
	data := string(file)
	return &data, nil
}
