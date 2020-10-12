package main

import (
	"fmt"
	"math"
)

func main() {

	const PI float32 = 3.1415926
	const PI2 = 3.1415926
	var count int
	println(count)
	var count2 int = 10
	println(count2)
	count3 := 10
	println(count3)

	//布尔型
	var b1, b2 bool
	b1 = true
	b2 = false
	fmt.Println(b1)
	fmt.Println(b2)

	//整形
	var sum int
	var a int = 100
	b := 28
	sum = a + b
	println(sum)

	var a2 int8 = 127
	fmt.Println(a2, a2+1)

	//整形数据溢出检查
	fmt.Printf("Min(int8) = %v, Max(int8)= %v\n", math.MaxInt8, math.MaxInt8)
	fmt.Printf("Min(int16) = %v, Max(int16)= %v\n", math.MaxInt16, math.MaxInt16)
	fmt.Printf("Min(int32) = %v, Max(int32)= %v\n", math.MaxInt32, math.MaxInt32)

	//浮点型数据
	var f1 float64 = 1.23
	f2 := 12.3
	fmt.Println(f2 - f1)

	//浮点数的舍入误差
	var f11, f12 float32
	f11 = 12356.789e5
	f12 = f11 + 20
	fmt.Println(f12)

}
