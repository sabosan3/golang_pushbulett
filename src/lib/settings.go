package lib

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/comail/colog"
	yaml "gopkg.in/yaml.v2"
)

/*
InitLogging return log.Logger
*/
func InitLogging() *log.Logger {

	cl := colog.NewCoLog(os.Stdout, "logger", log.LstdFlags)
	cl.SetDefaultLevel(colog.LDebug)
	cl.SetMinLevel(colog.LTrace)
	cl.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})

	logger := cl.NewLogger()

	return logger
}

/*
Config is struct
*/
type Config struct {
	Pushbulett PushBulett `yaml:"pushbulett"`
	Targetsite Target     `yaml:"target"`
}

/*
PushBulett is struct
*/
type PushBulett struct {
	Token string `yaml:"token"`
}

/*
Target is struct
*/
type Target struct {
	URL string `yaml:"url"`
}

/*
LoadConfig return
*/
func LoadConfig() Config {

	logger := InitLogging()

	exe, err := os.Executable()
	if err != nil {
		logger.Printf("error: failed to find path")
		logger.Printf(err.Error())
	}
	basePath := filepath.Dir(exe)

	buf, err := ioutil.ReadFile(basePath + "/conf/api.yaml")
	if err != nil {
		logger.Printf("error: failed to load config")
		logger.Printf(err.Error())
	}

	//conf := make([]Config, 1)
	conf := Config{}

	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		logger.Printf("error: failed to format yaml")
		logger.Printf(err.Error())
	}

	return conf
}
