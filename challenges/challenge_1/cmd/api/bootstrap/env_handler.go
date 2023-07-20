package bootstrap

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	httpPort           = 8000
	enableDatabase     = false
	dbUser             = "postgres_user"
	dbPass             = "postgres_pass"
	dbName             = "learning_go_db"
	dbHost             = "localhost"
	kafkaConsumerGroup = "any-consumer-group"
	kafkaConsumerTopic = "topic.test.1"
	kafkaBroker        = "localhost:9092"
)

type envHandler struct {
	httpPort           int
	enableDatabase     bool
	dbUser             string
	dbPass             string
	dbName             string
	dbHost             string
	kafkaConsumerGroup string
	kafkaConsumerTopic string
	kafkaBrokers       []string
}

func NewEnvHandler() *envHandler {
	// TODO REFACTOR SEVERAL ENV HANDLERS TO POSTGRES, KAFKA, HTTP SERVER AND SO ON
	envHandler := envHandler{
		httpPort:           stringOrIntDefault(os.Getenv("port"), httpPort),
		enableDatabase:     stringOrBoolDefault(os.Getenv("POSTGRES_ENABLE_DATABASE"), enableDatabase),
		dbUser:             stringOrDefault(os.Getenv("POSTGRES_USER"), dbUser),
		dbPass:             stringOrDefault(os.Getenv("POSTGRES_PASSWORD"), dbPass),
		dbName:             stringOrDefault(os.Getenv("POSTGRES_DB"), dbName),
		dbHost:             stringOrDefault(os.Getenv("POSTGRES_DB_HOST"), dbHost),
		kafkaConsumerGroup: stringOrDefault(os.Getenv("KAFKA_CONSUMER_GROUP"), kafkaConsumerGroup),
		kafkaConsumerTopic: stringOrDefault(os.Getenv("KAFKA_CONSUMER_TOPIC"), kafkaConsumerTopic),
		kafkaBrokers:       stringOrArrayDefault(os.Getenv("KAFKA_BROKERS"), []string{kafkaBroker}),
	}
	envHandler.PrintEnv()
	return &envHandler
}

func stringOrDefault(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	} else {
		return value
	}
}

func stringOrBoolDefault(value string, defaultValue bool) bool {
	boolValue, error := strconv.ParseBool(value)
	if error != nil {
		return defaultValue
	} else {
		return boolValue
	}
}

func stringOrIntDefault(value string, defaultValue int) int {
	intValue, error := strconv.Atoi(value)
	if error != nil {
		return defaultValue
	} else {
		return intValue
	}
}

func stringOrArrayDefault(value string, defaultValue []string) []string {
	if value == "" {
		return defaultValue
	} else {
		return strings.Split(value, ",")
	}
}

func (e *envHandler) HttpPort() int {
	return e.httpPort
}

func (e *envHandler) EnableDatabase() bool {
	return e.enableDatabase
}

func (e *envHandler) DbUser() string {
	return e.dbUser
}

func (e *envHandler) DbPass() string {
	return e.dbPass
}
func (e *envHandler) DbName() string {
	return e.dbName
}

func (e *envHandler) DbHost() string {
	return e.dbHost
}

func (e *envHandler) KafkaConsumerGroup() string {
	return e.kafkaConsumerGroup
}

func (e *envHandler) KafkaConsumerTopic() string {
	return e.kafkaConsumerTopic
}

func (e *envHandler) KafkaBrokers() []string {
	return e.kafkaBrokers
}

func (e *envHandler) PrintEnv() {
	fmt.Println(">>> Myads: env port: ", e.httpPort)
	fmt.Println(">>> Myads: env POSTGRES_ENABLE_DATABASE: ", e.enableDatabase)
	if e.enableDatabase {
		fmt.Println(">>> Myads: env POSTGRES_USER: ", e.dbUser)
		fmt.Println(">>> Myads: env POSTGRES_PASSWORD: ", e.dbPass)
		fmt.Println(">>> Myads: env POSTGRES_DB: ", e.DbName())
		fmt.Println(">>> Myads: env POSTGRES_DB_HOST: ", e.dbHost)
	}
	fmt.Println(">>> Myads: env KAFKA_BROKERS: ", e.kafkaBrokers)
	fmt.Println(">>> Myads: env KAFKA_CONSUMER_GROUP: ", e.KafkaConsumerGroup())
	fmt.Println(">>> Myads: env KAFKA_CONSUMER_TOPIC: ", e.KafkaConsumerTopic())
}

type EnvHandler interface {
	HttpPort() int
	EnableDatabase() bool
	DbUser() string
	DbPass() string
	DbName() string
	DbHost() string
	KafkaConsumerGroup() string
	KafkaConsumerTopic() string
	KafkaBrokers() []string
}
