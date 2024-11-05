package main

// #cgo LDFLAGS: -L. -lcstring_concat
// #include <stdio.h>
// #include <stdlib.h>
// char* string_concat(const char*, const char*);
import "C"
import(
	"fmt"
	"unsafe"
)
func main() {
    string1 := "I want to go Home"
    string2 := ", really!"

    cString1 := C.CString(string1)
    cString2 := C.CString(string2)
    defer C.free(unsafe.Pointer(cString1)) 
    defer C.free(unsafe.Pointer(cString2))

    result := C.string_concat(cString1, cString2)
    defer C.free(unsafe.Pointer(result))

    fmt.Println("Concatenated string:", C.GoString(result))
}