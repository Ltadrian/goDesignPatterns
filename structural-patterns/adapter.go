package main

import "fmt"

/**
  - Structural Pattern: Adapter
  - For classes that don't implement an interface, an adapter allows a way to
  - allows the classes to collaborate
  - Computer -> Interface with insertIntoLightningPort
  - Mac -> insertIntoLightningPort
  - Windows -> insertIntoUSBPort
  - WindowsAdapter is used to call insertIntoLightning port and implement interface
  - Client -> takes in computer and calls method
*/

type client struct {
}

func (c *client) insertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Client inserts lightning connector into computer")
	com.insertIntoLightningPort()
}

type computer interface {
	insertIntoLightningPort()
}

type mac struct {
}

func (m *mac) insertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine")
}

type windows struct {
}

func (w *windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine")
}

type windowsAdapter struct {
	windowsMachine *windows
}

func (w *windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB")
	w.windowsMachine.insertIntoUSBPort()
}

func main() {
	client := &client{}
	mac := &mac{}

	client.insertLightningConnectorIntoComputer(mac)

	windows := &windows{}

	adapter := &windowsAdapter{
		windowsMachine: windows,
	}

	client.insertLightningConnectorIntoComputer(adapter)
}
