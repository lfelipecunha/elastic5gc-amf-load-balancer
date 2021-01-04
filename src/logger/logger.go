package logger

import (
	"os"
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var AppLog *logrus.Entry
var NgapLog *logrus.Entry

var InitLog *logrus.Entry
var ContextLog *logrus.Entry
var HandlerLog *logrus.Entry
var HttpLog *logrus.Entry
var GmmLog *logrus.Entry
var MtLog *logrus.Entry
var ProducerLog *logrus.Entry
var LocationLog *logrus.Entry
var CommLog *logrus.Entry
var CallbackLog *logrus.Entry
var UtilLog *logrus.Entry
var NasLog *logrus.Entry
var ConsumerLog *logrus.Entry
var EeLog *logrus.Entry
var GinLog *logrus.Entry
var BalancerLog *logrus.Entry

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

	amfLoadBalancerLogHook, err := NewFileHook("/tmp/load_full.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err == nil {
		log.Hooks.Add(amfLoadBalancerLogHook)
	}

	selfLogHook, err := NewFileHook("/tmp/load_balancer.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err == nil {
		log.Hooks.Add(selfLogHook)
	}

	AppLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "App"})
	NgapLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "NGAP"})
	InitLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "Init"})
	ContextLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "Context"})
	HandlerLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "Handler"})
	HttpLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "HTTP"})
	GmmLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "Gmm"})
	MtLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "MT"})
	ProducerLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "Producer"})
	LocationLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "LocInfo"})
	CommLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "Comm"})
	CallbackLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "Callback"})
	UtilLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "Util"})
	NasLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "NAS"})
	ConsumerLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "Consumer"})
	EeLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "EventExposure"})
	GinLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "GIN"})
	BalancerLog = log.WithFields(logrus.Fields{"component": "LOAD BALANCER", "category": "Balancer"})
}

func SetLogLevel(level logrus.Level) {
	log.SetLevel(level)
}

func SetReportCaller(bool bool) {
	log.SetReportCaller(bool)
}
