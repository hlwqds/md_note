#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>
static void my_exit1(void);
static void my_exit2(void);

int main(int argc, char **argv){
    atexit(my_exit1);
    atexit(my_exit2);

    printf("main has done\n");
    return 0;
    exit(0);
    _Exit(0);
    _exit(0);
}

static void my_exit1(void){
    printf("my_exit1 has done\n");
}
static void my_exit2(void){
    printf("my_exit2 has done\n");
}