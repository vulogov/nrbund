package bund

import (
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/nrbund/internal/signal"
)



func Agitator() {
	Init()
	log.Debug("[ NRBUND ] bund.Agitator() is reached")
	InitEtcdAgent()
	UpdateConfigToEtcd()
	signal.ExitRequest()
}
