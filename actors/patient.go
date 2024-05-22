package actors

import (
	"github/kahunacohen/actor-demo/messages"
	"log"
	"time"
)

type Patient struct {
	Base
	Id int
}

func NewPatient(id int) Patient {
	return Patient{Base: NewBase(), Id: id}
}

func MakeHomeVisit(msg messages.Message) {
	log.Printf("Creating a home visit for patient: %d using %v\n", msg.Id, msg.Payload)
	time.Sleep(1 * time.Second)
}
