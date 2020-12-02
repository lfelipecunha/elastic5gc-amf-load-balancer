package logger

import (
	"os"
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"

	"amfLoadBalancer/lib/logger_conf"
	"amfLoadBalancer/lib/logger_util"
)

var log *logrus.Logger
var NudrDataRepositoryLog *logrus.Entry

func init() {
	log = logrus.New()
	log.SetReportCaller(false)

	log.Formatter = &formatter.Formatter{
		TimestampFormat: time.RFC3339,
		TrimMessages:    true,
		NoFieldsSpace:   true,
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
	}

	amfLoadBalancerLogHook, err := logger_util.NewFileHook(logger_conf.Free5gcLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err == nil {
		log.Hooks.Add(amfLoadBalancerLogHook)
	}

	selfLogHook, err := logger_util.NewFileHook(logger_conf.LibLogDir+"nudr_data_repository.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err == nil {
		log.Hooks.Add(selfLogHook)
	}

	NudrDataRepositoryLog = log.WithFields(logrus.Fields{"component": "OAPI", "category": "NudrDataRepo"})
}

func SetLogLevel(level logrus.Level) {
	NudrDataRepositoryLog.Infoln("set log level :", level)
	log.SetLevel(level)
}

func SetReportCaller(bool bool) {
	NudrDataRepositoryLog.Infoln("set report call :", bool)
	log.SetReportCaller(bool)
}