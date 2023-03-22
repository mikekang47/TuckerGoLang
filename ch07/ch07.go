package main

// // 반환 타입, 매개변수 타입 반드시 존재 해야함
// // 첫 글자가 대문자면 Public

// func Add(a int, b int) int {
// 	return a + b
// }

// func main() {
// 	c := Add(3, 6) // 함수 수행이 끝나면 함수 명령 포인터가 이 위치로 다시 돌아옴. 그래서 값을 c에 복사.
// 	fmt.Println(c)
// }

// func getAveragePoint(math int, eng int, history int) float64 {
// 	return (float64)(math+eng+history) / 3
// }

// func printPersonalPoint(name string, point float64) {
// 	fmt.Println(name, "님 평균 점수는", point, "입니다.")
// }

// func main() {
// 	slice := []string{"김일등", "송이등", "박삼등"}
// 	m := make(map[string][]int)
// 	m["김일등"] = []int{80, 74, 95}
// 	m["송이등"] = []int{88, 92, 53}
// 	m["박삼등"] = []int{78, 73, 78}

// 	for i := 0; i < len(slice); i++ {
// 		printPersonalPoint(slice[i], getAveragePoint(m[slice[i]][0], m[slice[i]][1], m[slice[i]][2]))
// 	}
// }

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}

// 연습문제
func Multiple(a, b int) int {
	return a * b
}
