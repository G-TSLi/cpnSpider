package mgo

import (
	"cpnSpider/common/pool"
	mgo "gopkg.in/mgo.v2"
)

type List struct {
	Dbs []string
}

func (self *List) Exec(resultPtr interface{}) (err error)  {
	//defer func() {
	//	if re := recover(); re != nil {
	//		err = fmt.Errorf("%v", re)
	//	}
	//}()
	err = Call(func(src pool.Src) error {

		resultPtr2 := resultPtr.(*map[string][]string)
		*resultPtr2 = map[string][]string{}


		s, err := mgo.Dial("192.168.1.104:27017")  //连接数据库
		var dbs []string

		if dbs, err = s.DatabaseNames(); err != nil {
			return err
		}

		if len(self.Dbs) == 0 {
			for _, dbname := range dbs {
				(*resultPtr2)[dbname], _ = s.DB(dbname).CollectionNames()
			}
			return err
		}
		for _, dbname := range self.Dbs {
			(*resultPtr2)[dbname], _ = s.DB(dbname).CollectionNames()
		}
		return err
	})
	return
}
