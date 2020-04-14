#include <stdio.h>

extern char **environ;

int main(int argc, char **argv){
    for(int i = 0; environ[i]; i++){
        printf("%s\n", *environ);
    }
    return 0;
}