package conf

import (
	"fmt"
	"os"
	"time"

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
	Debug   = App.Flag("debug", "Enable debug mode.").Default("false").Bool()
	CDebug  = App.Flag("core-debug", "Enable core debug mode.").Default("false").Bool()
	Color   = App.Flag("color", "--color : Enable colors on terminal --no-color : Disable colors .").Default("true").Bool()
	VBanner = App.Flag("banner", "Display [ NRBUND ] banner .").Default("false").Bool()
	NRAccount 		= App.Flag("nraccount", "New Relic account.").Envar("NEWRELIC_ACCOUNT").String()
	NRKey 				= App.Flag("nrkey", "New Relic API key.").Envar("NEWRELIC_API_KEY").String()
	NRLicenseKey 	= App.Flag("nrlicensekey", "New Relic API key.").Envar("NEWRELIC_LICENSE_KEY").String()
	NRIngestKey 	= App.Flag("nringestkey", "New Relic License key.").Envar("NEWRELIC_INGEST_KEY").String()
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

	Agitiator   = App.Command("agitator", "Run [ NRBUND ] Agitator")

	Agent   		= App.Command("agent", "Run [ NRBUND ] Agent")

)
