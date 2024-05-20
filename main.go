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
	fmt.Println("receive")
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
}

func main() {
	// Instantiate employee and patient actors.
	fmt.Println("main")
	employee := EmployeeActor{NewBaseActor()}
	patient := PatientActor{NewBaseActor()}
	patient.RegisterHandler(CreateHomeVisitNotes, func(msg Message) { fmt.Println("doing it") })

	// Start them listening
	go patient.Receive()

	// Send some messages
	employee.Send(Message{Id: 1, Payload: "saw him today", Type: CreateHomeVisitNotes})
	employee.Send(Message{Id: 2, Payload: "saw him today", Type: CreateHomeVisitNotes})
	employee.Send(Message{Id: 3, Payload: "saw him today", Type: CreateHomeVisitNotes})
	time.Sleep(5 * time.Second)
}
