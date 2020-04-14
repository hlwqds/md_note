#include <unistd.h>
#include <stdio.h>

int main(){
    printf("getpid %d\n", getpid());
    //获取进程ID
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