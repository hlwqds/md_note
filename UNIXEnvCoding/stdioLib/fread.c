#include <stdio.h>
#include <stdlib.h>
#include <string.h>
struct testData;
typedef struct testData TESTDATA;
struct testData{
    char name[64];
    int  age;
};

int main(){
    FILE *fp = NULL;
    fp = fopen("test.bits", "w+");
    if(fp == NULL){
        printf("fopen error.\n");
        exit(0);
    }

    float data[10];
    if(fwrite(&data[2], sizeof(float), 4, fp) != 4){
        printf("fwrite error\n");
        exit(0);
    }
    fclose(fp);
    putchar('[');
    for(int i = 0; i < 10; i++){
        if(i == 0){
            printf("%f", data[i]);
        }else{
            printf(",%f", data[i]);
        }
    }
    putchar(']');
    putchar('\n');
    printf("fp: %p\n", fp);
    printf("data: %lu\n", sizeof(data));

    fp = fopen("test.bits", "r");
    float data1[10] = {0};
    if(fread(data1, sizeof(float), 4, fp) != 4){
        printf("fread error\n");
    }
    putchar('[');
    for(int i = 0; i < 10; i++){
        if(i == 0){
            printf("%f", data1[i]);
        }else{
            printf(",%f", data1[i]);
        }
    }
    putchar(']');
    putchar('\n');
    fclose(fp);

    fp = fopen("testStruct.bits", "w+");
    if(fp == NULL){
        printf("fopen error.\n");
        exit(0);
    }
    TESTDATA data2 = {0};
    data2.age = 24;
    strncpy(data2.name, "huanglin", sizeof(data2.name));
    if(fwrite(&data2, sizeof(data2), 1, fp) != 1){
        printf("fwrite error.\n");
        exit(0);
    }
    fclose(fp);

    fp = fopen("testStruct.bits", "r+");
    if(fp == NULL){
        printf("fopen error.\n");
        exit(0);
    }
    TESTDATA data3 = {0};
    if(fread(&data3, sizeof(data3), 1, fp) != 1){
        printf("fread error.\n");
        exit(0);
    }
    printf("data3.name: %s\n", data3.name);
    printf("data3.age: %d\n", data3.age);

    return 0;
}