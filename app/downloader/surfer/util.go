package surfer

import "net/url"

func UrlEncode(urlStr string) (*url.URL, error) {
	urlObj, err := url.Parse(urlStr)
	urlObj.RawQuery = urlObj.Query().Encode()
	return urlObj, err
}