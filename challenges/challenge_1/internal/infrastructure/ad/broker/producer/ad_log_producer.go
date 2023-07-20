package producer

import (
	"challenges/challenge_1/internal/domain/ad"
	"encoding/json"
	"fmt"
)

type logAdProducer struct {
}

func NewLogAdProducer() ad.AdProducer {
	return &logAdProducer{}
}

func (logAdProducer *logAdProducer) SendMessage(message ad.BrokerMessage) (partition int32, offset int64, err error) {
	adMessage, err := json.Marshal(message.Ad)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(">>> Logging message...")
	fmt.Println(string(adMessage))
	return 0, 0, nil
}
