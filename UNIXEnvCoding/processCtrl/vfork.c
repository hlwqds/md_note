#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
int globvar = 7;

int main(){
    int var;

    pid_t pid;

    var = 99;

    printf("before vfork\n");
    if((pid = vfork()) < 0){
        printf("vfork error\n");
        return 0;
    }else if(pid == 0){
        printf("this is parent pid: %d\n", getpid());
        var++;
        globvar++;
        //在调用exec或者exit之前，子进程在父进程的空间内运行。
        printf("var: %d, globvar: %d\n", var, globvar);
        sleep(3);
        fclose(stdout);
        //模拟exit关闭标准I/O流的实现
        exit(0);
        //如果在exit的实现中关闭了标准I/O流，那么父进程将不会打印数据
    }else{
        printf("this is parent pid: %d\n", getpid());
        printf("sub pid:%d\n", pid);
        printf("var: %d, globvar: %d\n", var, globvar);
    }
    return 0;
}