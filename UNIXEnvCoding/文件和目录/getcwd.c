#include <apue.h>
#include <error.h>
int main(){
    char buf[256] = {0};
    if(chdir("/tmp") < 0){
        err_sys("chdir error");
    }

    if(getcwd(buf, sizeof(buf)) == NULL){
        err_sys("getcwd error");
    }
    printf("cwd = %s\n", buf);
    exit(0);
}