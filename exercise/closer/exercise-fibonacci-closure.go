package main

import "fmt"

// 实现一个 fibonacci 函数，它返回一个函数（闭包），该闭包返回一个斐波纳契数列 `(0, 1, 1, 2, 3, 5, ...)`。
// 返回一个“返回int的函数”
func fibonacci() func() int {
	a, b := 0, 1 //初始值为0,1
	return func() int {
		a, b = b, a+b //将a变为b,b变为a+b
		return b - a  //此时应该返回为改变时的a, 也就是a+b-b = b-a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
