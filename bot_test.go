package telegramApi

import (
	"testing"
)

// 尝试发送telegram的消息给机器人
// https://core.telegram.org/bots/api#available-methods
func TestSendHtmlContent(t *testing.T) {
	var param = map[string]interface{} {
		"Key":"test_notify_callback",
		"Msg":"12312321312312321323",
	}
	SendHtmlContent(param)
}

func TestSendMarkdownContent(t *testing.T) {
	var param = map[string]interface{} {
		"Key":"test_notify_callback",
		"Msg":"12312321312312321323",
	}
	SendMarkdownContent(param)
}
