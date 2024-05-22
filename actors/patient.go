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
	p.RegisterHandler(ms.CreateHomeVisit, CreateHomeVisitHandler)
	return p
}

func CreateHomeVisitHandler(msg ms.Message) {
	log.Printf("Creating a home visit for patient: %d using %v\n", msg.Id, msg.Payload)
	time.Sleep(1 * time.Second)
}
