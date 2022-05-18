package bund

import (
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
	log.Debugf("[ NRBUND ] sending %s", s.Name)
	script, err := tc.ReadFile(s.Uri)
	if err != nil {
		log.Errorf("[ NRBUND ] %v", err)
		return
	}
	if len(script) == 0 {
		log.Errorf("[ NRBUND ] script can not be a zero length")
		return
	}
	pkt, err := MakeScript("agitator", []byte(script), nil)
	NatsSend(pkt)
}


func AgitatorScheduleConfig() {
	for _, n := range(*conf.AConf) {
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
	}
}

func Agitator() {
	Init()
	log.Debug("[ NRBUND ] bund.Agitator() is reached")
	InitEtcdAgent("agitator")
	UpdateConfigToEtcd()
	InitNatsAgent()
	AgitatorScheduleConfig()
	Loop()
}
