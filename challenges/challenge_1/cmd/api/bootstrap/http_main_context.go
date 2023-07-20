package bootstrap

import (
	createad2 "challenges/challenge_1/internal/application/createad"
	"challenges/challenge_1/internal/application/fetchad"
	"challenges/challenge_1/internal/application/listads"
	"challenges/challenge_1/internal/domain/ad"
	"challenges/challenge_1/internal/infrastructure/ad/broker/producer"
	"challenges/challenge_1/internal/infrastructure/ad/repository"
	"challenges/challenge_1/internal/infrastructure/ad/server/handler/createad"
	fetchad2 "challenges/challenge_1/internal/infrastructure/ad/server/handler/fetchad"
	listads2 "challenges/challenge_1/internal/infrastructure/ad/server/handler/listads"
	"challenges/challenge_1/internal/infrastructure/kit/server/handler"
	"challenges/challenge_1/internal/infrastructure/system/server/handler/health"
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type mainContext struct {
	dbConnection *sql.DB
	handlers     map[string]handler.GinHandler
}

func NewMainContext(envHandler EnvHandler, dbConnection *sql.DB) MainContext {
	adRepository := createRepository(envHandler, dbConnection)
	adProducer, err := createAdProducer(envHandler)
	if err != nil {
		fmt.Println(">>> Impossible create the ad producer. Creating a logger implementation of it...")
		logAdProducer := producer.NewLogAdProducer()
		return &mainContext{
			dbConnection: dbConnection,
			handlers:     registerHandlers(adRepository, logAdProducer),
		}
	}
	return &mainContext{
		dbConnection: dbConnection,
		handlers:     registerHandlers(adRepository, adProducer),
	}
}

func createRepository(envHandler EnvHandler, dbConnection *sql.DB) ad.AdRepository {
	if envHandler.EnableDatabase() {
		// return createAdRepository(dbConnection)
		return createNewAdGormRepository(dbConnection)
	} else {
		return createAdInmemoryRepository()
	}
}

type MainContext interface {
	GetHandlers() map[string]handler.GinHandler
}

func (mainContext *mainContext) GetHandlers() map[string]handler.GinHandler {
	return mainContext.handlers
}

func registerHandlers(adRepository ad.AdRepository, adProducer ad.AdProducer) map[string]handler.GinHandler {
	handlers := map[string]handler.GinHandler{
		"healthCheck":        registerHealthCheckHandler(),
		"createAdHandler":    registerCreateAdHandler(adRepository, adProducer),
		"fetchAdHandler":     registerFetchAdHandler(adRepository),
		"fetchAllAdsHandler": registerFetchAllAdsHandler(adRepository),
	}
	return handlers
}

func createAdInmemoryRepository() ad.AdRepository {
	fmt.Println(">>> Init in memory repository...")
	return repository.NewAdInMemoryRepository()
}

func createAdRepository(dbConnection *sql.DB) ad.AdRepository {
	fmt.Println(">>> Init sql repository...")
	return repository.NewAdRepository(dbConnection)
}

func createNewAdGormRepository(dbConnection *sql.DB) ad.AdRepository {
	fmt.Println(">>> Init GORM repository...")
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbConnection,
	}), &gorm.Config{})
	if err != nil {
		log.Panic(">> Initialization Error, imposible to create the GORM repo. Shutdowning...", err)
	}
	return repository.NewAdGormRepository(gormDB)
}
func createAdProducer(envHandler EnvHandler) (ad.AdProducer, error) {
	brokers := envHandler.KafkaBrokers()
	topicName := envHandler.KafkaConsumerTopic()
	return producer.NewKafkaAdSyncProducer(brokers, topicName)
}

func registerCreateAdHandler(adRepository ad.AdRepository, adProducer ad.AdProducer) handler.GinHandler {
	postAdService := createad2.NewPostAdService(adRepository, adProducer)
	postAdHandler := createad.NewPostAdHandler(postAdService)
	return postAdHandler
}

func registerFetchAdHandler(adRepository ad.AdRepository) handler.GinHandler {
	fetchAdService := fetchad.NewFetchAdService(adRepository)
	fetchAdHandler := fetchad2.NewFetchAdHandler(fetchAdService)
	return fetchAdHandler
}

func registerFetchAllAdsHandler(adRepository ad.AdRepository) handler.GinHandler {
	fetchAllAds := listads.NewFetchAllAds(adRepository)
	fetchAdHandler := listads2.NewFetchAllAdsHandler(fetchAllAds)
	return fetchAdHandler
}

func registerHealthCheckHandler() handler.GinHandler {
	return health.NewHealthCheckHandler()
}
