#include <error.h>
#include <apue.h>
#include <fcntl.h>

int main(){
    int file = open("test.md", O_RDONLY | O_CREAT);
    close(file);
}