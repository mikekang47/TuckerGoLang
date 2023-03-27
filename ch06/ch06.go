package main

import "fmt"

func main() {
	var x = 7
	var y = 3

	fmt.Println("x + y = ", x+y)

	var s float32 = 3.14
	var t float32 = 5

	fmt.Println("s * t = ", s*t)

	// 비트 클리어 연산자
	var a uint8 = 10
	var b uint8 = 2

	fmt.Println("a &^ b = ", a&^b)

}
