package actors

import (
	ms "github/kahunacohen/actor-demo/messages"
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
	time.Sleep(1 * time.Second)
}
