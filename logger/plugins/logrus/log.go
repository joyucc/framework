package logrus

import (
	"bufio"
	"os"
	"time"

	"github.com/abaole/framework/logger/conf"
	"github.com/abaole/framework/logger/fileout"
	"github.com/olivere/elastic"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/verystar/logrus_sentry"
	"gopkg.in/sohlich/elogrus.v3"
)

type Log struct {
	logger      *logrus.Logger
	ProjectName string
}

//初始化日志
func New(opts ...conf.Option) *Log {
	log := &Log{}
	log.logger = logrus.New()

	o := &conf.Options{
		LogPath:     conf.LogPath,
		LogName:     conf.LogName,
		LogLevel:    conf.LogLevel,
		MaxSize:     conf.MaxSize,
		MaxAge:      conf.MaxAge,
		IsStdOut:    conf.IsStdOut,
		ProjectName: conf.ProjectName,
	}
	for _, opt := range opts {
		opt(o)
	}
	//设置项目名称
	log.ProjectName = o.ProjectName
	//所有日志都输出到文件
	lev, err := logrus.ParseLevel(o.LogLevel)
	if err != nil {
		panic(err.Error())
	}
	log.logger.Level = lev
	if o.IsStdOut != "yes" {
		log.logger.Out = setNull() //将日志写入空接口
	}
	writer := fileout.NewRollingFile(o.LogPath, o.LogName, o.MaxSize, o.MaxAge)

	log.logger.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.PanicLevel: writer,
			logrus.FatalLevel: writer,
			logrus.ErrorLevel: writer,
			logrus.WarnLevel:  writer,
			logrus.InfoLevel:  writer,
			logrus.DebugLevel: writer,
		},
		&logrus.JSONFormatter{},
	))

	if o.SentryDSN != "" {
		tags := map[string]string{
			"type": "logrus",
		}

		hook, err := logrus_sentry.NewWithTagsSentryHook(o.SentryDSN, tags, []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		})
		hook.Timeout = 1 * time.Second
		hook.StacktraceConfiguration.Enable = true

		if err == nil {
			log.logger.Hooks.Add(hook)
		}
	}

	if o.ElasticURL != "" {
		client, err := elastic.NewClient(elastic.SetURL(o.ElasticURL))
		if err != nil {
			log.logger.Panic(err)
		}
		hook, err := elogrus.NewElasticHook(client, "localhost", logrus.DebugLevel, "mylog")
		if err != nil {
			log.logger.Panic(err)
		}
		log.logger.Hooks.Add(hook)
	}

	return log
}

//设置一个空接口，将日志写入空接口

func setNull() *bufio.Writer {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(src)
	return writer
}
