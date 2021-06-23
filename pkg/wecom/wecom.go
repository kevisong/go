package wecom

import (
	"encoding/json"

	"github.com/KEVISONG/go/pkg/http"
)

// Robot interface
type Robot interface{}

// WCRobot WeComRobot
type WCRobot struct {
	Webhook string
}

// NewWeComRobot factory
func NewWeComRobot(webhook string) (r Robot) {
	return &WCRobot{webhook}
}

func send(webhook, content string) ([]byte, error) {
	messageJSON, err := json.Marshal(&Message{MsgType: "markdown", Markdown: MessageMarkdown{Content: content}})
	if err != nil {
		return nil, err
	}
	resp, err := http.SetHeaders(map[string]string{"content-type": "application/json"}).Post(webhook, messageJSON)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Send sends the content
func (w *WCRobot) Send(content string) ([]byte, error) {
	return send(w.Webhook, content)
}

// Send sends the content
func Send(webhook, content string) ([]byte, error) {
	return send(webhook, content)
}
