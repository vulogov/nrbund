package bund

import (
	"fmt"
	"github.com/pieterclaerhout/go-log"
)



func Agent() {
	Init()
	InitEtcdAgent("agent")
	fmt.Println(EtcdGetItems())
	log.Debugf("[ NRBUND ] bund.Agent(%v) is reached", ApplicationId)
}
