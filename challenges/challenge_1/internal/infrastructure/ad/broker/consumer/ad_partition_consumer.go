package consumer

import (
	"challenges/challenge_1/internal/domain/ad"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"time"
)

type kafkaAdPartitionConsumer struct {
	consumer          sarama.Consumer
	partitionConsumer sarama.PartitionConsumer
}

func NewKafkaAdPartitionConsumer(brokers []string, topicName string) (ad.AdConsumer, error) {
	consumer, err := sarama.NewConsumer(brokers, sarama.NewConfig())
	if err != nil {
		log.Println(">>> Error while creating a kafka consumer", err)
		return nil, err
	}
	partitionConsumer, err := consumer.ConsumePartition(topicName, 0, sarama.OffsetNewest)
	if err != nil {
		log.Println(">>> Error while consuming a kafka partition", err)
		return nil, err
	}
	return &kafkaAdPartitionConsumer{
		consumer:          consumer,
		partitionConsumer: partitionConsumer,
	}, nil
}

func (kafkaAdPartitionConsumer *kafkaAdPartitionConsumer) close() {
	if err := kafkaAdPartitionConsumer.consumer.Close(); err != nil {
		log.Fatalln(err)
	}
	if err := kafkaAdPartitionConsumer.partitionConsumer.Close(); err != nil {
		log.Fatalln(err)
	}
}

func (kafkaAdPartitionConsumer *kafkaAdPartitionConsumer) StartListening() {
	log.Printf(">>> kafkaAdPartitionConsumer Kafka StartListening...")
	time.Sleep(10 * time.Second)

	defer kafkaAdPartitionConsumer.close()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
	notKilled := true
	for notKilled {
		select {
		case msg := <-kafkaAdPartitionConsumer.partitionConsumer.Messages():
			log.Printf(">>> kafkaAdPartitionConsumer message offset %d: '%s': '%s'\n", msg.Offset, string(msg.Key), string(msg.Value))
			consumed++
		case <-signals:
			notKilled = false
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}
