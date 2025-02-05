package log

import (
	"log"
	"log/syslog"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var (
	logger  *log.Logger
	Println func(...any)
	Printf  func(string, ...any)
	Fatal   func(...any)
	Fatalf  func(string, ...any)
)

const (
	Error = "[ERROR]"
	Info  = "[INFO]"
)

type Deps struct {
	fx.In
	Config *viper.Viper
}

func New(deps Deps) *log.Logger {
	syslog, err := syslog.Dial("udp", deps.Config.GetString("log.host"), syslog.LOG_USER, "pocone")
	if err != nil {
		log.Fatalf("creating log file: %v", err)
	}

	logger = log.Default()
	logger.SetOutput(syslog)
	logger.SetFlags(log.LstdFlags | log.Llongfile | log.Lmicroseconds)

	Println = logger.Println
	Printf = logger.Printf
	Fatal = logger.Fatal
	Fatalf = logger.Fatalf

	return logger
}

func Noop() {
	logger = log.Default()
	Println = logger.Println
	Printf = logger.Printf
	Fatal = logger.Fatal
	Fatalf = logger.Fatalf
}
