package actors

import (
	ms "github/kahunacohen/actor-demo/messages"
	"log"
)

type System struct {
	Base
	registry map[string]Actor
}

func NewSystem() System {
	s := System{Base: NewBase()}
	s.registry = make(map[string]Actor)
	s.RegisterHandler(ms.CreateActorMessage, s.RegisterActorHandler)
	s.RegisterHandler(ms.PersistAllPatientsMessage, s.PersistAllPatientsHandler)
	return s
}

func (s System) RegisterActorHandler(msg ms.Message) {
	if patientData, ok := msg.Payload.(*PatientData); ok {
		patientActor := NewPatient(*patientData)
		s.registry[patientActor.Address] = &patientActor
	} else {
		log.Fatalln("failed to get patient data from payload")
	}
}

func (s System) PersistAllPatientsHandler(msg ms.Message) {
	for _, actor := range s.registry {
		patient, _ := actor.(*Patient)
		go patient.Receive()
		m, _ := ms.NewMessage(ms.PersistPatientMessage, 1)
		patient.Send(*m)
	}
}
