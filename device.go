package main

type Device struct {
	ID     string
	Name   string
	Online bool
}

func (d *Device) Connect() {
	d.Online = true
}

func (d *Device) Disconnect() {
	d.Online = false
}
