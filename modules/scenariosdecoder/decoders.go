package scenariosdecoder

import (
	"device-tasks/model"
	"strconv"
	"strings"
)

// scenarioItemDecoders will map each scenarioItem type to a handler
var scenarioItemDecoders = map[scenarioItem]func(string, *model.DeviceScenario) error{
	SCENARIO_ITEM_RESOURCE_CAPACITY: DecodeResourceCapacity,
	SCENARIO_ITEM_BACKGROUND_TASKS:  DecodeBackgroundTasks,
	SCENARIO_ITEM_FOREGROUND_TASK:   DecodeForegroundTasks,
}

// DecodeResourceCapacity will decode the type SCENARIO_ITEM_RESOURCE_CAPACITY
// it waits for a string with the format: 10
func DecodeResourceCapacity(item string, deviceScenario *model.DeviceScenario) error {
	itemAsInt, err := strconv.Atoi(item)
	if err != nil {
		return err
	}
	deviceScenario.ResourceCapacity = itemAsInt
	return nil
}

// DecodeBackgroundTasks will decode for SCENARIO_ITEM_BACKGROUND_TASKS
// Waits for a string with the format: (1,2),(2,2),(3,5)
func DecodeBackgroundTasks(item string, deviceScenario *model.DeviceScenario) error {
	pairs := strings.Split(item, "),(")
	for _, pair := range pairs {
		pair = strings.Trim(pair, "(")
		pair = strings.Trim(pair, ")")
		pairItems := strings.Split(pair, ",")
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
		deviceScenario.BackgroundTasks = append(deviceScenario.BackgroundTasks, model.BackgroundTask{
			ID:                  ID,
			ResourceConsumption: resourceConsumption,
		})
	}
	return nil
}

// DecodeForegroundTasks will decode for SCENARIO_ITEM_FOREGROUND_TASK
// Waits for a string with the format: (1,2),(2,2),(3,5)
func DecodeForegroundTasks(item string, deviceScenario *model.DeviceScenario) error {
	pairs := strings.Split(item, "),(")
	for _, pair := range pairs {
		pair = strings.Trim(pair, "(")
		pair = strings.Trim(pair, ")")
		pairItems := strings.Split(pair, ",")
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
		deviceScenario.ForegroundTasks = append(deviceScenario.ForegroundTasks, model.ForegroundTask{
			ID:                  ID,
			ResourceConsumption: resourceConsumption,
		})
	}
	return nil
}
