package distribute

// 任务仓库
type TaskJar struct {
	Tasks	chan *Task
}

func NewTaskJar() *TaskJar {
	return &TaskJar{
		Tasks: make(chan *Task, 1024),
	}
}

func (self *TaskJar) Push(task *Task) {
	self.Tasks <- task
}

func (self *TaskJar) Pull() *Task  {
	return <-self.Tasks
}