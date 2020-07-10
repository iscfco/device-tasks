package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Read scenarios from file
	scenariosStr, err := ReadScenariosFromFile()
	if err != nil {
		log.Fatal("Error at moment of read file", err)
	}
	// fmt.Println(*scenarios)

	// Scenarios Parser:
	// - Will read data and determine "Resource Capacity", "Background Tasks" and "Foreground Tasks"
	scenarios, err := DecodeScenarios(*scenariosStr)
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

// =====================================================
// 				STEP #2
// =====================================================

// DecodeScenarios will read the scenarios (from a parameter) and determine Resource
// Capacity, pairs of Background Tasks and pairs of Foreground Tasks
func DecodeScenarios(scenariosStr string) ([]DeviceScenario, error) {
	scenariosSliced := strings.Split(scenariosStr, "\n")
	scenarioItemToProccess := NewScenarioItem()

	var scenarios []DeviceScenario
	currentScenarioIndex := -1
	for _, line := range scenariosSliced {
		fmt.Println("reading: ...", line)
		if line == "" {
			continue // It just ignores the line, it doesn't contain anything
		}
		if scenarioItemToProccess == SCENARIO_ITEM_RESOURCE_CAPACITY {
			currentScenarioIndex++
			scenarios = append(scenarios, DeviceScenario{})
		}
		err := scenarioItemDecoders[scenarioItemToProccess](line, &scenarios[currentScenarioIndex])
		if err != nil {
			return nil, err
		}
		scenarioItemToProccess = scenarioItemToProccess.GetNext()
	}
	return scenarios, nil
}

// ======================= scenarioItem types ...
type scenarioItem int

func (i scenarioItem) GetNext() scenarioItem {
	switch i {
	case SCENARIO_ITEM_RESOURCE_CAPACITY:
		return SCENARIO_ITEM_BACKGROUND_TASKS
	case SCENARIO_ITEM_BACKGROUND_TASKS:
		return SCENARIO_ITEM_FOREGROUND_TASK
	case SCENARIO_ITEM_FOREGROUND_TASK:
		return SCENARIO_ITEM_RESOURCE_CAPACITY
	default:
		// Returning the first item to start the cycle of items
		return SCENARIO_ITEM_RESOURCE_CAPACITY
	}
}

func (i scenarioItem) String() string {
	return string(i)
}

const (
	SCENARIO_ITEM_RESOURCE_CAPACITY scenarioItem = 1
	SCENARIO_ITEM_BACKGROUND_TASKS  scenarioItem = 2
	SCENARIO_ITEM_FOREGROUND_TASK   scenarioItem = 3
)

func NewScenarioItem() scenarioItem {
	return SCENARIO_ITEM_RESOURCE_CAPACITY
}

// =============== scenarioItemDecoders .....
var scenarioItemDecoders = map[scenarioItem]func(string, *DeviceScenario) error{
	SCENARIO_ITEM_RESOURCE_CAPACITY: DecodeResourceCapacity,
	SCENARIO_ITEM_BACKGROUND_TASKS:  DecodeBackgroundTasks,
	SCENARIO_ITEM_FOREGROUND_TASK:   DecodeForegroundTasks,
}

func DecodeResourceCapacity(item string, deviceScenario *DeviceScenario) error {
	itemAsInt, err := strconv.Atoi(item)
	if err != nil {
		return err
	}
	deviceScenario.ResourceCapacity = itemAsInt
	return nil
}

// Waits for a string with the format: (1,2),(2,2),(3,5)
func DecodeBackgroundTasks(item string, deviceScenario *DeviceScenario) error {
	pairs := strings.Split(item, "),(")
	fmt.Println("*************** pairs...:", pairs)
	for _, pair := range pairs {
		fmt.Println("--- pair:", pair)
		pair = strings.Trim(pair, "(")
		pair = strings.Trim(pair, ")")
		fmt.Println("--- pair:", pair)

		pairItems := strings.Split(pair, ",")
		fmt.Println("--- pairItems:", pairItems)
		if len(pairItems) != 2 {
			return ErrWrongFormatOfBackgroundTask
		}

		ID, err := strconv.Atoi(pairItems[0])
		if err != nil {
			return err
		}
		resourceConsumption, err := strconv.Atoi(pairItems[1])
		if err != nil {
			return err
		}
		deviceScenario.BackgroundTasks = append(deviceScenario.BackgroundTasks, BackgroundTask{
			ID:                  ID,
			ResourceConsumption: resourceConsumption,
		})
	}
	return nil
}

// Waits for a string with the format: (1,2),(2,2),(3,5)
func DecodeForegroundTasks(item string, deviceScenario *DeviceScenario) error {
	pairs := strings.Split(item, "),(")
	for _, pair := range pairs {
		pair = strings.Trim(pair, "(")
		pair = strings.Trim(pair, ")")
		fmt.Sprintln("--- pair:", pair)

		pairItems := strings.Split(pair, ",")
		fmt.Sprintln("--- pairItems:", pairItems)
		if len(pairItems) != 2 {
			return ErrWrongFormatOfForegroundTask
		}

		ID, err := strconv.Atoi(pairItems[0])
		if err != nil {
			return err
		}
		resourceConsumption, err := strconv.Atoi(pairItems[1])
		if err != nil {
			return err
		}
		deviceScenario.ForegroundTask = append(deviceScenario.ForegroundTask, ForegroundTask{
			ID:                  ID,
			ResourceConsumption: resourceConsumption,
		})
	}
	return nil
}

// Decoder Errors:
var ErrWrongFormatOfBackgroundTask error = errors.New("Wrong format of Background Task items")
var ErrWrongFormatOfForegroundTask error = errors.New("Wrong format of Foreground Task items")

// ================= structs

// DeviceScenario is the struct to represent a device scenario with the format:
// 		12					<----- Represents the resource capacity of a device
// 		(1,2)(3,4)(3,6)		<----- Represents the Background Tasks
// 		(1,4)				<----- Represents the Foreground Tasks
type DeviceScenario struct {
	ResourceCapacity int
	BackgroundTasks  []BackgroundTask
	ForegroundTask   []ForegroundTask
}

// BackgroundTask is the struct to represent a Background Task of a Device
type BackgroundTask struct {
	ID                  int
	ResourceConsumption int
}

// ForegroundTask is the struct to represent a Foreground Task of a Device
type ForegroundTask struct {
	ID                  int
	ResourceConsumption int
}
