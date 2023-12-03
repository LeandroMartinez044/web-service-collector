package dependencies

import (
	"errors"
	"os"

	"github.com/LeandroMartinez044/lmenglish/collector/internal/core/ports"
	"github.com/LeandroMartinez044/lmenglish/collector/internal/core/services/collectorsrv"
	"github.com/LeandroMartinez044/lmenglish/collector/internal/handlers/collectorhdl"
	"github.com/LeandroMartinez044/lmenglish/collector/internal/repositories/collector"
	"github.com/LeandroMartinez044/lmenglish/collector/internal/repositories/mongodb"
	"github.com/LeandroMartinez044/lmenglish/collector/internal/repositories/youtube"
)

type Definition struct {

	//
	// Repositories
	//
	YoutubeRepository   ports.YoutubeRepository
	CollectorRepository ports.CollectorRepository

	//
	// Core
	//
	CollectorService ports.CollectorService

	//
	// Handlers
	//
	CollectorHandler *collectorhdl.Handler
}

func NewByEnvironment() Definition {

	var con *mongodb.Connection

	//
	// Obtains the environment
	//
	if os.Getenv("GO_ENVIRONMENT") == "production" {
		print("production")
	} else if os.Getenv("GO_ENV") == "test" {
		con = mongodb.New("mongodb://localhost:27017", "test")
	} else {
		panic(errors.New("can't init application in development mode"))
	}

	return initDependencies(con)
}

func initDependencies(con *mongodb.Connection) Definition {

	d := Definition{}

	//
	// Repositories
	//
	d.CollectorRepository = collector.New(con.Db.Collection("words"))
	d.YoutubeRepository = youtube.New()

	//
	// Core
	//
	d.CollectorService = collectorsrv.New(d.YoutubeRepository, d.CollectorRepository)

	//
	// Handlers
	//
	d.CollectorHandler = collectorhdl.New(d.CollectorService)

	return d
}
