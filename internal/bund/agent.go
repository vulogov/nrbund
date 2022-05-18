package bund

import (
	"github.com/nats-io/nats.go"
	"github.com/pieterclaerhout/go-log"
)

func NRBundAgent(m *nats.Msg) {
	msg := UnMarshal(m.Data)
	if msg == nil {
		log.Error("Invalid packet received")
	}
	IfSTOP(msg)
	if msg.PktKey == "Agitator" && len(msg.Value) > 0 {
		BundEvalExpression(string(msg.Value))
	}
}

func Agent() {
	Init()
	InitEtcdAgent("agent")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	log.Debugf("[ NRBUND ] bund.Agent(%v) is reached", ApplicationId)
	NatsRecv(NRBundAgent)
	Loop()
}
