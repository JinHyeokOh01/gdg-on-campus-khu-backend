package main

import (
	"fmt"
)

func main() {
	first := 0
	second := 0
	result := 0
	operator := ""
	fmt.Println("첫 번째 숫자 입력: ")
	fmt.Scan(&first)
	fmt.Println("두 번째 숫자 입력: ")
	fmt.Scan(&second)
	fmt.Println("연산자 입력(+,-,*,/)")
	fmt.Scan(&operator)

	if second == 0 {
		fmt.Println("나누는 숫자가 0입니다.")
		return
	}
	switch operator {
	case "+":
		result = first + second
	case "-":
		result = first - second
	case "*":
		result = first * second
	case "/":
		result = first / second
	default:
		fmt.Println("연산자 입력이 잘못되었습니다.")
	}
	fmt.Println(result)
}
