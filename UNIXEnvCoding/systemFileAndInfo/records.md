# 系统数据文件和信息

## 1 引言

UNIX系统的正常运行需要使用大量与系统有关的数据文件。例如口令文件/etc/passwd和组文件/etc/group

由于历史原因，这些数据文件都是ASCII文本文件，并且使用标准I/O库读这些文件。

## 2 口令文件

UNIX系统口令文件中的信息包含在<pwd.h>中定义的passwd结构中

```c
struct passwd
{
  char *pw_name;		/* Username.  */
  char *pw_passwd;		/* Password.  */
  __uid_t pw_uid;		/* User ID.  */
  __gid_t pw_gid;		/* Group ID.  */
  char *pw_gecos;		/* Real name.  */
  char *pw_dir;			/* Home directory.  */
  char *pw_shell;		/* Shell program.  */
};
```

由于历史原因，口令文件时/etc/password，而且是一个ASCII文件。格式如下：

```c
root:x:0:0:root:/root:/bin/bash
daemon:x:1:1:daemon:/usr/sbin:/usr/sbin/nologin
bin:x:2:2:bin:/bin:/usr/sbin/nologin
sys:x:3:3:sys:/dev:/usr/sbin/nologin
sync:x:4:65534:sync:/bin:/bin/sync
```

关于这些登录项，请注意下列各点：

## 下面的不想再记了，很有可能用不到，看一遍就够了

