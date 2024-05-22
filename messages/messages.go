package messages

type MessageType int

const (
	CancelAppointment MessageType = iota
	CreateHomeVisit
	ScheduleAppointment
	UpdatePatientRecord
)

type Message struct {
	Id      int
	Payload interface{}
	Type    MessageType
}
