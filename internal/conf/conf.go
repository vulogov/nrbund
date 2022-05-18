package conf

import (
	"fmt"
	"os"
	"time"
	"github.com/google/uuid"
	"gopkg.in/alecthomas/kingpin.v2"
)

type filelist []string

var Argv [][]string

func (i *filelist) Set(value string) error {
	_, err := os.Stat(value)
	if os.IsNotExist(err) {
		return fmt.Errorf("Script file '%s' not found", value)
	} else {
		*i = append(*i, value)
		return nil
	}
}

func (i *filelist) String() string {
	return ""
}

func (i *filelist) IsCumulative() bool {
	return true
}

func FileList(s kingpin.Settings) (target *[]string) {
	target = new([]string)
	s.SetValue((*filelist)(target))
	return
}

var (
	seed    = time.Now().UTC().UnixNano()
	App     = kingpin.New("NRBUND", fmt.Sprintf("[ NRBUND ] Language that is Functional and Stack-based: %v", BVersion))
	Name   	= App.Flag("name", "Application name.").Required().String()
	Id      = App.Flag("id", "Application name.").Default(uuid.New().String()).String()
	Debug   = App.Flag("debug", "Enable debug mode.").Default("false").Bool()
	CDebug  = App.Flag("core-debug", "Enable core debug mode.").Default("false").Bool()
	Color   = App.Flag("color", "--color : Enable colors on terminal --no-color : Disable colors .").Default("true").Bool()
	VBanner = App.Flag("banner", "Display [ NRBUND ] banner .").Default("false").Bool()
	Timeout = App.Flag("timeout", "Timeout for common NRBUND operations").Default("5s").Duration()
	NRAccount 		= App.Flag("nraccount", "New Relic account.").Envar("NEWRELIC_ACCOUNT").String()
	NRKey 				= App.Flag("nrkey", "New Relic API key.").Envar("NEWRELIC_API_KEY").String()
	NRLicenseKey 	= App.Flag("nrlicensekey", "New Relic API key.").Envar("NEWRELIC_LICENSE_KEY").String()
	NRIngestKey 	= App.Flag("nringestkey", "New Relic License key.").Envar("NEWRELIC_INGEST_KEY").String()
	Etcd				= App.Flag("etcd", "ETCD endpoint location").Default("127.0.0.1:2379").Strings()
	Gnats   		= App.Flag("gnats", "GNATS endpoint location").Default("0.0.0.0:4222").String()
	ShowResult 	= App.Flag("displayresult", "Display result of [ NRBUND ] expression evaluation").Default("false").Bool()
	Args    = App.Flag("args", "String of arguments passed to a script").String()


	Version = App.Command("version", "Display information about [ NRBUND ]")
	VTable  = Version.Flag("table", "Display [ NRBUND ] inner information .").Default("true").Bool()

	Shell      	= App.Command("shell", "Run [ NRBUND ] in interactive shell")
	ShowSResult = Shell.Flag("result", "Display result of expressions evaluated in [ NRBUND ] shell").Default("false").Short('r').Bool()
	SExpr 			= Shell.Arg("expression", "[ NRBUND ] expression passed to shell.").String()

	Run        	= App.Command("run", "Run NRBUND in non-interactive mode")
	Scripts    	= Run.Arg("Scripts", "[ NRBUND ] code to load").Strings()
	ShowRResult = Run.Flag("result", "Display result of scripts execution as it returned by [ NRBUND ]").Default("false").Short('r').Bool()

	Eval 				= App.Command("eval", "Evaluate a [ NRBUND ] expression")
	EStdin  		= Eval.Flag("stdin", "Read [ NRBUND ] expression from STDIN .").Default("false").Bool()
	Expr 				= Eval.Arg("expression", "[ NRBUND ] expression.").String()
	ShowEResult = Eval.Flag("result", "Display result of [ NRBUND ] expression evaluation").Default("false").Short('r').Bool()

	Agitator   	= App.Command("agitator", "Run [ NRBUND ] Agitator")
	UploadConf  = App.Flag("updateconf", "Update etcd configuration from local Agitator configuration").Default("false").Bool()
	AConf 			= Agitator.Flag("conf", "Configuration file for Agitator scheduler.").Required().Strings()

	Agent   		= App.Command("agent", "Run [ NRBUND ] Agent")

	Config   		= App.Command("config", "Send configuration to ETCD")

	Submit   		= App.Command("submit", "Schedule NRBUND script to be executed")
	SScript 		= Submit.Arg("script", "BUND URL to the script, submitted to NRBUND for execution").Default("-").String()

	Sync   			= App.Command("sync", "Send NRBUND SYNC event")

	Take   			= App.Command("take", "Take a single scheduled NRBUND script and execute it")

	Watch   		= App.Command("watch", "Watch for NRBUND event on message bus and print them to Stdout")

	Stop    		= App.Command("stop", "Send 'STOP' signal to a NRBUND bus")

)
