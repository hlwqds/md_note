#ifndef _List_H
#define ElementType int
struct Node;
typedef struct Node *PtrToNode;

//对于表头和表元素，我们将返回不同的别名，同时将指针进行隐藏
typedef PtrToNode List;
typedef PtrToNode Position;

List MakeEmpty(List L);
int isEmpty(List L);
int IsLast(Position P, List L);
Position Find(ElementType X, List L);
void Delete(ElementType X, List L);
Position FindPrevious(ElementType X, List L);
void Insert(ElementType X, List L, Position P);
void DeleteList(List L);
Position Header(List L);
Position First(List L);
Position Addvance(Position P);
Position Retrieve(Position P);
#endif

struct Node{
    ElementType Element;
    Position    Next;
};