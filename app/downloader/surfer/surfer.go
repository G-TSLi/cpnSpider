package surfer

import (
	"net/http"
	"sync"
)


var (
	surf         Surfer
	phantom      Surfer
	once_surf    sync.Once
	once_phantom sync.Once
	tempJsDir    = "./tmp"
	phantomjsFile = `./phantomjs`
)


func Download(req Request) (resp *http.Response, err error) {
	switch req.GetDownloaderID() {
	case SurfID:
		once_surf.Do(func() { surf = New() })
		resp, err = surf.Download(req)
	}
	return
}

type Surfer interface {
	Download(Request) (resp *http.Response, err error)
}