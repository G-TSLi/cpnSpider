package spider

import (
	"fmt"
)

type SpiderSpecies struct {
	list   []*Spider
	hash   map[string]*Spider
}

var Species = &SpiderSpecies{
	list: []*Spider{},
	hash: map[string]*Spider{},
}

func (self *SpiderSpecies) Add(sp *Spider) *Spider {
	name :=sp.Name
	for i := 2; true; i++ {
		if _, ok := self.hash[name]; !ok {
			sp.Name = name
			self.hash[sp.Name] = sp
			break
		}
		name = fmt.Sprintf("%s(%d)", sp.Name, i)
	}
	sp.Name = name
	self.list = append(self.list, sp)
	return sp
}

func (self *SpiderSpecies) Get() []*Spider {
	return self.list
}

func (self *SpiderSpecies) GetByName(name string) *Spider {
	return self.hash[name]
}