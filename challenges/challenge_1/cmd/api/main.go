package main

import (
	"challenges/challenge_1/cmd/api/bootstrap"
	"fmt"
	"log"
	"time"
)

func main() {
	time.Sleep(2 * time.Second)
	fmt.Println(">>> Running my-app...")
	time.Sleep(3 * time.Second)
	envHandler := bootstrap.NewEnvHandler()
	execute(envHandler)
}

func execute(envHandler bootstrap.EnvHandler) {
	dbConnection := bootstrap.GetDBConnectionOrNil(envHandler)
	if dbConnection != nil {
		defer dbConnection.Close()
	}

	kafkaContext := bootstrap.NewKafkaContext(envHandler)
	if kafkaContext != nil {
		startListeners(kafkaContext)
	}

	mainContext := bootstrap.NewMainContext(envHandler, dbConnection)
	executeHttpServer(envHandler, mainContext)
}

func executeHttpServer(envHandler bootstrap.EnvHandler, mainContext bootstrap.MainContext) {
	err := bootstrap.NewHttpServer(envHandler.HttpPort(), mainContext).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func startListeners(kafkaContext bootstrap.KafkaContext) {
	fmt.Println(">>> Starting Kafka Consumers...")
	for key, element := range kafkaContext.KafkaConsumers() {
		fmt.Println(">>> Running consumer:", key)
		go element.StartListening()
	}
}
