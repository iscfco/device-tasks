package configurator

import "device-tasks/model"

// GetOptimalConfig will define the optimal configuration for each element from the
// array of scenarios
func GetOptimalConfig(scenarios []model.DeviceScenario) []model.DeviceConfigList {
	var result []model.DeviceConfigList
	for _, scenario := range scenarios {
		result = append(result, model.DeviceConfigList{
			List: optimizeDevice(scenario),
		})
	}
	return result
}

// optimizeDevice for a given scenario of a device this funcion define the optimal config
func optimizeDevice(scenario model.DeviceScenario) []model.DeviceConfig {
	lastOptimalConfigValue := -1
	var optimalConfigs []model.DeviceConfig

	for _, backgroundTask := range scenario.BackgroundTasks {
		for _, foregroundTask := range scenario.ForegroundTasks {
			configToAnalize := backgroundTask.ResourceConsumption + foregroundTask.ResourceConsumption

			switch {
			case surpassesCapacity(configToAnalize, scenario.ResourceCapacity):
				continue
			case isTheOptimalConfig(configToAnalize, lastOptimalConfigValue):
				lastOptimalConfigValue = configToAnalize
				optimalConfigs = []model.DeviceConfig{}
				optimalConfigs = append(optimalConfigs, model.DeviceConfig{
					BackgroundTaskID: backgroundTask.ID,
					ForegroundTaskID: foregroundTask.ID,
				})
			case doWeHaveMoreThanOneOptimalConfig(configToAnalize, lastOptimalConfigValue):
				optimalConfigs = append(optimalConfigs, model.DeviceConfig{
					BackgroundTaskID: backgroundTask.ID,
					ForegroundTaskID: foregroundTask.ID,
				})
			}

		}
	}
	return optimalConfigs
}

func surpassesCapacity(configToAnalize, resourceCapacity int) bool {
	if configToAnalize > resourceCapacity {
		return true
	}
	return false
}

func isTheOptimalConfig(configToAnalize, lastOptimalConfigValue int) bool {
	if configToAnalize > lastOptimalConfigValue {
		return true
	}
	return false
}

func doWeHaveMoreThanOneOptimalConfig(configToAnalize, lastOptimalConfigValue int) bool {
	if configToAnalize == lastOptimalConfigValue {
		return true
	}
	return false
}
