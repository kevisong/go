package common

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// HTTPRequester interface
type HTTPRequester interface {
	Post(url string, body []byte) ([]byte, error)
	Get(url string) ([]byte, error)
}

// Req sends http request
type Req struct {
	c       *http.Client
	headers map[string]string
	cookies []http.Cookie
}

var defaultReq = &Req{}

// NewHTTPRequester Factory
func NewHTTPRequester(timeout int, headers map[string]string, cookies []http.Cookie) HTTPRequester {
	return &Req{
		c:       &http.Client{Timeout: time.Duration(timeout) * time.Second},
		headers: headers,
		cookies: cookies,
	}
}

// SetHeaders sets http request headers, which overwrites existing headers
func (r *Req) SetHeaders(headers map[string]string) {
	r.headers = headers
}

// SetCookies sets cookies, which overwrites existing cookies
func (r *Req) SetCookies(cookies []http.Cookie) {
	r.cookies = cookies
}

// Post by default requester
func Post(url string, body []byte) ([]byte, error) {
	return defaultReq.Post(url, body)
}

// Get by dfault requester
func Get(url string) ([]byte, error) {
	return defaultReq.Get(url)
}

// Post implements HTTPRequester interface
func (r *Req) Post(url string, body []byte) ([]byte, error) {

	req, err := r.newReq("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	r.addHeaders(req)
	r.addCookies(req)

	resp, err := r.doRequest(req)
	if err != nil {
		return nil, err
	}

	return r.readBody(resp)

}

// Get implements HTTPRequester interface
func (r *Req) Get(url string) ([]byte, error) {

	req, err := r.newReq("GET", url, nil)
	if err != nil {
		return nil, err
	}

	r.addHeaders(req)
	r.addCookies(req)

	resp, err := r.doRequest(req)
	if err != nil {
		return nil, err
	}

	return r.readBody(resp)

}

func (r *Req) newReq(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		errorInfo := fmt.Sprintf("http.NewRequest failed for %s %s, error: %s", "GET", url, err)
		log.Error(errorInfo)
		return nil, errors.New(errorInfo)
	}
	defer req.Body.Close()
	return req, nil
}

func (r *Req) addHeaders(req *http.Request) {
	if r.headers != nil {
		for key, value := range r.headers {
			req.Header.Add(key, value)
		}
	}
}

func (r *Req) addCookies(req *http.Request) {
	if r.cookies != nil {
		for _, cookie := range r.cookies {
			req.AddCookie(&cookie)
		}
	}
}

func (r *Req) doRequest(req *http.Request) (*http.Response, error) {
	resp, err := r.c.Do(req)
	if err != nil {
		errorInfo := fmt.Sprintf("http.Client.Do failed, error: %s", err)
		log.Error(errorInfo)
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}

func (r *Req) readBody(resp *http.Response) ([]byte, error) {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errorInfo := fmt.Sprintf("ioutil.ReadAll failed, error: %s", err)
		log.Error(errorInfo)
		return nil, errors.New(errorInfo)
	}
	return bodyBytes, nil
}
