package bund

import (
	"fmt"
	tc "github.com/vulogov/ThreadComputation"
	"github.com/bamzi/jobrunner"
	"github.com/vulogov/nrbund/internal/conf"
	"github.com/pieterclaerhout/go-log"
)

type TheScript struct {
	Name    string
	Uri     string
}

func (s TheScript) Run() {
	txn := NRapp.StartTransaction(fmt.Sprintf("[%v]%v", s.Name, s.Uri))
	log.Debugf("[ NRBUND ] sending %s", s.Name)
	segment := txn.StartSegment(fmt.Sprintf("[%v] Read %v", s.Name, s.Uri))
	script, err := tc.ReadFile(s.Uri)
	segment.End()
	if err != nil {
		log.Errorf("[ NRBUND ] %v", err)
		txn.NoticeError(err)
		txn.End()
		return
	}
	if len(script) == 0 {
		log.Errorf("[ NRBUND ] script can not be a zero length")
		txn.NoticeError(err)
		txn.End()
		return
	}
	pkt, err := MakeScript("agitator", []byte(script), nil)
	NatsSend(pkt)
	txn.End()
}


func AgitatorScheduleConfig() {
	txn := NRapp.StartTransaction(fmt.Sprintf("%s loading schedule configuration", ApplicationId))
	defer txn.End()
	for _, n := range(*conf.AConf) {
		segment := txn.StartSegment(fmt.Sprintf("%s loading %s", ApplicationId, n))
		cfg := HJsonLoadConfig(n)
		if cfg != nil {
			if jobs, ok := (*cfg)["jobs"]; ok {
				for _, j := range(jobs.([]interface{})) {
					job := j.(map[string]interface{})
					if name, ok := job["name"]; ok {
						if schedule, ok := job["schedule"]; ok {
							if uri, ok := job["uri"]; ok {
								log.Debugf("Scheduling (%v)[%v]=%v", name.(string), schedule.(string), uri.(string))
								jobrunner.Schedule(schedule.(string), TheScript{Name: name.(string), Uri: uri.(string)})
							}
						}
					}
				}
			}
		}
		segment.End()
	}
}

func Agitator() {
	Init()
	log.Debug("[ NRBUND ] bund.Agitator() is reached")
	InitEtcdAgent("agitator")
	if *conf.UploadConf {
		log.Info("Updating ETCD from local Agitator configuration")
		UpdateConfigToEtcd()
	} else {
		log.Info("Updating local Agitator configuration from ETCD")
		UpdateLocalConfigFromEtcd()
	}
	InitNatsAgent()
	InitNewRelicAgent()
	AgitatorScheduleConfig()
	Loop()
}
