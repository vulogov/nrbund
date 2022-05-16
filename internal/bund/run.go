package bund

import (
	"github.com/pieterclaerhout/go-log"
	tc "github.com/vulogov/ThreadComputation"
	"github.com/vulogov/nrbund/internal/conf"
	"github.com/vulogov/nrbund/internal/signal"
	"github.com/vulogov/nrbund/internal/stdlib"
)

func RunFile(core *stdlib.BUNDEnv, name string) {
	log.Debugf("Running: %v", name)
	code, err := tc.ReadFile(name)
	if err != nil {
		log.Fatalf("Error loading file: %v", err)
	}
	core.Eval(code)
}

func Run() {
	Init()
	log.Debug("[ NRBUND ] bund.Run() is reached")
	if *conf.CDebug {
		log.Info("BUND core debug is on")
		tc.SetVariable("tc.Debuglevel", "debug")
		log.Infof("[ NRBUND ] core version: %v", tc.VERSION)
	} else {
		log.Debug("BUND core debug is off")
		tc.SetVariable("tc.Debuglevel", "info")
		log.Debugf("[ NRBUND ] core version: %v", tc.VERSION)
	}
	core := stdlib.InitBUND()
	for _, f := range *conf.Scripts {
		RunFile(core, f)
	}
	signal.ExitRequest()
}
