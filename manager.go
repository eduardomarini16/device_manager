package main

import "fmt"

type DeviceManager struct {
	devices []Device
}

func (m *DeviceManager) AddDevice(device Device) {
	m.devices = append(m.devices, device)
}

func (m *DeviceManager) ListDevices() {
	for _, device := range m.devices {
		fmt.Println(device.ID, device.Name)
	}
}
