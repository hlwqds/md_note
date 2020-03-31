#include <stdio.h>

void PrintDigit(int d){
    printf("%d", d);
}

void PrintOut(int N){
    if(N >= 10){
        PrintOut(N / 10);
    }

    PrintDigit(N % 10);
}

void PrintR(int R){
    if(R < 0){
        printf("%c", '-');
        R = -R;
    }
    PrintOut(R);
}

int main(){
    int test = -993212;

    PrintR(test);
}