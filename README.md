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
7. c代码与golang不在同一级代码，使用libxx.a的方式引用(未完)
8. golang与c进行struct类型的交互
9. golang调用c/c++的时候，c/c++有依赖lib(如果在c/c++中可以直接使用#pragma comment ( lib, "D3D11.lib"))(未完)
10. golang中通过使用dll中的方法(未完)
11. golang如果设定给c中的callback方法(未完)
12. c如何调用go函数(未完)

### 提交git
1. git init
2. git remote add origin https://github.com/zj360202/TestCpp.git
3. git pull origin master
4. git add src/main.js
5. git commit -m "first commit"
6. git push -u origin master