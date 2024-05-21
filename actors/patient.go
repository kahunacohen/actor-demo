package actors

import (
	"fmt"
	"time"
)

type Patient struct {
	Base
	Id int
}

func MakeHomeVisit(msg Message) {
	fmt.Printf("Creating a home visit for patient: %d using %v\n", msg.Id, msg.Payload)
	time.Sleep(1 * time.Second)
}
