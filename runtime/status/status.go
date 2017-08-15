package status

// 运行模式
const (
	UNSET int = iota - 1
	OFFLINE
	SERVER
	CLIENT
)

// 运行状态
const (
	STOPPED = iota - 1
	STOP
	RUN
	PAUSE
)
