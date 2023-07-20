package bootstrap

type kafkaConsumerConfig struct {
	brokers       []string
	topicName     string
	consumerGroup string
}

type KafkaConsumerConfig interface {
	Brokers() []string
	TopicName() string
	ConsumerGroup() string
}

func NewKafkaConsumerConfig(envHandler EnvHandler) KafkaConsumerConfig {
	return &kafkaConsumerConfig{
		brokers:       envHandler.KafkaBrokers(),
		topicName:     envHandler.KafkaConsumerTopic(),
		consumerGroup: envHandler.KafkaConsumerGroup(),
	}
}

func (k *kafkaConsumerConfig) Brokers() []string {
	return k.brokers
}

func (k *kafkaConsumerConfig) TopicName() string {
	return k.topicName
}

func (k *kafkaConsumerConfig) ConsumerGroup() string {
	return k.consumerGroup
}

type kafkaProducerConfig struct {
	brokers   []string
	topicName string
}

type KafkaProducerConfig interface {
	Brokers() []string
	TopicName() string
}

func NewKafkaProducerConfig(envHandler EnvHandler) KafkaProducerConfig {
	return &kafkaProducerConfig{
		brokers:   envHandler.KafkaBrokers(),
		topicName: envHandler.KafkaConsumerTopic(),
	}
}

func (k *kafkaProducerConfig) Brokers() []string {
	return k.brokers
}

func (k *kafkaProducerConfig) TopicName() string {
	return k.topicName
}
