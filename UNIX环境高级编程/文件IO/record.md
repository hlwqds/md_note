# 文件I/O

## 1 文件描述符

对于内核而言，所有打开的文件都通过文件描述符引用。文件描述符是一个非负整数。当打开一个现有文件或者创建一个新文件时，内核向进程返回一个文件描述符。当读写一个文件时，使用open或creat返回的文件描述符标示该文件，将其作为参数传递给read或write。

```c
#define  STDIN_FILENO   0  /* Standard input.  */

#define  STDOUT_FILENO  1  /* Standard output.  */

#define  STDERR_FILENO  2  /* Standard error output.  */
```

依照惯例0,1,2分别表示标准输入，标准输出和标准错误。

文件描述符的变化范围是0～OPEN_MAX -1，即_SC_OPEN_MAX。对于linux来说，文件描述符的变化范围几乎是无限的，只会受制于系统配置，字长限制等。

## 2 函数open和openat

```c
#include <fcntl.h>

__fortify_function int open (const char *__path, int __oflag, ...)；
__fortify_function int openat (int __fd, const char *__path, int __oflag, ...)；
```

调用open或openat函数可以打开或创建一个文件。

我们将最后一个参数写为...，ISO C用这种方法表明剩下的参数的数量及其类型是可变的。对于open函数而言，仅当创建新文件时才使用最后这个参数。

path参数是要打开或创建文件的名字。oflag参数可用来说明此函数的多个选项。用下列一个或多个常量进行“或”运算构成oflag参数。

```c
#include <fcntl.h>

open("./", O_RDONLY);		//只读打开
open("./", O_WRONLY);		//只写打开
open("./", O_RDWR);			//读写打开
//这几个常量中必须且只能指定一个。下面的常量则是可选的

open("record.md", O_WRONLY | O_APPEND);	//每次写时都追加到文件的尾端
open("record.md", O_WRONLY | O_CLOEXEC);	//把FD_CLOEXEC常量设置为文件描述符标志
open("test.md", O_WRONLY | O_CREAT);		//若此文件不存在则创建他。使用此选项时，open函数应该同时说明第三个参数mode，用mode制定该新文件的访问权限位
open("test.md", O_WRONLY | O_DIRECTORY);	//如果文件不是目录类型，则出错

```

