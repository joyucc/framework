package logger

import (
	"context"
	"testing"
	"time"

	"github.com/abaole/framework/logger/conf"
	"github.com/abaole/framework/logger/plugins/logrus"
	"github.com/abaole/framework/logger/plugins/zaplog"
)

func TestSetLogger(t *testing.T) {
	//设置为当前目录下 //设置级别
	SetLogger(zaplog.New(
		conf.WithProjectName("zap test"),
		conf.WithLogPath("tmp"),
		conf.WithLogLevel("info"),
	))
	Debug("this is zap")
	Debug("hello", context.Background())
	Infof("hello %s", "world", context.Background())
	Infof("hello %s", "world")
	Infof("hello %s,%d", "world", 2018, context.Background())
	Errorf("hello %s,%d", "world", 2018)
	l2 := logrus.New(conf.WithLogPath("tmp"), conf.WithLogName("logrus"), conf.WithProjectName("logrus test"))
	SetLogger(l2)
	Debugf("this is logrus %s", "test", context.Background())
	time.Sleep(time.Second * 5)
}
