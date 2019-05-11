package config

type Log struct {
	LogPath     string //日志保存路径
	LogName     string //日志保存的名称，不些随机生成
	LogLevel    string //日志记录级别
	MaxSize     int    //日志分割的尺寸 MB
	MaxAge      int    //分割日志保存的时间 day
	Stacktrace  string //记录堆栈的级别
	IsStdOut    string //是否标准输出console输出
	ProjectName string //项目名称
	SentryDSN   string //SentryDSN 地址
	ElasticURL  string //elastic 地址
}
