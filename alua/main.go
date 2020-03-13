package main

import (
	"fmt"
	"learn-go/alua/core"
	"net/http"
)

func main() {
	jobs := make(map[string]func([]byte))
	jobs["http://baidu.com"] = baiduParse
	task := core.Task{}
	param := make(map[string]string)
	param["test"] = "123"
	task.Build("http://www.baidu.com", http.MethodGet, "", param, nil)
	core.SubmitJob(task, baiduParse)

}

func baiduParse(b []byte) {
	fmt.Print(string(b))
}
