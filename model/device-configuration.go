package model

type DeviceConfig struct {
	BackgroundTaskID int
	ForegroundTaskID int
}

type DeviceConfigList struct {
	List []DeviceConfig
}
