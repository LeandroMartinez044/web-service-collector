package main

import (
	"github.com/LeandroMartinez044/lmenglish/collector/cmd/dependencies"
	"github.com/gin-gonic/gin"
)

func main() {

	d := dependencies.NewByEnvironment()
	router := gin.Default()
	router.POST("/generate", d.CollectorHandler.GenerateSubtitles)
	router.GET("/videos/:word", d.CollectorHandler.GetVideosByWord)
	router.GET("/check", d.CollectorHandler.Check)
	router.Run("52.207.222.195:8080")
}
