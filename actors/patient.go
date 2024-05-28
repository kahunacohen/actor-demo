package actors

import (
	ms "github/kahunacohen/actor-demo/messages"
	"log"
	"time"
)

type Patient struct {
	Base
	FirstName string
	LastName  string
	LocalID   string
	Id        int
}

func NewPatient(id int) Patient {
	p := Patient{Base: NewBase(), Id: id}
	p.RegisterHandler(ms.CreatePatientMessage, p.CreatePatientHandler)
	return p
}

func (p Patient) CreatePatientHandler(msg ms.Message) {
	log.Printf("Creating patient actor: %v\n", p.Address)
	time.Sleep(1 * time.Second)
}
