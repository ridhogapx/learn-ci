package main

import (
	"learn-ci/config"
	"learn-ci/controller"
	"learn-ci/repository"
	"learn-ci/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	conn := config.NewDBConnection()
	memoRepository := repository.NewMemoRepository(conn)
	memoService := service.NewMemoService(memoRepository)
	memoController := controller.NewMemoController(memoService)

	memoController.Route(r)

	r.Run(":80")

}
