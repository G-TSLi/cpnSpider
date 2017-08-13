package config

// 软件信息
const (
	VERSION   string = "v1.0.0"                                      // 软件版本号
	AUTHOR    string = "lishuntao"                                 // 软件作者
	NAME      string = "企业数据采集"                              // 软件名
	FULL_NAME string = NAME + "_" + VERSION + " （by " + AUTHOR + "）" // 软件全称
	TAG       string = "cpn"                                     // 软件标识符
)


var (
	DB_NAME                  string = dbname                                             // 数据库名称
	MGO_CONN_STR             string = mgoconnstring                                    // mongodb连接字符串
	MGO_CONN_CAP             int    = mgoconncap                       // mongodb连接池容量
	MGO_CONN_GC_SECOND       int64  = mgoconngcsecond           // mongodb连接池GC时间，单位秒
)