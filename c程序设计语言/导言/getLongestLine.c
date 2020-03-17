#include <stdio.h>
#define MAXLINE 1000    //允许的输入行的最大长度

int getline(char line[], int maxline);
void copy(char to[], char from[]);

int main(){
    int len = 0;
    int max = 0;
    char line[MAXLINE] = {0};
    char longest[MAXLINE] = {0};

    while((len = getline(line, MAXLINE) > 0)){
        if(len > max){
            max = len;
            copy(longest, line);
        }
    }
    if(max > 0){
        printf("%s\n", longest);
    }
    return 0;
}

void copy(char to[], char from[]){
    int i = 0;
    while((to[i] = from[i]) != '\0'){
        i++;
    }
}

//getline函数，将一行读入s中并返回其长度
int getline(char line[], int maxline){
    int c, i;
    for(i = 0; i<maxline-1 && (c = getchar())!=EOF && c!='\n'; i++){
        line[i] = c;
    }

    if (c == '\n'){
        i++;
    }
    s[i] = '\0';
    return i;
}