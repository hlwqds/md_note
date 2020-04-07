#include <apue.h>
#include <error.h>
#include <fcntl.h>

int main(){
    int fd = 0;
    fd = open("test.md", O_RDONLY);

    int fd1 = dup(fd);
    int fd2 = dup2(fd1, fd1);
    fd1 = dup(fd);

    printf("fd1: %d\n", fd1);
    printf("fd2: %d\n", fd2);
    fcntl(fd, F_DUPFD, 0);
}