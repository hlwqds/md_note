#include <stdio.h>
#include <error.h>

int main(){

    FILE *fp = NULL;
    fp = fopen("test.md", "r");
    char c = getc(fp);
    printf("%c", c);
    ferror(fp);
    ungetc(fp, c);

    putc(c, fp);
    fputc(c, fp);
    putchar(c);
    
    feof(fp);
    fclose(fp);
}