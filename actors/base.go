package actors

import (
	"fmt"
	ms "github/kahunacohen/actor-demo/messages"
	"log"
	"os"

	"github.com/google/uuid"
)

type Actor interface {
	Receive()
	Send(msg ms.Message)
}

type Base struct {
	Address  string
	Inbx     chan ms.Message
	handlers map[ms.MessageType]func(msg ms.Message)
}

func NewBase() Base {
	host, _ := os.Hostname()
	uuid, _ := uuid.NewRandom()
	return Base{
		Address:  fmt.Sprintf("%s@%s", uuid, host),
		Inbx:     make(chan ms.Message, 16),
		handlers: make(map[ms.MessageType]func(msg ms.Message)),
	}
}

func (b *Base) RegisterHandler(msgType ms.MessageType, handler func(msg ms.Message)) {
	// TODO make threadsafe
	b.handlers[msgType] = handler
}

func (b *Base) Receive() {
	for msg := range b.Inbx {
		handler, found := b.handlers[msg.Type]
		if found {
			handler(msg)
		} else {
			log.Printf("handler not found for type %v", msg.Type)
		}
	}
}

func (b *Base) Send(msg ms.Message) {
	b.Inbx <- msg
}
