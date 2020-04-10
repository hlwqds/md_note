# 标准I/O库

## 1 引言

标准库处理了很多细节，如缓冲区分配、以优化的块长度执行I/O等。这使得它便于用户使用，但是如果我们不深入了解I/O库函数的操作，也会带来一些问题。

## 2 流和FILE对象

在I/O库中，所有函数都是围绕文件描述符的。当打开一个文件时，即返回一个文件描述符，然后将文件描述符用于后续的I/O操作。而对于标准I/O库，他们的操作是围绕流进行的。当我们用标准I/O库打开或创建一个文件时，我们已使一个文件和流相关联。

对于ASCII字符集，一个字符用一个字节表示。对于国际字符集，一个字符可用于多个字节表示。标准I/O文件流可用于单字节或多字节字符集。流的定向决定了所读、写的字符是单字节还是多字节的。当一个流最初被创建时，它没有定向。如果在未定向的流上使用一个多字节I/O函数(见<wchar.h>)，则将该流的定向设置为宽定向的。如果在未定向的流上使用一个单字节I/O函数，则将流的定向设为字节定向的。只有两个函数可以改变流的定向。freopen函数清除一个流的定向；fwide函数可用于设置流的定向。

```c
#include <stdio.h>
#include <wchar.h>
extern int fwide (__FILE *__fp, int __mode) __THROW;
```

根据mode参数的不同值，fwide函数执行不同的工作。

- 如果mode参数值为负，fwide将试图使指定的流是字节定向的。
- 如果mode参数值为正，fwide将试图使指定的流是宽定向的。
- https://code.visualstudio.com/docs/cpp/config-mingw

- 

注意，fwide并不改变已定向流的定向。而且fwide无出错返回，所以我们唯一可靠的是，在调用fwide前先清除errno，返回时检查errno的值。

当打开一个流时，标准I/O函数fopen返回一个指向FILE对象的指针。该对象通常是一个结构，它包含了标准I/O库为管理该流需要的所有信息。

## 3 标准输入、输出和错误流

对一个进程预定义了3个流。他们是标准输入、输出和错误流，这些流引用的文件与文件描述符STDIN_FILENO，STDOUT_FILENO，STDERR_FILENO所引用的相同

## 4 缓冲

标准I/O库提供缓冲的目的是尽可能减少使用read和write的调用次数。它也对每个I/O流自动地进行缓冲管理。从而避免了应用程序需要考虑这一点所带来的麻烦。

标准I/O提供了以下的3种类型的缓冲。

1. 全缓冲。这种情况下，在填满标准I/O缓冲区后才进行实际I/O操作。对于驻留在磁盘上的文件通常是由标准I/O库实施全缓冲的。在一个流上执行第一次I/O操作时，相关标准函数通常使用malloc获得所需的缓冲区。

   术语冲洗(flush)说明标准I/O缓冲区的写操作。缓冲区可由标准I/O例程自动地冲洗(例如，当填满一个缓冲区时)，或者调用函数fflush冲洗一个流。

2. 行缓冲。在这种情况下，当在输入和输出中遇到换行符时，标准I/O库执行I/O操作，这允许我们一次输出一个字符(用标准函数fputc)，但只有在写了一行之后才进行实际I/O操作。当涉及一个终端时(如标准输出和标准输入)，通常使用行缓冲。

   对于行缓冲有两个限制。第一，因为标准I/O库用来收集每一行的缓冲区的长度是固定的，所以只要填满了缓冲区，那么即使还没有写一个换行符，也进行I/O操作。第二，任何时候只要通过标准I/O库要求从一个不带缓冲的流或者一个行缓冲的流中得到数据，那么就会冲洗所有行缓冲输出流。

3. 不带缓冲。标准I/O库不对字符进行缓冲处理。处理相当于文件I/O中的write函数

一般情况下，标准错误是不带缓冲的，若是指向终端设备的流，是行缓冲的，否则是无缓冲的。

对于一个给定的流，如果我们不喜欢系统默认的流，则可以通过下列函数中的一个更改其类型。

```c
#include <stdio.h>
extern void setbuf (FILE *__restrict __stream, char *__restrict __buf) __THROW;
extern int setvbuf (FILE *__restrict __stream, char *__restrict __buf,
		    int __modes, size_t __n) __THROW;
```

这些函数必须在流打开后调用，而且也应该在对该流执行任何一个其他操作执行调用。

可以使用setbuf函数打开或者关闭缓冲机制。为了带缓冲进行I/O，buf必须指向一个长度为BUFSIZ的缓冲区。为了关闭缓冲，将buf设置为NULL。

使用setvbuf函数，我们可以精确的说明所需的缓冲类型。这是通过mode参数说明的。

_IOFBF	全缓冲

_IOLBF	行缓冲

_IONBF	无缓冲

如果指定一个无缓冲的流，忽略buf和n参数。如果流是带缓冲的，而buf是NULL，则标准I/O将自动地为该流分配适当长度的缓冲区。适当长度指的是由常量BUFSIZ所指定的值。

要了解如果在一个函数内分配一个自动变量类的标准I/O缓冲区，则从该函数返回之前必须关闭该流。

在任何情况下，我们都可以强制冲洗一个流。

```c
#include <stdio.h>
extern int fflush (FILE *__stream);
```

此函数将使该流所有未写入的数据都被传送至内核。作为一种特殊情况，如果fp是NULL，函数将导致所有输出流都被冲洗。

## 5 打开流

下列3个函数打开一个标准I/O流。

```c
#inlcude <stdio.h>
extern FILE *fopen (const char *__restrict __filename,
		    const char *__restrict __modes) __wur;
extern FILE *freopen (const char *__restrict __filename,
		      const char *__restrict __modes,
		      FILE *__restrict __stream) __wur;
extern FILE *fdopen (int __fd, const char *__modes) __THROW __wur;
```

这3个函数的区别如下：

1. fopen函数打开路径名为pathname的一个指定的文件
2. freopen函数在一个指定的流上打开一个指定的文件，如果该流已经打开，则先关闭该流。如果该流已经定向，则使用freopen清除该定向。该函数一般用于讲一个指定的文件打开为一个预定义的流：标准输入，标准输出或者标准错误
3. fdopen函数取一个已有的文件描述符，并使一个标准的I/O流与该描述符相结合。此函数常用于由创建管道和网络通信通道函数返回的描述符。因为这些特殊类型的文件不能用标注I/O函数打开。

在以读和写类型打开一个文件时，具有下列限制：

- 如果中间没有fflush、fseek、fsetpo或者rewind，则在输出的后面不能直接跟随输入
- 如果中间没有fseek、fsetpos或者rewind，或者一个输入操作没有达到文件尾端，则在输入操作之后不能直接跟随输出

```c
#include <stdio.h>
extern int fclose (FILE *__stream);
```

调用fclose关闭一个打开的流

在该文件被关闭之前，冲洗缓冲区的输出数据。缓冲区 的任何输入数据被丢弃。如果标准I/O库已经为该流自动分配了一个缓冲区，则释放此缓冲区。

当一个进程正常终止时，则所有带未写缓冲数据的标准I/O流都被冲洗，所有打开的标准I/O流都被关闭。

## 6 读和写流

一旦打开了流，则可以在3个不同类型的非格式话I/O中进行选择，对其进行读写操作：

1. 每次一个字符的I/O。一次读取或写一个字符，如果流是带缓冲的，则标准I/O处理所有缓冲
2. 每次一行的I/O。使用fgets和fputs。当调用fgets时，应说明能处理的最大行长。
3. 直接I/O。fread和fwrite函数支持这种类型的I/O。每次I/O操作读或写某种数量的对象，而每个对象具有指定的长度。这两个函数常用于从二进制文件中每次读或写一个结构。

### 6.1 输入函数

以下3个函数可用于一次读一个字符。

```c
#include <stdio.h>
#define getc(_fp) _IO_getc (_fp)
extern int fgetc (FILE *__stream);
extern int getchar (void);
```

函数getchar等同于getc(stdin)。前两个函数的区别是，getc可被实现为宏，而fgetc不能实现为宏。这意味着：

1. getc的参数不应当时有副作用的表达式，因为它有可能被计算多次
2. 调用fgetc的时间可能比getc要长

这三个函数在返回下一个字符时，将其unsigned char 类型转换为int类型。这样就可以用-1表示错误了，而且可以表示所有的字符。当函数执行错误或者到达文件末尾时，会返回EOF=-1。

注意不管是出错还是到达文件尾端，这3个函数都返回同样的值。为了区分这两种不同的情况必须调用ferror和feof

```c
#include <stdio.h>
int __cdecl ferror(FILE *_File);
int __cdecl feof(FILE *_File);

void __cdecl clearerr(FILE *_File);
```

在大多数实现中，为每个流在FILE对象中维护了两个标志：

- 出错标志
- 文件结束标志

调用cleaerr可以清楚这两个标志。

从流中读取数据以后，还可以调用ungetc将字符再压回流中。

```c
#include <stdio.h>
int __cdecl ungetc(int _Ch,FILE *_File);
```

每次只支持压入一个字符，且不能将EOF压入。

当正在读取一个输入流，并进行某种形式的切词或记号切分操作时，会经常用到回送字符操作。有时需要先看一下下一个字符，以决定如何处理当前字符。然后就需要方便地将刚查看的字符回送，以便下一次调用getc时返回该字符。

### 6.2 输出函数

 对应于上面所述的每个输入函数都有一个输出函数。

```c

```

## 专有名词

流  		stream

流定向		stream's orientation

限制		restrict

