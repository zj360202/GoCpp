package main

// 6. golang与c进行集合类型的交互

/*
// 注* 上面只能有一个星号
//第一种，直接在代码上面引入c代码，
#include <stdio.h> // c
#include <stdlib.h>  // c
#include <string.h>  // memcpy
#include <stdbool.h> // 支持bool类型

void arrayIn(int** data, int lenth) {	// 第一个举例，传入slice, 打印
    int* dataArray = (int*)data;
	for (int i=0; i< lenth; i++){
		printf("data: %d %d\n", i, dataArray[i]);
	}
}
void mapIn(int** keys, bool** values, int lenth) { // 第二个举例，传入map, 打印
    int* km = (int*)keys;
    bool* vm = (bool*)values;
	for (int i=0; i< lenth; i++){
        printf("data: k: %d v: %d\n", km[i], vm[i]);
    }
}
int* arrayOut(int lenth) { // 第三个举例，返回slice
	int* ao = (int*)malloc(lenth*sizeof(int));
    for (int i=0; i< lenth; i++){
  		ao[i] = i+1;
        printf("data: index: %d v: %d intLen: %d, addr: %d\n", i, i+1, sizeof(int), &ao[i]);
    }
    return ao;
}
int mapOut() { // 第四个举例，返回map
	return 1;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// 第一个举例，传入slice, 打印
	ar := make([]int32, 8)
	for i := 0; i < len(ar); i++ {
		ar[i] = int32(i) + 8
	}
	fmt.Println("ar:", ar)
	C.arrayIn((**C.int)(unsafe.Pointer(&ar[0])), C.int(len(ar)))

	// 第二个举例，传入map, 打印
	mr := make(map[int32]bool, 8)
	for i := 0; i < len(ar); i++ {
		if i%2 == 0 {
			mr[int32(i)+2] = true
		} else {
			mr[int32(i)+2] = false
		}
	}
	ka, va := make([]int32, 0), make([]bool, 0)
	for k, v := range mr {
		ka = append(ka, k)
		va = append(va, v)
	}
	fmt.Println("ka:", ka)
	fmt.Println("va:", va)
	C.mapIn((**C.int)(unsafe.Pointer(&ka[0])), (**C.bool)(unsafe.Pointer(&va[0])), C.int(len(mr)))
	// 第三个举例，返回slice
	lenth := 4
	lo := C.arrayOut(C.int(lenth))
	fmt.Println("lo:", lo)
	defer C.free(unsafe.Pointer(lo))
	gslice := make([]int32, lenth)
	C.memcpy(unsafe.Pointer(&gslice[0]), unsafe.Pointer(lo), C.size_t(lenth*4))
	fmt.Println("gslice:", gslice)
	// 第四个举例，返回map，由于c无默认的map，这里不在说明，要使用c的map，可以考虑下面的项目
	// https://github.com/Broadroad/map , 将对应文件复制到这个项目中，并引用即可，对于struct的使用，会在后面的内容说明，这里不在讨论，作为后续讲解内容的作业
	a := C.mapOut()
	fmt.Println(a)
}
