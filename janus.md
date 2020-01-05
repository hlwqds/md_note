# janus交互流程
我们的目的是通过janus实现网页工具拨打电话，通过janus消息转发，我们不必知道sip的内部交互方式就可以实现媒体流的传输
虽然使用janus的客户端不需要知道这些信息，但是对于维护janus服务器的人员，其中牵扯到的消息类型和交互协议需要有一定的了解

## 1 与janus的交互流程
### 1.1 janus的注册
```sequence
client->janus: 发送create请求
janus->janus: 创建新的janus会话
janus->client: 返回session_id
client->janus: 发视attach请求，使用特定的模块
janus->janus: 创建新的plugin会话
janus->client: 返回plugin_id
```
至此客户端已在janus的plugin中注册了信息，接下来的交互将针对plugin进行。需要指定session_id和plugin_id来使plugin验证通话。
### 1.2 通过janus与plugin进行交互
```sequence
client->janus: 发送message,消息格式由plugin规定\n交互时需要携带session_id和plugin_id
janus->janus: 解析并且处理，将\n处理结果保存在队列中
janus->client: 返回成功
client->janus: 发送get请求
janus->client: 从队列中取出一条处理结果返回
```
## 2 与plugin的消息交互
目前我希望做到的是利用janus现有的sip插件实现客户端外呼，因此只涉及到janus的sip demo

### 2.1 与sip demo的消息交互
消息格式为json
{
	"janus": "message",
	"body": {
        "request": "register"	//请求事件
        ......	//扩展内容
	},	//消息载体
	"transaction": transaction,		//标识这次请求
}

#### 2.1.1 register(sip register)
data
{
	"request": "register",	//注册
	"username": username,	//
	“authuser”: authuser,	//
	"display_name": display_name,
	"secret": secret,
	"proxy": proxy,
}

sip demo在收到请求后会进一步将注册消息转换成sip消息并且将消息发往sip proxy

#### 2.1.2 call(sip invite)
data
{
	"body":{
		"request": "call",
		"uri":"sip:2150_00014748@cc-jskf02-cr-test-sp.emic.com.cn:9500"
	}
	"jsep_dict" : {
		  "sdp":"v=0\r\n"\
          "o=mozilla...THIS_IS_SDPARTA-71.0 193135304282057553 0 IN IP4 0.0.0.0\r\n"\
          "s=-\r\n"\
          "t=0 0\r\n"\
          "a=fingerprint:sha-256 7A:59:CD:7A:79:8A:3C:E4:A0:2D:D7:59:63:3D:17:FD:85:F1:5D:8D:94:3C:60:C3:82:52:63:B2:96:BC:91:FB\r\n"\
          "a=group:BUNDLE 0\r\n"\
          "a=ice-options:trickle\r\n"\
          "a=msid-semantic:WMS *\r\n"\
          "m=audio 9 UDP/TLS/RTP/SAVPF 109 9 0 8 101\r\n"\
          "c=IN IP4 0.0.0.0\r\n"\
          "a=sendrecv\r\n"\
          "a=extmap:1 urn:ietf:params:rtp-hdrext:ssrc-audio-level\r\n"\
          "a=extmap:2/recvonly urn:ietf:params:rtp-hdrext:csrc-audio-level\r\n"\
          "a=extmap:3 urn:ietf:params:rtp-hdrext:sdes:mid\r\n"\
          "a=fmtp:109 maxplaybackrate=48000;stereo=1;useinbandfec=1\r\n"\
          "a=fmtp:101 0-15\r\n"\
          "a=ice-pwd:printf`1\r\n"\
		  "a=ice-ufrag:huanglin\r\n"\
          "a=mid:0\r\n"\
          "a=msid:{96e3a11e-ecda-499a-ad79-fc9216d99d2b} {16c327d4-157a-4d08-8e63-e05a89031dc2}\r\n"\
          "a=rtpmap:109 opus/48000/2\r\n"\
          "a=rtpmap:9 G722/8000/1\r\n"\
          "a=rtpmap:0 PCMU/8000\r\n"\
          "a=rtpmap:8 PCMA/8000\r\na=rtpmap:101 telephone-event/8000/1\r\n"\
          "a=setup:actpass\r\n"\
          "a=ssrc:743106987 cname:{2ccb8d2c-e04b-4fb0-a27b-42084a25bc58}\r\n",
          "type":"offer"
}

##### 2.1.2.1 sdp消息解析：
1. v=0
   sdp版本号，一直为0,rfc4566规定
   
2. o=mozilla...THIS_IS_SDPARTA-71.0 193135304282057553 0 IN IP4 0.0.0.0
    RFC 4566 o=<username> <sess-id> <sess-version> <nettype> <addrtype> <unicast-address>
    username如果没有就用-代替
    sess-id是整个会话的编号
    sess-version指的是会话的版本，如果会话中sdp发生改变，则需要将sess-version+1
    nettype指的是网络类型，这里指的是IN互联网
    addrtype指的是地址类型
    unicast-address指的是具体地址

3. s=-
    会话名称，没有就用-代替

4. t=0 0
    起始时间和结束时间，0 0表示没有限制

5. a=fingerprint:sha-256 7A:59:CD:7A:79:8A:3C:E4:A0:2D:D7:59:63:3D:17:FD:85:F1:5D:8D:94:3C:60:C3:82:52:63:B2:96:BC:91:FB
    这行是dtls协商过程中需要的认证信息
    dtls即数据包传输层安全性协议(Datagram Transport Layer Security)，保证udp传输过程中数据的安全性

6. a=group:BUNDLE 0
    需要共用一个传输通道传输的媒体，如果没有这一行，音视频，数据就会分别单独用一个udp端口来发送，示例：将audio、video和data共用一个传输通道  a=group:BUNDLE audio video data

7. a=ice-options:trickle

8. a=ice-ufrag:khLS

   a=ice-pwd:cxLzteJaJBou3DspNaPsJhlQ

   ice协商过程中所需要的安全验证信息

9. 