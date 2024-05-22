package main

import (
	"log"
	"time"

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
	time.Sleep(1 * time.Second)
}
