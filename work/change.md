# 工作日志

# 2019/12/24以前

### 1、错误码新增

```c
HTTP_API_RESPONSE_CM_SEATS_BELONGED_GROUP_COUNT_LIMITED,    //坐席归属的技能组数量限制 102340
HTTP_API_RESPONSE_CM_GROUP_CONTAINING_SEAT_COUNT_LIMITED,   //技能组容纳坐席数量限制 102341
```
```c
HTTP_API_RESPONSE_CALLCENTER_SEAT_MOBILE_IN_CALL,  //移动坐席在通话中 102527
HTTP_API_RESPONSE_CALLCENTER_RE_LISTEN,  //重复监听 102528
```
## 2019/12/24

尝试向janus发送消息，但是返回invalid URI

更改sofia-sip的日志等级，尝试深入日志看下问题出在哪里

看不出问题，于是到网上查找解决方法，开发者回复可能需要最新版本或者初始化url_type为0。官网已经不再维护，于是到github找到了社区维护的最新版本。

现在invite消息也成功发出，可以调试接下来的流程。目前我都是向我们公司的测试用sip_proxy发送的sip消息，接下来需要自己搭建一个sip_proxy，决定使用freeswitch

## 2019/12/25

freeswitch已经成功安装，但是还是有不少报错，需要对进行相关的配置。

