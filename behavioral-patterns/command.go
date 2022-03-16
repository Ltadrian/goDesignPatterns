package main

import "fmt"

/**
Command is behavioral design pattern that converts requests
or simple operations into objects

conversion allows deferred or remote execution of commands,
storing command history, etc

ex: tv can be turned on by button on the remote or actual tv
*/

// Invoker
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

// Command interface
type command interface {
	execute()
}

// Concrete command
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

// Receiver interface
type device interface {
	on()
	off()
}

// Concrete interface
type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Tuning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	tv := &tv{}

	onCommand := &onCommand{
		device: tv,
	}

	offCommand := &offCommand{
		device: tv,
	}

	onButton := &button{
		command: onCommand,
	}

	onButton.press()

	offButton := &button{
		command: offCommand,
	}
	offButton.press()

}
