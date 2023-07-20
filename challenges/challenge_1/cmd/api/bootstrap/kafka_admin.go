package bootstrap

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

func BoostrapKafkaCustomConfig(kafkaConsumerConfig KafkaConsumerConfig) {
	fmt.Println(">>> Creating Kafka Consumers...")
	time.Sleep(3 * time.Second)
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	fmt.Println(">>> Creating Kafka cluster...")
	admin, err := sarama.NewClusterAdmin(kafkaConsumerConfig.Brokers(), config)
	if err != nil {
		log.Printf(">>> Error while creating cluster admin: ", err.Error())
	}
	defer func() { _ = admin.Close() }()
	fmt.Println(">>> Creating Kafka topic...", kafkaConsumerConfig.TopicName())
	err = admin.CreateTopic(kafkaConsumerConfig.TopicName(), &sarama.TopicDetail{
		NumPartitions:     4,
		ReplicationFactor: 1,
	}, false)
	if err != nil {
		log.Printf(">>> Warning while creating topic: ", err.Error())
	}
}
