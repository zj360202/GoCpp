package main

// 11. golang与c进行callback交互， 通过//export进行对应，可以同时多个export

/*
extern void go_callback_int(int);
extern void go_callback_obj(void*);
static void CallMyFunction1() {
    go_callback_int(5);
}
static void CallMyFunction2(void* foo) {
    go_callback_obj(foo);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export go_callback_int
func go_callback_int(p1 C.int) { // 函数名称与export后面的名称必须一样
	fmt.Println("go callback int:", p1)
}

//export go_callback_obj
func go_callback_obj(pTest unsafe.Pointer) {
	test := (*Test)(pTest)
	test.cb()
}

type Test struct {
}

func (s *Test) cb() {
	fmt.Println("Test callback with")
}

func main() {
	C.CallMyFunction1()
	data := &Test{}
	C.CallMyFunction2(unsafe.Pointer(&data))
}
