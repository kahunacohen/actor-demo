package main

import (
	"time"

	"github/kahunacohen/actor-demo/actors"
	"github/kahunacohen/actor-demo/messages"
)

func main() {
	// Create a few patient actors
	for i := 1; i <= 6; i++ {
		patient := actors.Patient{Base: actors.NewBase(), Id: i}
		patient.RegisterHandler(messages.CreateHomeVisit, actors.MakeHomeVisit)
		go patient.Receive()
		patient.Send(messages.Message{Id: i, Payload: "arbitrary data", Type: messages.CreateHomeVisit})
	}
	time.Sleep(3 * time.Second)
}
