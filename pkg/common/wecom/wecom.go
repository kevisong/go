package wecom

import (
	"encoding/json"

	"github.com/KEVISONG/go/pkg/common/http"
)

type Robot interface{}

type WeComRobot struct {
	Webhook string
}

func NewWeComRobot(webhook string) (r Robot) {
	return &WeComRobot{webhook}
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

func (w *WeComRobot) Send(content string) ([]byte, error) {
	return send(w.Webhook, content)
}

func Send(webhook, content string) ([]byte, error) {
	return send(webhook, content)
}
