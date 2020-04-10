#include <apue.h>
#include <unistd.h>
#include <fcntl.h>
int main(){
    int file = open("test.md", O_RDONLY);
    off_t currpos = lseek(file, 10, SEEK_SET);
    currpos = lseek(file, 20, SEEK_CUR);
    currpos = lseek(file, 20, SEEK_END);

    printf("currpos : %ld\n", currpos);
}