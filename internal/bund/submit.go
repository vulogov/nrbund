package bund

import (
	"fmt"
	"github.com/pieterclaerhout/go-log"
)



func Submit() {
	Init()
	log.Debug("[ NRBUND ] bund.Submit() is reached")
	InitEtcdAgent()
	fmt.Println(EtcdGetItems())
}
