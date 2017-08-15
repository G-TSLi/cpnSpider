package distribute

type Task struct {
	ThreadNum      int                 // 全局最大并发量
	Pausetime      int64               // 暂停时长
	Limit          int64               // 采集上限，0为不限
}