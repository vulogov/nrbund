package bund

import (
	"github.com/pieterclaerhout/go-log"
)



func Agitator() {
	Init()
	log.Debug("[ NRBUND ] bund.Agitator() is reached")
	InitEtcdAgent("agitator")
	UpdateConfigToEtcd()
}
