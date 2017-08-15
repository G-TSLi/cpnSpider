package teleport


type Teleport interface {
	SetAPI(api API) Teleport
}

type TP struct {
	api API
}

type API map[string]Handle

type Handle interface {
	Process(data *NetData) *NetData
}

func New() Teleport {
	return &TP{
		api:           API{},
	}
}

func (self *TP) SetAPI(api API) Teleport {
	self.api = api
	return self
}







