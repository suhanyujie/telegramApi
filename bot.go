package telegramApi

import (
	"html/template"
	"log"
	"telegramApi/apiImpl"
)

type Wr struct {
	Content []byte
}

func (wr *Wr)Write(p []byte) (int,error) {
	copy(wr.Content, p)
	return len(p),nil
}

func SendMarkdownContent()  {
	//mdCon := "```\n " +
	//	"hello world" +
	//	"\n```"

	//apiImpl.SendMessage(mdCon)
}

func SendHtmlContent(assignVar map[string]interface{})  {
	ContentBox := new(Wr)
	t,err := template.ParseFiles("template/html/messageStyle1.html")
	if err!= nil {
		log.Println(err)
	}
	t.Execute(ContentBox, assignVar)
	apiImpl.SendMessage(string(ContentBox.Content),"HTML")
}
