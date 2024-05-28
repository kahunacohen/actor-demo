package actors

import (
	ms "github/kahunacohen/actor-demo/messages"
	"log"
	"time"
)

type PatientData struct {
	FirstName string
	LastName  string
	LocalID   string
	Id        int
}

type Patient struct {
	Base
	PatientData
}

func NewPatient(data PatientData) Patient {
	p := Patient{Base: NewBase(), PatientData: data}
	//p.RegisterHandler(ms.CreatePatientMessage, p.CreatePatientHandler)
	return p
}

func (p Patient) CreatePatientHandler(msg ms.Message) {
	log.Printf("Creating patient actor: %v\n", p.Address)
	time.Sleep(1 * time.Second)
}
