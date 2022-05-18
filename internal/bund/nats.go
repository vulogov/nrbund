package bund

import (
	"os"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/vulogov/nrbund/internal/conf"
	"github.com/vulogov/nrbund/internal/signal"
	"github.com/pieterclaerhout/go-log"
)

var Nats *nats.Conn
var QueueName string


func InitNatsAgent() {
	var err error

	log.Debugf("Connecting to NATS")
	Nats, err = nats.Connect(
		*conf.Gnats,
		nats.Name(ApplicationId),
		nats.ReconnectWait(*conf.Timeout),
		nats.PingInterval(*conf.Timeout),
		nats.Timeout(*conf.Timeout),
	)
	if err != nil {
		log.Errorf("[ NATS ] %v", err)
		signal.ExitRequest()
		os.Exit(10)
	}
	QueueName = fmt.Sprintf("%s:%s", *conf.Id, *conf.Name)
	log.Debugf("Queue name: %v", QueueName)
}

func NatsSend(data []byte) {
	if DoContinue {
		Nats.Publish(QueueName, data)
	}
}

func NatsRecv(fun nats.MsgHandler) {
	Nats.QueueSubscribe(QueueName, *conf.Id, fun)
}

func CloseNatsAgent() {
	log.Debugf("Terminating and draining NATS session")
	Nats.Flush()
}
