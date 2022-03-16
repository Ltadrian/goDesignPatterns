package main

/**
Bridge is a structural design pattern that
divides business logic or huge class into separate class
hierachies that can be developed independently
Ex:
Have two type computers -> Mac/Windows
Two printers -> Epson / HP
Both printers need to work with each other
in any combination
The client doesn't want to worry about the details of connecting
printers to computers
Abstraction hiearchy: computers
Implementation hiearchy: printers
*/

import "fmt"

// Abstraction
type computer interface {
	print()
	setPrinter()
}

// Refined abstraction
type mac struct {
	printer printer
}

func (m *mac) print() {
	fmt.Println("Print request for mac")
	m.printer.printFile()
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}

type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}

type printer interface {
	printFile()
}

type epson struct{}

func (e *epson) printFile() {
	fmt.Println("print file by epson")
}

type hp struct{}

func (h *hp) printFile() {
	fmt.Println("print file by hp")
}

func main() {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	macComputer := &mac{}
	macComputer.setPrinter(hpPrinter)
	macComputer.print()
	fmt.Println()

	macComputer.setPrinter(epsonPrinter)
	macComputer.print()
	fmt.Println()

	winComputer := &windows{}
	winComputer.setPrinter(hpPrinter)
	winComputer.print()
	fmt.Println()

	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
	fmt.Println()
}
