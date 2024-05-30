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
	p.RegisterHandler(ms.PersistPatientMessage, p.PersistPatientHandler)
	return p
}

func (p Patient) PersistPatientHandler(msg ms.Message) {
	log.Printf("saving %s, %s to db...\n", p.PatientData.LastName, p.PatientData.FirstName)
	time.Sleep(1 * time.Second)
	log.Printf("done saving %s, %s to db\n", p.PatientData.LastName, p.PatientData.FirstName)
}
