package actors

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

type MessageType int

const (
	CreatePatientMessage MessageType = iota
	CreateActorMessage
	RequestAllPatientsMessage
)

type Message struct {
	Id      string
	Payload interface{}
	Type    MessageType
	ReplyTo Actor
}

func NewMessage[T any](messageType MessageType, payload T) (*Message, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &Message{
		Id:      uuid.String(),
		Payload: payload,
		Type:    messageType,
	}, nil

}

type Actor interface {
	Receive()
	Send(msg Message)
}

type Base struct {
	Address  string
	Inbx     chan Message
	handlers map[MessageType]func(msg Message)
}

func NewBase() Base {
	host, _ := os.Hostname()
	uuid, _ := uuid.NewRandom()
	return Base{
		Address:  fmt.Sprintf("%s@%s", uuid, host),
		Inbx:     make(chan Message, 16),
		handlers: make(map[MessageType]func(msg Message)),
	}
}

func (b *Base) RegisterHandler(msgType MessageType, handler func(msg Message)) {
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

func (b *Base) Send(msg Message) {
	b.Inbx <- msg
}
