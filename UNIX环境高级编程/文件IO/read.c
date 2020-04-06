#include <apue.h>
#include <error.h>
#include <fcntl.h>

int main(void){
    int fd;
    if((fd = open("test.md", O_RDONLY)) < 0){
        err_sys("open error");
    }
    char buf[10];

    while(read(fd, buf, sizeof(buf)) > 0){
        printf("%s", buf);
    }
    exit(0);
}