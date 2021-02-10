package main

// 3. c代码包含.c与.h文件，.h文件只包含声明，.c文件是实现；此方式需要引入hello.c文件

/*
// 注* 上面只能有一个星号
#include <stdio.h> // c
#include <stdlib.h>  // c
#include <stdbool.h> // 支持bool类型

#include "hello.c"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// 由于go test不支持cgo，这里还是传统方式测试
func TestPrintIn() { // 给c传入基础类型数据，进行打印
	name := C.CString("张三")
	defer C.free(unsafe.Pointer(name)) // C接口返回的指针对象，都需要释放，防止内存泄漏
	age := 23
	C.hello(name, C.int(age))
}

func main() {
	// 第一个举例，能调用此段代码输出Hello World
	C.sayHi() // 输出 hello world.
	// 第二个举例，测试string与int类型
	TestPrintIn() // 输出 Hello 张三, your age is 23.
	// 第三个举例，测试bool类型
	C.helloBool(C.bool(true)) // 输出 say 1.
	// 第四个举例，返回基础类型
	a := int(C.getInt())
	b := C.GoString(C.getString())
	c := bool(C.getBool())
	fmt.Println(a, b, c) // 1 ni hao true
}
