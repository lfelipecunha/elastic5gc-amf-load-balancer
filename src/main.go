package main

import (
	"amfLoadBalancer/lib/path_util"
	"amfLoadBalancer/src/context"
	"amfLoadBalancer/src/factory"
	"amfLoadBalancer/src/logger"
	"amfLoadBalancer/src/service"
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
	app.Usage = "-balancercfg configuration file"
	app.Action = action
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "free5gccfg",
			Usage: "common config file",
		},
		cli.StringFlag{
			Name:  "balancercfg",
			Usage: "amf config file",
		},
	}

	//app.Flags = AMF.GetCliCmd()
	if err := app.Run(os.Args); err != nil {
		logger.AppLog.Errorf("AMF Load Balancer Run error: %v", err)
	}
}

func action(c *cli.Context) {
	configFile := c.String("balancercfg")

	if configFile != "" {
		factory.InitConfigFactory(configFile)
	} else {
		DefaultAmfConfigPath := path_util.GoamfLoadBalancerPath("amfLoadBalancer/config/load_balancer.conf")
		factory.InitConfigFactory(DefaultAmfConfigPath)
	}

	util.InitAmfContext(context.AMF_Self())
	service.UpdateAmfList()
	service.RunHttp(context.AMF_Self().NgapIpList[0], 8080, context.AMF_Self().NrfUri)
	service.RunSCtp(context.AMF_Self().NgapIpList, 38412, service.Dispatch)
}
