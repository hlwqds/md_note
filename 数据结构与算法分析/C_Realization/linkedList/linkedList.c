#include "linkedList.h"

/*return true if the list L is empty*/
int isEmpty(List L){
    return L->Next == NULL;
}

/*return true if P is the last position in list L*/
/*parameter L is unused in this implementation*/
//参数L在这次执行中不会用到
int IsLast(Position P, List L){
    return P->Next == NULL;
}

/*return the position of X in L, NULL if not found*/
Position Find(ElementType X, List L){
    Position P;
    P = L->Next;
    while(P != NULL && P->Element != X){
        P = P->Next;
    }
    return P;
}

Position FindPrevious(ElementType X, List L){
    Position P;
    P = L;
    while(P->Next != NULL && P->Next->Element != X){
        P = P->Next;
    }
    return P;
}

/*delete X in L*/
void Delete(ElementType X, List L){
    Position P, TmpCall;

    P = FindPrevious(X, L);

    /*if previous element is the last, it means there is no X in List*/
    if(!IsLast(P, L)){
        TmpCall = P->Next;
        P->Next = TmpCall->Next;
        free(TmpCall);
    }
}

/*insert X into L after position P*/
void Insert(ElementType X, List L, Position P){
    Position TmpCall;

    TmpCall = malloc(sizeof(*TmpCall));
    if(TmpCall == NULL){
        FataError("Out of space!!!");
    }

    TmpCall->Element = X;
    TmpCall->Next = P->Next;
    P->Next = TmpCall;
}

void DeleteList(List L){
    Position P, TmpCall;

    P = L->Next;
    L->Next = NULL;
    while(P != NULL){
        TmpCall = P->Next;
        free(P);
        P = TmpCall->Next;
    }
}

Position Header(List L){
    return L;
}

Position First(List L){
    return L->Next;
}
