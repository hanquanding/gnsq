package gnsq

import (
	"github.com/bitly/go-nsq"
	"log"
	"sync"
)

// 消费者
type Consumer struct {
	Topic   string
	Channel string
	Network string
	Config  *nsq.Config
	Handler nsq.Handler
}

func (this *Consumer) Listen() {
	c, err := nsq.NewConsumer(this.Topic, this.Channel, this.Config)
	if err != nil {
		log.Println(err)
		return
	}

	c.AddHandler(this.Handler)
	if err := c.ConnectToNSQD(this.Network); err != nil {
		log.Println(err)
	}
}

func Listen(topic string, channel string, network string, handle nsq.Handler) {
	var wg sync.WaitGroup
	wg.Add(1)

	retentionConsumer := &Consumer{
		Topic:   topic,
		Channel: channel,
		Network: network,
		Config:  nsq.NewConfig(),
		Handler: handle,
	}
	retentionConsumer.Listen()
	wg.Wait()
}
