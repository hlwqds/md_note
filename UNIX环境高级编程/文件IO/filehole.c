#include <apue.h>
#include <fcntl.h>
#include <error.h>

char buf1[] = "abcdefghijk";
char buf2[] = "ABCDEFGHIJK";

int main(){
    int fd;
    if((fd = creat("file.hole", FILE_MODE)) < 0){
        err_sys("creat error");
    }

    if(write(fd, buf1, sizeof(buf1)) != sizeof(buf1)){
        err_sys("write error");
    }

    if(lseek(fd, 20, SEEK_CUR) == -1){
        err_sys("lseek error");
    }

    if(write(fd, buf2, sizeof(buf2)) != sizeof(buf2)){
        err_sys("write error");
    }
    exit(0);
}