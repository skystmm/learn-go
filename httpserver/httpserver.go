package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(rps http.ResponseWriter, req *http.Request) {
	fmt.Print(req.Form)
	fmt.Fprint(rps, "test")
}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}

}
