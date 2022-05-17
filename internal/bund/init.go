package bund

import (
	"github.com/cosiner/argv"
	"github.com/pieterclaerhout/go-log"
	"github.com/bamzi/jobrunner"
	"github.com/vulogov/nrbund/internal/conf"
	"github.com/vulogov/nrbund/internal/stdlib"
	tlog "github.com/vulogov/nrbund/internal/log"
	"github.com/vulogov/nrbund/internal/signal"
)

func Init() {
	tlog.Init()
	stdlib.InitStdlib()
	log.Debug("[ NRBUND ] bund.Init() is reached")
	signal.InitSignal()
	if len(*conf.Args) > 0 {
		Argv, err := argv.Argv(*conf.Args, func(backquoted string) (string, error) {
			return backquoted, nil
		}, nil)
		if err != nil {
			log.Fatalf("Error parsing ARGS: %v", err)
		}
		log.Debugf("ARGV: %v", Argv)
		conf.Argv = Argv
	}
	log.Debugf("[ NRBUND ] Id: %v", *conf.Id)
	log.Debugf("[ NRBUND ] Name: %v", *conf.Name)
	jobrunner.Start()
	log.Debugf("[ NRBUND ] Job runner started")
	stdlib.StoreArgs()
	CheckNewRelic()
	InitNewRelicAgent()
}
