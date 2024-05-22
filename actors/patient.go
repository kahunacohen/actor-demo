package actors

import (
	ms "github/kahunacohen/actor-demo/messages"
	"log"
	"time"
)

type Patient struct {
	Base
	Id int
}

func NewPatient(id int) Patient {
	p := Patient{Base: NewBase(), Id: id}
	p.RegisterHandler(ms.CreatePatientMessage, CreatePatientHandler)
	return p
}

func CreatePatientHandler(msg ms.Message) {
	log.Printf("Creating patient with message ID: %s using %v\n", msg.Id, msg.Payload)
	time.Sleep(1 * time.Second)
}
