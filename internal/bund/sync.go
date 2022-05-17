package bund

import (
	"fmt"
	"github.com/pieterclaerhout/go-log"
)



func Sync() {
	Init()
	log.Debug("[ NRBUND ] bund.Sync() is reached")
	InitEtcdAgent()
	fmt.Println(EtcdGetItems())
}
