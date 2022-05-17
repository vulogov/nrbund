package bund

import (
	"time"
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/nrbund/internal/banner"
)

func Fin() {
	banner.Banner("[ Zay Gezunt ]")
	log.Debug("[ NRBUND ] bund.Fin() is reached")
	CloseEtcdAgent()
	log.Debug("Wait while NR application is shut down")
	NRapp.Shutdown(60 * time.Second)
	log.Debug("NR Application is down")
}
