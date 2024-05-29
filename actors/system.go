package actors

import (
	"log"
)

type System struct {
	Base
	registry map[string]Actor
}

func NewSystem() System {
	s := System{Base: NewBase()}
	s.registry = make(map[string]Actor)
	s.RegisterHandler(CreateActorMessage, s.RegisterActorHandler)
	s.RegisterHandler(RequestAllPatientsMessage, s.RequestAllPatientsHandler)
	return s
}

func (s System) RegisterActorHandler(msg Message) {
	if patientData, ok := msg.Payload.(*PatientData); ok {
		patientActor := NewPatient(*patientData)
		s.registry[patientActor.Address] = &patientActor
	} else {
		log.Fatalln("failed to get patient data from payload")
	}
}

func (s System) RequestAllPatientsHandler(msg Message) {
	for address, actor := range s.registry {
		patient, _ := actor.(*Patient)
		log.Printf("address: %s, actor: %v\n", address, patient.LastName)
	}
}
