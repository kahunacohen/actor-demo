package actors

import (
	ms "github/kahunacohen/actor-demo/messages"
	"log"
)

type Base struct {
	Inbx     chan ms.Message
	handlers map[ms.MessageType]func(msg ms.Message)
}

func NewBase() Base {
	return Base{
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
