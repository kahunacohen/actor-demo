package main

import (
	"fmt"
	"log"
	"time"
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

type Actor interface {
	Send(inbx chan Message)
	Receive()
}

type BaseActor struct {
	Inbx     chan Message
	handlers map[MessageType]func(msg Message)
}

func NewBaseActor() BaseActor {
	return BaseActor{
		Inbx:     make(chan Message, 10),
		handlers: make(map[MessageType]func(msg Message)),
	}
}

func (ba *BaseActor) RegisterHandler(msgType MessageType, handler func(msg Message)) {
	// TODO make threadsafe
	ba.handlers[msgType] = handler
}

func (ba *BaseActor) Receive() {
	for msg := range ba.Inbx {
		handler, found := ba.handlers[msg.Type]
		if found {
			handler(msg)
		} else {
			log.Printf("handler not found for type %v", msg.Type)
		}
	}
}

func (ba *BaseActor) Send(msg Message) {
	ba.Inbx <- msg
}

type EmployeeActor struct {
	BaseActor
}

type PatientActor struct {
	BaseActor
	Id int
}

func main() {
	// Instantiate employee and patient actors.
	// employee := EmployeeActor{NewBaseActor()}
	patient := PatientActor{BaseActor: NewBaseActor(), Id: 1}
	patient.RegisterHandler(CreateHomeVisitNotes, func(msg Message) {
		fmt.Printf("Doing some expensive operation on patient: %d using %v\n", msg.Id, msg.Payload)
		time.Sleep(1 * time.Second)
	})
	patient2 := PatientActor{BaseActor: NewBaseActor(), Id: 2}
	patient2.RegisterHandler(CreateHomeVisitNotes, func(msg Message) {
		fmt.Printf("Doing some expensive operation on patient: %d using %v\n", msg.Id, msg.Payload)
		time.Sleep(1 * time.Second)
	})

	go patient.Receive()
	go patient2.Receive()

	// Send some messages
	patient.Send(Message{Id: 1, Payload: "arbitrary data", Type: CreateHomeVisitNotes})
	patient2.Send(Message{Id: 2, Payload: "arbitrary data", Type: CreateHomeVisitNotes})
	time.Sleep(3 * time.Second)
}
