#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>

int main ()  
{  
	pid_t fpid; //fpid表示fork函数返回的值 
	int count=0;
	char *ptr = NULL;

	putenv("HUANGLIN=77");
	ptr = getenv("HUANGLIN");

        if(ptr)
                printf("%s\n",ptr); 
	fpid=fork();
	if (fpid < 0)  
        	printf("error in fork!\n");  
	else if (fpid == 0) { 
		printf("i am the child process, my process id is %d\n",getpid());  
        	printf("我是爹的儿子\n");//对某些人来说中文看着更直白。 
        	count++;
		ptr = getenv("HUANGLIN");

        	if(ptr)
                	printf("%s\n",ptr);
	} 
	else {
		printf("i am the parent process, my process id is %d\n",getpid());  
        	printf("我是孩子他爹\n"); 
        	count++; 
	} 
	printf("统计结果是: %d\n",count); 
	return 0; 
}  
