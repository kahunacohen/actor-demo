package main

import (
	"fmt"

	"github/kahunacohen/actor-demo/actors"
	ms "github/kahunacohen/actor-demo/messages"
)

func main() {
	// Create 25 patients
	system := actors.NewSystem()
	go system.Receive()

	msg, _ := ms.NewMessage(ms.CreateActorMessage, actors.PatientData{FirstName: "Aaron", LastName: "Cohen", LocalID: "341077360"})
	msg2, _ := ms.NewMessage(ms.CreateActorMessage, actors.PatientData{FirstName: "Courtney", LastName: "Cohen", LocalID: "341077361"})
	msg3, _ := ms.NewMessage(ms.CreateActorMessage, actors.PatientData{FirstName: "Yochanan", LastName: "Harel", LocalID: "341077362"})

	system.Send(*msg)
	system.Send(*msg2)
	system.Send(*msg3)

	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}

// package main
//
// import (
// 	"fmt"
// )
//
// // Message represents a message that can be sent to actors
// type Message interface{}
//
// // CreatePatient is a message to create a patient actor
// type CreatePatient struct {
// 	Name string
// }
//
// // PatientActor represents a patient actor
// type PatientActor struct {
// 	Name string
// }
//
// // Receive handles incoming messages for PatientActor
// func (patient *PatientActor) Receive(msg Message) {
// 	switch m := msg.(type) {
// 	case string:
// 		fmt.Printf("Patient %s received message: %s\n", patient.Name, m)
// 	}
// }
//
// // SystemActor represents a system actor that creates patient actors
// type SystemActor struct{}
//
// // CreatePatientActor creates a new patient actor with the given name
// func (system *SystemActor) CreatePatientActor(name string) *PatientActor {
// 	patient := &PatientActor{Name: name}
// 	go patient.Receive("Hello from SystemActor")
// 	return patient
// }
//
// func main() {
// 	// Create the system actor
// 	system := &SystemActor{}
//
// 	// Create a patient actor
// 	patient := system.CreatePatientActor("John Doe")
//
// 	// Wait for user input to exit
// 	fmt.Println("Press Enter to exit...")
// 	fmt.Scanln()
// }
//
