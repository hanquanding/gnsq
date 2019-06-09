package main

import (
	"flag"
	"github.com/bitly/go-nsq"
	"gnsq"
	"log"
	"sync"
)

var channel string

type ComsumerHandler struct {
	Topic string
}

func init() {
	flag.StringVar(&channel, "c", "default", "channel name")
	flag.Parse()
}

func (ch *ComsumerHandler) HandleMessage(msg *nsq.Message) error {
	log.Println("msg: ", string(msg.Body))
	return nil
}

func main() {
	var wg sync.WaitGroup
	gnsq.Listen("test", channel, "127.0.0.1:4150", &ComsumerHandler{Topic: "test"})
	wg.Add(1)
	wg.Wait()
}
