package main

import "fmt"

func main() {
	device := Device{
		ID:   1,
		Name: "modem",
	}

	device2 := Device{
		ID:   2,
		Name: "antena",
	}

	device3 := Device{
		ID:   3,
		Name: "marthe",
	}

	device.Connect()

	manager := DeviceManager{}

	fmt.Println(device.Online)

	manager.AddDevice(device)
	manager.AddDevice(device2)
	manager.AddDevice(device3)

	manager.ListDevices()
}
