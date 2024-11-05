package main

import(
	"fmt"
)

func main(){
	inputChannel1 := make(chan int)
	inputChannel2 := make(chan int)
	exitChannel := make(chan int)

	sum := 0
	mul := 1

	go func(inputChannel1, inputChannel2, exitChannel chan int){
		for{
			select{
			case x := <- inputChannel1:
				sum += x
				mul *= x
			case y := <- inputChannel2:
				sum += y
				mul *= y
			case <-exitChannel:
				return
			}
		}
	}(inputChannel1, inputChannel2, exitChannel)

	inputNum1 := 0
	inputNum2 := 0
	fmt.Print("첫 번째 정수를 입력하세요 : ")
	fmt.Scan(&inputNum1)
	fmt.Print("두 번째 정수를 입력하세요 : ")
	fmt.Scan(&inputNum2)

	inputChannel1 <- inputNum1;
	inputChannel2 <- inputNum2;

	fmt.Println("덧셈 결과는 :", sum)
	fmt.Println("곱셈 결과는 :", mul)

	exitChannel <- 0
}