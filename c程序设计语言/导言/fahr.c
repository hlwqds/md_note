#include <stdio.h>
int main(){
    float fahr = 0.0, celsius = 0.0;
    int lower = 0, upper = 600, step = 20;

    printf("fahr   \tcelsius\n");
    fahr = lower;
    while(fahr <= upper){
        celsius = 5 * (fahr - 32) / 8;
        printf("%6.1f\t%6.1f\t\n", fahr, celsius);
        fahr += step;
    }    
}