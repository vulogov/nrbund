package bund

import (
	"fmt"
	"github.com/pieterclaerhout/go-log"
)



func Agent() {
	Init()
	log.Debug("[ NRBUND ] bund.Agent() is reached")
	InitEtcdAgent()
	fmt.Println(EtcdGetItems())
}
