package main

import (
	"fmt"
)

type MessageType int

const (
	CreateHomeVisitNotes MessageType = iota
	UpdatePatientRecord
	ScheduleAppointment
	CancelAppointment
)

type Message struct {
	Id      int
	Payload interface{}
	Type    MessageType
}

type Actor struct {
	inbx chan Message
}

func NewActor() Actor {
	return Actor{inbx: make(chan Message)}
}

func (a *Actor) Receive() {
	for msg := range a.inbx {
		fmt.Printf("receiving: %v", msg)
	}
}

func (a *Actor) Send(msg Message) {
	a.inbx <- msg
}

type EmployeeActor struct {
	Actor
	Name string
}

func (e *EmployeeActor) UpdateHomeVisitNotes(patient PatientActor, msg Message) {
	//patient.Send(fmt.Sprintf("updating home visit notes: %s\n", notes))
}

type PatientActor struct {
	Actor
	Name string
}

func main() {
	employee := EmployeeActor{Actor: NewActor(), Name: "Shai"}
	patient := PatientActor{Actor: NewActor(), Name: "Aaron"}
	go patient.Receive()

	// Maybe an actor should be an interface that satisfies Receives and dispatches. But how to avoid duplication?
	employee.Send(patient, Message{Id: 1, Payload: "saw him today", Type: CreateHomeVisitNotes})
	employee.Send(patient, Message{Id: 2, Payload: "she wasn't feeling well", Type: CreateHomeVisitNotes})
}
