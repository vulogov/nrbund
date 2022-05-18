package bund

import (
	"github.com/google/uuid"
	"github.com/vmihailenco/msgpack"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/nrbund/internal/conf"
	"github.com/vulogov/nrbund/internal/signal"
)

type NRBundPacket struct {
	Id        string
	PktId     string
	OrigName  string
	OrgRole   string
	PktClass  string
	PktKey    string
	Args      []interface{}
	Value     []byte
}

func Marshal(orole string, pktclass string, pktkey string, args []interface{}, value []byte) ([]byte, error) {
	res := new(NRBundPacket)
	res.PktId = uuid.New().String()
	res.Id 				= *conf.Id
	res.OrigName 	= *conf.Name
	res.OrgRole   = orole
	res.PktClass  = pktclass
	res.PktKey  	= pktkey
	res.Args      = args
	res.Value 		= value
	return msgpack.Marshal(res)
}

func UnMarshal(data []byte) *NRBundPacket {
	res := new(NRBundPacket)
	err := msgpack.Unmarshal(data, res)
	if err != nil {
		log.Errorf("[ PACKET ] %v", err)
		return nil
	}
	return res
}

func IfSTOP(msg *NRBundPacket) bool {
	if msg.PktClass == "SYS" && msg.PktKey == "STOP" {
		signal.ExitRequest()
		DoContinue = false
		return true
	}
	return false
}

func IfSYNC(msg *NRBundPacket) bool {
	if msg.PktClass == "SYS" && msg.PktKey == "SYNC" {
		return true
	}
	return false
}

func MakeSync(orole string) ([]byte, error) {
	return Marshal(orole, "SYS", "SYNC", nil, nil)
}

func MakeStop(orole string) ([]byte, error) {
	return Marshal(orole, "SYS", "STOP", nil, nil)
}

func MakeScript(orole string, script []byte, args []interface{}) ([]byte, error) {
	return Marshal(orole, "SYS", "Agitator", args, script)
}
