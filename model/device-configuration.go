package model

import "fmt"

type DeviceConfig struct {
	BackgroundTaskID int
	ForegroundTaskID int
}

type DeviceConfigList struct {
	List []DeviceConfig
}

func (d DeviceConfigList) String() string {
	var result string
	template := "(%v,%v)"
	for index, config := range d.List {
		result += fmt.Sprintf(template, config.ForegroundTaskID, config.BackgroundTaskID)
		if index+1 < len(d.List) {
			result += ","
		}
	}
	return result
}
