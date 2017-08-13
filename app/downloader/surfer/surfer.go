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
	// GET @param url string, header http.Header, cookies []*http.Cookie
	// HEAD @param url string, header http.Header, cookies []*http.Cookie
	// POST PostForm @param url, referer string, values url.Values, header http.Header, cookies []*http.Cookie
	// POST-M PostMultipart @param url, referer string, values url.Values, header http.Header, cookies []*http.Cookie
	Download(Request) (resp *http.Response, err error)
}