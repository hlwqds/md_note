#include <fcntl.h>
#include <apue.h>

int main(){
    int fd = open("test.md", O_RDONLY);

    int fd1 = fcntl(fd, F_DUPFD, 0);
}