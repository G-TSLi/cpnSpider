package surfer

import (
	"net/http/cookiejar"
	"net/http"
	"net"
	"time"
)

type Surf struct {
	cookieJar *cookiejar.Jar
}

func New() Surfer {
	s := new(Surf)
	s.cookieJar, _ = cookiejar.New(nil)
	return s
}

func (self *Surf) Download(req Request) (resp *http.Response, err error) {
	param, err := NewParam(req)
	if err != nil {
		return nil, err
	}
	param.client = self.buildClient(param)
	resp, err = self.httpRequest(param)
	return
}

func (self *Surf) buildClient(param *Param) *http.Client {
	client := &http.Client{

	}
	if param.enableCookie {
		client.Jar = self.cookieJar
	}
	transport := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(network, addr, param.dialTimeout)
			if err != nil {
				return nil, err
			}
			if param.connTimeout > 0 {
				c.SetDeadline(time.Now().Add(param.connTimeout))
			}
			return c, nil
		},
	}

	if param.proxy != nil {
		transport.Proxy = http.ProxyURL(param.proxy)
	}

	client.Transport = transport
	return client
}

func (self *Surf) httpRequest(param *Param) (resp *http.Response, err error) {
	req, err := http.NewRequest(param.method, param.url.String(), param.body)
	if err != nil {
		return nil, err
	}

	req.Header=param.header

	if param.tryTimes <= 0 {
		for {
			resp, err = param.client.Do(req)
			if err != nil {
				time.Sleep(param.retryPause)
			}
			break
		}
	} else {
		for i := 0; i < param.tryTimes; i++ {
			resp, err = param.client.Do(req)
			if err != nil {
				time.Sleep(param.retryPause)
				continue
			}
			break
		}
	}

	return resp, err
}