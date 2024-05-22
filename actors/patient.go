package actors

import (
	"fmt"
	"github/kahunacohen/actor-demo/messages"
	"time"
)

type Patient struct {
	Base
	Id int
}

func MakeHomeVisit(msg messages.Message) {
	fmt.Printf("Creating a home visit for patient: %d using %v\n", msg.Id, msg.Payload)
	time.Sleep(1 * time.Second)
}
