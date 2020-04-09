#include <apue.h>
#include <error.h>
#include <fcntl.h>

int main(int argc, char **argv){
    int i, fd;
    struct stat statbuf;
    struct timespec times[2];

    for(i = 1; i < argc; i++){
        if(stat(argv[i], &statbuf) < 0){
            err_sys("stat error");
        }
        if((fd = open(argv[i], O_RDWR | O_TRUNC)) < 0){
            err_ret("open error");
            continue;
        }

        times[0] = statbuf.st_atim;
        times[1] = statbuf.st_mtim;
        if(futimens(fd, times) < 0){
            err_ret("futimens error");
        }
        close(fd);
    }
    exit(0);
}