package bund

import (
	"time"
	"github.com/vulogov/nrbund/internal/signal"
	"github.com/bamzi/jobrunner"
	"github.com/pieterclaerhout/go-log"
)

type NATSSync struct {
}

func (s NATSSync) Run() {
	log.Debug("[ NRBUND ] sending Sync")
	SendSync()
}

func SendSync() {
	data, err := MakeSync("sync")
	if err != nil {
		log.Errorf("[ NRBUND ] SYNC: %v", err)
	}
	NatsSend(data)
}

func Sync() {
	Init()
	InitEtcdAgent("sync")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	log.Debugf("[ NRBUND ] bund.Sync(%v) is reached", ApplicationId)
	jobrunner.Schedule("@every 5s", NATSSync{})
	for ! signal.ExitRequested() {
		time.Sleep(1 * time.Second)
	}
}
