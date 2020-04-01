#include <stdio.h>
#include <string.h>
#define C_LIB  "/usr/include/"
#define MAXN 2048
int processFile(char *fileName);

int processFile(char *fileName){
    int ret = 0;
    char *newFile = NULL;
    char line[MAXN] = {0};
    FILE *fptr = NULL;
    fptr = fopen(fileName, "r");
    if(fptr == NULL){
        ret = -1;
        goto _exit;
    }
    
    printf("%s\n", fileName);
    while(fgets(line, MAXN, fptr)){
        line[strlen(line) - 1] = '\0';
        if(newFile = strstr(line, "#include")){
            char *start = NULL;
            char *end = NULL;
            char *path = NULL;
            char realPath[MAXN] = {0};
            char tmp[MAXN] = {0};
            if(start = strchr(newFile, '<')){
                path = C_LIB;
                if((end = strchr(start+1, '>')) == NULL){
                    continue;
                }
            } else if(start = strchr(newFile, '"')){
                path = "./";
                if((end = strchr(start+1, '"')) == NULL){
                    continue;
                }
            } else{
                    continue;
            }
            int fileNameLen = end - start - 1;
            strncpy(tmp, start + 1, fileNameLen);

            tmp[fileNameLen] = '\0';

            snprintf(realPath, sizeof(realPath), "%s%s", path, tmp);

            processFile(realPath);
        }
    }

_exit:
    if(fptr != NULL){
       fclose(fptr);
    }
    return ret;
} 

int main(){
    processFile("1-4.c");
}