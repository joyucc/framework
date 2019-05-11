package framework

import (
	"github.com/abaole/framework/config"
	"github.com/abaole/framework/database"
	"github.com/abaole/framework/database/orm"
	"github.com/abaole/framework/logger"
	logCfg "github.com/abaole/framework/logger/conf"
	"github.com/abaole/framework/logger/plugins/logrus"
	"github.com/abaole/framework/pprof"
)

//保存需要关闭的选项
var (
	closeArgs []string
	cfg       *config.Config
)

func configInit(path string) {
	config.SetConfigurationPath(path)
	cfg = config.GetConfig()
}

//初始化选项
//logger:日志(必须) trace:链路跟踪 mysql:mysql数据库 mongo:MongoDB
func InitOption(path string, args ...string) {
	//开启pprof
	go pprof.Run()
	//保存需要关闭的参数
	closeArgs = args
	//1.初始化配置参数
	configInit(path)
	//2.初始化日志
	logInit()
	//3.其他服务
	for _, o := range args {
		switch o {
		case "mysql":
			database.SetDatabase(orm.New(
				orm.WithDialect(cfg.Database.Adapter),
				orm.WithDSN(cfg.Database.URL()),
				orm.WithActive(cfg.Database.Active),
				orm.WithIdle(cfg.Database.Idle),
				orm.WithIdleTimeout(cfg.Database.IdleTimeout),
			))
		}
	}
}

//关闭打开的服务
func Close() {
	for _, o := range closeArgs {
		switch o {
		case "mysql":
			//关闭mysql
			database.Close()
		}
	}
}

func logInit() {
	logger.SetLogger(logrus.New(
		logCfg.WithProjectName(""),
		logCfg.WithLogPath(""),
		logCfg.WithLogName(""),
		logCfg.WithMaxAge(""),
		logCfg.WithMaxSize(""),
		logCfg.WithIsStdOut(""),
		logCfg.WithLogLevel(""),
		logCfg.WithSentryDSN(""),
		logCfg.WithElasticURL(""),
	))
}
