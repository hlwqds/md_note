



# 新架构ivr交互查询

## 1 需求描述

新架构ivr交互查询兼容老架构98,99查询被叫协议，并在此基础上新增了三种协议，分别是95,96,97。

通用的修改需求：

1. 新架构企业，通过ccgeid获取企业信息
2. 新架构企业，通过CR-Web提供的查询响应接口返回数据
3. 新架构企业，确定用户请求和用户响应格式

### 1.1 95(交互收键模式)

### 1.2 96(被叫查询模式)

### 1.3 97(通用交互模式)

### 1.4 98(AMNB或者AMXNB模式)

### 1.5 99(AMB或者AMXB模式)

## 2 交互流程

```sequence
apd->api: 交互查询请求
api->apd: 接收成功
api->客户服务器: 根据apd的推送封装消息\n并且发给客户
客户服务器->api: 返回对应的ivr参数
api-->cr_web: 封装消息，将对应的消息交给cr_web\n并由其转发给cr
cr_web-->api:cr_web响应成功（响应时机？）
```

## 3 消息定义

### 3.1 apd->api

#### 3.1.1 95(交互收键模式)

```json
{
    "type" : 95,
    "ccgeid" : 1576,
    "callId" : "apixxxxxxxxxxxx",
    "ccNumber" : "13260278209conf0_1519439781636",
    "callType" : 0,
    "ivrFlowId" : 245,
    "ivrQueryId" : 1,
    "caller" : "13260278209",
    "callerType" : 1,
    "switchNumber" : "02566699794",
    "called" : "1001",
    "calledType" : 2,
    "timestamp" : "1519439787",
    "userQueryId" : "id_0000001",
    "inputKeys" : "1000",
    "variables" : [
	    {"id_number" : "110108198703127621" },
		{"name" :"" },
		{"address":"" }
	]
}

```

#### 3.1.2 96(被叫查询模式)

```json
{
    "type" : 96,
    "ccgeid" : 1576,
    "callId" : "apixxxxxxxxxxxx",
    "ccNumber" : "13260278209conf0_1519439781636",
    "callType" : 0,
    "ivrFlowId" : 245,
    "ivrQueryId" : 1,
    "caller" : "13260278209",
    "callerType" : 1,
    "switchNumber" : "02566699794",
    "called" : "1001",
    "calledType" : 2,
    "timestamp" : "1519439787",
    "userQueryId" : "id_0000001",
    "inputKeys" : "1000",
    "variables" : [
	    {"id_number" : "110108198703127621" },
		{:"name" :"" },
		{"address":"" }
	]
}
```

#### 3.1.3 97(通用交互模式)

```json
{
    "type" : 97,
    "ccgeid" : 1576,
    "callId" : "apixxxxxxxxxxxx",
    "ccNumber" : "13260278209conf0_1519439781636",
    "callType" : 0,
    "ivrFlowId" : 245,
    "ivrQueryId" : 1,
    "caller" : "13260278209",
    "callerType" : 1,
    "switchNumber" : "02566699794",
    "called" : "1001",
    "calledType" : 2,
    "timestamp" : "1519439787",
    "userQueryId" : "id_0000001",
    "inputKeys" : "1000",
    "variables" : [
	    {"id_number" : "110108198703127621" },
		{:"name" :"" },
		{"address":"" }
	]
}
```

#### 3.1.4 98(AMNB或者AMXNB模式)

```json
{
    "type" : 98,
    "ccgeid" : 1576,
    "callId" : "apixxxxxxxxxxxx",
    "ccNumber" : "13260278209conf0_1519439781636",
    "ivrFlowId" : 34,
    "ivrQueryId" : 35,
    "switchNumber" : "02566699794",
    "useNumber" : "02566699794",
    "caller" : "13260278209",
    "callerType" : 1,
    "timestamp" : "1519439787",
    "path" : "1000"
}
```

#### 3.1.5 99(AMB或者AMXB模式)

```json
{
    "type" : 98,
    "ccgeid" : 1576,
    "callId" : "apixxxxxxxxxxxx",
    "ccNumber" : "13260278209conf0_1519439781636",
    "ivrFlowId" : 34,
    "ivrQueryId" : 35,
    "switchNumber" : "02566699794",
    "useNumber" : "02566699794",
    "caller" : "13260278209",
    "callerType" : 1,
    "timestamp" : "1519439787",
    "path" : "1000"
}
```

### 3.2 api内部消息转换

根据前面的[需求描述](#1 需求描述)

1. 根据ccgeid从企业表中获取企业信息

根据ccgeid获取企业信息，我们不需要通过callId或者总机号才能获取到企业信息，需要注意的是。后面处理时还会尝试获取一次企业的详细信息，对于新架构，我们需要将企业的信息发送过去，以减少一次数据库操作（enterprise结构体的大小显然对于推送的消息队列来说过于庞大了，我还是妥协后面再一次获取企业的信息）。



2. 新架构的通用ivr交互新增了部分通用字段，这些字段api必须转发给客户。需要新增一个结构体用来存储这些信息

| 通话类型 | IVR流程      | callType | caller     | callerType | called     | calledType |
| -------- | ------------ | -------- | ---------- | ---------- | ---------- | ---------- |
| 呼入     | 正常流程     | 5        | 客户号码   | 1          | -          | -          |
| 呼入     | 子IVR流程    | 5        | 客户号码   | 1          | 坐席话机号 | 1/2        |
| 呼出     | 子IVR流程    | 0或7     | 坐席话机号 | 1/2        | 客户号码   | 1          |
| 语音通知 | 语音通知流程 | 3        | 总机号码   | 1          | 客户号码   | 1          |

```c
typedef enum
{
    IVR_CALLER_AND_CALLLED_TYPE_OUTLINE = 1,
    IVR_CALLER_AND_CALLLED_TYPE_INLINE = 2，
}IVRCallerAndCalledType;

typedef struct __push_post_data_ivr_moudle__
{
    DB_CALL_RECORD_CALL_TYPE    calltype;     //联合calltype caller called，可以得知在此次ivr请求中，哪个是客户，哪个是坐席
    /*
    | 通话类型 | IVR流程      | callType | caller     | callerType | called     | calledType |
    | -------- | ------------ | -------- | ---------- | ---------- | ---------- | ---------- |
    | 呼入     | 正常流程     | 5        | 客户号码   | 1          | -          | -          |
    | 呼入     | 子IVR流程    | 5        | 客户号码   | 1          | 坐席话机号 | 1/2        |
    | 呼出     | 子IVR流程    | 0或7     | 坐席话机号 | 1/2        | 客户号码   | 1          |
    | 语音通知 | 语音通知流程 | 3        | 总机号码   | 1          | 客户号码   | 1          |
    */
    unsigned long ivrFlowId;    //ivr流程id
    unsigned long ivrQueryId;   //ivr查询节点id
    char          userQueryId[64];  //用于向客户服务器确定id对应下一步的操作是什么
    char          inputKeys[64];    //推送上次输入的按键
    char          *variables;
    /*
    用户自定义查询请求变量值集合，是{"variables" : [ {"id_number" : "110108198703127621"}, {:"name" :""}, {"address":""} ]}，
    存入对空间中，用完释放
    */
}ModCallPushPostDataIVRMoudle;

typedef struct __push_post_data__
{
    unsigned char               callId[CALL_RECORD_CALLID_MAX_LEN];
    unsigned char               ccNumber[CALL_RECORD_CALLID_MAX_LEN];
    unsigned char               caller[PHONE_NUMBER_MAX_LEN];
    unsigned char               called[PHONE_NUMBER_MAX_LEN];
    int                         xferTimes;
    BOOL                        extCaller;
    BOOL                        extCalled;
    BOOL                        isCaller;
    DB_CALL_RECORD_CALL_TYPE    type;
    DB_CALL_RECORD_STATUS       status;
    unsigned char               useNumber[PHONE_NUMBER_MAX_LEN];
    unsigned char               switchNumber[PHONE_NUMBER_MAX_LEN];
    unsigned char               subNumber[PHONE_NUMBER_MAX_LEN];
    unsigned char               virtNumber[PHONE_NUMBER_MAX_LEN];
    unsigned long               enterpriseId;
    unsigned long               ccgeid;
    unsigned long               ringTime;
    unsigned long               startTime;
    unsigned long               endTime;
    unsigned long               timestamp;
    unsigned long               duration;
    unsigned long               reason;
    unsigned long               gid;
    unsigned long               pbxCallLogId;
    char                        feedback[CALL_RECORD_FEEDBACK_MAX_LEN];
    ModCallPushPostDataIVRMoudle    *ivr_argv;      //新架构通用ivr扩展字段

    BOOL                        realtimeData;
    BOOL                        isCheckData;
    BOOL                        isCommonEnterprise;
    BOOL                        hangup2calling;
    int                         index;
    int                         lwpid;
    int                         mes_type;
    int                         failed_delay_time;
    unsigned char               path[CALL_RECORD_FEEDBACK_MAX_LEN];   //99协议 ，按键 (可能包含二级按键，比如:2-9)

    unsigned long               app_id;
    unsigned long               provinceId;
    char                        number[USER_NUMBER_MAX_LEN];
    char                        mobile[PHONE_NUMBER_MAX_LEN];
    char                        destNumber[USER_NUMBER_MAX_LEN];
    unsigned long               ngnReason;                              //201708-N02细化通话失败和挂断原因
    char                        batchCallId[CALL_RECORD_CALLID_MAX_LEN];
    char                        batchCallUserData[USER_DATA_MAX_LEN];
    char                        batchCallTaskId[BATCH_TASK_ID_MAX_LEN];
} ModCallPushPostData;
```



### 3.2 api->用户服务器

#### 3.2.1 请求客户

1. 99/98 99和98向客户的请求信息保持不变

```json
{
	'appId':'b23abb6d451346efa13370172d1921ef',
	'callId':'api1234059445aDbbJxIdbT',
    'accountSid':'c5dc4b87f33ef2ef37c8e974793ad8e5',
    'caller':'18769874345',
    'path':'1-2',
    'callType':99,
    'type':0,
    'useNumber':'02566687987',
    'switchNumber' : "02566699794",
    'userData':FE87D3
}
```

2. 95,96和97消息需要使用新架构新的消息格式

```json
{
    "type" : 95/96/97,
    "ccgeid" : 1576,
    "callId" : "apixxxxxxxxxxxx",
    "ccNumber" : "13260278209conf0_1519439781636",
    "callType" : 0,
    "ivrFlowId" : 245,
    "ivrQueryId" : 1,
    "caller" : "13260278209",
    "callerType" : 1,
    "switchNumber" : "02566699794",
    "called" : "1001",
    "calledType" : 2,
    "timestamp" : "1519439787",
    "userQueryId" : "id_0000001",
    "inputKeys" : "1000",
    "variables" : [
	    {"id_number" : "110108198703127621" },
		{"name" :"" },
		{"address":"" }
	]
}
说明：在上例中，有3个全局变量：id_number、name和address，id_number已经赋值，name和address未赋值，需要用户服务器返回，并在IVR其它节点中引用。
```



#### 3.2.2 客户响应
actionType

```c
typedef enum
{
  	0,	invalid
  	1,	放音响应
  	2,	放音按键响应
  	3,	转技能组响应
  	4,	转坐席响应
  	5,	转外线响应
  	6,	转其他IVR流程响应
  	7,	流程结束响应
}IvrActionType;
```

##### 3.2.2.1 放音响应
json to cr

```json
{
    "rspCode" : 0,
    “ccgeid” : “123”,
    “ccNumber” : :21212”,
	"userQueryId" : "id_0000001",
	"variables" : [
	    { "id_number" : "110108198703127621" },
		{ "name" :"张三" },
		{ "address":"江苏省南京市江宁区" }
	]
	"nextAction" : {
	    "action" : 1,
		"paras" : {
		    "voiceId" : "播放语音文件id",
			"voiceName" : "播放语音文件唯一名称",
			"allowBreak" : "是否允许打断: 0-不允许 1-允许"
		}
	}
	"reason" : "test",
	"userdata" : "test"
}
```

##### 3.2.2.2 放音收键响应

json to cr

```json
{
    "rspCode" : 0,
    “ccgeid” : “123”,
    “ccNumber” : :21212”,
	"userQueryId" : "id_0000001",
	"variables" : [
	    { "id_number" : "110108198703127621" },
		{ "name" :"张三" },
		{ "address":"江苏省南京市江宁区" }
	]
	"nextAction" : {
	    "action" : 2,
		"paras" : {
		    "voiceId" : "播放语音文件id",
			"voiceName" : "播放语音文件唯一名称",
			"allowBreak" : "是否允许打断: 0-不允许 1-允许",
			"getKeyNumber" : "获取按键位数",
			"getKeyTimeout" : "收键超时时间",
			"endWithHashKey" : "是否以#号键结束, 0-不是, 1-是"
		}
	}
	"reason" : "test",
	"userdata" : "test"
}

```

##### 3.2.2.3 转技能组响应

json to cr

```json
{
    "rspCode" : 0,
    “ccgeid” : “123”,
    “ccNumber” : :21212”,
	"userQueryId" : "id_0000001",
	"variables" : [
	    { "id_number" : "110108198703127621" },
		{ "name" :"张三" },
		{ "address":"江苏省南京市江宁区" }
	]
	"nextAction" : {
	    "action" : 3,
		"paras" : {
		    "acdId" : "技能组id",
			"acdName" : "技能组名称",
			"useAcdValue" : "0-不使用技能组配置 1-使用技能组配置",
			"queueTime" : "排队超时时长",
			"switchTimes" : "坐席流转次数",
			"ringTimeout" : "坐席振铃超时时长",
			"customerMemory" : "0-不记忆 1-优先熟客记忆 2-强制熟客记忆"
		}
	}
	"reason" : "test",
	"userdata" : "test"
}

```

##### 3.2.2.4 转座席响应

json to cr

```json
{
    "rspCode" : 0,
    “ccgeid” : “123”,
    “ccNumber” : :21212”,
	"userQueryId" : "id_0000001",
	"variables" : [
	    { "id_number" : "110108198703127621" },
		{ "name" :"张三" },
		{ "address":"江苏省南京市江宁区" }
	]
	"nextAction" : {
	    "action" : 4,
		"paras" : {
		    "workNumber" : "1001,1002,1003",
			"number" : "1001,1002,1003",
			"queueTime" : "坐席忙时排队时长",
			"ringTimeout" : "多坐席情况下，坐席振铃超时时长"
		}
	}
	"reason" : "test",
	"userdata" : "test"
}
```

##### 3.2.2.5 转外线响应

json to cr

```json
{
    "rspCode" : 0,
    “ccgeid” : “123”,
    "ccNumber" : :"21212”,
	"userQueryId" : "id_0000001",
	"variables" : [
	    { "id_number" : "110108198703127621" },
		{ "name" :"张三" },
		{ "address":"江苏省南京市江宁区" }
	]
	"nextAction" : {
	    "action" : 5,
		"paras" : {
		    "called" : "外线被叫号码",
			"outNumber" : "呼出总机号码"
		}
	}
	"reason" : "test",
	"userdata" : "test"
}

```

##### 3.2.2.6 转其他IVR流程响应

json to cr

```json
{
    "rspCode" : 0,
    “ccgeid” : “123”,
    “ccNumber” : "21212”,
	"userQueryId" : "id_0000001",
	"variables" : [
	    { "id_number" : "110108198703127621" },
		{ "name" :"张三" },
		{ "address":"江苏省南京市江宁区" }
	]
	"nextAction" : {
	    "action" : 6,
		"paras" : {
		    "ivrFlowId" : "IVR流程id",
			"ivrFlowName" : "IVR流程名称"
		}
	}
	"reason" : "test",
	"userdata" : "test"
}
```

##### 3.2.2.7 流程结束响应

```json
{
    "rspCode" : 0,
    “ccgeid” : “123”,
    “ccNumber” : “21212”,
	"userQueryId" : "id_0000001",
	"variables" : [
	    { "id_number" : "110108198703127621" },
		{ "name" :"张三" },
		{ "address":"江苏省南京市江宁区" }
	]
	"nextAction" : {
	    "action" : 7
	}
	"reason" : "test",
	"userdata" : "test"
}

```



### 3.3 api->cr_web

api到cr_web，只需要将json数据外包装一个头，然后透传给cr_web就可以了

## 4 处理流程

cr给api的推送目前分为通话推送和坐席状态推送。

查询被叫默认按照通话推送进行处理，并且按照振铃推送的处理逻辑进行处理。

之前的查询被叫协议只会出现在呼入场景中，而现在的交互式ivr呼入呼出都有涉及，不仅仅只是查询被叫，同样的会有和通话无关的查询请求，首先考察是否会对通话记录产生影响

- api callpush模块时序图
```sequence
asynccallpush->appcallback: 封装数据交给appcallback处理
appcallback->appcallback: 根据数据类型进行\n不同的数据封装
appcallback->appserver: 将推送数据包发送给客户服务器
appserver->appcallback: 返回数据给appcallback
appcallback->appcallback: 处理客户服务器返回的数据
```

- asynccallpush流程图(status为0的分支)
```flow
st=>start: 开始框
findrecord=>operation: 查询通话记录
ifrecordexist=>condition: 通话记录是否存在？
createnewrecord=>operation: 新建通话记录(准备数据)
completecallpushdata=>operation: 完善callpushdata
ifrecordexist2=>condition: 通话记录是否存在？
chargeold=>operation: 检测是否是已经推送过的detail
ifdetailexist=>condition: detail是否存在
createnewdetail=>operation: 新建detail(准备数据)
ifdatapush=>condition: 是否是通话过程中的内容推送
datapushstatuschange=>operation: 内容推送可能也需要改变通话记录状态
ifstatuscalling=>condition: 通话记录状态为振铃中？
updateringtime=>operation: 更新振铃时间
sub=>subroutine: ....
sub1=>subroutine: ....
subroutine=>operation: 进入不同的子流程

e=>end: 结束

st->findrecord->ifrecordexist
ifrecordexist(yes)->completecallpushdata
ifrecordexist(no)->createnewrecord->completecallpushdata
completecallpushdata->ifrecordexist2
ifrecordexist2(yes)->chargeold->ifdetailexist
ifrecordexist2(no)->ifdetailexist
ifdetailexist(yes)->ifdatapush
ifdetailexist(no)->createnewdetail
ifdatapush(yes)->datapushstatuschange
ifdatapush(no)->ifstatuscalling
ifstatuscalling(yes)->updateringtime
ifstatuscalling(no)->sub
datapushstatuschange->sub1
updateringtime->sub1
createnewdetail->sub1
sub1->subroutine
subroutine->e
```



- 通话振铃推送子流程图
```flow
st=>start: 开始
confirmurl=>operation: 确定推送url
pack=>operation: 打包数据，确定是否进行推送
up_db=>operation: record和detail入库
ifdocallback=>condition: 是否给客户推送
docallback=>operation: 发送至推送处理队列
e=>end: 结束框
st->confirmurl->pack->up_db->ifdocallback
ifdocallback(no)->e
ifdocallback(yes)->docallback->e
```

- appcallback请求消息处理流程
```flow
st=>start: 开始
genquerydata=>operation: 封装请求信息
ifquery=>condition: 是否需要向客户服务器查询数据
commoncallback=>operation: 正常推送流程
getappsetting=>operation: 获取应用配置，并且获取配置对应的回调查询函数
querydata=>operation: 向客户服务器查询数据
notify=>operation: 通过消息通道将消息返回给请求方
e=>end: 结束
st->ifquery
ifquery(yes)->genquerydata->getappsetting->notify->e
ifquery(no)->commoncallback->e
```

经过考察我们主要的修改是：

1. 发向appcallback队列的数据会进行扩展，存储ivr扩展信息
2. 回调接口扩展，新建新的协议和对应的回调函数
3. 查询结果返回

### 4.1 95(交互收键模式)

### 4.2 96(被叫查询模式)

### 4.3 97(通用交互模式)

### 4.4 98(AMNB或者AMXNB模式)

被叫查询api会直接根据消息数据入库，和之前的处理方式保持一致即可。

### 4.5 99(AMB或者AMXB模式)

被叫查询api会直接根据消息数据入库，和之前的处理方式保持一致即可。

## 5 涉及代码

```c
if( data->type == DB_CALL_RECORD_TYPE_QUERY_CALLED_BASE ||
   data->type == DB_CALL_RECORD_TYPE_QUERY_CALLED_BY_PBX ||
   data->type == DB_CALL_RECORD_TYPE_QUERY_CALLED_BY_USENUMBER ||
   data->type == DB_CALL_RECORD_TYPE_QUERY_CALLED_BY_VIRTNUMBER )
{
    strcpy(call_record->useNumber, data->useNumber);
}
else if(data->type == DB_CALL_RECORD_TYPE_QUERY_CALLED_BY_SUBNUMBER)
    strcpy(call_record->subNumber, data->subNumber);
else if(data->type == DB_CALL_RECORD_TYPE_QUERY_CALLED_BY_VIRTNUMBER)
    strcpy(call_record->virtNumber, data->virtNumber);
```

```c
else if(IsOnlineCallPushType(data->type))
{
    if(call_record->status < DB_CALL_RECORD_STATUS_CALL_ESTABLISHED)
        data->status = call_record->status;	//如果是查被叫的话重新改为振铃
    else
        data->status = DB_CALL_RECORD_STATUS_CALL_ESTABLISHED;
}
```
    else if(NULL != callreqUrl ||
            ( call_record->type >= DB_CALL_RECORD_TYPE_QUERY_CALLED_BASE &&
              DB_APPLICATION_CALLBACK_COMMON_PROTOCOL == appInfo->protocol) )
```c
if(call_record->type >= DB_CALL_RECORD_TYPE_QUERY_CALLED_BASE)
{
```

