#include <apue.h>
#include <error.h>
#include <fcntl.h>

int main(int argc, char **argv){
    int val;
    if(argc != 2){
        err_quit("argc error");
    }

    if((val = fcntl(atoi(argv[1]), F_GETFL, 0)) < 0){
        err_sys("fcntl error");
    }

    switch(val & O_ACCMODE){
        case O_RDONLY:
            printf("read only\n");
            break;
        case O_WRONLY:
            printf("write only\n");
            break;
        case O_RDWR:
            printf("read write only\n");
            break;
        default:
            err_dump("unknown access mode\n");
    }

    if(val & O_APPEND){
        printf("append\n");
    }

    if(val & O_NONBLOCK){
        printf("nonblock\n");
    }

    if(val & O_SYNC){
        printf("sync\n");
    }

#if !defined(_POSIX_C_SOURCE) && defined(O_FSYNC) && (O_FSYNC != O_SYNC)
    if(val & O_FSYNC){
        printf("fsync\n");
    }
#endif

    exit(0);
}