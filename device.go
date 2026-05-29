package main

type Device struct {
	ID     int
	Name   string
	Online bool
}

func (d *Device) Connect() {
	d.Online = true
}

func (d *Device) Disconnect() {
	d.Online = false
}
