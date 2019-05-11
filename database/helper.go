package database

//默认
var l IDatabase

//设置
func SetDatabase(ll IDatabase) {
	l = ll
}

func Close() {
	l.Close()
}
