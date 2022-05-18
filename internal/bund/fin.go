package bund

import (
	"time"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/nrbund/internal/signal"
	"github.com/vulogov/nrbund/internal/banner"
)

func Fin() {
	banner.Banner("[ Zay Gezunt ]")
	log.Debugf("[ NRBUND ] bund.Fin(%v) is reached", ApplicationId)
	CloseNatsAgent()
	CloseEtcdAgent()
	log.Debug("Wait while NR application is shut down")
	NRapp.Shutdown(60 * time.Second)
	log.Debug("NR Application is down")
	log.Infof("[ NRBUND ] %s is down", ApplicationId)
	signal.ExitRequest()
}
