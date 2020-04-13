#include<stdio.h>
#include<unistd.h>

int main()
{
	execl("commandline","commandline","-a huanglin -b haha",NULL);
}
