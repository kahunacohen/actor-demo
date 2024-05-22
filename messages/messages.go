package messages

type MessageType int

const (
	CreatePatientMessage MessageType = iota
)

type Message struct {
	Id      int
	Payload interface{}
	Type    MessageType
}
