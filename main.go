package main

import (
	"time"

	"github/kahunacohen/actor-demo/actors"
	ms "github/kahunacohen/actor-demo/messages"
)

func main() {
	// Create 25 patients
	for i := 1; i <= 25; i++ {
		patient := actors.NewPatient(i)
		go patient.Receive()
		patient.Send(ms.Message{Id: i, Payload: "arbitrary data", Type: ms.CreatePatientMessage})
	}
	time.Sleep(1 * time.Second)
}
