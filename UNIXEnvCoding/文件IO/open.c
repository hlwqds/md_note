#include <apue.h>
#include <fcntl.h>
#include <error.h>
int main(){
    int file = open("record.md", O_RDONLY);
    printf("file: %d\n", file);

    int fileat = open(".", O_RDONLY);
    printf("file: %d\n", fileat);

    file = openat(fileat, "record.md", O_RDONLY);
    printf("file: %d\n", file);
    
    file = open("record.md", O_WRONLY | O_APPEND);
    printf("file: %d\n", file);

    file = open("record.md", O_RDONLY | O_CLOEXEC);
    printf("file: %d\n", file);

    file = open("test.md", O_WRONLY | O_CREAT);
    err_ret("open");
    printf("file: %d\n", file);

    file = open("test.md", O_RDONLY | O_DIRECTORY);
    printf("file: %d\n", file);

    file = open("atomic.md", O_RDONLY | O_CREAT | O_EXCL);
    printf("file: %d\n", file);
    err_ret("atomic");

    file = open("record.md", O_WRONLY | O_NOCTTY);
    printf("file: %d\n", file);

    file = open("record.md", O_WRONLY | O_NOFOLLOW);
    printf("file: %d\n", file);

    file = open("record.md", O_WRONLY | O_NONBLOCK);
    printf("file: %d\n", file);

    file = open("record.md", O_WRONLY | O_SYNC);
    printf("file: %d\n", file);

    file = open("record.md", O_WRONLY | O_TRUNC);
    printf("file: %d\n", file);

    file = open("record.md", O_WRONLY | O_DSYNC);
    printf("file: %d\n", file);

    file = open("record.md", O_WRONLY | O_RSYNC);
    printf("file: %d\n", file);

    file = openat(AT_FDCWD, "record.md", O_WRONLY | O_RSYNC);
    printf("file: %d\n", file);
}