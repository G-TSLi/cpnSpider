package app

import (
	"cpnSpider/runtime/cache"
	"cpnSpider/app/spider"
	"cpnSpider/app/distribute"
	"cpnSpider/app/crawler"
	"reflect"
	"strings"
	"cpnSpider/teleport"
)

type (
	App interface {
		Init(mode int, port int, master string) App
		ReInit(mode int, port int, master string) App
		GetAppConf(k ...string) interface{}
		SetAppConf(k string, v interface{}) App
		SpiderPrepare(original []*spider.Spider) App
	}
	Logic struct {
		*cache.AppConf
		*spider.SpiderSpecies
		crawler.SpiderQueue
		*distribute.TaskJar
		teleport.Teleport
	}
)

// 全局唯一的核心接口实例
var LogicApp = New()

func New() App {
	return newLogic()
}

func newLogic() *Logic {
	return &Logic{
		AppConf:       cache.Task,
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

func (self *Logic) Run()  {
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

