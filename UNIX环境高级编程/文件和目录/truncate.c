#include <apue.h>
#include <error.h>
#include <fcntl.h>
#define RWRWRW (S_IRUSR | S_IWUSR | S_IRGRP | S_IWGRP | S_IROTH | S_IWOTH)

int main(){
    int fd = 0;
    if((fd = open("test.md", O_RDWR | O_CREAT, RWRWRW)) < 0){
        err_sys("open error");
    }

    char buf[] = "huanglin test";
    if(write(fd, buf, sizeof(buf)) < 0){
        err_sys("write error");
    }

    ftruncate(fd, 5);

    close(fd);
}