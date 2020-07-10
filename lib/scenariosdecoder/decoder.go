package scenariosdecoder

import (
	"fmt"
	"strings"
)

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
