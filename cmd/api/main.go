package main

import (
	"github.com/LeandroMartinez044/lmenglish/collector/cmd/dependencies"
	"github.com/gin-gonic/gin"
)

func main() {

	d := dependencies.NewByEnvironment()
	router := gin.Default()
	router.POST("/generate", d.CollectorHandler.GenerateSubtitles)
	router.Run("localhost:8080")
}
