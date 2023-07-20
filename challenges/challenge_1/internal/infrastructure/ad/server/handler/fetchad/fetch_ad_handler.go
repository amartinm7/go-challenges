package fetchad

import (
	"challenges/challenge_1/internal/application/fetchad"
	"challenges/challenge_1/internal/domain/ad"
	"challenges/challenge_1/internal/infrastructure/kit/server/handler"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type fetchAdHandler struct {
	fetchAdService fetchad.FetchAdService
}

func NewFetchAdHandler(fetchAdService fetchad.FetchAdService) handler.GinHandler {
	return &fetchAdHandler{
		fetchAdService: fetchAdService,
	}
}

type fetchAdRequest struct {
	Id uuid.UUID
}

type fetchAdResponse struct {
	Id          string
	Title       string
	Description string
	Price       int
	TimeStamp   string
}

func (handler *fetchAdHandler) NewGinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		foundAd, err := handler.fetchAdService.Execute(uuid) // TODO REFACTOR TO fetchAdRequest
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, mapAdToResponse(*foundAd))
	}
}

func mapAdToResponse(ad ad.Ad) *fetchAdResponse {
	return &fetchAdResponse{
		Id:          ad.Id.String(),
		Title:       ad.Title,
		Description: ad.Description.Value,
		Price:       ad.Price,
		TimeStamp:   ad.Timestamp.Format("2006-01-02"),
	}
}
