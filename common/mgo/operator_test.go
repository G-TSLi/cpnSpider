package mgo

import "testing"

func TestMgo(t *testing.T) {
	var li = map[string][]string{}
	Mgo(&li, "list", map[string]interface{}{
		"Dbs": []string{"test"},
	})
	t.Log(li)
}