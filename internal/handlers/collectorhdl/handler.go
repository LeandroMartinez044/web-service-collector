package collectorhdl

import (
	"fmt"
	"net/http"

	"github.com/LeandroMartinez044/lmenglish/collector/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	collectorSrv ports.CollectorService
}

func New(collectorSrv ports.CollectorService) *Handler {
	return &Handler{collectorSrv: collectorSrv}
}

type RequestData struct {
	Url string `json:"url"`
}

// getAlbums responds with the list of all albums as JSON.
func (h *Handler) GenerateSubtitles(c *gin.Context) {

	// Get data from the request body
	var requestData RequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.collectorSrv.StoreSubtitlesByVideoId(requestData.Url)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d", err)})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data received successfully"})
}
