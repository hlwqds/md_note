#include <stdio.h>
#include <string.h>
void reverseT(char *string);

void reverseT(char *string){
	if(string == NULL){
		return;
	}

	char tmp;
	int i = 0;
	int len = strlen(string);
	for(i = 0; i < len/2; i++){
		tmp = string[i];
		string[i] = string[len-1-i];
		string[len-1-i] = tmp;
	}
	return;
}

int main(){
	char test[20] = {'1','h','3','4','6','9'};

	reverseT(test);
}
