# TestCpp
接受如何通过golang调用c/c++的(部分c++接口还是无法使用的)
#### 概要
1. golang调用c/c++主要通过cgo, cgo原本只是为c设计的， 调用c代码，无需预编译，可以直接调用
2. 如果是调用c++, 需要编译成libxx.a的编译包，然后进行调用
#### 课程内容
1. 基础用法，在golang中直接调用c代码
2. 将c代码写到h文件中，在go文件中调用
3. c代码包含.c与.h文件，.h文件只包含声明，.c文件是实现；此方式需要引入hello.c文件
4. c代码包含.c与.h文件，.h文件只包含声明，.c文件是实现；这里我们开始通过libxx.a的方式进行引入
5. c代码包含.cpp与.h文件，.h文件只包含声明，.cpp文件是实现；这里只能通过libxx.a的方式进行引入
6. golang与c进行集合类型的交互
7. c代码与golang不在同一级代码，使用libxx.a的方式引用
8. golang与c进行struct类型的交互
9. golang调用c/c++的时候，c/c++有依赖lib(如果在c/c++中可以直接使用#pragma comment ( lib, "D3D11.lib"))
10. golang中通过使用dll中的方法
11. golang与c进行callback交互， 通过//export进行对应，可以同时多个export
12. c调用golang的函数，与callback完全是一回事

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

|C语言类型	|CGO类型	|Go语言类型|
|  ----  | ----  | ----  |
|int8_t	|C.int8_t	|int8|
|uint8_t	|C.uint8_t	|uint8|
|int16_t	|C.int16_t	|int16|
|uint16_t	|C.uint16_t	|uint16|
|int32_t	|C.int32_t	|int32|
|uint32_t	|C.uint32_t	|uint32|
|int64_t	|C.int64_t	|int64|
|uint64_t	|C.uint64_t	|uint64|

### 提交git
1. git init
2. git remote add origin https://github.com/zj360202/TestCpp.git
3. git pull origin master
4. git add src/main.js
5. git commit -m "first commit"
6. git push -u origin master