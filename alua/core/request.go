package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Task struct {
	url         string
	method      string
	contentType string
	header      map[string]string
	params      map[string]string
}

//构建任务对象
func (task *Task) Build(url string, method string, contentType string, params map[string]string, header map[string]string) {
	//TODO check
	task.url = url
	task.params = params
	task.contentType = contentType
	task.method = method
	task.header = header
}

//根据任务进行数据请求
func (task Task) request() []byte {

	var resp *http.Response
	var err error
	resp, err = doRequest(task)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return body
}

//请求处理
func doRequest(task Task) (*http.Response, error) {
	client := http.Client{}
	var req *http.Request
	var body io.Reader
	url := task.url
	switch task.method {
	case http.MethodGet, http.MethodPut:
		if len(task.params) > 0 {
			url = buildUrl(url, task.params)
		}
	case http.MethodPost:
		body = buildBody(task)
	default:
		log.Fatal("unsurport http method")
		return nil, errors.New("unsupport http method")
	}
	req, err := http.NewRequest(task.method, url, body)
	if err != nil {
		log.Fatal("request build fail ")
		return nil, errors.New("request build fail ")
	}
	buildHeader(req, task.header)
	return client.Do(req)
}

//构建header
func buildHeader(req *http.Request, header map[string]string) {
	if len(header) > 0 {
		for key, value := range header {
			req.Header.Add(key, value)
		}
	}
}

//构建http body
func buildBody(task Task) io.Reader {
	var body io.Reader
	switch task.contentType {
	case "application/text", "text/html":
		res := ""
		for key, value := range task.params {
			res += fmt.Sprintf("%s=%s&", key, value)
		}
		body = strings.NewReader(res[:len(res)-1])
	case "application/json":
		jsonParam, _ := json.Marshal(task.params)
		body = bytes.NewReader(jsonParam)
	default:
		log.Fatal("build body error ,unsupport Content-Type")
	}
	return body
}

//get请求参数拼接
func buildUrl(url string, params map[string]string) string {
	if !strings.Contains(url, "?") {
		url += "?"
	}
	for key, value := range params {
		url += fmt.Sprintf("%s=%s&", key, value)
	}
	return url[:len(url)-1]
}
