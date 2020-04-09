#include <apue.h>
#include <error.h>

int main(){

    FILE *fp = NULL;
    fp = fopen("test.md", "r");
    char c = getc(fp);
    printf("%c", c);
    fclose(fp);
}