package producer

import (
	"challenges/challenge_1/internal/domain/ad"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
)

type kafkaAdSyncProducer struct {
	producer  sarama.SyncProducer
	brokers   []string
	topicName string
}

func NewKafkaAdSyncProducer(brokers []string, topicName string) (ad.AdProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		fmt.Println(">>> Error while consuming a kafka partition", err)
		return nil, err
	}
	return &kafkaAdSyncProducer{
		producer:  producer,
		brokers:   brokers,
		topicName: topicName,
	}, nil
}

func (kafkaAdSyncProducer *kafkaAdSyncProducer) SendMessage(message ad.BrokerMessage) (partition int32, offset int64, err error) {
	// TODO I'm going to create an async producer similar to this
	// TODO remember if I have to use the context
	adMessage, err := json.Marshal(message.Ad)
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := &sarama.ProducerMessage{
		Topic: kafkaAdSyncProducer.topicName,
		Value: sarama.StringEncoder(adMessage),
		Key:   sarama.StringEncoder(message.Ad.Id.String()),
	}
	fmt.Printf(">>> Sending Kafka message to topic %s.... %s\n", kafkaAdSyncProducer.topicName, string(adMessage))
	return kafkaAdSyncProducer.producer.SendMessage(msg)
}
