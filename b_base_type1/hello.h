#ifndef HELLO_H_
#define HELLO_H_

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
	return "ni hao";
}
bool getBool() { // 第四个举例，测试bool类型
	return true;
}
 
#endif // HELLO_H_