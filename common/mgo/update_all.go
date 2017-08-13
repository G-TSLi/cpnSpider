package mgo

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"../pool"
)

type UpdateAll struct {
	Database 	string
	Collection 	string
	Selector 	map[string]interface{}
	Change		map[string]interface{}
}

func (self *UpdateAll) Exce(resultPtr interface{}) (err error)  {

	resultPtr2 := resultPtr.(*map[string]interface{})
	*resultPtr2 = map[string]interface{}{}


	err =Call(func(src pool.Src) error {
		c := src.(*MgoSrc).DB(self.Database).C(self.Collection)

		if id, ok := self.Selector["_id"]; ok {
			if idStr, ok2 := id.(string); !ok2 {
				return fmt.Errorf("%v", "参数 _id 必须为 string 类型！")
			} else {
				self.Selector["_id"] = bson.ObjectIdHex(idStr)
			}
		}

		info, err := c.UpdateAll(self.Selector, self.Change)
		if err != nil {
			return err
		}

		(*resultPtr2)["Updated"] = info.Updated
		(*resultPtr2)["Removed"] = info.Removed
		(*resultPtr2)["UpsertedId"] = info.UpsertedId

		return err
	})
	return
}