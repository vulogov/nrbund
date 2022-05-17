package bund

import (
	"fmt"
	"github.com/pieterclaerhout/go-log"
)



func Watch() {
	Init()
	log.Debug("[ NRBUND ] bund.Watch() is reached")
	InitEtcdAgent()
	fmt.Println(EtcdGetItems())
}
