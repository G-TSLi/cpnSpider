package config

const (
	dbname                string = TAG                         // 数据库名称
	mgoconnstring         string = "192.168.1.104:27017"           // mongodb连接字符串
	mgoconncap            int    = 1024                        // mongodb连接池容量
	mgoconngcsecond       int64  = 600                         // mongodb连接池GC时间，单位秒
)