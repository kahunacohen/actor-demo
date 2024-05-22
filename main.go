package main

import (
	"time"

	"github/kahunacohen/actor-demo/actors"
	ms "github/kahunacohen/actor-demo/messages"
)

func main() {
	// Create a few patient actors
	for i := 1; i <= 25; i++ {
		patient := actors.NewPatient(i)
		patient.RegisterHandler(ms.CreateHomeVisit, actors.MakeHomeVisit)
		go patient.Receive()
		patient.Send(ms.Message{Id: i, Payload: "arbitrary data", Type: ms.CreateHomeVisit})
	}
	time.Sleep(1 * time.Second)
}
