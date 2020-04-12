# 进程环境

## 1 引言

下一章将介绍进程控制原语，在此之前需先了解进程的环境。

## 2 main函数

C程序总是从main函数开始执行。main函数的原型是:

```c
int main(int argc, char **argv)；
```

其中argc是命令行参数的数目，argv是指向命令行参数列表的指针。

当内核执行C程序时，在调用main前先调用一个特殊的启动例程。可执行程序文件将此启动例程指定为程序的起始地址，这是有连接编辑器设置的，而连接编辑器则有c编译器调用。启动例程从内核取得命令行参数和环境变量，然后为按上述方式调用main函数做好安排。

## 3 进程终止

有8中方式使得进程终止，其中5中为正常终止

1. 从main返回
2. 调用exit
3. 调用_exit 或 _Exit
4. 最后一个线程从其启动例程返回
5. 从最后一个线程调用pthread_exit

异常终止有3中方式

6. 调用abort
7. 接到一个信号
8. 最后一个线程对取消请求做出响应

线程相关的内容暂且不谈。

上面提到的启动例程是这样编写的，使得从main返回后立即调用exit函数。如果将启动例程以C代码形式表示，则它调用main函数的形式可能是

exit(main(argc, argv))

### 3.1 退出函数

3个函数用于正常终止一个程序：_exit 和 _Exit立即进入内核，exit则先执行一些清理操作，然后返回内核。

```c
#include <stdlib.h>
extern void _Exit (int __status) __THROW __attribute__ ((__noreturn__));
extern void exit (int __status) __THROW __attribute__ ((__noreturn__));

#include <unistd.h>
extern void _exit (int __status) __attribute__ ((__noreturn__));
```

由于历史原因，exit函数总是执行一个标准I/O库的清理关闭操作。这导致输出缓冲区的所有数据都会被冲洗到文件上。

### 3.2 函数atexit

按照ISO C的规定，一个进程可以登记多至32个函数，这些函数将由exit自动调用。我们称这些函数为终止处理程序，并调用atexit函数来登记他们

```c
extern int atexit (void (*__func) (void)) __THROW __nonnull ((1));
```

其中，atexit的参数是一个函数地址，当调用此函数时无需向他传递任何参数，也不期待它返回一个值。exit调用这些函数的顺序和他们登记时候的顺序相反。同一个函数如果登记多次，也会被调用多次。



## 专有名词

终止 		termination