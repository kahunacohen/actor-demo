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

	// Here is some patient data.
	patientsToCreate := []actors.PatientData{
		{FirstName: "Aaron", LastName: "Cohen", LocalID: "341077360"},
		{FirstName: "Courtney", LastName: "Cohen", LocalID: "341077361"},
		{FirstName: "Yochanan", LastName: "Harel", LocalID: "341077362"},
		{FirstName: "Harvey", LastName: "Weinstein", LocalID: "341077363"},
		{FirstName: "Donald", LastName: "Trump", LocalID: "341077364"},
		{FirstName: "Ronald", LastName: "Reagan", LocalID: "341077365"},
		{FirstName: "Richard", LastName: "Nixon", LocalID: "341077364"},
	}

	// Create the messages and send them to the system
	for i := range patientsToCreate {
		msg, _ := ms.NewMessage(ms.CreateActorMessage, &patientsToCreate[i])
		system.Send(*msg)
	}

	var payload interface{}
	requestPatientsMsg, _ := ms.NewMessage(ms.PersistAllPatientsMessage, payload)
	system.Send(*requestPatientsMsg)

	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
