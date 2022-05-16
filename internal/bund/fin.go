package bund

import (
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/nrbund/internal/banner"
)

func Fin() {
	banner.Banner("[ Zay Gezunt ]")
	log.Debug("[ NRBUND ] bund.Fin() is reached")
}
