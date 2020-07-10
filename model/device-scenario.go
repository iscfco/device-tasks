package model

// DeviceScenario is the struct to represent a device scenario with the format:
// 		12					<----- Represents the resource capacity of a device
// 		(1,2)(3,4)(3,6)		<----- Represents the Background Tasks
// 		(1,4)				<----- Represents the Foreground Tasks
type DeviceScenario struct {
	ResourceCapacity int
	BackgroundTasks  []BackgroundTask
	ForegroundTasks  []ForegroundTask
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
