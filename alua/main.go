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
	task.Build("http://baidu.com", http.MethodGet, "", nil, nil)
	core.SubmitJob(task, baiduParse)

}

func baiduParse(b []byte) {
	fmt.Print(string(b))
}
