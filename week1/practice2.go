package main

import(
	"fmt"
)

func main(){
	var a int
	fmt.Print("숫자를 입력하세요: ")
	fmt.Scan(&a)
	if((a % 2) == 1){
		fmt.Print("홀수입니다.")
	} else{
		fmt.Print("짝수입니다.")
	}
}