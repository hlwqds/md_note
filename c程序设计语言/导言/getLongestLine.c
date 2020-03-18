#include <stdio.h>
#include <stdlib.h>
#define MAXLINE 1000    //允许的输入行的最大长度

int getLine(char **line);
void copy(char to[], char from[]);
#define FREE_IF_NOT_NULL(x) {if(x){free(x);x=NULL;}}
int main(){
    int len = 0;
    int max = 0;
    char *line = NULL;
    char *longest = NULL;

    while((len = getLine(&line)) > 0){
        if(len > max){
            max = len;
			FREE_IF_NOT_NULL(longest);
			longest = line;
        }else{
			FREE_IF_NOT_NULL(line);
		}
    }
    if(max > 0){
        printf("%s:len-%d\n", longest, max);
    }
	FREE_IF_NOT_NULL(longest);
	FREE_IF_NOT_NULL(line);
    return 0;
}

void copy(char to[], char from[]){
    int i = 0;
    while((to[i] = from[i]) != '\0'){
        i++;
    }
}

//getline函数，将一行读入s中并返回其长度
int getLine(char **line){
    int c, i = 0, size = 0;
	size = 16;
	*line = malloc(sizeof(char) * size);
	while((c = getchar())!=EOF && c!='\n'){
		i++;
		if(i >= size){
			size *= 2;
			*line = realloc(*line, sizeof(char) * size);
		}
        (*line)[i-1] = c;
    }

    if (c == '\n'){
        i++;
    }
    (*line)[i] = '\0';
    return i;
}

