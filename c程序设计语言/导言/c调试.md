## linux环境下生成corefile文件

```bash
ulimit -c unlimited

#指定corefile生成目录
echo "/var/crash/coredump.%e.%p" > /proc/sys/kernel/core_pattern

#编译时指定-g调试
```



## valgrind检查运行时问题

https://blog.csdn.net/caijiwyj/article/details/99188644