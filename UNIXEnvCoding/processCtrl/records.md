# 进程控制

## 1 引言

本章介绍UNIX系统的进程控制，包括创建新进程、执行程序和进程终止。还将说明进程属性的各种ID，以及他们如何受到进程控制原语的影响。

## 2 进程标识

每个进程都有一个非负整数表示的唯一进程ID。

虽然是惟一的，但是进程ID是可复用的。当一个进程终止后，其进程ID就成为复用的候选。大多数UNIX系统实现了延迟复用算法，使得赋予新建进程的ID不同于最近终止进程所使用的的ID。

系统中有一些专用进程，但具体细节随实现而不同。

ID为0的进程通常是调度进程，常常被称为交换进程(swapper)。该进程是内核的一部分，它并不执行任何磁盘上的程序，因此也被称为系统进程。

ID为1的进程通常是init进程，在自举过程结束后由内核调用。该进程的程序文件存在于/sbin/init。此进程负责在自举内核后启动一个UNIX系统。init通常读取与系统有关的初始化文件，并将系统引导到一个状态。init进程绝不会终止。他是一个普通用户进程，但是以超级用户特权运行。

除了进程ID，每个进程还有一些其他标识符

```c
#include <unistd.h>
#include <stdio.h>
int main(){
    printf("getpid %d\n", getpid());
    //获取进程IDc
    printf("getppid %d\n", getppid());
    //获取父进程ID

    printf("getuid %d\n", getuid());
    //获取进程实际用户ID
    printf("geteuid %d\n", geteuid());
    //获取进程有效用户ID

    printf("getgid %d\n", getgid());
    //获取进程实际组ID
    printf("getegid %d\n", getegid());
    //获取进程有效组ID

    return 0;
}
```

## 3 函数fork

一个现有进程可以调用fork函数创建一个新进程。

```c
#include <unistd.h>
pid_t fork(void);
返回值：子进程返回0， 父进程返回子进程ID；如果出错，返回-1
```

由fork创建的新进程被称为子进程。fork函数被调用一次，但返回两次。两次返回的区别是子进程的返回值是0，而父进程的返回值则是新建子进程的进程ID。

子进程和父进程继续执行fork调用之后的命令。子进程是父进程的副本。父进程和子进程会共享正文段。

```c
#include <unistd.h>
#include <stdio.h>

int globvar = 6;
char buf[] = "a write to stdout\n";

int main(){
    int var;
    pid_t pid;

    var = 88;
    if(write(STDERR_FILENO, buf, sizeof(buf) - 1) != sizeof(buf) - 1){
        printf("write error!\n");
        return 0;
    }

    printf("before fork!\n");

    if((pid = fork()) < 0){
        printf("fork error!\n");
    }else if(pid == 0){
        globvar++;
        var++;
    }else{
        sleep(2);
    }

    printf("pid = %ld, glob = %d, var = %d\n", (long)getpid(), globvar, var);
    return 0;
}
```

一般来说，在fork之后是父进程先执行还是子进程先执行是不确定的，这取决于内核所使用的调度算法。在交互式运行程序时，只得到before fork!\n一次，其原因是标准输出缓冲区由换行符冲洗。但当标准输出重定向到一个文件时，却得到before fork!\n两次，这是因为当调用fork时，before fork!\n仍然在缓冲区中，子进程同样将其拷贝。在exit之前的第二个printf将其数据追加到已有的缓冲区中，当进程终止时，其缓冲区的内容都被写到响应的文件中。

