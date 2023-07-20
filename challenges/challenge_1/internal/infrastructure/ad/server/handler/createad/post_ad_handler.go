package createad

import (
	"challenges/challenge_1/internal/application/createad"
	"challenges/challenge_1/internal/domain/ad"
	"challenges/challenge_1/internal/infrastructure/kit/server/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type postAdHandler struct {
	postAdService createad.PostAdService
}

func NewPostAdHandler(postAdService createad.PostAdService) handler.GinHandler {
	return &postAdHandler{
		postAdService: postAdService,
	}
}

type PostAdRequest struct {
	Id          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	TimeStamp   string `json:"timeStamp" binding:"required"`
}

type PostAdResponse struct {
	Id string `json:"id" binding:"required"`
}

func NewPostAdResponse(id string) PostAdResponse {
	return PostAdResponse{
		Id: id,
	}
}

func (adRequest *PostAdRequest) toAd() (ad.Ad, error) {
	return ad.NewAd(adRequest.Id, adRequest.Title, adRequest.Description, adRequest.Price, adRequest.TimeStamp)
}
func (handler *postAdHandler) NewGinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var adRequest PostAdRequest
		if err := ctx.BindJSON(&adRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ad, err := adRequest.toAd()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		_, err = handler.postAdService.Execute(ad)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, NewPostAdResponse(ad.Id.String()))
	}
}
