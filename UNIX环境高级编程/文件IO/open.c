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

    file = open("test.md", O_RDONLY | O_DIRECTORY);
    printf("file: %d\n", file);
}