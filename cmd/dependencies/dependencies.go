package dependencies

import (
	"github.com/LeandroMartinez044/lmenglish/collector/internal/core/ports"
	"github.com/LeandroMartinez044/lmenglish/collector/internal/core/services/collectorsrv"
	"github.com/LeandroMartinez044/lmenglish/collector/internal/handlers/collectorhdl"
	"github.com/LeandroMartinez044/lmenglish/collector/internal/repositories/youtube"
)

type Definition struct {

	//
	// Repositories
	//
	YoutubeRepository ports.YoutubeRepository

	//
	// Core
	//
	CollectorService ports.CollectorService

	//
	// Handlers
	//
	CollectorHandler *collectorhdl.Handler
}

func InitDependencies() Definition {

	d := Definition{}

	//
	// Repositories
	//
	d.YoutubeRepository = youtube.New()

	//
	// Core
	//
	d.CollectorService = collectorsrv.New(d.YoutubeRepository)

	//
	// Handlers
	//
	d.CollectorHandler = collectorhdl.New(d.CollectorService)

	return d
}
