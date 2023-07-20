package consumer

import (
	"challenges/challenge_1/internal/domain/ad"
	"context"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type kafkaAdConsumerGroup struct {
	client  sarama.ConsumerGroup
	topic   string
	groupId string
}

type KafkaAdConsumer struct {
	ready chan bool
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (kafkaAdConsumer *KafkaAdConsumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(kafkaAdConsumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (kafkaAdConsumer *KafkaAdConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (kafkaAdConsumer *KafkaAdConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message := <-claim.Messages():
			log.Printf(">>> kafkaAdConsumerGroup message claimed: value = %s, timestamp = %v, topic = %s, , partition = %v", string(message.Value), message.Timestamp, message.Topic, message.Partition)
			session.MarkMessage(message, "")

		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/Shopify/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}

func NewKafkaAdConsumerGroup(brokers []string, topic, groupId string) (ad.AdConsumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	client, err := sarama.NewConsumerGroup(brokers, groupId, config)
	if err != nil {
		log.Println("Error creating consumer group client: ", err)
		return nil, err
	}
	return &kafkaAdConsumerGroup{
		client:  client,
		topic:   topic,
		groupId: groupId,
	}, nil
}

func (kafkaAdConsumerGroup *kafkaAdConsumerGroup) close() {
	if err := kafkaAdConsumerGroup.client.Close(); err != nil {
		log.Fatalln(err)
	}
}

func (kafkaAdConsumerGroup *kafkaAdConsumerGroup) StartListening() {
	log.Printf(">>> kafkaAdConsumerGroup Kafka StartListening...")
	kafkaAdConsumer := KafkaAdConsumer{
		ready: make(chan bool),
	}
	ctx, cancel := context.WithCancel(context.Background())

	consumptionIsPaused := false
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := kafkaAdConsumerGroup.client.Consume(ctx, []string{kafkaAdConsumerGroup.topic}, &kafkaAdConsumer); err != nil {
				log.Println("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			kafkaAdConsumer.ready = make(chan bool)
		}
	}()

	<-kafkaAdConsumer.ready // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	keepRunning := true

	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
		case <-sigusr1:
			toggleConsumptionFlow(kafkaAdConsumerGroup.client, &consumptionIsPaused)
		}
	}
	cancel()
	wg.Wait()
	if err := kafkaAdConsumerGroup.client.Close(); err != nil {
		log.Println("Error closing client: %v", err)
	}
}

func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		log.Println("Resuming consumption")
	} else {
		client.PauseAll()
		log.Println("Pausing consumption")
	}

	*isPaused = !*isPaused
}
