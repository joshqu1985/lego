package broker

// Message 消息结构.
type Message struct {
	Properties map[string]string `json:"properties"`
	Key        string            `json:"key"`
	Payload    []byte            `json:"payload"`
	topic      string            `json:"-"`
	messageId  string            `json:"-"`
}

// 内存消息队列和redis stream消息队列的最大消息数量
var MaxMessageCount = int64(10000)

func NewMessage(payload []byte) *Message {
	return &Message{
		Payload: payload,
	}
}

// SetKey 设置消息key
func (m *Message) SetKey(key string) {
	m.Key = key
}

// GetKey 获取消息key
func (m *Message) GetKey() string {
	return m.Key
}

// SetTopic 设置消息topic
func (m *Message) SetTopic(topic string) {
	m.topic = topic
}

// GetTopic 获取消息topic
func (m *Message) GetTopic() string {
	return m.topic
}

// SetMsgId 设置消息id
func (m *Message) SetMsgId(msgId string) {
	m.messageId = msgId
}

// GetMsgId 获取消息id
func (m *Message) GetMsgId() string {
	return m.messageId
}

// SetPayload 设置消息payload
func (m *Message) SetPayload(payload []byte) {
	m.Payload = payload
}

// GetPayload 获取消息payload
func (m *Message) GetPayload() []byte {
	return m.Payload
}

// SetTag 设置消息tag
func (m *Message) SetTag(tag string) {
	if m.Properties == nil {
		m.Properties = make(map[string]string)
	}
	m.Properties["tag"] = tag
}

// GetTag 获取消息tag
func (m *Message) GetTag() string {
	tag, ok := m.Properties["tag"]
	if !ok {
		return ""
	}
	return tag
}

// SetProperty 设置消息属性
func (m *Message) SetProperty(key, value string) {
	if m.Properties == nil {
		m.Properties = make(map[string]string)
	}
	m.Properties[key] = value
}

// GetProperty 获取消息属性
func (m *Message) GetProperty(key string) string {
	return m.Properties[key]
}

// SetProperties 设置消息属性
func (m *Message) SetProperties(properties map[string]string) {
	m.Properties = properties
}

// GetProperties 获取消息属性
func (m *Message) GetProperties() map[string]string {
	return m.Properties
}

// RangeProperty 遍历消息属性
func (m *Message) RangeProperty(f func(key, val string)) {
	for k, v := range m.Properties {
		f(k, v)
	}
}
