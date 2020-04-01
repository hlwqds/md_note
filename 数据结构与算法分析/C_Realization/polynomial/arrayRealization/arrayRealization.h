#ifndef _POLYNOMIAL_ARRAY_REALIZATION_
#include <stdlib.h>
#include <stdio.h>
#define MAXDEGREE 2048
typedef struct{
    int CoffArray[MAXDEGREE];
    int HighPower;
} *Polynomail;
void ZeroPolynomial(Polynomail Poly);
void AddPolynomial(Polynomail Poly1, Polynomail Poly2, Polynomail PolySum);
void MultPolynomial(Polynomail Poly1, Polynomail Poly2, Polynomail PolyMult);
#endif