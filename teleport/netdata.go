package teleport

type NetData struct {
	// 消息体
	Body interface{}
	// 操作代号
	Operation string
	// 发信节点uid
	From string
	// 收信节点uid
	To string
}

func NewNetData(from, to, operation string, flag string, body interface{}) *NetData {
	return &NetData{
		From:      from,
		To:        to,
		Body:      body,
		Operation: operation,
	}
}