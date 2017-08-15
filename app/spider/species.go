package spider

type SpiderSpecies struct {
	list   []*Spider
}


var Species = &SpiderSpecies{
	list: []*Spider{},
}

func (self *SpiderSpecies) Get() []*Spider {
	return self.list
}