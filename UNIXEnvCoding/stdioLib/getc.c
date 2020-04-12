#include <stdio.h>
#include <error.h>

int main(){

    FILE *fp = NULL;
    fp = fopen("test.md", "r+");
    if(fp == NULL){
        printf("fopen error\n");
        exit(0);
    }

    char c = getc(fp);
    printf("c: %d\n", c);
    ferror(fp);
    ungetc(c, fp);
    printf("afer ungetc c: %d\n", c);

    putc(c, fp);
    printf("afer putc c: %d\n", c);
    printf("putc c: %d\n", c);

    fputc(c, fp);
    printf("fputc c: %c\n", c);

    c = 'h';
    putchar(c);
    printf("putchar c: %c\n", c);

    feof(fp);
    fclose(fp);
}