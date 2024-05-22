package actors

import (
	"github/kahunacohen/actor-demo/messages"
	"log"
)

type Base struct {
	Inbx     chan messages.Message
	handlers map[messages.MessageType]func(msg messages.Message)
}

func NewBase() Base {
	return Base{
		Inbx:     make(chan messages.Message, 10),
		handlers: make(map[messages.MessageType]func(msg messages.Message)),
	}
}

func (b *Base) RegisterHandler(msgType messages.MessageType, handler func(msg messages.Message)) {
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

func (b *Base) Send(msg messages.Message) {
	b.Inbx <- msg
}
