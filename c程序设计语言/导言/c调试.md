## linux环境下生成corefile文件

```bash
ulimit -c unlimited

#指定corefile生成目录
echo "/var/crash/coredump.%e.%p" > /proc/sys/kernel/core_pattern

#编译时指定-g调试
```

## gcc调试选项

https://blog.csdn.net/weixin_42615308/article/details/83151569

## valgrind检查运行时问题

https://blog.csdn.net/caijiwyj/article/details/99188644

## CMAKE

https://www.hahack.com/codes/cmake/

https://cmake.org/cmake/help/cmake2.4docs.html

