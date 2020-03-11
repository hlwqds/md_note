# 新架构子ivr交互查询

对于api来说子ivr和交互式ivr不会有任何区别，详情见交互式ivr文档：

[新架构ivr交互查询](./新架构ivr交互查询.md)

### 注意：

子ivr根据文档定义每次查询都会发送type为60的推送消息：

```json
{
	"ccNumber": "cc_pc_101_1007_0001483c_1577157149_7037041", 
    "switchNumber": "902566687523", 
    "eid": 84028, 
    "type": 60, //注意
    "callId": "cc_pc_101_1007_0001483c_1577157149_7037041", 
    "caller": "02566687523", 
    "isCaller": 1, 
    "called": "1007", 
    "number": "1007", 
    "ivrFlowId":101,
    "step": 0, 
    "status": 0, 
    "timestamp": 1577157149
}
```

api目前并不需要处理这个消息，因为api目前是按照振铃处理的查询请求，这个消息会造成details状态混乱，直接忽略就好

```c
//api暂时处理不了的消息
if(CALL_PUSH_NOT_NEED_HANDLE(type))
{
    goto __exit;
}
```