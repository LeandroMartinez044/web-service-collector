package collectorhdl

import (
	"fmt"
	"net/http"

	"github.com/LeandroMartinez044/web-service-collector/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	collectorSrv ports.CollectorService
	finderSrv    ports.FinderService
}

func New(collectorSrv ports.CollectorService,
	finderSrv ports.FinderService) *Handler {
	return &Handler{collectorSrv: collectorSrv, finderSrv: finderSrv}
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

type VideoDto struct {
	Url       string `json:"url"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// Get Video
func (h *Handler) GetVideosByWord(c *gin.Context) {

	// Get the ID parameter from the URL

	word := c.Param("word")

	videos, err := h.finderSrv.FindVideosBy(word)

	var videoDtoList []VideoDto

	for _, video := range videos {
		videoDto := VideoDto{video.ID, video.StartTime, video.EndTime}
		videoDtoList = append(videoDtoList, videoDto)
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d", err)})
	}

	c.JSON(http.StatusOK, videoDtoList)
}

func (h *Handler) Check(c *gin.Context) {

	c.JSON(http.StatusOK, "Hello world")
}
