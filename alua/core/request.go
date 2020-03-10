package core

import (
	"bytes"
	"encoding/json"
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

func (task *Task) Build(url string, method string, contentType string, params map[string]string, header map[string]string) {
	//TODO check
	task.url = url
	task.params = params
	task.contentType = contentType
	task.method = method
	task.header = header
}

func request(task Task) []byte {
	client := http.Client{}
	var resp *http.Response
	var err error
	//TODO  待优化 统一使用Do方法，可控制header信息
	if task.method == http.MethodGet {
		url := task.url
		if len(task.params) > 0 {
			url = buildUrl(url, task.params)
		}
		resp, err = client.Get(url)
	} else {
		body := buildBody(task)
		resp, err = client.Post(task.url, task.contentType, body)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return body
}

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
