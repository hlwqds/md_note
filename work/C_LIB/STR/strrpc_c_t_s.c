#include "stdio.h"
#include "stdlib.h"
#include "string.h"

//src_size 系统实际分配的空间，而不是字符串
//如果返回错误，证明空间不足，需要将空间扩展
int strrpl_c_t_s(char *src, int src_size, char c, char *string, int stringLen){
	int act_len = 0, ret = 0;
	char *ptr = src;
	char *tmp = NULL, *start = NULL;
	int i = 0;

	tmp = malloc(src_size);
	start = tmp;

	while(1){
		if(*ptr == c){
			act_len += stringLen;

			if(act_len >= src_size){
				ret = -1;
				goto _exit;
			}

			for(i = 0; i < stringLen; i++){
				*tmp = *(string + i);
				tmp++;
			}
		}
		else{
			*tmp = *ptr;
			tmp++;
			act_len++;
		}

		if(*ptr == '\0'){
			break;
		}

		if(act_len >= src_size){
			ret = -1;
			goto _exit;
		}

		ptr++;
	}

	strncpy(src, start, act_len);

_exit:
	free(start);
	return ret;
}


int main(){
	int ret = 0;
	
	char list[19] = "huanglin,lin,l,";
	char *sss = "\"\"";

	ret = strrpl_c_t_s(list, sizeof(list), ',', sss, strlen(sss));
	printf("%s\n", list);
}
