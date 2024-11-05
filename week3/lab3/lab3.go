package main

// #cgo LDFLAGS: -L. -lcfibonacci
// int fibonacci(int);
import "C"
import "fmt"

func main(){
    userValue := 0
    result := 0
    fmt.Print("사용자 정수 입력 : ")
    fmt.Scan(&userValue)
    result = int(C.fibonacci(C.int(userValue)))
    fmt.Println(userValue, "번째 피보나치 수 : ", result)
}