package bund

import (
	"fmt"

	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/nrbund/internal/banner"
	"github.com/vulogov/nrbund/internal/conf"
)

func Version() {
	Init()
	log.Debug("[ NRBUND ] bund.Version() is reached")
	banner.Banner(fmt.Sprintf("[ BUND %v ]", conf.EVersion))
	banner.Table(true)
}
