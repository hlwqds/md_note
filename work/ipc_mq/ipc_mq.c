#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/msg.h>
#include <string.h>

#define MQ_FILE "/usr/local/EmicallApp/comm/smsSend"

typedef struct
{
    long mtype;
    char mdata[2048];
}MQ_DATA;

int main()
{
    int ret = 0;
    int msgid = 0;
    msgid = msgget(ftok(MQ_FILE, 0), IPC_CREAT|0666);
    printf("%d\n", msgid);
    
    MQ_DATA buf = {0};
    int snd_size = 0;
    char *json_data = "";
    int data_len = strlen(json_data);

    buf.mtype = 1;
    snd_size = sizeof(long) + sizeof(int);
    memcpy(buf.mdata, (void *)json_data, data_len + 1);
    snd_size += data_len + 1;

    ret = msgsnd(msgid, (void *)&buf, snd_size, IPC_NOWAIT);
    printf("ret = %d\n", ret);
    return ret;
}
