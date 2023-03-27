package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력하세요.")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자를 입력하세요.")

			stdin.ReadString('\n') // input의 delimeter를 받아서 , delim이 나올 때까지만 처리함.
			continue
		}
		fmt.Printf("%d\n", number)
		if number%2 == 0 {
			break
		}
	}
	fmt.Println("for문이 종료되었습니다.")
}
