package main

// 4. c代码包含.c与.h文件，.h文件只包含声明，.c文件是实现；这里我们开始通过libxx.a的方式进行引入

/*
// 生成命令 在.c文件目录下面执行
//gcc -Wall --machine-64 -c hello.c -o hello.o  -lwsock32
//ar -rv libhello.a hello.o
// 注* 上面只能有一个星号
// CFLAGS: -I.  将当前目录加入目录索引
#cgo CFLAGS: -I.
// LDFLAGS: -L. -lhello 在当前目录下面，检索libhello.a
#cgo LDFLAGS: -L. -lhello -ld3d11 -lstdc++
#include "hello.h"
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
