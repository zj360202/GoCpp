package main

// 9. golang调用c/c++的时候，c/c++有依赖lib(如果在c/c++中可以直接使用#pragma comment ( lib, "D3D11.lib"))
//    解决办法: 如果是windows，则在系统中找到Windows Kits，这个目录中，主要存放了系统中x86与x64的包，一般系统lib都在这下面，可以将系统包复制到当前目录，然后改名字为libxxx.a，使用cgo就可以引用成功
//    此处不在演示

/*
#include <stdio.h>
#include <stdlib.h>
#cgo CFLAGS: -I.
// 即可引用成功libxxx.a
#cgo LDFLAGS: -L. -lxxx
*/
import "C"

func main() {

}
