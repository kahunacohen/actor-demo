package main

import (
	"fmt"
	"time"

	"github/kahunacohen/actor-demo/actors"
)

func main() {
	// Create a few patient actors
	for i := 0; i <= 5; i++ {
		p := actors.PatientActor{BaseActor: actors.NewBaseActor(), Id: i}
		p.RegisterHandler(actors.CreateHomeVisitNotes, func(msg actors.Message) {
			fmt.Printf("Doing some expensive operation on patient: %d using %v\n", msg.Id, msg.Payload)
			time.Sleep(1 * time.Second)
		})
		go p.Receive()
		p.Send(actors.Message{Id: i, Payload: "arbitrary data", Type: actors.CreateHomeVisitNotes})
	}
	time.Sleep(3 * time.Second)
}
