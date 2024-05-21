package main

import (
	"time"

	"github/kahunacohen/actor-demo/actors"
)

func main() {
	// Create a few patient actors
	for i := 1; i <= 6; i++ {
		patient := actors.Patient{Base: actors.NewBase(), Id: i}
		patient.RegisterHandler(actors.MHomeVisit, actors.MakeHomeVisit)
		go patient.Receive()
		patient.Send(actors.Message{Id: i, Payload: "arbitrary data", Type: actors.MHomeVisit})
	}
	time.Sleep(3 * time.Second)
}
