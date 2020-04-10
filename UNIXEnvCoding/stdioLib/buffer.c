#include <stdio.h>

int main(){
    FILE *fp = NULL;
    char buf[BUFSIZ] = {0};
    setbuf(fp, NULL);

    setvbuf(fp, buf, _IOFBF, sizeof(buf));
    buf[3] = 'h';
    printf("huanglin\n");

    fflush(fp);
}