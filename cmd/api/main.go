package main

import (
	"net/http"

	"github.com/LeandroMartinez044/web-service-collector/cmd/dependencies"
	"github.com/gin-gonic/gin"
)

func main() {

	d := dependencies.NewByEnvironment()
	router := gin.Default()
	router.POST("/generate", d.CollectorHandler.GenerateSubtitles)
	router.GET("/videos/:word", d.CollectorHandler.GetVideosByWord)
	router.GET("/check", d.CollectorHandler.Check)

	http.ListenAndServe(":8080", router)

}
