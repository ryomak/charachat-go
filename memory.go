package charachat

type Memory interface {
	GetAllMessage() []Message
	AddMessage(Message)
	DeleteAll()
}
type LocalMemory struct {
	Messages []Message
}

func NewLocalMemory() *LocalMemory {
	return &LocalMemory{}
}

func (m *LocalMemory) GetAllMessage() []Message {
	return m.Messages
}

func (m *LocalMemory) AddMessage(msg Message) {
	m.Messages = append(m.Messages, msg)
}

func (m *LocalMemory) DeleteAll() {
	m.Messages = []Message{}
}
