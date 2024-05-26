package main

import (
	"fmt"
	"log"

	"github/kahunacohen/actor-demo/actors"
	ms "github/kahunacohen/actor-demo/messages"
)

func main() {
	// Create 25 patients
	for i := 1; i <= 25; i++ {
		patient := actors.NewPatient(i)
		go patient.Receive()
		msg, err := ms.NewMessage(ms.CreatePatientMessage, "arbitrary data")
		if err != nil {
			log.Fatal("Could not create message")
		}
		patient.Send(*msg)
	}
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
