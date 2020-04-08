# 文件和目录

## 1 引言

上一章我们说明了执行I/O操作的基本函数，其中的讨论是围绕普通文件I/O进行的。本章将描述文件系统的其他特征和文件的性质。我们将从stat函数开始逐个说明stat结构的每一个成员以了解文件的所有属性。在此过程中我们将说明修改这些属性的各个函数，还将更详细的说明unix文件系统的结构以及符号链表本章最后介绍对目录进行操作的各个函数，并且开发了一个以降序遍历目录层次结构的函数

## 2 函数stat、fstat、fstatat和lstat

```c
#include <sys/stat.h>
extern int stat (const char *__restrict __file,
		 struct stat *__restrict __buf) __THROW __nonnull ((1, 2));
//一旦给出pathname，stat函数将返回与此命名文件相关的信息结构
extern int fstat (int __fd, struct stat *__buf) __THROW __nonnull ((2));
//fstat获得在描述符fd上打开的文件的相关信息
extern int fstatat (int __fd, const char *__restrict __file,
		    struct stat *__restrict __buf, int __flag);
//fstatat函数为一个相对与当前打开目录（由fd指定）的路径名返回文件统计信息。flag控制着是否跟随一个符号链接。当AT_SYMLINK_NOFOLLOW标志被设置时，fstatat不会跟随符号链接，而是返回符号链接本身的信息。否则将会返回符号链接指向的实际文件的信息。如果fd时AT_FDCWD，表现和fstat相同，如果path是绝对路径fd将被忽略。
extern int lstat (const char *__restrict __file,
		  struct stat *__restrict __buf) __THROW __nonnull ((1, 2));
//lstat与stat类似，但是对链接有特殊处理。如果是链接，返回的是链接的属性，而不是引用的文件的信息

struct stat
  {
    __dev_t st_dev;		/* Device.  */
#ifndef __x86_64__
    unsigned short int __pad1;
#endif
#if defined __x86_64__ || !defined __USE_FILE_OFFSET64
    __ino_t st_ino;		/* File serial number.	*/
#else
    __ino_t __st_ino;			/* 32bit file serial number.	*/
#endif
#ifndef __x86_64__
    __mode_t st_mode;			/* File mode.  */
    __nlink_t st_nlink;			/* Link count.  */
#else
    __nlink_t st_nlink;		/* Link count.  */
    __mode_t st_mode;		/* File mode.  */
#endif
    __uid_t st_uid;		/* User ID of the file's owner.	*/
    __gid_t st_gid;		/* Group ID of the file's group.*/
#ifdef __x86_64__
    int __pad0;
#endif
    __dev_t st_rdev;		/* Device number, if device.  */
#ifndef __x86_64__
    unsigned short int __pad2;
#endif
#if defined __x86_64__ || !defined __USE_FILE_OFFSET64
    __off_t st_size;			/* Size of file, in bytes.  */
#else
    __off64_t st_size;			/* Size of file, in bytes.  */
#endif
    __blksize_t st_blksize;	/* Optimal block size for I/O.  */
#if defined __x86_64__  || !defined __USE_FILE_OFFSET64
    __blkcnt_t st_blocks;		/* Number 512-byte blocks allocated. */
#else
    __blkcnt64_t st_blocks;		/* Number 512-byte blocks allocated. */
#endif
#ifdef __USE_XOPEN2K8
    /* Nanosecond resolution timestamps are stored in a format
       equivalent to 'struct timespec'.  This is the type used
       whenever possible but the Unix namespace rules do not allow the
       identifier 'timespec' to appear in the <sys/stat.h> header.
       Therefore we have to handle the use of this header in strictly
       standard-compliant sources special.  */
    struct timespec st_atim;		/* Time of last access.  */
    struct timespec st_mtim;		/* Time of last modification.  */
    struct timespec st_ctim;		/* Time of last status change.  */
# define st_atime st_atim.tv_sec	/* Backward compatibility.  */
# define st_mtime st_mtim.tv_sec
# define st_ctime st_ctim.tv_sec
#else
    __time_t st_atime;			/* Time of last access.  */
    __syscall_ulong_t st_atimensec;	/* Nscecs of last access.  */
    __time_t st_mtime;			/* Time of last modification.  */
    __syscall_ulong_t st_mtimensec;	/* Nsecs of last modification.  */
    __time_t st_ctime;			/* Time of last status change.  */
    __syscall_ulong_t st_ctimensec;	/* Nsecs of last status change.  */
#endif
#ifdef __x86_64__
    __syscall_slong_t __glibc_reserved[3];
#else
# ifndef __USE_FILE_OFFSET64
    unsigned long int __glibc_reserved4;
    unsigned long int __glibc_reserved5;
# else
    __ino64_t st_ino;			/* File serial number.	*/
# endif
#endif
  };
```

使用stat函数最多的地方可能就是ls -l命令，用其可以获得有关一个文件的所有信息。

## 3 文件类型

至此我们已经介绍了两种不同的文件类型：普通文件和目录。UNIX系统的大多数文件时普通文件和目录，但是也有一些其他文件类型。

1. 普通文件(regular file)

这是最常用的文件类型，这种文件包含了某种形式的数据。

2. 目录文件(directory file)

这种文件包含了其他文件的名字以及指向这些文件相关信息的指针。对一个目录文件具有读权限的任一进程都可以读取该目录的内容，但只有内核可以直接写目录文件。进程必须使用对应的函数才能修改目录。

3. 块特殊文件(block special file)

这种类型的文件提供对设备带缓冲的访问，每次访问以固定长度为单位进行。

4. 字符特殊文件(character special file)

这种类型的文件提供对设备不带缓冲的访问，每次访问长度可变。系统中的所有设备要么是字符特殊文件，要么是块特殊文件。

5. FIFO

这种类型的文件用于进程间通信，有时也被称为命名管道(name pipe)

6. 套接字(socket)

这种类型的文件用于进程间的网络通信。套接字也可用于在一台宿主机上进程之间的非网络通信。

7. 符号链接(symbolic link)

这种类型的文件指向另一个文件。

我们可以通过以下的宏定义来确定文件类型

```c
#define	S_ISDIR(mode)	 __S_ISTYPE((mode), __S_IFDIR)
#define	S_ISCHR(mode)	 __S_ISTYPE((mode), __S_IFCHR)
#define	S_ISBLK(mode)	 __S_ISTYPE((mode), __S_IFBLK)
#define	S_ISREG(mode)	 __S_ISTYPE((mode), __S_IFREG)
#ifdef __S_IFIFO
# define S_ISFIFO(mode)	 __S_ISTYPE((mode), __S_IFIFO)
#endif
#ifdef __S_IFLNK
# define S_ISLNK(mode)	 __S_ISTYPE((mode), __S_IFLNK)
#endif

#if defined __USE_MISC && !defined __S_IFLNK
# define S_ISLNK(mode)  0
#endif

#if (defined __USE_XOPEN_EXTENDED || defined __USE_XOPEN2K) \
    && defined __S_IFSOCK
# define S_ISSOCK(mode) __S_ISTYPE((mode), __S_IFSOCK)
#elif defined __USE_XOPEN2K
# define S_ISSOCK(mode) 0
#endif

//POSIX.1允许将进程间通信(IPC)对象说明为文件
# define S_TYPEISMQ(buf) __S_TYPEISMQ(buf)
# define S_TYPEISSEM(buf) __S_TYPEISSEM(buf)
# define S_TYPEISSHM(buf) __S_TYPEISSHM(buf)
```

实例

```c
#include <apue.h>
#include <error.h>
#include <fcntl.h>
int main(int argc, char **argv){
    int i;
    struct stat buf;
    char *ptr;
    for(i = 1; i < argc; i++){
        printf("%s: \n", argv[i]);
        if(lstat(argv[i], &buf) < 0){
            err_ret("lstat error");
            continue;
        }
        if(S_ISREG(buf.st_mode)){
            ptr = "regular";
        }else if(S_ISDIR(buf.st_mode)){
            ptr = "directory";
        }else if(S_ISCHR(buf.st_mode)){
            ptr = "character special";
        }else if(S_ISFIFO(buf.st_mode)){
            ptr = "fifo";
        }else if(S_ISBLK(buf.st_mode)){
            ptr = "block special";
        }else if(S_ISLNK(buf.st_mode)){
            ptr = "symbolic link";
        }else if(S_ISSOCK(buf.st_mode)){
            ptr = "socket";
        }else{
            ptr = "unknown mode";
        }
        printf("%s\n",ptr);
    }
    exit(0);
}
```

## 4 设置用户ID和设置组ID

与一个进程相关联的ID有6个或者更多

| 实际用户ID<br />实际组ID               | 我们实际上是谁       |
| :------------------------------------- | -------------------- |
| 有效用户ID<br />有效组ID<br />附属组ID | 用于文件访问权限检查 |
| 保存的设置用户ID<br />保存的设置组ID   | 由exec函数保存       |

- 实际用户ID和实际组ID标识我们究竟是谁。这两个字段在登录时取自口令文件中的登录项。通常，在一个登录回话期间这些值不会改变，但是超级用户进程有方法改变他们。
- 有效用户ID、有效组ID以及附属组ID决定了我们的文件访问权限
- 保存的设置用户ID组ID在执行一个程序时包含了有效用户ID和有效组ID的副本

通常，有效用户ID等于实际用户ID，有效组ID等于实际组ID

每个文件有一个所有者和一个组所有者，所有者由stat结构中的st_uid指定，组所有者ID由stat结构中的st_gid指定

st_mode中有两个标志位，可以将进程的有效用户ID或者组ID替换为文件所有者的ID。分别称为替换用户ID位和替换组ID位

## 5 文件访问权限

st_mode值也包含了对文件的访问权。所有文件都有访问权限。

每个文件有9个访问权限位

| st_mode屏蔽                       | 含义                             |
| --------------------------------- | -------------------------------- |
| S_IRUSR<br />S_IWUSR<br />S_IXUSR | 用户读<br />用户写<br />用户执行 |
| S_IRGRP<br />S_IWGRP<br />S_IXGRP | 组读<br />组写<br />组执行       |
| S_IROTH<br />S_IWOTH<br />S_IXOTH | 其他读<br />其他写<br />其他执行 |

图中的3类访问权限以各种方式由不同的函数使用。我们将这些不同的方式汇总在下面。

- 第一个规则是，我们用名字打开任一类型的文件时，对该名字中包含的每一个目录，包括它可能隐藏的当前工作目录都应具有执行权限。
- 对于一个文件的读权限决定了我们是否可以打开现有文件进行读操作，这与O_RDONLY和O_RDWR标志有关
- 对于一个文件的写权限决定了我们是否可以打开现有文件进行写操作，这与O_WRONLY和O_RDWR标志有关
- 为了在open函数中对一个文件指定O_TRUNC标志，必须对该文件具有写权限
- 为了在一个目录中新建一个文件，必须对该目录具有写权限和执行权限
- 为了删除一个现有文件，必须对该文件的目录具有写权限和执行权限，对该文件则不需要写权限
- 如果用7个exec函数中的任何一个执行某个文件，都必须对该文件具有执行权限。该文件还必须是一个普通文件

进程每次打开、创建或删除一个文件时，内核就进行文件访问权限测试，而这种测试可能涉及文件的所有者、进程的有效ID以及进程的附属组ID。两个文件所有者ID是文件的性质，有效ID和附属组ID则是进程的性质。内核进行的测试具体如下：

1. 若进程的有效ID是0，则允许访问。这给予了超级用户对整个文件系统进行处理的最充分的自由
2. 若进程的有效用户ID等于文件的所有者ID，那么如果所有者适当的访问权限被置位，则允许访问，否则拒绝访问
3. 如果进程的有效组ID等于文件的所有者组ID，那么如果所有者适当的访问权限被置位，则允许访问，否则拒绝访问
4. 若其他用户的适当访问权限被置位，则允许访问，否则拒绝访问

按顺序执行

## 6 新文件和目录的所有权

新文件的用户ID设置为进程的有效用户ID

关于组ID

1. 新文件的组ID可以是进程的有效组ID
2. 新文件的组ID可以是它所在目录的组ID

对于linux，新文件的组ID取决于它所在的目录的设置组ID位是否被设置。如果这一位已经被设置，则新文件组ID设置为目录ID，否则设置为进程的有效组ID

## 7 函数access和faccessat

access和faccessat是根据实际用户ID进行权限测试的函数

```c
#include <unistd.h>
extern int access (const char *__name, int __type) __THROW __nonnull ((1));
extern int faccessat (int __fd, const char *__file, int __type, int __flag) __THROW __nonnull ((2)) __wur;
```

其中，如果测试文件是否已经存在，mode就为F_OK；否则mode是权限常量的按位与

```c
#include <apue.h>
#include <error.h>
#include <fcntl.h>

int main(int argc, char **argv){
    if(argc != 2){
        err_quit("argc error");
    }

    if(access(argv[1], R_OK) < 0){
        err_ret("access error for %s", argv[1]);
    }else{
        printf("read access OK\n");
    }

    if(open(argv[1], O_RDONLY) < 0){
        err_ret("open error for %s", argv[1]);
    }else{
        printf("open %s OK\n", argv[1]);
    }
    exit(0);
}
```

```sh
./a.out /etc/shadow
sudo chown root a.out
sudo chmod u+s a.out
./a.out /etc/shadow
```

## 8 函数umask

umask函数为进程设置文件模式创建屏蔽字，并返回之前的值。

```c
#include <sys/stat.h>
extern __mode_t umask (__mode_t __mask) __THROW;
```

cmask由9个权限位按位或构成。

在进程创建一个新文件时，就一定会使用文件模式创建屏蔽字。比如open和creat函数，通过参数mode创建屏蔽字。在文件模式创建屏蔽字中为1的为，在文件mode中对应位一定会被关闭。

```c
#include <apue.h>
#include <fcntl.h>
#include <error.h>

#define RWRWRW (S_IRUSR | S_IWUSR | S_IRGRP | S_IWGRP | S_IROTH | S_IWOTH)

int main(int argc, char **argv){
    umask(0);

    open("apple1", O_WRONLY | O_CREAT);

    if (creat("foo", RWRWRW) < 0){
        err_sys("creat foo error");
    }
    umask(S_IRGRP | S_IWGRP | S_IROTH | S_IWOTH);
    if(creat("bar", RWRWRW) < 0){
        err_sys("creat bar error");
    }
    open("apple2", O_WRONLY | O_CREAT);

    exit(0);
}
```

更改进程的文件模式创建屏蔽字并不影响其父进程的屏蔽字。

## 9 函数chmod、fchmod和fchmodat

函数chmod、fchmod和fchmodat这3个函数使我们可以更改现有文件的访问权限

```c
#include <sys/stat.h>
extern int chmod (const char *__file, __mode_t __mode)
     __THROW __nonnull ((1));
extern int fchmod (int __fd, __mode_t __mode) __THROW;
extern int fchmodat (int __fd, const char *__file, __mode_t __mode,
		     int __flag)
     __THROW __nonnull ((2)) __wur;
```

chmod函数可以让我们更改现有文件的权限。三者之间的区别上面已有类似的说明。

为了改变一个文件的权限位，进程的有效用户ID必须等于文件的所有者ID，或者该进程必须具有超级用户权限

```c
# define S_IRUSR	__S_IREAD       /* Read by owner.  */
# define S_IWUSR	__S_IWRITE      /* Write by owner.  */
# define S_IXUSR	__S_IEXEC       /* Execute by owner.  */
/* Read, write, and execute by owner.  */
# define S_IRWXU	(__S_IREAD|__S_IWRITE|__S_IEXEC)

# define S_IRGRP	(S_IRUSR >> 3)  /* Read by group.  */
# define S_IWGRP	(S_IWUSR >> 3)  /* Write by group.  */
# define S_IXGRP	(S_IXUSR >> 3)  /* Execute by group.  */
/* Read, write, and execute by group.  */
# define S_IRWXG	(S_IRWXU >> 3)

# define S_IROTH	(S_IRGRP >> 3)  /* Read by others.  */
# define S_IWOTH	(S_IWGRP >> 3)  /* Write by others.  */
# define S_IXOTH	(S_IXGRP >> 3)  /* Execute by others.  */
/* Read, write, and execute by others.  */
# define S_IRWXO	(S_IRWXG >> 3)
```

```c
#include <apue.h>
#include <fcntl.h>
#include <error.h>

int main(){
    struct stat statbuf;
    if(stat("foo", &statbuf) < 0){
        err_sys("stat error for foo");
    }

    if(chmod("foo", (statbuf.st_mode & ~S_IXGRP) | S_ISGID) < 0){
        err_sys("chmod error for foo");
    }
	//组id位，对应文件权限中组执行位的S。如果表示进程执行此文件时，有效用户为该文件的所有者。同理用户id位
    
    if(chmod("bar", S_IRUSR | S_IWUSR |  S_IRGRP | S_IROTH) < 0){
        err_sys("chmod error for bar");
    }

    exit(0);
}
```

## 10 函数chown、fchown、fchownat和lchown

下面几个chown函数可用于更改文件的用户ID和组ID。如果两个参数owner或group中任意一个是-1，则对应的ID不变。

除了引用的文件时符号链接外，这4个函数的操作类似。在符号链接的情况下，lchown和fchownat(设置了AT_SYMLINK_NOFOLLOW)更改符号链接本身的所有者，而不改变链接指向的文件的所有者。

fchown函数改变fd参数指向的打开文件的所有者，所以不可用来改变链接符号本身的所有者。

若_POSIX_CHOWN_RESTRICTED对指定的文件生效。则：

1. 只有超级用户进程能改变该文件的用户ID
2. 如果进程拥有此文件（其有效ID等于该文件的用户ID），参数等于-1或者文件的用户ID，并且参数group等于进程的有效组ID或进程的附属组ID之一，那么一个非超级用户可以更改该文件的组ID。

这意味这，当_POSIX_CHOWN_RESTRICTED有效时，不能更改其他用户文件的用户ID。你可以更改你所拥有的文件的组ID，但只能改到你所在的组。

## 11 文件长度

stat结构成员st_size表示以字节为单位的文件的长度。此字段只对普通文件、目录文件和符号链接有意义。

对于普通文件，其文件长度可以是0，在开始读这种文件时，将得到文件结束指示。对于目录，文件长度通常是一个数的整数倍。

对于符号链接，文件长度是在文件名中的实际字节数。即它指向的文件的名称的总字节数。（注意，因为符号链接文件长度总是由st_size指示，所以它并不包含通常C语言用作名字结尾的NULL字节）

现今，大多数现代的UNIX系统提供字段st_blksize和st_blocks。其中第一个是对文件I/O较合适的块长度，第二个是所分配的实际512字节块块数。

## 12 文件截断

有时我们需要在文件尾端截取一些数据以缩短文件。讲一个文件的长度截断为0是一个特例，在打开文件时使用O_TRUC标志可以做到这一点。为了截断文件可以调用函数truncate和ftruncate。

```c
#include <unistd.h>
extern int ftruncate (int __fd, __off_t __length) __THROW __wur;
extern int truncate (const char *__file, __off_t __length)
     __THROW __nonnull ((1)) __wur;
```

## 13 文件系统

- 每个i节点中都有一个链接计数，其值是指向该i节点的目录项数。只有当链接计数减少至0时，才可以删除该文件。这就是为什么“接触对一个文件的链接”操作并不总意味着释放该文件占用的磁盘块的原因。这也是为什么删除一个目录项的函数被称为unlink而不是delete的原因。在stat结构中，连接计数包含在st_nlink成员中。这种链接称为硬链接。LINK_MAX指定了一个文件链接数的最大值。硬链接多个最常见的是目录..和.
- 另外一种链接类型称为符号链接。符号链接文件的实际内容包含了该符号链接所指向的文件的名字。
- i节点中包含了文件有关的所有信息：文件类型、文件访问权限、文件长度和指向文件数据块的指针。stat结构中的大多数数据取自i节点。只有两项重要数据存放在目录项中：文件名和文件i节点编号。
- 因为目录项中的i节点编号指向同一文件系统中的相应i节点，一个目录项不能指向另一个文件系统的i节点。这就是为什么ln命令不能跨越文件系统的原因
- 当在不更新文件系统的情况下为一个文件重命名时，该文件的实际内容并没有移动，只需构造一个指向现有i节点的新目录项，并删除老的目录项。链接计数不会改变

## 14 函数link、linkat、unlink、unlinkat和remove

创建一个指向现有文件的链接的方法是使用link函数或linkat函数。

```c
#include <unistd.h>
extern int link (const char *__from, const char *__to)
     __THROW __nonnull ((1, 2)) __wur;
extern int linkat (int __fromfd, const char *__from, int __tofd,
		   const char *__to, int __flags)
     __THROW __nonnull ((2, 4)) __wur;
```

这两个函数创建一个新的目录项 to，它引用现有的文件 from。如果to已经存在，则返回出错。只创建to中的最后一个文件，这意味之前的路径应该是真实存在的。



linkat比较灵活，可以通过fd和path的组合确定是相对路径还是绝对路径，是相对fd还是相对工作目录。

flags可以在为符号链接创建链接时。可以通过设置flagAT_SYMLINK_FOLLOW确定是指向符号链接本身，还是指向符号链接指向的文件。

创建新目录项和增加链接计数应当是一个原子操作。

很多操作系统不允许对目录的硬链接，因为这可能造成文件系统中形成循环



为了删除一个现有的目录项，可以调用unlink函数。

```c
#include <unistd.h>
extern int unlink (const char *__name) __THROW __nonnull ((1));
extern int unlinkat (int __fd, const char *__name, int __flag)
     __THROW __nonnull ((2));
```

这两个函数删除目录项，并将pathname所引用文件的链接计数减一。如果对该文件还有其他链接，则仍可以通过其他链接访问该文件的数据。如果出错，不对文件进行任何修改。

只有当链接计数达到0时，该文件的内容才可以被删除。另一个条件也会阻止文件的删除--只要有进程打开了该文件，其内容也不能删除。关闭一个文件时，内核首先检查打开该文件的进程个数；如果这个计数达到0，内核再去检查其链接计数，如果也是0，则删除文件内容。

unlinkat可以通过指定flag参数AT_REMOVEDIR来删除目录，类似于rmdir

unlink这种特性经常被程序用来确保即使是在程序崩溃时，它所创建的临时文件也不会遗留下来。进程用open创建一个文件，然后立即调用unlink。只有当进程关闭该文件或终止时，该文件内容才会被删除。

如果是符号链接，unlink只会删除符号链接的数据，而不会删除符号链接指向的文件数据



也可以使用remove删除文件，对于文件效果和unlink相同，对于目录，和rmdir相同。

## 特有名词

普通文件		regular file

目录文件		directory file

块特殊文件		block special file

字符特殊文件		character special file

命名管道			name pipe

文件结束		end-of-file

符号链接		symbolic link