#include <stdio.h>

#define NULL ((void *)0)

static inline int formatSeconds(unsigned long totalSec, unsigned int *hour, unsigned int *min, unsigned int *sec)
{
    if (hour == NULL || min == NULL || sec == NULL)
    {
        return -1;
    }

    *hour = totalSec / 3600;
    *min = (totalSec % 3600) / 60;
    *sec = (totalSec % 3600) % 60;

    return 0;
}

int main(){
    unsigned long duration = 8888;
    unsigned long hour = 0, min = 0, sec = 0;
    formatSeconds(duration, &hour, &min, &sec);
    printf("%lu, %lu, %lu\n", hour, min, sec);
    return 0;
}