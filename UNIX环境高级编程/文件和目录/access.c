#include <apue.h>
#include <error.h>
#include <fcntl.h>

int main(int argc, char **argv){
    if(argc != 2){
        err_quit("argc error");
    }

    if(access(argv[1], R_OK) < 0){
        err_ret("access error for %s", argv[1]);
    }else{
        printf("read access OK\n");
    }

    if(open(argv[1], O_RDONLY) < 0){
        err_ret("open error for %s", argv[1]);
    }else{
        printf("open %s OK\n", argv[1]);
    }
    exit(0);
}