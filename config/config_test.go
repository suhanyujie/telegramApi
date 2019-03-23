package config

import (
	"log"
	"testing"
)

func TestParsConfig(t *testing.T) {
	var path = "./config.json"
	data,err := ParseConfig(path)
	if err!= nil {
		t.Error(err)
	}
	log.Println(data)
}
