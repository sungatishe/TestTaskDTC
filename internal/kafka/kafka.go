package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(address string, topic string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(address),
			Topic:        topic,
			Balancer:     &kafka.LeastBytes{},
			Async:        true,
			RequiredAcks: kafka.RequireAll,
		},
	}
}

func (p *Producer) Publish(key, value []byte) error {
	message := kafka.Message{
		Key:   key,
		Value: value,
	}
	err := p.writer.WriteMessages(context.Background(), message)
	if err != nil {
		return err
	}
	log.Printf("Message published: key=%s value=%s\n", key, value)
	return nil
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
