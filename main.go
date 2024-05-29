package main

import (
	"fmt"

	"github/kahunacohen/actor-demo/actors"
)

func main() {
	// Create 25 patients
	system := actors.NewSystem()
	go system.Receive()

	// Here is some patient data.
	patientsToCreate := []actors.PatientData{
		{FirstName: "Aaron", LastName: "Cohen", LocalID: "341077360"},
		{FirstName: "Courtney", LastName: "Cohen", LocalID: "341077361"},
		{FirstName: "Yochanan", LastName: "Harel", LocalID: "341077362"},
		{FirstName: "Harvey", LastName: "Weinstein", LocalID: "341077363"},
	}

	// Create the messages and send them to the system
	for i := range patientsToCreate {
		msg, _ := actors.NewMessage(actors.CreateActorMessage, &patientsToCreate[i])
		system.Send(*msg)
	}

	// Now get all actors back and persist each one
	var payload interface{}
	requestPatientsMsg, _ := actors.NewMessage(actors.RequestAllPatientsMessage, payload)
	system.Send(*requestPatientsMsg)

	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
