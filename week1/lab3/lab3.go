package main

import (
	"fmt"
)

func SumArray(arr []int) int {
	sum := 0
	for i := range len(arr) {
		sum += arr[i]
	}
	return sum
}

func findMaxMin(arr []int) (int, int) {
	maxNum := arr[0]
	minNum := arr[0]

	for i := 1; i < len(arr); i++ {
		if maxNum < arr[i] {
			maxNum = arr[i]
		}
		if minNum > arr[i] {
			minNum = arr[i]
		}
	}
	return maxNum, minNum
}

func main() {
	defer fmt.Println("프로그램이 종료되었습니다.")
	arr := []int{3, 5, 1, 2, 0}
	fmt.Println(arr)
	sum := SumArray(arr)
	fmt.Println("배열의 총합:", sum)
	maxNum := 0
	minNum := 0
	maxNum, minNum = findMaxMin(arr)
	fmt.Println("최대값:", maxNum)
	fmt.Println("최소값:", minNum)

	switch {
	case len(arr) < 3:
		fmt.Println("배열의 길이가 짧습니다.")
	case len(arr) == 3:
		fmt.Println("배열의 길이가 적당합니다.")
	case len(arr) > 3:
		fmt.Println("배열의 길이가 깁니다.")
	default:
	}
}
