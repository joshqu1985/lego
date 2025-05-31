package broker

var (
	MaxMessageCount = int64(10000)
)

// Message 消息结构
type Message struct {
	Payload    []byte            `json:"payload"`
	Key        string            `json:"key"`
	Properties map[string]string `json:"properties"`

	MessageId string `json:"-"` // 不需要传 发送后会自动填入
	Topic     string `json:"-"` // 不需要传 内部使用
}

func (this *Message) SetTag(tag string) {
	if this.Properties == nil {
		this.Properties = make(map[string]string)
	}
	this.Properties["tag"] = tag
}

func (this *Message) GetTag() string {
	tag, ok := this.Properties["tag"]
	if !ok {
		return ""
	}
	return tag
}
