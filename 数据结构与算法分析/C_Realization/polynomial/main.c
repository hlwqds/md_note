#include "arrayRealization/arrayRealization.h"

int main(int argc, char **argv){
    int i;
    Polynomail Poly1 = malloc(sizeof(*Poly1));
    Polynomail Poly2 = malloc(sizeof(*Poly2));
    Polynomail PolySum = malloc(sizeof(*PolySum));
    Polynomail PolyMult = malloc(sizeof(*PolyMult));
    ZeroPolynomial(Poly1);
    ZeroPolynomial(Poly2);
    Poly1->HighPower = 5;
    Poly1->CoffArray[5] = 4;
    Poly1->CoffArray[4] = 3;
    Poly1->CoffArray[3] = 2;
    Poly1->CoffArray[2] = 1;
    Poly1->CoffArray[0] = -1;

    Poly2->HighPower = 3;
    Poly2->CoffArray[3] = 7;
    Poly2->CoffArray[2] = 8;
    Poly2->CoffArray[1] = 9;
    Poly2->CoffArray[0] = 6;
    AddPolynomial(Poly1, Poly2, PolySum);
    for(i = 0; i <= PolySum->HighPower; i++){
        printf("%d ", PolySum->CoffArray[i]);
    }
    printf("\n");
    MultPolynomial(Poly1, Poly2, PolyMult);
    for(i = 0; i <= PolyMult->HighPower; i++){
        printf("%d ", PolyMult->CoffArray[i]);
    }
    printf("\n");
}