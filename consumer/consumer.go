package consumer

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

type activityConsumer struct {
	Reader *kafka.Reader
}

func NewActivityConsumer(url, topic, groupId string) *activityConsumer {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{url},
		Topic:   topic,
		GroupID: groupId,
		Dialer:  dialer,
		// MinBytes: 10e3, // 10KB
		// MaxBytes: 10e6, // 10MB
		// MaxWait:  time.Microsecond * 10,
	})
	return &activityConsumer{
		Reader: reader,
	}
}

func (c *activityConsumer) ReadMessage(ctx context.Context) ([]byte, []byte, error) {
	msg, err := c.Reader.ReadMessage(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("could not read the message:%s\n", err.Error())
	}
	fmt.Printf("received message: %s - %s\n", string(msg.Key), string(msg.Value))
	return msg.Key, msg.Value, nil
}
