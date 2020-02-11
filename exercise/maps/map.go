package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

/*
实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。
*/
func WordCount(s string) map[string]int {
	res := make(map[string]int)
	tar := strings.Fields(s)
	for i := 0; i < len(tar); i++ {
		tmp := tar[i]
		value, contain := res[tmp]
		if contain {
			res[tmp] = value + 1
		} else {
			res[tmp] = 1
		}
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
