package gnsq

import "github.com/bitly/go-nsq"

// 生产者
func getProducer(network string, config *nsq.Config) (*nsq.Producer, error) {
	p, err := nsq.NewProducer(network, nsq.NewConfig())
	if err != nil {
		return nil, err
	}
	return p, nil
}

// 同步
func Publish(network, topic string, body []byte) error {
	producer, err := getProducer(network, nsq.NewConfig())
	if err != nil {
		return err
	}
	producer.Publish(topic, body)
	producer.Stop()
	return nil
}

// 异步
func PublishAsync(network, topic string, body []byte, done chan *nsq.ProducerTransaction) error {
	producer, err := getProducer(network, nsq.NewConfig())
	if err != nil {
		return err
	}
	err = producer.PublishAsync(topic, body, done)
	if err != nil {
		return err
	}
	producer.Stop()
	return nil
}
