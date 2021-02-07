#include <stdio.h> // c
#include <stdlib.h>  // c
//#include <iostream> // c++ 这个使用不了
//#include <Windows.h>
#include <d3d11.h>
//#include <dxgi1_2.h>
#include "hello.h"
// 这样的模式是不支持的，可以在系统中找到对应的lib包(此处是d3d11.lib)，然后复制到libhello.a的同级目录,并改名libd3d11.a，在go中通过#cgo LDFLAGS: -L./test -lhello -ld3d11 即可
//#pragma comment ( lib, "D3D11.lib")

int hello(char *name, int age) {
    printf("Hello %s, your age is %d.\n", name, age);
    return age;
}