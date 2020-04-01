#ifndef _LINKED_LIST_REALIZATION_
#include <stdlib.h>
typedef struct Node *List;
typedef struct Node *PtrToNode;
struct Node{
    int Coefficient;
    int Exponent;
    PtrToNode Next;
};

List InitHeader();
void InsertAfterNode(List l, PtrToNode node, int Coefficient, int Exponent);
void InsertIntoList(List l, int Coefficient, int Exponent);
void SumLinkedPolynomial(List poly1, List poly2, List sum);
void MultLinkedPolynomial(List poly1, List poly2, List mult);
#endif