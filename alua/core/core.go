package core

import (
	"log"
	"os"
)

var ch = make(chan struct{})

type Server struct {
}

func (server Server) run(config Config) {

}

func (server Server) SubmitJobs(config Config, jobMap map[string]func([]byte)) {
	if len(jobMap) == 0 {
		log.Fatal("no job submit")
		os.Exit(-1)
	}

	//for key,value := range jobMap{
	//	go spiderSingle(key,value)
	//}
}

func SubmitJob(task Task, f func([]byte)) {
	go spiderSingle(task, f)
	ch <- struct{}{}
}

// 基础请求
func spiderSingle(task Task, f func([]byte)) {
	f(request(task))
	<-ch
}
