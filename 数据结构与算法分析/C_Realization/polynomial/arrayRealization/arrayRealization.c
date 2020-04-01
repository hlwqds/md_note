#include "arrayRealization.h"

int Max(int a, int b){

    return (a > b) ? a : b;
}

void ZeroPolynomial(Polynomail Poly){
    int i;
    for(i = 0; i < MAXDEGREE; i++){
        Poly->CoffArray[i] = 0;
    }
    Poly->HighPower = 0;
}

void AddPolynomial(Polynomail Poly1, Polynomail Poly2, Polynomail PolySum){
    int i;
    ZeroPolynomial(PolySum);
    PolySum->HighPower = Max(Poly1->HighPower, Poly2->HighPower);
    for(i = 0; i <= PolySum->HighPower; i++){
        PolySum->CoffArray[i] = Poly1->CoffArray[i] + Poly2->CoffArray[i];
    }
}

void MultPolynomial(Polynomail Poly1, Polynomail Poly2, Polynomail PolyMult){
    int i, j;
    ZeroPolynomial(PolyMult);
    PolyMult->HighPower = Poly1->HighPower + Poly2->HighPower;
    if(PolyMult->HighPower >= MAXDEGREE){
        printf("Exceeded array size\n");
    }else{
        for(i = 0; i <= Poly1->HighPower; i++){
            for(j = 0; j <= Poly2->HighPower; j++){
                PolyMult->CoffArray[i+j] += Poly1->CoffArray[i] * Poly2->CoffArray[j];
            }
        }
    }
}