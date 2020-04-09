#include <apue.h>
#include <error.h>

int main(){
    FILE *fp = NULL;
    char buf[BUFSIZ] = {0};
    setbuf(fp, NULL);

    setvbuf(fp, buf, _IOFBF, sizeof(buf));

    fflush(fp);
}