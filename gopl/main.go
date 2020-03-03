// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
)

func main() {

	li := [...]int{1, 2, 3}
	fmt.Println(len(li))
	const (
		USD = iota
		CHN
		EUP
	)
	dict := [...]string{USD: "$", CHN: "￥", EUP: "€"} //数组
	fmt.Printf("dict type : %T", dict)
}
