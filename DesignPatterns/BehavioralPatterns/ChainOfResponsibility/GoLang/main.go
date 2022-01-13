package main

import "fmt"

//DEPARTMENT
type derpartment interface {
	execute(*patient)
	setNext(derpartment)
}

//RECEPTION
type reception struct {
	next derpartment
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *reception) setNext(next derpartment) {
	r.next = next
}

//DOCTOR

type doctor struct {
	next derpartment
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next derpartment) {
	d.next = next
}

//MEDICAL
type medical struct {
	next derpartment
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already give to patient")
		m.next.execute(p)
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *medical) setNext(next derpartment) {
	m.next = next
}

//CASHIER
type cashier struct {
	next derpartment
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *cashier) setNext(next derpartment) {
	c.next = next
}

//PATIENT
type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func main() {

	cashier := &cashier{}

	//Set next for medical department
	medical := &medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}
