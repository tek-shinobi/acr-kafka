package producer

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type activityProducer struct {
	Writer *kafka.Writer
}

func CreateTopic(url, topic string) error {
	conn, err := kafka.Dial("tcp", url)
	if err != nil {
		return err
	}
	defer conn.Close()
	topicConfig := kafka.TopicConfig{Topic: topic, NumPartitions: 1, ReplicationFactor: 1}
	return conn.CreateTopics(topicConfig)
}

func NewActivityProducer(url, topic string) *activityProducer {

	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{url},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		Dialer:   dialer,
	})
	return &activityProducer{
		Writer: kafkaWriter,
	}
}

func (ap *activityProducer) WriteMessage(ctx context.Context, topic string, key, value []byte) error {
	msg := kafka.Message{
		Key:   []byte(key),
		Value: value,
	}
	return ap.Writer.WriteMessages(ctx, msg)

}
