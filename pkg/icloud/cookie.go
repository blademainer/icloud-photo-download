package icloud

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/publicsuffix"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type CookieJar struct {
	wrapped http.CookieJar
}

func (c *CookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	logrus.Debugf("setting cookie: %v cookies: %v", u.String(), cookies)
	c.wrapped.SetCookies(u, cookies)
}

func (c *CookieJar) Cookies(u *url.URL) []*http.Cookie {
	cookies := c.wrapped.Cookies(u)
	logrus.Debugf("get cookie: %v cookies: %v", u.String(), cookies)
	return cookies
}

func NewCookieJar() (http.CookieJar, error) {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}
	c := &CookieJar{wrapped: jar}
	return c, nil
}
