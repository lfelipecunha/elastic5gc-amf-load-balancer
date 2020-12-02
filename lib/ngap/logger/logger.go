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
var NgapLog *logrus.Entry

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

	selfLogHook, err := logger_util.NewFileHook(logger_conf.LibLogDir+"ngap.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err == nil {
		log.Hooks.Add(selfLogHook)
	}

	NgapLog = log.WithFields(logrus.Fields{"component": "LIB", "category": "NGAP"})
}

func SetLogLevel(level logrus.Level) {
	NgapLog.Infoln("set log level :", level)
	log.SetLevel(level)
}

func SetReportCaller(bool bool) {
	NgapLog.Infoln("set report call :", bool)
	log.SetReportCaller(bool)
}
