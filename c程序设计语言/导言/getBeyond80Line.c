#include <stdio.h>
#include <stdlib.h>
#define MAXLINE 1000    //允许的输入行的最大长度

int getLine(char **line);
void copy(char to[], char from[]);
#define FREE_IF_NOT_NULL(x) {if(x){free(x);x=NULL;}}
int main(){
    int len = 0;
	char **lineList = NULL;
    char *line = NULL;
	int listSize = 16;
	int lineNum = 0, i = 0;

	lineList = malloc(sizeof(char *) * listSize);
    while((len = getLine(&line)) > 0){
        if(len > 80){
			lineNum++;
			if(lineNum > listSize){
				listSize *= 2;
				lineList = realloc(lineList, sizeof(char *) * listSize);
			}
			lineList[lineNum - 1] = line;
        }else{
			FREE_IF_NOT_NULL(line);
		}
	}
	
	printf("List:\n");
	for(i = 0; i < lineNum; i++){
		printf("%s-index:%d\n",*(lineList+i) ,i);
		FREE_IF_NOT_NULL(*(lineList+i));
	}
	
	FREE_IF_NOT_NULL(lineList);
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

