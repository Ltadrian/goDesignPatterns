package main

import "fmt"

/**
- Behavioral Design Pattern: Chain of Responsibility
- allows passing request along chain of potential handlers
- until one of them handles request
- Ex: Patient goes to hospital
- Needs to see reception, doctor, medicine room, then cashier
*/

// Interface

type department interface {
	execute(*patient)
	setNext(department)
}

// Concrete Handler

type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}

	fmt.Println("Reception done registering")
	p.registrationDone = true
	r.next.execute(p)

}

func (r *reception) setNext(next department) {
	r.next = next
}

// Concrete handler

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor check up already done")
		d.next.execute(p)
		return
	}

	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}

// Concrete handler
type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *medical) setNext(next department) {
	m.next = next
}

// Concrete handler
type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
	p.paymentDone = true
}

func (c *cashier) setNext(next department) {
	c.next = next
}

// patient
type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func main() {
	cashier := &cashier{}

	// next for medical department
	medical := &medical{}
	medical.setNext(cashier)

	doctor := &doctor{}
	doctor.setNext(medical)

	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "abc"}

	// execute chain of responsibility
	reception.execute(patient)
}
