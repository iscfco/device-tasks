package scenariosdecoder

import (
	"device-tasks/model"
	"strings"
)

// DecodeScenarios will read the scenarios (from a parameter) and determine Resource
// Capacity, pairs of Background Tasks and pairs of Foreground Tasks
func DecodeScenarios(scenariosStr string) ([]model.DeviceScenario, error) {
	scenariosSliced := strings.Split(scenariosStr, "\n")
	scenarioItemToProccess := NewScenarioItem()

	var scenarios []model.DeviceScenario
	currentScenarioIndex := -1
	for _, line := range scenariosSliced {
		if line == "" {
			continue // It just ignores the line, it doesn't contain anything
		}
		if scenarioItemToProccess == SCENARIO_ITEM_RESOURCE_CAPACITY {
			currentScenarioIndex++
			scenarios = append(scenarios, model.DeviceScenario{})
		}
		err := scenarioItemDecoders[scenarioItemToProccess](line, &scenarios[currentScenarioIndex])
		if err != nil {
			return nil, err
		}
		scenarioItemToProccess = scenarioItemToProccess.GetNext()
	}
	return scenarios, nil
}
