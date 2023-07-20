package main

import (
	"challenges/challenge_1/internal/application/createad"
	"challenges/challenge_1/internal/application/fetchad"
	"challenges/challenge_1/internal/application/listads"
	"challenges/challenge_1/internal/domain/ad"
	"challenges/challenge_1/internal/infrastructure/ad/broker/producer"
	"challenges/challenge_1/internal/infrastructure/ad/repository"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Running my first app!...")
	var motorAd, error = ad.NewAd(uuid.New().String(), "Opel Omega", "como nuevo.", 15000, "2023-03-01")
	if error != nil {
		fmt.Println(">>> Error to create ad...")
		return
	}
	adRepository := repository.NewAdInMemoryRepository()
	adProducer := producer.NewLogAdProducer()
	opelAstraAd, _ := createad.NewPostAdService(adRepository, adProducer).Execute(motorAd)
	fmt.Println("Posting new ad...")
	fmt.Println(opelAstraAd)
	fmt.Println("")
	foundOpelAstraAd, _ := fetchad.NewFetchAdService(adRepository).Execute(opelAstraAd.Id)
	fmt.Println("Find Ad by id...", opelAstraAd.Id)
	fmt.Println(*foundOpelAstraAd)
	fmt.Println("")
	foundAds, _ := listads.NewFetchAllAds(adRepository).Execute()
	fmt.Println("Find All Ads...")
	fmt.Println(*foundAds)
	fmt.Println("Ending process...")
}
