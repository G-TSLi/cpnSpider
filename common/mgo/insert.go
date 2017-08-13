package mgo

import (
	"../pool"
	"gopkg.in/mgo.v2/bson"
)

type Insert struct {
	Database 	string
	Collection 	string
	Docs 		[]map[string]interface{}
}



const (
	MaxLen int=5000
)

func (self *Insert) Exce(resultPtr interface{}) (err error) {


	var (
		resultPtr2  = new([]string)
		count		= len(self.Docs)
		docs       = make([]interface{}, count)
	)

	if resultPtr != nil {
		resultPtr2 = resultPtr.(*[]string)
	}
	*resultPtr2 = make([]string, count)

	return Call(func(src pool.Src) error {
		c := src.(*MgoSrc).DB(self.Database).C(self.Collection)
		for i, doc := range self.Docs {
			var _id string
			if doc["_id"] == nil || doc["_id"] == interface{}("") || doc["_id"] == interface{}(0) {
				objId := bson.NewObjectId()
				_id = objId.Hex()
				doc["_id"] = objId
			} else {
				_id = doc["_id"].(string)
			}

			if resultPtr != nil {
				(*resultPtr2)[i] = _id
			}
			docs[i] = doc
		}
		loop := count / MaxLen
		for i := 0; i < loop; i++ {
			err := c.Insert(docs[i*MaxLen : (i+1)*MaxLen]...)
			if err != nil {
				return err
			}
		}
		if count%MaxLen == 0 {
			return nil
		}
		return c.Insert(docs[loop*MaxLen:]...)
	})
}
