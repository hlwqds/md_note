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
open("atomic.md", O_RDONLY | O_CREAT | O_EXCL);	//O_EXCL如果同时指定了O_CREAT，而文件已经存在，则出错。用此可以测试一个文件是否存在，如果不存在，则创建此文件，这使测试和创建是一个原子操作
open("record.md", O_WRONLY | O_NOCTTY);		//如果引用的是终端设备，则不将该设备分配作为此进程的控制终端
open("record.md", O_WRONLY | O_NOFOLLOW);	//如果引用的是一个符号链接，则出错
open("record.md", O_WRONLY | O_NONBLOCK);	//如果引用的是一个FIFO、一个块特殊文件或一个字符特殊文件，则此选项为文件的本次打开操作和后续的I/O操作设置非阻塞模式
open("record.md", O_WRONLY | O_SYNC);		//使每次write等待物理I/O完成，包括由该write操作引起的文件属性更新所需的I/O
open("record.md", O_WRONLY | O_TRUNC);		//如果此文件存在，而且为只写或读写成功打开，则将其长度截为0
open("record.md", O_WRONLY | O_DSYNC);		//使每次write要等待物理I/O操作完成，但是如果该写操作不影响读取刚写入的数据，则不需要等待文件属性被更新
open("record.md", O_WRONLY | O_RSYNC);		//使每一个以文件描述符作为参数进行的read操作等待，直至所有对文件同一部分挂起的写操作完成。
```

由open和openat函数返回的文件描述符一定是最小的未用描述符数值。这一点被某些应用程序用来在标准输入，标准输出或标准错误上打开新的文件。例如可以先关闭标准错误，再打开另一个文件，该文件就能在文件描述符2上打开。



fd参数将open和openat区分开，共有三种可能：

1、path参数指定的是绝对路径，这种情况下，fd参数被忽略，openat就相当于open

2、path指定的是相对路径，fd参数指定了相对路径名再文件系统中的起始位置。fd参数是通过打开相对路径名所在的目录获取的。

3、path参数指定了相对路径名，fd参数具有特殊值AT_FDCWD。这种情况下，路径名在当前工作目录中获取，openat和open函数再操作上类似。

## 3 函数creat



