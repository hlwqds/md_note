#include <stdio.h>
#include <unistd.h>
#include <stdlib.h> 
int main(int argc,char *argv[])
{
    int ch;
    opterr=0;
    int a_int = 0;
    while((ch=getopt(argc,argv,"a:b::cde"))!=-1)
    {
        printf("optind:%d\n",optind);
        printf("optarg:%s\n",optarg);
        printf("ch:%c\n",ch);
        switch(ch)
        {
        case 'a':
	    a_int = atoi(optarg);
	    printf("a_int:%d\n", a_int);
            printf("option a:'%s'\n",optarg);
            break;
        case 'b':
            printf("option b:'%s'\n",optarg);
            break;
        case 'c':
            printf("option c\n");
            break;
        case 'd':
            printf("option d\n");
            break;
        case 'e':
            printf("option e\n");
            break;
        default:
            printf("other option:%c\n",ch);
        }
        printf("optopt+%c\n",optopt);
    }
}
