package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/szpinc/study-go/util"
)

var logger *logrus.Logger

type conf struct {
	log logConf
}

type logConf struct {
	formatter string // text,json
	level     string // 日志级别: trace,debug,info,wran,error
}

func init() {

	conf := conf{}
	// 读取配置文件
	err := util.ReadYaml("/conf/conf.yaml", &conf)

	if err != nil {
		panic(err)
	}

	logger = logrus.New()
	// 设置formatter
	logger.Formatter = &logrus.TextFormatter{}
	// 控制台输出
	logger.SetOutput(os.Stdout)

	level, err := logrus.ParseLevel(conf.log.level)

	if err != nil {
		panic(err)
	}
	// 日志级别
	logger.SetLevel(level)
}

func Logger() *logrus.Logger {
	return logger
}
