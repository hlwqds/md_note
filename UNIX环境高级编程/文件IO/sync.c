#include <apue.h>
#include <fcntl.h>
int main(){
    int fd = open("test.md", O_WRONLY | O_APPEND);
    char buf[] = "syncfd";
    write(fd, buf, sizeof(buf));
    sync();
    fsync(fd);
    fdatasync(fd);
}