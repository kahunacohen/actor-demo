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
	return s
}

func (s System) RegisterActorHandler(msg ms.Message) {
	patientData, _ := msg.Payload.(PatientData)
	patientActor := NewPatient(patientData)
	log.Printf("Creating and registring patient actor: %s, %s, address: %s\n", patientActor.LastName, patientActor.FirstName, patientActor.Address)
	s.registry[patientActor.Address] = &patientActor
}
