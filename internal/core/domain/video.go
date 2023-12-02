package domain

type Video struct {
	ID        string
	StartTime string
	EndTime   string
}

func NewVideo(id string, startTime string, endTime string) *Video {
	return &Video{id, startTime, endTime}
}
