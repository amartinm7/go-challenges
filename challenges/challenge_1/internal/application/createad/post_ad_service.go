package createad

import (
	"challenges/challenge_1/internal/domain/ad"
	"fmt"
)

type postAd struct {
	repository ad.AdRepository
	producer   ad.AdProducer
}

func NewPostAdService(repository ad.AdRepository, producer ad.AdProducer) PostAdService {
	return &postAd{repository: repository, producer: producer}
}

func (service *postAd) Execute(_ad ad.Ad) (ad.Ad, error) {
	savedAd, err := service.repository.Save(_ad)
	if err != nil {
		return _ad, err
	}
	err = service.emitDomainEvent(savedAd)
	if err != nil {
		return _ad, err
	}
	return _ad, nil
}

func (service *postAd) emitDomainEvent(_ad ad.Ad) error {
	// evolve this part to use the eventBus to publish a message
	partition, offset, err := service.producer.SendMessage(ad.NewBrokerMessage(_ad))
	fmt.Printf(">>> Sent message to kafka, partition %d offset: %d\n", partition, offset)
	return err
}

type PostAdService interface {
	Execute(ad ad.Ad) (ad.Ad, error)
}
