package bund

import (
	"github.com/vulogov/nrbund/internal/signal"
	"github.com/nats-io/nats.go"
	"github.com/pieterclaerhout/go-log"
)

func NRBundExecuteScript(m *nats.Msg) {
	msg := UnMarshal(m.Data)
	if msg == nil {
		log.Error("Invalid packet received")
	}
	IfSTOP(msg)
	if msg.PktKey == "Agitator" && len(msg.Value) > 0 {
		BundGlobalEvalExpression(string(msg.Value))
	}
	signal.ExitRequest()
	DoContinue = false
}

func Take() {
	Init()
	InitEtcdAgent("take")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	InitNewRelicAgent()
	log.Debugf("[ NRBUND ] bund.Take(%v) is reached", ApplicationId)
	NatsRecv(NRBundExecuteScript)
	Loop()
}
