package actors

import (
	"fmt"
	ms "github/kahunacohen/actor-demo/messages"
)

type ActorAddress string

type System struct {
	Base
	registry map[ActorAddress]Actor
}

func NewSystem() System {
	s := System{Base: NewBase()}
	s.RegisterHandler(ms.CreateActorMessage, s.CreateActorHandler)
	return s
}

func (s System) CreateActorHandler(msg ms.Message) {
	fmt.Println("Creating actor")
}
