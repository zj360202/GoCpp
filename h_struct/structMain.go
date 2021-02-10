package main

// 8. golang与c进行struct类型的交互

/*
// 注* 上面只能有一个星号
//第一种，直接在代码上面引入c代码，
#include <stdio.h> // c
#include <stdlib.h>  // c
#include <string.h>  // memcpy
#include <stdbool.h> // 支持bool类型

//typedef struct{ // 此种声明方式在golang中调用，无法找到
//    int dataLen;
//    unsigned char* data;
//} Test;
typedef struct Test{ // 如果是在golang中使用，则必须在前面也要声明
    int dataLen;
    unsigned char* data;
} Test;

Test* initStruct(){
    return (Test *)malloc(sizeof(Test));
}

void structIn(Test* testPtr) {	// 第一个举例，传入struct
    //Test* testPtr = (Test*)stIn;
	printf("dataLen: %d\n", testPtr->dataLen);
    for (int i=0; i<testPtr->dataLen; i++){
        printf("data: k: %d v: %d\n", i, testPtr->data[i]);
    }
}
Test* structOut() { // 第二个举例，返回struct
    Test* test = (Test *)malloc(sizeof(Test));
    memset(test, 0 , sizeof(test));
    test->dataLen = 4*sizeof(unsigned char*);
    test->data = (unsigned char*)malloc(test->dataLen);
    short a = 6;
	for (int i=0; i< 4; i++){
        test->data[i] = (unsigned char)(i+4);
        printf("data: k: %d v: %d\n", i, test->data[i]);
    }
    return test;
}
void structInOut(Test** stInOut) { // 第三个举例，直接对象修改
	Test* test = (Test *)malloc(sizeof(Test));
    memset(test, 0 , sizeof(test));
    test->dataLen = 5*sizeof(unsigned char*);
    test->data = (unsigned char*)malloc(test->dataLen);
    short a = 6;
	for (int i=0; i< 5; i++){
        test->data[i] = (unsigned char)(i+2);
    }
    *stInOut = test;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// 第一个举例，传入struct
	a := C.initStruct() // 如果struct包含指针或struct，则当前struct必须在c中初始化，不然在golang中初始化的field将完全或部分是go Memory的，在传给c使用会包cgo argument has Go pointer to Go pointer
	d := make([]int8, 4)
	a.dataLen = C.int(4)
	for i := 0; i < int(a.dataLen); i++ {
		d[i] = int8(i) + 3
	}
	//dPtr := uintptr(unsafe.Pointer(&d[0]))
	a.data = (*C.uchar)(unsafe.Pointer(&d[0]))
	C.structIn(a)

	// 第二个举例，返回struct
	b := C.structOut()
	fmt.Println("dataLen:", b.dataLen/8)
	gslice := make([]int8, b.dataLen/8)
	C.memcpy(unsafe.Pointer(&gslice[0]), unsafe.Pointer(b.data), C.size_t(b.dataLen/8))
	fmt.Println("gslice:", gslice)

	// 第三个举例，直接对象修改
	c := C.initStruct()
	C.structInOut((**C.struct_Test)(unsafe.Pointer(&c))) // CGO: 在c中定义的struct在golang中引用的方式是struct_xxx
	fmt.Println("dataLen1:", c.dataLen/8)
	gslice = make([]int8, c.dataLen/8)
	C.memcpy(unsafe.Pointer(&gslice[0]), unsafe.Pointer(c.data), C.size_t(c.dataLen/8))
	fmt.Println("gslice1:", gslice)
	/** 执行结果
	dataLen: 4
	gslice: [4 5 6 7]
	dataLen1: 5
	gslice1: [2 3 4 5 6]
	dataLen: 4
	data: k: 0 v: 3
	data: k: 1 v: 4
	data: k: 2 v: 5
	data: k: 3 v: 6
	data: k: 0 v: 4
	data: k: 1 v: 5
	data: k: 2 v: 6
	data: k: 3 v: 7
	*/
}
