package actors

import "log"

type MessageType int

const (
	MHomeVisit MessageType = iota
	UpdatePatientRecord
	ScheduleAppointment
	CancelAppointment
)

type Message struct {
	Id      int
	Payload interface{}
	Type    MessageType
}

type Base struct {
	Inbx     chan Message
	handlers map[MessageType]func(msg Message)
}

func NewBase() Base {
	return Base{
		Inbx:     make(chan Message, 10),
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
