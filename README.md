# TestCpp

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