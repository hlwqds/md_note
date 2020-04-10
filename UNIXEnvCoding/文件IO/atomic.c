#include <apue.h>
#include <error.h>
#include <fcntl.h>

int main(){
    int fd = 0;
    char buf[] = "pwrite";
    fd = open("test.md", O_WRONLY);

    if (lseek(fd, 0, SEEK_END) == -1){
        err_sys("lseek error");
    }

    if(write(fd, buf, sizeof(buf)) != sizeof(buf)){
        err_sys("write error");
    }

    int seek = 0;
    if((seek = lseek(fd, 0, SEEK_END)) == -1){
        err_sys("lseek error");
    }
    printf("seek: %d\n", seek);
    if(pwrite(fd, buf, sizeof(buf), seek) != sizeof(buf)){
        err_sys("pwrite error");
    }

    if(pwrite(fd, buf, sizeof(buf), seek) != sizeof(buf)){
        err_sys("pwrite error");
    }

    if((seek = lseek(fd, 0, SEEK_END)) == -1){
        err_sys("lseek error");
    }
    printf("seek: %d\n", seek);
    pread(fd, buf, sizeof(buf), 0);
    exit(0);
}