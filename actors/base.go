package actors

import "log"

type MessageType int

const (
	CreateHomeVisitNotes MessageType = iota
	UpdatePatientRecord
	ScheduleAppointment
	CancelAppointment
)

type Message struct {
	Id      int
	Payload interface{}
	Type    MessageType
}

type BaseActor struct {
	Inbx     chan Message
	handlers map[MessageType]func(msg Message)
}

func NewBaseActor() BaseActor {
	return BaseActor{
		Inbx:     make(chan Message, 10),
		handlers: make(map[MessageType]func(msg Message)),
	}
}

func (ba *BaseActor) RegisterHandler(msgType MessageType, handler func(msg Message)) {
	// TODO make threadsafe
	ba.handlers[msgType] = handler
}

func (ba *BaseActor) Receive() {
	for msg := range ba.Inbx {
		handler, found := ba.handlers[msg.Type]
		if found {
			handler(msg)
		} else {
			log.Printf("handler not found for type %v", msg.Type)
		}
	}
}

func (ba *BaseActor) Send(msg Message) {
	ba.Inbx <- msg
}
