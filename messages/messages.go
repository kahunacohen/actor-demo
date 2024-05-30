package messages

import "github.com/google/uuid"

type MessageType int

const (
	CreatePatientMessage MessageType = iota
	CreateActorMessage
	PersistAllPatientsMessage
	PersistPatientMessage
)

type Message struct {
	Id      string
	Payload interface{}
	Type    MessageType
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
