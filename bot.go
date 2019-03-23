package telegramApi

import (
	"bytes"
	"html/template"
	"log"
	"telegramApi/apiImpl"
)

type Wr struct {
	Content []byte
}

func (wr *Wr)Write(p []byte) (int,error) {
	copy(wr.Content, p)
	wr.Content = bytes.TrimSpace(wr.Content)
	return len(p),nil
}

func SendMarkdownContent(assignVar map[string]interface{})  {
	var docBuf bytes.Buffer
	t,err := template.ParseFiles("template/markdown/mdStyle1.md")
	if err!= nil {
		log.Println(err)
	}
	t.Execute(&docBuf, assignVar)
	contentStr := docBuf.String()
	apiImpl.SendMessage(contentStr,"Markdown")
}

func SendHtmlContent(assignVar map[string]interface{})  {
	// 默认分配2048字节的缓冲区
	var doc bytes.Buffer
	t,err := template.ParseFiles("template/html/messageStyle1.html")
	if err!= nil {
		log.Println(err)
	}
	t.Execute(&doc, assignVar)
	contentStr := doc.String()
	apiImpl.SendMessage(contentStr,"HTML")
}
