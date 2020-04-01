#define ElementType int

int binarySearch(const ElementType *A, ElementType X, int N){
    int Low, Mid, High;
    Low = 0;
    High = N-1;
    while(Low <= High){
        Mid = (Low + High) / 2;
        if (X < A[Mid]){
            High = Mid;
        }else if(X > A[Mid]){
            Low = Mid;
        }else{
            return Mid;
        }
    }

    return -1;
}