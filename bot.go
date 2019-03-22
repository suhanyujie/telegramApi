package telegramApi

import (
	"log"
	"html/template"
	"github.com/suhanyujie/telegramApi/apiImpl"
	"bytes"
)

func test()  {
	log.Println("test...")
}

type Wr struct {
	Content []byte
}

func (wr *Wr)Write(p []byte) (int,error) {
	copy(wr.Content, p)
	wr.Content = bytes.TrimSpace(wr.Content)
	return len(p),nil
}

func SendMarkdownContent()  {
	//mdCon := "```\n " +
	//	"hello world" +
	//	"\n```"

	//apiImpl.SendMessage(mdCon)
}

func SendHtmlContent(assignVar map[string]interface{})  {
	box := make([]byte, 1024)
	ContentBox := &Wr{
		box,
	}
	t,err := template.ParseFiles("template/html/messageStyle1.html")
	if err!= nil {
		log.Println(err)
	}
	t.Execute(ContentBox, assignVar)
	log.Println(string(ContentBox.Content))
	apiImpl.SendMessage(string(ContentBox.Content),"HTML")
}
