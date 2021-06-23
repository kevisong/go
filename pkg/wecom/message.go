package wecom

// Message to be sent by robot
type Message struct {
	MsgType  string          `json:"msgtype"`
	Markdown MessageMarkdown `json:"markdown"`
}

// MessageMarkdown markdown content
type MessageMarkdown struct {
	Content string `json:"content"`
}
