package listads

import (
	"challenges/challenge_1/internal/application/listads"
	"challenges/challenge_1/internal/domain/ad"
	"challenges/challenge_1/internal/infrastructure/kit/server/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type fetchAllAdsHandler struct {
	fetchAllAdsService listads.FetchAllAdsService
}

func NewFetchAllAdsHandler(fetchAllAdsService listads.FetchAllAdsService) handler.GinHandler {
	return &fetchAllAdsHandler{fetchAllAdsService: fetchAllAdsService}
}

func (handler *fetchAllAdsHandler) NewGinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		foundAds, err := handler.fetchAllAdsService.Execute()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, mapAdToResponse(*foundAds))
	}
}

func mapAdToResponse(ads []ad.Ad) fetchAllAdsResponse {

	var adsResponse []fetchAdResponse
	for _, ad := range ads {
		adResponse := fetchAdResponse{
			Id:          ad.Id.String(),
			Title:       ad.Title,
			Description: ad.Description.Value,
			Price:       ad.Price,
			TimeStamp:   ad.Timestamp.Format("2006-01-02"),
		}
		adsResponse = append(adsResponse, adResponse)
	}
	return fetchAllAdsResponse{Ads: adsResponse}
}

type fetchAdResponse struct {
	Id          string
	Title       string
	Description string
	Price       int
	TimeStamp   string
}

type fetchAllAdsResponse struct {
	Ads []fetchAdResponse
}
