package main

import (
	"day1/util"
	"fmt"
	"runtime"
	"strconv"
)

func atc() {
	var a float32 = 8
	var b float32 = 3
	var c float32 = a + b
	var d float32 = a - b
	var e float32 = a * b
	var f float32 = a / b

	var m int = 8
	var n int = 3
	var g int = m % n

	fmt.Printf("c=%f\n", c) //11
	fmt.Println(d)          //5
	fmt.Println(e)          //24
	fmt.Println(f)          //2.66
	fmt.Println(g)          //2
}

func rel() {
	var a float32 = 8
	var b float32 = 3
	var c float32 = 5
	fmt.Println(a > b && b > c)
	fmt.Println(a > b || b > c)
	fmt.Println(!(b > c))
}

func bit_op() {
	fmt.Printf("os arh %s, int size %d\n", runtime.GOARCH, strconv.IntSize)
	fmt.Printf("260  %d\n", 260) // 十进制
	fmt.Printf("260  %b\n", 260) // 二进制
	fmt.Printf("260  %s\n", util.BinaryFormart(260))
	fmt.Printf("-260  %s\n", util.BinaryFormart(-260))
	fmt.Printf("-260>>23  %s\n", util.BinaryFormart(-260>>23))
	fmt.Printf("260>>23  %s\n", util.BinaryFormart(260>>23))
	fmt.Printf("%b\n", 6&5)  // 110 & 101 = 100
	fmt.Printf("%b\n", 6|5)  // 110 | 101 = 111
	fmt.Printf("%b\n", 6^5)  // 110 ^ 101 = 011
	fmt.Printf("%b\n", 5>>1) // 101 >> 1 = 10
	fmt.Printf("%b\n", 5<<1) // 101 << 1 = 1010
}

func test1() {
	var a int = 7 // 7=4+2+1 111>>1 = 11
	var b int = a / 2
	var c int = a >> 1 // 3
	fmt.Println(b, c)
}

func test2() {
	a := 5 //5
	a += 1 // 5+1=6
	b := 7 //7
	b *= a //b=b*a=7*6=42
	b ^= a // b=b^a
	fmt.Println(a, b)
}

func main() {
	// atc()
	// rel()
	// bit_op()
	// test1()
	test2()
}
