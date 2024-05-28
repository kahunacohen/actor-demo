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
	s.RegisterHandler(ms.RequestAllPatientsMessage, s.RequestAllPatientsHandler)
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

func (s System) RequestAllPatientsHandler(msg ms.Message) {
	for address, actor := range s.registry {
		patient, _ := actor.(*Patient)
		log.Printf("address: %s, actor: %v\n", address, patient.LastName)
	}
}
