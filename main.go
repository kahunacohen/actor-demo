package main

import (
	"fmt"
)

type Actor struct {
	inbx chan string
}

func NewActor() Actor {
	return Actor{inbx: make(chan string)}
}

func (a *Actor) Receive() {
	for msg := range a.inbx {
		fmt.Printf("receiving: %v", msg)
	}
}

func (a *Actor) Send(msg string) {
	a.inbx <- msg
}

type EmployeeActor struct {
	Actor
	Name string
}

func (e *EmployeeActor) UpdateHomeVisitNotes(patient PatientActor, notes string) {
	patient.Send(fmt.Sprintf("updating home visit notes: %s\n", notes))
}

type PatientActor struct {
	Actor
	Name string
}

func main() {
	employee := EmployeeActor{Actor: NewActor(), Name: "Shai"}
	patient := PatientActor{Actor: NewActor(), Name: "Aaron"}
	go patient.Receive()
	employee.UpdateHomeVisitNotes(patient, "saw him today")
}
