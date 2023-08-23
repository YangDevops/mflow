package wechat

import "github.com/infraboard/mflow/apps/task"

const (
	URL_PREFIX = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"
)

const (
	MarkdownMessage = "markdown"
)

type MessageType string

func NewStepMarkdownMessage(s *task.JobTask) *Message {
	return &Message{
		MsgType: MarkdownMessage,
		Markdown: &MarkDownContent{
			Content: s.ShowTitle() + "\n" + s.MarkdownContent(),
		},
	}
}

// 群机器人配置说明: https://work.weixin.qq.com/api/doc/90000/90136/91770
type Message struct {
	MsgType  MessageType      `json:"msgtype"`
	Markdown *MarkDownContent `json:"markdown"`
}

type MarkDownContent struct {
	Content string `json:"content"`
}
