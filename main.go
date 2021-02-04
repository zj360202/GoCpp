package main

import (
	"fmt"
	"log"
	"unsafe"
)

/*
#include <stdio.h>
#include <stdlib.h>
#include "test/hello.h"
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L./test -lhello
*/
import "C"

func main() {
	name := C.CString("jack")
	defer C.free(unsafe.Pointer(name))
	age := C.int(28)
	fmt.Println(name)
	result := C.hello(name, age)
	log.Println("cgo----------->age %d ", result)
}
