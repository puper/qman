package producer

import (
	"log"

	"github.com/Shopify/sarama"
)

type Config struct {
	Brokers []string
}

type Producer struct {
	config *Config
	client sarama.SyncProducer
}

func New(cfg *Config) (*Producer, error) {
	producer := &Producer{
		config: cfg,
	}
	sc := sarama.NewConfig()
	sc.Producer.Return.Successes = true
	log.Println(cfg.Brokers)
	client, err := sarama.NewSyncProducer(cfg.Brokers, sc)
	if err != nil {
		return nil, err
	}
	producer.client = client
	return producer, nil
}

func (this *Producer) Stop() error {
	return nil
}

func (this *Producer) Put(msg *Message) error {
	pm := &sarama.ProducerMessage{
		Topic:     msg.Topic,
		Partition: -1,
	}
	if msg.Key != "" {
		pm.Key = sarama.StringEncoder(msg.Key)
	}
	pm.Value = sarama.ByteEncoder(msg.Encode())
	_, _, err := this.client.SendMessage(pm)
	return err
}
