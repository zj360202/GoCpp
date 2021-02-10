#include "hello.h"
#include <cstdio> // c++    // c++中的很多包，需要用#cgo CFLAGS: -Iinclude目录  但是c++包引用的c++类的写法可能不满足c语言的风格，导致引入失败，所以尽量使用c包或封装c包的c++包

void sayHi() {	// 第一个举例，能调用此段代码输出Hello World
	printf("hello world.\n");
}
void hello(char* name, int age) { // 第二个举例，测试string与int类型
	printf("Hello %s, your age is %d.\n", name, age);
}
void helloBool(bool b) { // 第三个举例，测试bool类型
	printf("say %d.\n", b);
}
int getInt() { // 第四个举例，测试bool类型
	return 1;
}
char* getString() { // 第四个举例，测试bool类型
	return (char*)("ni hao");       // c++中，不允许常量字符串直接转换为char*，常量字符串的默认类型是const char* ，需要显示转换
}
bool getBool() { // 第四个举例，测试bool类型
	return true;
}