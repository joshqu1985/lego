package broker

// Message 消息结构.
type Message struct {
	Properties map[string]string `json:"properties"`
	Key        string            `json:"key"`
	MessageId  string            `json:"-"`
	Topic      string            `json:"-"`
	Payload    []byte            `json:"payload"`
}

var MaxMessageCount = int64(10000)

func (m *Message) SetTag(tag string) {
	if m.Properties == nil {
		m.Properties = make(map[string]string)
	}
	m.Properties["tag"] = tag
}

func (m *Message) GetTag() string {
	tag, ok := m.Properties["tag"]
	if !ok {
		return ""
	}

	return tag
}
