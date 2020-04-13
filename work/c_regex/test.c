#include <stdio.h>
#include <regex.h>
#include <string.h>
#include <stdlib.h>
static char* substr(const char*str, unsigned start, unsigned end)
{
    unsigned n = end - start;
    static char stbuf[256];
    strncpy(stbuf, str + start, n);
    stbuf[n] = 0;
    return stbuf;
}

static int gen_reg_string(char *source_string, char *reg_string, int reg_size, char **param_list, int *num)
{
    char *aim_ptr = NULL;
    char *begin_ptr = NULL;
    char tmp[128] = {0};
    char *next_ptr = NULL;
    begin_ptr = tmp;
    aim_ptr = tmp;
    next_ptr = tmp;
    
    *num = 0;
    strncpy(tmp, source_string, sizeof(tmp));

    while((aim_ptr = strchr(next_ptr, '{')) != NULL)
    {   
        *aim_ptr = '\0';
        if(reg_size > strlen(reg_string) + strlen(next_ptr))
        {
            strcat(reg_string, next_ptr);
            next_ptr = aim_ptr + 1;
        }
        else
            goto __error;
        

        if((aim_ptr = strchr(next_ptr, '}')) != NULL)
        {
            *aim_ptr = '\0';
            *param_list = realloc(*param_list, 100 * ((*num) + 1));
            strncpy(*param_list + (*num * 100), next_ptr, 100);
            printf("param:%s\n", *param_list + (*num * 100));
            (*num)++;
            if(reg_size > strlen(reg_string) + strlen("(.*)"))
            {
                strcat(reg_string, "(.*)");
                next_ptr = aim_ptr + 1;
            }
        }
        else
            goto __error;

    }

    return 0;
__error:
    printf("error\n");
    return -1;
}

static int reg_match_with_expect_num(char *matched_string, char *reg_string, int num, char **match_list)
{
    int ret = 0;
    int i = 0;
    regex_t reg;
    regmatch_t *pm = NULL;

    ret = regcomp(&reg, reg_string, REG_EXTENDED|REG_ICASE);
    pm = malloc(sizeof(regmatch_t) * (num + 1));
    
    ret = regexec(&reg, matched_string, num + 1, pm, 0);
    if(ret != 0)
    {
        return 2;
    }

    *match_list = malloc(num * 100);

    for(i = 0; i < num + 1; i++)
    {
        snprintf(*match_list + i * 100, 100, "%s", substr(matched_string, pm[i].rm_so, pm[i].rm_eo));
    }
    regfree(&reg);
    free(pm);
    return 0;
}

int main(int argc, char **argv)
{
    char *string = "hello,{xxx},i am{xxxx}, qunide {xxxxx}";
    char reg_string[100] = {0};
    char *param_list = NULL;
    int num = 0;
    int i = 0;
    gen_reg_string(string, reg_string, sizeof(reg_string), &param_list, &num);
    printf("reg_string:%s\n", reg_string);
    printf("num:%d\n", num);
    for(i = 0; i < num; i++)
    {
        printf("param:%s\n", param_list + i * 100);
    }
    if(param_list != NULL)
        free(param_list);

    char *value_list = NULL;
    char *matched_string = "hello,huanglin,i amaaaa, qunide peipeipei";

    reg_match_with_expect_num(matched_string, reg_string, num, &value_list);
      
    for(i = 0; i < num; i++)
    {
        printf("value:%s\n", value_list + i * 100);
    }
    if(value_list)
        free(value_list);

    //regex_t reg;
    //regmatch_t pm[4];

    //regcomp(&reg, reg_string, REG_EXTENDED|REG_ICASE);
    //pm = malloc(sizeof(regmatch_t) * (num + 1));
    //regexec(&reg, matched_string, num + 1, pm, 0);

    //for(i = 0; i < num + 1; i++)
    //{
        //strncpy(value_list + i * 100, substr(matched_string, pm[i].rm_so, pm[i].rm_eo), 100);
    //    printf("value:%s\n", substr(matched_string, pm[i].rm_so, pm[i].rm_eo));
    //}
    //regfree(&reg);
    return 0;
}
