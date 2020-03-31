#include "linkedListRealization.h"

List InitHeader(){
    List header = malloc(sizeof(*header));
    return header;
}


//insert after node
void InsertAfterNode(List l, PtrToNode node, int Coefficient, int Exponent){
    if(node == NULL || l == NULL){
        return;
    }
    PtrToNode curNode = malloc(sizeof(*curNode));
    if(curNode == NULL){
        return;
    }

    curNode->Next = node->Next;
    node->Next = curNode;
}

void InsertIntoList(List l, int Coefficient, int Exponent){
    if(l == NULL){
        return;
    }
    PtrToNode curNode = NULL;
    PtrToNode preNode = NULL;
    for(curNode = l->Next; curNode != NULL; preNode = curNode,curNode = curNode->Next){
        if(curNode->Next == NULL){
            //找到链表末尾，证明没有比他次数更小的项
            InsertAfterNode(l, curNode, Coefficient, Exponent);
            break;
        }else if(curNode->Exponent < Exponent){
            //此时节点次数比他小，证明应该插在前一个节点后面
            InsertAfterNode(l, preNode, Coefficient, Exponent);
            break;
        }else if(curNode->Exponent == Exponent){
            curNode->Coefficient += Coefficient;
            break;
        }
        
    }
}

void SumLinkedPolynomial(List poly1, List poly2, List sum){
    if(poly1 == NULL || poly2 == NULL || sum == NULL){
        return;
    }


    PtrToNode tmp = poly2->Next;
    PtrToNode node1, node2;
    PtrToNode curSumNode = sum->Next;
    for(node1 = poly1->Next; node1 != NULL; node1=node1->Next){
        for(node2 = tmp; node2 != NULL; node2=node2->Next){
            if(node2->Exponent == node1->Exponent){
                //找到了一样大的数，下次从下一个节点开始查找
                InsertAfterNode(sum, curSumNode, node1->Coefficient + node2->Coefficient, node1->Exponent);
                curSumNode = curSumNode->Next;
                tmp = node2->Next;
                break;
            }else if(node2->Exponent < node1->Exponent || node2->Next == NULL){
                //出现比node1还要小的数，此次循环中并没有找到想要的数，下次循环从这个节点开始查找
                InsertAfterNode(sum, curSumNode, node1->Coefficient, node1->Exponent);
                tmp = node2->Next;
                break;
            }
        }
    }
}

void MultLinkedPolynomial(List poly1, List poly2, List mutl){
    if(poly1 == NULL || poly2 == NULL){
        return;
    }

    PtrToNode node1, node2;
    for(node1 = poly1->Next; node1 != NULL; node1=node1->Next){
        for(node2 = poly2->Next; node2 != NULL; node2=node2->Next){
            InsertIntoList(mutl, node1->Coefficient * node2->Coefficient,
            node1->Exponent + node2->Exponent);
        }
    }
    return;
}

int main(int argc, char **argv){

}