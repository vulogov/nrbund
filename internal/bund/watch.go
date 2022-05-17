package bund

import (
	"fmt"
	"github.com/pieterclaerhout/go-log"
)



func Watch() {
	Init()
	InitEtcdAgent("watch")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	log.Debugf("[ NRBUND ] bund.Watch(%v) is reached", ApplicationId)

}
