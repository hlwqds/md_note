#include <fcntl.h>
#include <apue.h>
#include <error.h>

int main(){
    int file = creat("test.md", 7);
    file = open("test.md", O_WRONLY | O_CREAT | O_TRUNC);
    file = open("test.md", O_RDWR | O_CREAT | O_TRUNC);

}