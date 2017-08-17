package app

import (
	"cpnSpider/runtime/cache"
	"cpnSpider/app/spider"
	"cpnSpider/app/distribute"
	"cpnSpider/app/crawler"
	"reflect"
	"strings"
	"cpnSpider/teleport"
	"sync"
	"cpnSpider/runtime/status"
)

type (
	App interface {
		Init(mode int, port int, master string) App
		ReInit(mode int, port int, master string) App
		GetAppConf(k ...string) interface{}
		SetAppConf(k string, v interface{}) App
		SpiderPrepare(original []*spider.Spider) App
		GetSpiderQueue() crawler.SpiderQueue
		GetSpiderLib() []*spider.Spider
		Status() int
		Run()
	}
	Logic struct {
		*cache.AppConf
		*spider.SpiderSpecies
		crawler.SpiderQueue
		*distribute.TaskJar
		teleport.Teleport
		status		int
		sync.RWMutex
	}
)

// 全局唯一的核心接口实例
var LogicApp = New()

func New() App {
	return newLogic()
}

func newLogic() *Logic {
	return &Logic{
		AppConf:		cache.Task,
		SpiderSpecies: 	spider.Species,
		status:        	status.STOPPED,
		SpiderQueue:   crawler.NewSpiderQueue(),
	}
}

// 获取全局参数
func (self *Logic) GetAppConf(k ...string) interface{} {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	if len(k) == 0 {
		return self.AppConf
	}
	key := strings.Title(k[0])
	acv := reflect.ValueOf(self.AppConf).Elem()
	return acv.FieldByName(key).Interface()
}

func (self *Logic) SetAppConf(k string, v interface{}) App {
	acv := reflect.ValueOf(self.AppConf).Elem()
	key := strings.Title(k)
	if acv.FieldByName(key).CanSet() {
		acv.FieldByName(key).Set(reflect.ValueOf(v))
	}
	return self
}

// 通过名字获取某蜘蛛
func (self *Logic) GetSpiderByName(name string) *spider.Spider {
	return self.SpiderSpecies.GetByName(name)
}

// 返回当前运行状态
func (self *Logic) Status() int {
	self.RWMutex.RLock()
	defer self.RWMutex.RUnlock()
	return self.status
}

// 获取蜘蛛队列接口实例
func (self *Logic) GetSpiderQueue() crawler.SpiderQueue {
	return self.SpiderQueue
}

func (self *Logic) Init(mode int, port int, master string) App {
	self.AppConf.Mode, self.AppConf.Port, self.AppConf.Master = mode, port, master
	self.Teleport = teleport.New()
	self.TaskJar = distribute.NewTaskJar()
	self.SpiderQueue = crawler.NewSpiderQueue()
	return self
}

func (self *Logic) ReInit(mode int, port int, master string) App{
	// 重新开启
	self = newLogic().Init(mode, port, master).(*Logic)
	return self
}

func (self *Logic) SpiderPrepare(original []*spider.Spider) App {
	for _, sp := range original {
		self.SpiderQueue.Add(sp)
	}
	return self
}


// 获取全部蜘蛛种类
func (self *Logic) GetSpiderLib() []*spider.Spider {
	return self.SpiderSpecies.Get()
}

func (self *Logic) addNewTask() (tasksNum int)  {
	t := distribute.Task{}
	// 从配置读取字段
	self.setTask(&t)

	// 存入
	one := t
	self.TaskJar.Push(&one)
	tasksNum++
	return
}

func (self *Logic) setTask(task *distribute.Task) {
	task.ThreadNum = self.AppConf.ThreadNum
	task.Pausetime = self.AppConf.Pausetime
	task.Limit = self.AppConf.Limit
}

func (self *Logic) Run() {
	self.offline()
}

// 离线模式运行
func (self *Logic) offline() {
	self.exec()
}

// 开始执行任务
func (self *Logic) exec() {
	count := self.SpiderQueue.Len()
	go self.goRun(count)
}

func (self *Logic) goRun(count int) {


	var i int


	for i=0;i < count && self.Status() != status.STOP;i++{
		// 从爬行队列取出空闲蜘蛛，并发执行
		c :=crawler.New()
		if c != nil {
			go func(i int, c crawler.Crawler) {
				// 执行并返回结果消息
				c.Init(self.SpiderQueue.GetByIndex(i)).Run()
			}(i, c)
		}
	}

}