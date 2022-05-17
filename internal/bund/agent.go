package bund

import (
	"fmt"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/nrbund/internal/signal"
)



func Agent() {
	Init()
	log.Debug("[ NRBUND ] bund.Agent() is reached")
	InitEtcdAgent()
	fmt.Println(EtcdGetItems())
	signal.ExitRequest()
}
