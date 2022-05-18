package bund

import (
	"bufio"
	"os"
	"github.com/vulogov/nrbund/internal/conf"
	tc "github.com/vulogov/ThreadComputation"
	"github.com/pieterclaerhout/go-log"
)



func Submit() {
	var err error

	Init()
	InitEtcdAgent("submit")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	log.Debugf("[ NRBUND ] bund.Submit(%v) is reached", ApplicationId)
	script := ""
	if *conf.SScript == "--" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
        script += scanner.Text()
				script += "\n"
    }
	} else {
		script, err = tc.ReadFile(*conf.SScript)
		if err != nil {
			log.Errorf("[ NRBUND ] %v", err)
			return
		}
	}
	if len(script) == 0 {
		log.Errorf("[ NRBUND ] script can not be a zero length")
		return
	}
	pkt, err := MakeScript("submit", []byte(script), nil)
	if err != nil {
		log.Errorf("[ NRBUND ] %v", err)
		return
	}
	log.Debugf("[ NRBUND ] Sending script for execution len()=%v", len(pkt))
	NatsSend(pkt)
}
