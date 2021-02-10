package main

// 12. c调用golang的函数，与callback完全是一回事

/*
extern int goAdd(int, int);
static int cAdd(int a, int b) {
    return goAdd(a, b);
}
*/
import "C"
import "fmt"

//export goAdd
func goAdd(a, b C.int) C.int { // 函数名称与export后面的名称必须一样
	return a + b
}

func main() {
	a := C.cAdd(C.int(5), C.int(6))
	fmt.Println(int(a))
}
