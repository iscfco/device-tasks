package scenariosdecoder

// scenarioItem represents an element of a DeviceScenario struct
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
