package controller

import (
	"learn-ci/model"
	"learn-ci/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemoController struct {
	Service service.MemoService
}

func NewMemoController(service *service.MemoService) MemoController {
	return MemoController{Service: *service}
}

func (controller *MemoController) Route(r *gin.Engine) {
	r.POST("/api/v1/memo")
	r.GET("/api/v1/memo/:author")
}

func (controller *MemoController) Create(ctx *gin.Context) {
	var request model.AddMemoRequest
	err := ctx.BindJSON(&request)

	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
	}

	_, err = controller.Service.AddMemo(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:    int32(http.StatusInternalServerError),
			Message: "Failed to add memo!",
			Data:    nil,
		})
	}

	ctx.JSON(http.StatusCreated, model.WebResponse{
		Code:    int32(http.StatusCreated),
		Message: "Successfully adding Memo!",
		Data:    request,
	})

}
