#ifndef HELLO_H_
#define HELLO_H_

#include <stdbool.h> // 支持bool类型  c语言的包，可以放这里，也可以放{}里面

#ifdef __cplusplus
extern "C" {
#endif

// 第一个举例，能调用此段代码输出Hello World
void sayHi();

// 第二个举例，测试string与int类型
void hello(char* name, int age);
// 第三个举例，测试bool类型
void helloBool(bool b);

// 第四个举例，测试bool类型
int getInt();
// 第四个举例，测试bool类型
char* getString();
// 第四个举例，测试bool类型
bool getBool();

#ifdef __cplusplus
}
#endif

#endif // HELLO_H_