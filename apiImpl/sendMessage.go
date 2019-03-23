package apiImpl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"telegramApi/apiClient"
	"telegramApi/common"
	"telegramApi/config"
	"time"
)

// conType是内容类型，值为Markdown,HTML
func SendMessage(msgContent string,conType string) string {
	telegramConfig := config.ConfigData.TelegramBot
	params := map[string]string{
		"chat_id":    telegramConfig.ChannelRoom, //
		"text":       msgContent,
		"parse_mode": conType,
	}
	req, err := apiClient.GetRequest("POST", "sendMessage", params)
	common.CheckError(err, 2)
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(req)
	common.CheckError(err, 2)
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	common.CheckError(err, 1)
	resStr := string(resBody)
	fmt.Println(resStr)
	return resStr
}
