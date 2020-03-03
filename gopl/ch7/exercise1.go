package main

import (
	"bufio"
	"fmt"
	"io"
)

type Count int

//练习7.1
func (count *Count) count(data []byte) {
	c, token, error := bufio.ScanWords(data, true)
	if error != nil {
		fmt.Println("word count error %s", error)
	} else {
		fmt.Println(token)
		*count += Count(c)
	}
}

func (count *Count) countLine(data []byte) {
	c, token, error := bufio.ScanLines(data, true)
	if error != nil {
		fmt.Println("word count error %s", error)
	} else {
		fmt.Println(token)
		*count += Count(c)
	}
}

//练习7.2
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c, error := w.Write([]byte("123"))
	if error != nil {
		return nil, nil
	} else {
		res := int64(c)
		return io.Writer(w), &res
	}
}

//练习7.4

func main() {
	wc := Count(0)
	wc.count([]byte("abc222222 qwe qwe"))
	//	wc.countLine([]byte(`abd
	//egf
	//ttt
	//aaa`))
	fmt.Println(wc)
}
