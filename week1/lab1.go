package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var operator string
	result := 0
	fmt.Println("첫 번째 숫자 입력: ")
	fmt.Scan(&a)
	fmt.Println("두 번째 숫자 입력: ")
	fmt.Scan(&b)
	fmt.Println("연산자 입력(+,-,*,/): ")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b != 0 {
			result = a / b
		} else {
			fmt.Println("나누는 숫자가 0입니다.")
			return
		}
	}
	fmt.Println(result)
}
