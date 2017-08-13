package mgo

import (
	mgo "gopkg.in/mgo.v2"
	"cpnSpider/common/pool"
	"cpnSpider/config"
	"time"
)

type MgoSrc struct {
	*mgo.Session
}



var (
	connGcSecond = time.Duration(config.MGO_CONN_GC_SECOND) * 1e9
	session      *mgo.Session
	err          error
	MgoPool      = pool.ClassicPool(
		config.MGO_CONN_CAP,
		config.MGO_CONN_CAP/5,
		func() (pool.Src, error) {
			return &MgoSrc{session.Clone()}, err
		},
		connGcSecond)
)



// 调用资源池中的资源
func Call(fn func(pool.Src) error) error {
	return MgoPool.Call(fn)
}


