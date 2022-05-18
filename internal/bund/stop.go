package bund

import (
	"time"
	"github.com/vulogov/nrbund/internal/signal"
	"github.com/pieterclaerhout/go-log"
)

func SendStop() {
	data, err := MakeStop("stop")
	if err != nil {
		log.Errorf("[ NRBUND ] STOP: %v", err)
	}
	NatsSend(data)
	signal.ExitRequest()
}

func Stop() {
	Init()
	InitEtcdAgent("stop")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	InitNewRelicAgent()
	log.Debugf("[ NRBUND ] bund.Stop(%v) is reached", ApplicationId)
	SendStop()
	for ! signal.ExitRequested() {
		time.Sleep(1 * time.Second)
	}
}
