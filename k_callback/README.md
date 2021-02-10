# k_callback
#### 概要
1. golang与c进行callback交互， 通过//export进行对应，可以同时多个export
#### c中几个常见问题
1. go中包含的c内容，只能使用// 或 /* .. \*/ (第一个斜杠后面只能有一个*号，否则会语法错误)
2. import “C” 与注释的c代码之间不能有空格
3. 无bool类型，c++中有，如果c中需要使用，则需要引入#include <stdbool.h>包 
4. C接口返回的指针对象，都需要释放，防止内存泄漏
5. CGO后面不能使用注释(//); 否则会出现：invalid flag in #cgo CFLAGS: //
6. c++中的很多包，需要用#cgo CFLAGS: -Iinclude目录  但是c++包引用的c++类的写法可能不满足c语言的风格，导致引入失败，所以尽量使用c包或封装c包的c++包
7. 类型上面c与golang的int是相互对应的，不过如果对于数组或其他有序集合来说，golang的int是64位(与系统相关)， c的int是32位的，64位的是long; 所以如果golang是slice的话用int32与c的int进行匹配
8. c中的struct定义，必须保证typedef struct 后面也需要跟着名称 typedef struct Test{..}Test;
9. 如果struct包含指针或struct，则当前struct必须在c中初始化，不然在golang中初始化的field将完全或部分是go Memory的，在传给c使用会包cgo argument has Go pointer to Go pointer
10. CGO: 在c中定义的struct在golang中引用的方式是struct_xxx
11. golang中定义的方法，要在c中被使用，就需要//export
12. go中可以同时声明多一个//export
13. go中//export后面的名称与go的函数名称必须一样
14. 使用回调函数的地方，必须使用static
15. go中被//export标记的函数，输入和输出都是cgo类型(C.int, C.bool, C.struct_xx)

#### C、CGO、golang 基础数据对应关系
|  C   | CGO  | golang |
|  ----  | ----  | ----  |
|char |  C.char |  byte |
|signed char |  C.schar |  int8 |
|unsigned char |  C.uchar |  uint8 |
|short int |  C.short |  int16 |
|short unsigned int |  C.ushort |  uint16 |
|int |  C.int |  int |
|unsigned int |  C.uint |  uint32 |
|long int |  C.long |  int32 or int64 |
|long unsigned int |  C.ulong |  uint32 or uint64 |
|long long int |  C.longlong |  int64 |
|long long unsigned int |  C.ulonglong |  uint64 |
|float |  C.float |  float32 |
|double |  C.double |  float64 |
|wchar_t |  C.wchar_t  |  |
|char* |    | string |
|bool |  C.bool  | bool |
|void * | |unsafe.Pointer |
- 注*： string(golang) --> char*(c)  C.CString("张三")
-   char*(c) --> string(golang)  C.GoString(C.getString()) 详情看代码baseMain.go

```
// 脚步在根目录执行，无需到test目录下执行，不然无法调用目录中的内容
gcc -Wall -c test/hello.cpp -o test/hello.o  -lstdc++
ar -rv test/libhello.a test/hello.o

调用的时候
/*
#include <stdio.h>
#include <stdlib.h>
#include "test/hello.h"     // 在创建的时候，带目录，此处应用就带目录
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L./test -lhello   // -L接目录 -l后面接libxxxx.a后面的名字
*/
import "C"
```

#### 常见问题
1. pragma comment []
```
//这样的模式是不支持的，可以在系统中找到对应的lib包(此处是d3d11.lib)，然后复制到libhello.a的同级目录,并改名libd3d11.a，在go中通过#cgo LDFLAGS: -L./test -lhello -ld3d11 即可
```

### 提交git
1. git init
2. git remote add origin https://github.com/zj360202/TestCpp.git
3. git pull origin master
4. git add src/main.js
5. git commit -m "first commit"
6. git push -u origin master