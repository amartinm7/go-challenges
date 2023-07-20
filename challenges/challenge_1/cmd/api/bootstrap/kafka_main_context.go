package bootstrap

import (
	"challenges/challenge_1/internal/domain/ad"
	"challenges/challenge_1/internal/infrastructure/ad/broker/consumer"
	"golang.org/x/exp/maps"
)

type kafkaContext struct {
	kafkaConsumerConfig KafkaConsumerConfig
	kafkaProducerConfig KafkaProducerConfig
	kafkaConsumers      map[string]ad.AdConsumer
}

func (k *kafkaContext) KafkaProducerConfig() KafkaProducerConfig {
	return k.kafkaProducerConfig
}

func (k *kafkaContext) KafkaConsumerConfig() KafkaConsumerConfig {
	return k.kafkaConsumerConfig
}

func (k *kafkaContext) KafkaConsumers() map[string]ad.AdConsumer {
	return k.kafkaConsumers
}

func NewKafkaContext(envHandler EnvHandler) KafkaContext {
	consumerConfig := NewKafkaConsumerConfig(envHandler)
	BoostrapKafkaCustomConfig(consumerConfig)
	return &kafkaContext{
		kafkaConsumerConfig: consumerConfig,
		kafkaProducerConfig: NewKafkaProducerConfig(envHandler),
		kafkaConsumers:      registerConsumers(consumerConfig),
	}
}

func registerKafkaAdConsumerGroup(kafkaConsumerConfig KafkaConsumerConfig) map[string]ad.AdConsumer {
	newKafkaAdConsumer, err := consumer.NewKafkaAdConsumerGroup(kafkaConsumerConfig.Brokers(), kafkaConsumerConfig.TopicName(), kafkaConsumerConfig.ConsumerGroup())
	if err != nil {
		return map[string]ad.AdConsumer{}
	}
	return map[string]ad.AdConsumer{"kafkaAdConsumerGroup": newKafkaAdConsumer}
}

func registerKafkaAdPartitionConsumer(kafkaConsumerConfig KafkaConsumerConfig) map[string]ad.AdConsumer {
	newKafkaAdConsumer, err := consumer.NewKafkaAdPartitionConsumer(kafkaConsumerConfig.Brokers(), kafkaConsumerConfig.TopicName())
	if err != nil {
		return map[string]ad.AdConsumer{}
	}
	return map[string]ad.AdConsumer{"kafkaAdPartitionConsumer": newKafkaAdConsumer}
}

func registerConsumers(kafkaConsumerConfig KafkaConsumerConfig) map[string]ad.AdConsumer {
	kafkaAdPartitionConsumer := registerKafkaAdPartitionConsumer(kafkaConsumerConfig)
	kafkaAdConsumerGroup := registerKafkaAdConsumerGroup(kafkaConsumerConfig)
	maps.Copy(kafkaAdPartitionConsumer, kafkaAdConsumerGroup)
	return kafkaAdPartitionConsumer
}

type KafkaContext interface {
	KafkaConsumerConfig() KafkaConsumerConfig
	KafkaProducerConfig() KafkaProducerConfig
	KafkaConsumers() map[string]ad.AdConsumer
}
