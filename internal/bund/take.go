package bund

import (
	"fmt"
	"github.com/pieterclaerhout/go-log"
)



func Take() {
	Init()
	log.Debug("[ NRBUND ] bund.Take() is reached")
	InitEtcdAgent()
	fmt.Println(EtcdGetItems())
}
