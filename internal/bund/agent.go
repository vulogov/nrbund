package bund

import (
	"fmt"
	"github.com/nats-io/nats.go"
	tc "github.com/vulogov/ThreadComputation"
	"github.com/vulogov/nrbund/internal/conf"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/nrbund/internal/stdlib"
	"github.com/mgutz/ansi"
)

func GlobalDisplayResult(core *stdlib.BUNDEnv) {
	var out string
	if core.TC.Ready() {
		e := core.TC.Get()
		core.TC.Set(e)
		fun := tc.GetConverterCallback(e)
		if fun == nil {
			out = fmt.Sprintf("%v", e)
		} else {
			out_add := fun(e, tc.String)
			if out_add == nil {
				out += fmt.Sprintf("%v", e)
			} else {
				out += out_add.(string)
			}
		}
		if *conf.ShowResult {
			if *conf.Color {
				out = ansi.Color(out, "yellow")
				fmt.Println(out)
			} else {
				fmt.Println(out)
			}
		} else {
			log.Debugf("Result: %v", out)
		}
	} else {
		log.Debug("Stack is too shallow for result display")
	}
}

func BundGlobalEvalExpression(code string) {
	if *conf.CDebug {
		log.Info("BUND core debug is on")
		tc.SetVariable("tc.Debuglevel", "debug")
		log.Infof("[ NRBUND ] core version: %v", tc.VERSION)
	} else {
		log.Debug("BUND core debug is off")
		tc.SetVariable("tc.Debuglevel", "info")
		log.Debugf("[ NRBUND ] core version: %v", tc.VERSION)
	}
	log.Debugf("BUND core display result %v", *conf.ShowResult)
	core := stdlib.InitBUND()
	core.Eval(code)
	GlobalDisplayResult(core)
}


func NRBundAgent(m *nats.Msg) {
	if ! HadSync {
		log.Warn("Request received but agent not in SYNC state. Request ignored.")
		return
	}
	msg := UnMarshal(m.Data)
	if msg == nil {
		log.Error("Invalid packet received")
	}
	if msg.PktKey == "Agitator" && len(msg.Value) > 0 {
		BundGlobalEvalExpression(string(msg.Value))
	}
}

func Agent() {
	Init()
	InitEtcdAgent("agent")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	InitNewRelicAgent()
	log.Debugf("[ NRBUND ] bund.Agent(%v) is reached", ApplicationId)
	NatsRecv(NRBundAgent)
	Loop()
}
