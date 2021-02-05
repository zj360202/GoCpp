#include <stdio.h> // c
#include <stdlib.h>  // c
//#include <iostream> // c++ 这个使用不了
//#include <Windows.h>
#include <d3d11.h>
//#include <dxgi1_2.h>
#include "hello.h"
#pragma comment ( lib, "D3D11.lib")

int hello(char *name, int age) {
    printf("Hello %s, your age is %d.\n", name, age);
    return age;
}