package bund

import (
	"fmt"
	"github.com/pieterclaerhout/go-log"
)



func Stop() {
	Init()
	log.Debug("[ NRBUND ] bund.Stop() is reached")
	InitEtcdAgent("stop")
	fmt.Println(EtcdGetItems())
}
