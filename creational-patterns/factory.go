package main

import (
	"fmt"
)

/**
- It is impossible to implement the classic factory method pattern in Go due to lack of OOP features
- Can still implement a simple factory
- iCar -> interface which defines all methods a car should have
- car struct -> implements the iCar interface
- two concrete cars -> tesla, ford
- both embed car struct and inrectly implement all iCar methods
- the carFactory serves as a factory, which creates cars
- of the desired type based on incoming argument
- main.go acts as client, uses car factory to create instances of various car
*/

// Interface
type iCar interface {
	setName(name string)
	setYear(year int)
	getName() string
	getYear() int
}

// Implements car interface
type car struct {
	name string
	year int
}

func (c *car) setName(name string) {
	c.name = name
}

func (c *car) getName() string {
	return c.name
}

func (c *car) setYear(year int) {
	c.year = year
}

func (c *car) getYear() int {
	return c.year
}

// Concrete Types
type tesla struct {
	car
}

func newTesla() iCar {
	return &tesla{
		car: car{
			name: "Tesla model Y",
			year: 2021,
		},
	}
}

type ford struct {
	car
}

func newFord() iCar {
	return &ford{
		car: car{
			name: "Ford model E",
			year: 2022,
		},
	}
}

// Factory Method
func getCar(carType string) (iCar, error) {
	if carType == "tesla" {
		return newTesla(), nil
	}
	if carType == "ford" {
		return newFord(), nil
	}
	return nil, fmt.Errorf("Wrong car type passed")
}

func main() {
	tesla, _ := getCar("tesla")
	ford, _ := getCar("ford")

	printDetails(tesla)
	printDetails(ford)
}

func printDetails(c iCar) {
	fmt.Printf("Car: %s", c.getName())
	fmt.Println()
	fmt.Printf("Year: %d", c.getYear())
	fmt.Println()
}
