package pool

import (
	"time"
	"sync"
	"errors"
)

type (

	Pool interface {
		// 调用资源池中的资源
		Call(func(Src) error) error
	}

	classic struct {
		srcs     chan Src      // 资源列表(Src须为指针类型)
		capacity int           // 资源池容量
		maxIdle  int           // 资源最大空闲数
		len      int           // 当前资源数
		factory  Factory       // 创建资源的方法
		gctime   time.Duration // 空闲资源回收时间
		closed   bool          // 标记是否已关闭资源池
		sync.RWMutex
	}

	Src interface {
	}

	// 创建资源的方法
	Factory func() (Src, error)
)

const (
	GC_TIME = 60e9
)

var (
	closedError = errors.New("资源池已关闭")
)

func ClassicPool(capacity, maxIdle int,factory Factory, gctime ...time.Duration) Pool {
	if len(gctime) == 0 {
		gctime = append(gctime, GC_TIME)
	}
	pool := &classic{
		srcs:     make(chan Src, capacity),
		capacity: capacity,
		maxIdle:  maxIdle,
		factory:  factory,
		gctime:   gctime[0],
		closed:   false,
	}
	return pool
}

func (self *classic) Call(callback func(Src) error) (err error) {
	var src Src
	err = callback(src)
	return err
}