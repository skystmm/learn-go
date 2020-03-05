package main

import (
	"fmt"
)

type IPAddr [4]byte

type firstInterface interface {
	test(a string)
	test2()
}
type testStruct struct {
	firstInterface
	A string
}

func (t testStruct) test(a string) {
	fmt.Println(a + t.A)
}

// TODO: 给 IPAddr 添加一个 "String() string" 方法
func (ipAddr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}
func main() {
	t := testStruct{A: "asd"}
	t.test("dsa")
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
