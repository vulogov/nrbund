package bund

import (
	"fmt"
	"os"
	"time"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/nrbund/internal/conf"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var NRapp newrelic.Application

func InitNewRelicAgent() {
	NRapp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(*conf.Name),
		newrelic.ConfigLicense(*conf.NRLicenseKey),
		// newrelic.ConfigDebugLogger(os.Stdout),
	)
	if err != nil {
		log.Errorf("[ NEWRELIC ] %v", err)
		os.Exit(10)
	}
	if err := NRapp.WaitForConnection(5 * time.Second); nil != err {
		log.Errorf("[ NEWRELIC ] %v", err)
		os.Exit(10)
	}
	log.Debug("NR application has been initialized and connection established")
	NRapp.RecordCustomEvent("BundApplication", map[string]interface{}{
		"msg": fmt.Sprintf("Application %s started", *conf.Name),
	})
}
