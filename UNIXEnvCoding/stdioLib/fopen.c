#include <apue.h>
#include <error.h>

int main(){
    FILE *fp = NULL;
    fp = fopen("test.md", "r");
    fclose(fp);
}