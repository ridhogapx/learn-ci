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

func NewMemoController(service service.MemoService) MemoController {
	return MemoController{Service: service}
}

func (controller MemoController) Route(r *gin.Engine) {
	r.POST("/api/v1/memo", controller.AddMemo)
	r.GET("/api/v1/memo/:author", controller.GetMemo)
	r.GET("/", controller.Home)
}

func (MemoController) Home(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "welcome to my memo",
	})
}

func (controller MemoController) AddMemo(ctx *gin.Context) {
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
		return
	}

	ctx.JSON(http.StatusCreated, model.WebResponse{
		Code:    int32(http.StatusCreated),
		Message: "Successfully adding Memo!",
		Data:    request,
	})

}

func (controller MemoController) GetMemo(ctx *gin.Context) {
	param := ctx.Param("author")

	res, err := controller.Service.GetMemo(model.GetMemoRequest{
		Author: param,
	})

	if err != nil {
		ctx.JSON(http.StatusNotFound, model.WebResponse{
			Code:    int32(http.StatusNotFound),
			Message: "Not found",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusFound, model.WebResponse{
		Code:    int32(http.StatusFound),
		Message: "Successfully retrieve data!",
		Data:    res,
	})
}
