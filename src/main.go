package main

import (
	"amfLoadBalancer/lib/path_util"
	"amfLoadBalancer/src/context"
	"amfLoadBalancer/src/factory"
	"amfLoadBalancer/src/loadbalancer"
	"amfLoadBalancer/src/logger"
	"amfLoadBalancer/src/util"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var appLog *logrus.Entry

func init() {
	appLog = logger.AppLog
	logger.SetLogLevel(logrus.DebugLevel)
}

func main() {
	app := cli.NewApp()
	app.Name = "Amf Load Balancer"
	appLog.Infoln(app.Name)
	//appLog.Infoln("AMF version: ", version.GetVersion())
	//app.Usage = "-amfLoadBalancercfg common configuration file -amfcfg amf configuration file"
	app.Action = action
	//app.Flags = AMF.GetCliCmd()
	if err := app.Run(os.Args); err != nil {
		logger.AppLog.Errorf("AMF Load Balancer Run error: %v", err)
	}
}

func action(c *cli.Context) {
	var hosts []string
	hosts = append(hosts, "10.100.200.250")
	factory.InitConfigFactory(path_util.GoamfLoadBalancerPath("amfLoadBalancer/config/load_balancer.conf"))
	util.InitAmfContext(context.AMF_Self())
	loadbalancer.Run(hosts, 38412, loadbalancer.Dispatch)
	//AMF.Initialize(c)
	//AMF.Start()
}
