#include <stdio.h>

typedef enum{
    OUT = 0,
    IN = 1,
}CursorState;

int main() {
    int c, nw, nc, nl, state;
    nw = nc = nl = 0;
    state = OUT;
    while((c = getchar()) != EOF){
        nc++;
        if (c == '\n'){
            ++nl;
        }

        //出现频率‘ ’ > ‘\n’ > '\t'
        if (c == ' ' || c == '\n' || c == '\t')
        {
            state = OUT;
        }
        else if (state == OUT){
            state = IN;
            ++nw;
            putchar('\n');
            putchar(c);
        }
        else{
            putchar(c);
        }
    }
    printf("nc: %d, nw: %d, nl: %d\n", nc, nw, nl);
    return 0;
}