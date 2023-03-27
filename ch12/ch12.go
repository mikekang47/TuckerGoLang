package main

import "fmt"

func main() {
	// var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// float의 경우 따로 지정하지 않으면 최소 소숫점 자리로 표시됨.
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(arr[i])
	// }

	// // var temps [5]float64 = [5]float64{24.3, 26.7}
	// var s = [5]int{1: 14, 3: 30}
	// fmt.Println(s[1])

	// x := [...]int{10, 20, 30}
	// fmt.Println(x[0])

	// 배열 크기는 무조건 상수
	// const X int = 10
	// var arr [X]int
	// arr[0] = 10

	// fmt.Println(arr[0])

	// 	배열 복사
	// a := [5]int{1, 2, 3, 4, 5}
	// b := [5]int{500, 400, 300, 200, 100}

	// for i, v := range a {
	// 	fmt.Printf("a[%d] = %d\n", i, v)
	// }
	// fmt.Println()

	// for i, v := range b {
	// 	fmt.Printf("b[%d] = %d\n", i, v)
	// }
	// b = a // 바로 깊은 복사 가능.
	// fmt.Println()
	// for i, v := range b {
	// 	fmt.Printf("b[%d] = %d\n", i, v)
	// }

	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9},
	}
	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

}
