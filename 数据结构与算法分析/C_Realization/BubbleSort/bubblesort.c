#include <stdio.h>
#include <time.h>
#define ARR_LEN 200000 /*数组长度上限*/
#define elemType int /*元素类型*/

void bubbleSort(elemType *arr, int len){
    elemType i = 0, j = 0, k = 0;
    elemType tmp = 0;
    for(i = 0; i < len; i++){
        for(j = 0; j < len - i - 1; j++){
            if(arr[j] > arr[j+1]){
                tmp = arr[j];
                arr[j] = arr[j+1];
                arr[j+1] = tmp;
            }
        }
    }
}
 
int main (void) {
    elemType arr[ARR_LEN];
    int i;
    time_t timeStart = 0;
    timeStart = time(NULL);
    bubbleSort (arr, ARR_LEN);
    printf("%ld\n", time(NULL) - timeStart);
    printf("%d\n", arr[ARR_LEN/2]);
    
    return 0;
}