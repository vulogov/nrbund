package bund

import (
	"os"
	"context"
	"fmt"
	"strings"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/nrbund/internal/conf"
	"github.com/vulogov/nrbund/internal/signal"
	"go.etcd.io/etcd/client/v3"
)

var Etcd *clientv3.Client

func InitEtcdAgent() {
	var err error

	log.Debugf("Connecting to ETCD: %v", *conf.Etcd)
	Etcd, err = clientv3.New(
		clientv3.Config{
			Endpoints: *conf.Etcd,
			DialTimeout: *conf.Timeout,
		})
	if err != nil {
		log.Errorf("[ ETCD ] %v", err)
		signal.ExitRequest()
		os.Exit(10)
	}
	log.Debug("Sync ETCD endpoints")
	ctx, _ := context.WithTimeout(context.Background(), *conf.Timeout)
	err = Etcd.Sync(ctx)
	if err != nil {
		log.Errorf("[ ETCD ] %v", err)
		signal.ExitRequest()
		os.Exit(10)
	}
}

func EtcdSetItem(key string, value string) {
	ctx, _ := context.WithTimeout(context.Background(), *conf.Timeout)
	_, err := Etcd.Put(ctx, fmt.Sprintf("NRBUND/%s/%s", *conf.NRAccount, key), value)
	if err != nil {
		log.Errorf("[ ETCD ] %v", err)
		signal.ExitRequest()
		os.Exit(10)
	}
}

func EtcdGetItems()  *map[string]string {
	ctx, _ := context.WithTimeout(context.Background(), *conf.Timeout)
	value, err := Etcd.Get(ctx, fmt.Sprintf("NRBUND/%s/", *conf.NRAccount), clientv3.WithPrefix())
	if err != nil {
		log.Errorf("[ ETCD ] %v", err)
		signal.ExitRequest()
		os.Exit(10)
	}
	res := make(map[string]string)
	for _, v := range(value.Kvs) {
		key := strings.Split(string(v.Key), "/")
		res[key[len(key)-1]] = string(v.Value)
	}
	return &res
}

func UpdateConfigToEtcd() {
	if len(*conf.Etcd) > 0 {
		log.Debugf("Upload NRBUND configuration to ETCD")
		addr := (*conf.Etcd)[0]
		log.Debugf("Update ETCD endpoints with %s", addr)
		EtcdSetItem("etcd", addr)
		log.Debugf("Update GNATS endpoints with %s", *conf.Gnats)
		EtcdSetItem("gnats", addr)
		log.Debug("Upload NR keys")
		EtcdSetItem("NEWRELIC_ACCOUNT", *conf.NRAccount)
		EtcdSetItem("NEWRELIC_API_KEY", *conf.NRKey)
		EtcdSetItem("NEWRELIC_INGEST_KEY", *conf.NRIngestKey)
		EtcdSetItem("NEWRELIC_LICENSE_KEY", *conf.NRLicenseKey)
	}
}

func CloseEtcdAgent() {
	log.Debug("Closing ETCD agent")
	if Etcd != nil {
		Etcd.Close()
	}
}
