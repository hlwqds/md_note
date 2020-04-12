#include <stdio.h>

int main(){
    FILE *fp = NULL;
    fp = fopen("test.md", "r+");
    if(fp == NULL){
        printf("fopen error\n");
    }
    char buf[256];

    fgets(buf, sizeof(buf), fp);
    
    printf("buf: %s", buf);

    fputs(buf, fp);
    
    fclose(fp);
}