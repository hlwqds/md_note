#include <apue.h>
#include <fcntl.h>
#include <error.h>

#define RWRWRW (S_IRUSR | S_IWUSR | S_IRGRP | S_IWGRP | S_IROTH | S_IWOTH)

int main(int argc, char **argv){
    umask(0);

    open("apple1", O_WRONLY | O_CREAT);

    if (creat("foo", RWRWRW) < 0){
        err_sys("creat foo error");
    }
    umask(S_IRGRP | S_IWGRP | S_IROTH | S_IWOTH);
    if(creat("bar", RWRWRW) < 0){
        err_sys("creat bar error");
    }
    open("apple2", O_WRONLY | O_CREAT);

    exit(0);
}