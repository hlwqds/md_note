#include <apue.h>
#include <error.h>
#include <fcntl.h>

int main(void){
    int input = STDIN_FILENO;
    int output = STDOUT_FILENO;
    int err = STDERR_FILENO;
    int openMax = _SC_OPEN_MAX;

    printf("%d\n", openMax);
}