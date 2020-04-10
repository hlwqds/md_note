#include <apue.h>
#include <error.h>
#include <fcntl.h>

int main(){
    char buf[] = "12345";
    int fd = open("test.md", O_WRONLY | O_APPEND);
    if(write(fd, buf, sizeof(buf)) < 0){
        err_sys("write error");
    }

    exit(0);
}