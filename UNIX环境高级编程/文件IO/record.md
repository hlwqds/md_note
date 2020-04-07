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

也可以调用creat函数创建一个新文件

```c
creat("test.md", 7);
//相当于
open("test.md", O_WRONLY | O_CREAT | O_TRUNC);
//如果要创建一个临时文件，并且要先写文件然后又读该文件，我们可以使用以下方法
open("test.md", O_RDWR | O_CREAT | O_TRUNC);
```

## 4 函数close

可调用close关闭一个打开的文件

```c
#include <fcntl.h>
#include <unistd.h>
int file = open("test.md", O_RDONLY | O_CREAT);
close(file);
```

关闭一个文件时还会释放该进程加在该文件上的所有记录锁。

当一个进程终止时，内核自动关闭它所有的打开文件

## 5 函数lseek

每个打开文件都有一个与之相关联的"当前文件偏移量"。它通常是一个非负整数，用以度量从文件开始计算的字节数。通常，读写操作都从当前文件偏移量处开始，并使偏移量增加所读写的字节数。按系统默认的情况，当打开一个文件时，除非制定O_APPEND选项，否则该偏移量被设置为0。

```c
#include <unistd.h>
extern __off_t lseek (int __fd, __off_t __offset, int __whence) __THROW;
```

对参数offset的解释与参数whence的值相关。

若whence是SEEK_SET，则将该文件的偏移量设置为距文件开始处offset个字节。

若whence是SEEK_CUR，则将该文件的偏移量设置为其当前值加offset，offset可为正或负。

若whence是SEEK_END，则将该文件的偏移量设置为文件长度加offset，offset可为正或负。

若lseek执行成功，则返回新的文件偏移量。

```c
#include <apue.h>
#include <unistd.h>
#include <fcntl.h>
int main(){
    int file = open("test.md", O_RDONLY);
    off_t currpos = lseek(file, 10, SEEK_SET);
    currpos = lseek(file, 20, SEEK_CUR);
    currpos = lseek(file, 20, SEEK_END);

    printf("currpos : %ld\n", currpos);
}
```

这种方法也可以用来确认所涉及的文件是否能设置偏移量。如果文件描述符指向的是一个管道、FIFO或网络套接字，则lseek返回-1，并将errno设置为ESPIPE

```c
#include <apue.h>
#include <fcntl.h>

int main(){
    if(lseek(STDIN_FILENO, 0, SEEK_CUR) == -1){
        printf("can't seek\n");
    }else{
        printf("seek OK\n");
    }

    exit(0);
}
```

```sh
./a.out  < /etc/passwd
cat < /etc/passwd | ./a.out
```

通常，文件的当前偏移量应当是一个非负整数，但是，某些设备也可能允许负的偏移量。但对于普通文件，其偏移量必须是非负值。所以在比较lseek的返回值时，我们最好将其与-1比较，而不是判断它是否小于0。

文件偏移量可以大于文件的当前长度，在这种情况下，对文件的下一次写将会加长该文件，并在文件中构成一个空洞，这一点是允许的。位于文件中但没有写过的字节都被读为0。空洞数据不占用磁盘空间

```c
#include <apue.h>
#include <fcntl.h>
#include <error.h>

char buf1[] = "abcdefghijk";
char buf2[] = "ABCDEFGHIJK";

int main(){
    int fd;
    if((fd = creat("file.hole", FILE_MODE)) < 0){
        err_sys("creat error");
    }

    if(write(fd, buf1, sizeof(buf1)) != sizeof(buf1)){
        err_sys("write error");
    }

    if(lseek(fd, 20, SEEK_CUR) == -1){
        err_sys("lseek error");
    }

    if(write(fd, buf2, sizeof(buf2)) != sizeof(buf2)){
        err_sys("write error");
    }
    exit(0);
}
```

```sh
od -c file.hole
```

## 6 函数read

```c
#include <unistd.h>
__fortify_function __wur ssize_t read (int __fd, void *__buf, size_t __nbytes)
```

如read成功，则返回读到的字节数。如果已经到达文件的尾端，则返回0。

有多种情况会使实际读到的字节数小于要求的字节数

- 读普通文件时，在读到要求字节数时已经到达了文件的尾端。
- 从终端设备读时，通常一次最多读一行
- 当从网络中读时，网络中的缓冲机制可能造成返回值小于要求读的字节数
- 当从管道或者FIFO读时，如果管道中包含的字节数小于所需的数量，那么read将只返回实际可用的字节数
- 当从某些面向记录的设备读时，一次最多返回一个记录
- 当一个信号造成中断，而已经读了部分数据量时

读操作从文件的当前偏移量处开始，在成功返回之前，该偏移量将增加实际读取的字节数

## 7 函数write

调用函数write向打开的文件写数据

```c
#include <unistd.h>
extern ssize_t write (int __fd, const void *__buf, size_t __n) __wur;
```

其返回值通常与参数__n相同，否则表示出错。write出错的一个常见原因是磁盘已经写满，或者超过了一个给定进程的文件长度限制。

## 8 I/O的效率

指定一个合理的缓冲区大小，会对性能有较高的提升。合理的大小和磁盘块长度有关。

大多数文件系统为改善性能都采用某种预读技术。当检测到正在顺序读取时，系统就试图读入比应用所要求的更多数据。

## 9 文件共享

内核使用3中数据结构表示打开文件，他们之间的关系决定了在文件共享方面一个进程对另一个进程可能产生的影响。

1. 每个进程在进程表中都有一个记录项，记录项中包含一张打开文件描述符表，可将其视为一个矢量，每个描述符占用一项。与每个描述符相关联的是
   1. 文件描述符标志
   2. 指向一个文件表项的指针
2. 内核为所有打开文件维持一张文件表。每个文件表项包含：
   1. 文件状态标志
   2. 当前文件偏移量
   3. 指向该文件v节点表项的指针
3. 每个打开文件都有一个v节点结构，v节点包含了文件类型和对此文件进行各种操作的函数指针。对于大多数文件，v节点还包含了该文件的i节点，i节点包含了打开文件时从磁盘中读取的文件的相关信息。

## 10 原子操作

当多个进程写同一个文件时，可能产生预想不到的结果。为了说明如何避免这种情况需要理解原子操作的概念。

### 10.1 追加到一个文件

考虑一个进程，它要将数据追加到一个文件尾端。早起的UNIX不支持open的O_APPEND选项，所以程序被编写成下列的形式。

```c
#include <apue.h>
#include <error.h>
#include <fcntl.h>

int main(){
    int fd = 0;
    char buf[] = "awdawd";
    fd = open("test.md", O_WRONLY);

    if (lseek(fd, 0, SEEK_END) == -1){
        err_sys("lseek error");
    }

    if(write(fd, buf, sizeof(buf)) != sizeof(buf)){
        err_sys("write error");
    }
}
```

对单个进程来说，这段程序能够正常工作，但如果有多个进程同时使用这个方法将数据追加到同一文件，则会产生问题

### 10.2 函数pread和pwrite

```c
#include <unistd.h>
__fortify_function __wur ssize_t pread (int __fd, void *__buf, size_t __nbytes, __off_t __offset)
extern ssize_t pwrite (int __fd, const void *__buf, size_t __n,
		       __off_t __offset) __wur;
```

1 **用于多线程磁盘文件操作**

2 APUE中的那段话所谓的“原子性”从语义上讲与O_APPEND模式的write相同。

3 用这两个调用模拟O_APPEND解决竞争是不可能的，必然导致新的竞争——你怎么知道文件现在的文件长度？

4 处理多线程磁盘文件操作时，不用这两个调用很麻烦。同情“很少用”的人。

原子性是指多步操作要么全部执行完毕，要么不会执行

### 10.3 创建一个文件

对open函数的O_CREAT和O_EXCL选项进行说明时，我们已见到另一个和原子操作有关的例子。当同时指定这两个选项，而文件又存在时，open将会失败。

## 11 函数dup和dup2

这两个函数都可以用来复制一个现有的文件描述符

```c
#include <unistd.h>
extern int dup (int __fd) __THROW __wur;
extern int dup2 (int __fd, int __fd2) __THROW;
```

对于dup返回的新文件描述符一定是当前可用文件描述符中的最小数值。对于dup2可以用fd2参数指定新描述符的值。如果fd2已经打开，则先将其关闭。如果fd等于fd2，则dup2返回fd2而不关闭它。

这些函数返回的新文件描述符与参数fd共享同一个文件表项

dup(fd)等效于fcntl(fd, F_DUPFD, 0);

dup2(fd, fd1)相当于close(fd1);fcntl(fd, F_DUPFD, fd2);但是又有区别：

dup2是一个原子操作，而close和fcntl包括两个函数调用。有可能在close和fcntl函数之间调用了信号处理函数，修改了文件描述符。如果多线程修改一个文件描述符也会出现问题。

dup2和fcntl有一些不同的errno

## 12 函数sync、fsync和fdatasync

传统的unix系统实现在内核中设有缓冲区高速缓存或页高速缓存，大多数磁盘I/O都通过缓冲区进行。当我们向文件写入数据时，内核通常先将数据复制到缓冲区中，然后排入队列，晚些时候再写入磁盘。这种方式被称为延迟写。

通常，当内核需要重用缓冲区来存放其他磁盘块数据时，它会把所有延迟写数据块写入磁盘。为了保证磁盘上实际文件系统与缓冲区中内容的一致性，UNIX系统提供了sync、fsync和fdatasync函数。



## 专业名词

current file offset 当前文件偏移量

read ahead 预读

delayed write	延迟写