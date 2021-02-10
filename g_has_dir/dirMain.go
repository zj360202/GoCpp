package main

// 7. c代码与golang不在同一级代码，使用libxx.a的方式引用

/*
#include <stdio.h>
#include <stdlib.h>
#include "dir_c/hello.h"
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L./dir_c -lhello
*/
import "C"
import (
	"fmt"
	"log"
	"unsafe"
)

func main() {
	name := C.CString("jack")
	defer C.free(unsafe.Pointer(name))
	age := C.int(28)
	fmt.Println(name)
	result := C.hello(name, age)
	log.Println("cgo----------->age %d ", result)
}
