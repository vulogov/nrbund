package bund

import (
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/nrbund/internal/signal"
)



func Agent() {
	Init()
	log.Debug("[ NRBUND ] bund.Agent() is reached")

	signal.ExitRequest()
}
