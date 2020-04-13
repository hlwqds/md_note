#/usr/bin/python3
from api_post_class import *
from sys import exit
from time import *
import json

auth_info = {}
auth_info["accountSid"] = "420d7b0d024b981da73c744ea887f3d2"
auth_info["accountToken"] = "9855e41e2ef2305c720c451540840a7c"
domain = "58.240.254.106:4080"
version = "20161021"
dataType = "json"
seat = "2143"
switchNumber = "02566687671"
client = "15861800293"

def pushup(f, post_dict):
	except_num = 0
	while(True):
		operation = "CallPush"
		function = ""
		authType = 0
		f.set_func_mode(operation,function,authType)
		rsp = f.post_info_to_api(post_dict)
		print(rsp)
		rsp_dic = json.loads(rsp)
		try:
			if(rsp_dic['resp']['respCode'] != 0):
				print("Error, please check and retry")
				except_num += 1
				if(except_num >= 3):
					print("Max except num,break")
					return False
				continue
		except Exception as e:
			print(e)
			return False
		break
	return True

class CommonPostModel():
	def __init__(self, callId, ccNumber, eId, switchNumber, isCcgeid=False):
		self.status = 0
		self.step = 0
		self.ccNumber = ccNumber
		self.ccgeid = self.eid = 0

		if isCcgeid:
			self.ccgeid = eId
		else:
			self.eid = eId
			
		self.switchNumber = switchNumber
		self.type = 0	#在逻辑中，这个type是会随着请求类型改变而改变的，通话中不会改变的另一个属性应该是calltype，但是以前并没有这个字段
		self.callId = callId
		self.caller = ""
		self.isCaller = 0
		self.number = ""
		self.called = ""
		self.timestamp = 0
	def _genFormatDict(self):
		formatDict = {}
		formatDict["ccNumber"] = self.ccNumber
		if self.ccgeid != 0:
			formatDict["ccgeid"] = self.ccgeid
		else:
			formatDict["eid"] = self.eid
		formatDict["type"] = self.type
		formatDict["callId"] = self.callId
		formatDict["switchNumber"] = self.switchNumber
		formatDict["caller"] = self.caller
		formatDict["isCaller"] = self.isCaller
		formatDict["number"] = self.number
		formatDict["step"] = self.step
		formatDict["called"] = self.called
		formatDict["status"] = self.status
		formatDict["timestamp"] = self.timestamp
		return formatDict

	def genCallingInfo(self, step, type, caller, isCaller, number, called, timestamp):
		self.status = 0	#振铃请求的status为0
		self.step = step
		self.type = type
		self.caller = caller
		self.isCaller = isCaller
		self.number = number
		self.called = called
		self.timestamp = timestamp
  
		return self._genFormatDict()

	def genQueryCalledInfo(self, step, type, caller, isCaller, number, called, timestamp):
		self.status = 0	#振铃请求的status为0
		self.step = step
		self.type = type	#查被叫请求的可能是98或99
		self.caller = caller
		self.isCaller = isCaller
		self.number = number
		self.called = called
		self.timestamp = timestamp

		return self._genFormatDict()

	def genEstablishedInfo(self, step, type, caller, isCaller, number, called, timestamp):
		self.status = 2	#建立请求的status为2
		self.step = step
		self.type = type
		self.caller = caller
		self.isCaller = isCaller
		self.number = number
		self.called = called
		self.timestamp = timestamp

		return self._genFormatDict()

	def genHangupInfo(self, step, type, caller, isCaller, number, called, timestamp):
		self.status = 3	#挂机请求的status为3
		self.step = step
		self.type = type
		self.caller = caller
		self.isCaller = isCaller
		self.number = number
		self.called = called
		self.timestamp = timestamp

		return self._genFormatDict()

	def genFaliedInfo(self, step, type, caller, isCaller, number, called, timestamp):
		self.status = 1	#失败请求的status为1
		self.step = step
		self.type = type
		self.caller = caller
		self.isCaller = isCaller
		self.number = number
		self.called = called
		self.timestamp = timestamp
		return self._genFormatDict()

class CommonIvrPostModel(CommonPostModel):
	def __init__(self, callId, ccNumber, ccgeid, switchNumber, callType, ivrFlowId):
		self.callType = callType
		self.ivrFlowId = ivrFlowId
		self.ivrQueryId = 0
		self.callerType = 0
		self.calledType = 0
		self.usrQueryId = 0
		self.inputKeys = ""
		self.variables = []
		super().__init__(callId, ccNumber, ccgeid, switchNumber, True)

	def _genCommonIvrFormatDict(self):
		dic = self._genFormatDict()
		dic["callType"] = self.callType
		dic["ivrFlowId"] = self.ivrFlowId
		dic["ivrQueryId"] = self.ivrQueryId
		dic["callerType"] = self.callerType
		dic["calledType"] = self.calledType
		dic["usrQueryId"] = self.usrQueryId
		dic["inputKeys"] = self.inputKeys
		dic["variables"] = self.variables
		return dic

	def genCommonIvrInfo(self, step, caller, callerType, isCaller, number, called, calledType,
		timestamp, ivrQueryId, usrQueryId, inputKeys, variables):
		self.ivrQueryId = ivrQueryId
		self.usrQueryId = usrQueryId
		self.inputKeys = inputKeys
		self.variables = variables
		self.type = 97	#通用交互式ivr类型是97
		self.step = step
		self.caller = caller
		self.callerType = callerType
		self.isCaller = isCaller
		self.number = number
		self.called = called
		self.calledType = calledType
		self.timestamp = timestamp
		return self._genCommonIvrFormatDict()

	def genKeyBackInfo(self, step, caller, callerType, isCaller, number, called, calledType,
		timestamp, ivrQueryId, usrQueryId, inputKeys, variables):
		self.ivrQueryId = ivrQueryId
		self.usrQueryId = usrQueryId
		self.inputKeys = inputKeys
		self.variables = variables
		self.type = 95	#按键反馈ivr类型是95
		self.step = step
		self.caller = caller
		self.callerType = callerType
		self.isCaller = isCaller
		self.number = number
		self.called = called
		self.calledType = calledType
		self.timestamp = timestamp
		return self._genCommonIvrFormatDict()

def initPostModel():
	postModel = {
		"ccNumber": "",
		"ccgeid": 0,
		"type": 0,
		"callId": "",
		"switchNumber":"",
		"caller": "",
		"isCaller": 0,
		"number": "",
		"step": 0,
		"called": "",
		"status": 0,
		"timestamp": int(time())
	}
	return postModel

def initPostCalling(ccNumber, callId, switchNumber, type, ccgeid, caller, isCaller, called, number, step):
	postModel = initPostModel()
	postModel["ccNumber"] = ccNumber
	postModel["callId"] = callId
	postModel["type"] = type
	postModel["ccgeid"] = ccgeid
	postModel["caller"] = caller
	postModel["isCaller"] = isCaller
	postModel["number"] = number
	postModel["step"] = step
	postModel["status"] = 0
	postModel["switchNumber"] = switchNumber
	postModel["called"] = called
	return postModel

def initPostEstablish(ccNumber, callId, switchNumber, type, ccgeid, caller, isCaller, called, number, step):
	postModel = initPostModel()
	postModel["ccNumber"] = ccNumber
	postModel["callId"] = callId
	postModel["type"] = type
	postModel["ccgeid"] = ccgeid
	postModel["caller"] = caller
	postModel["isCaller"] = isCaller
	postModel["number"] = number
	postModel["step"] = step
	postModel["status"] = 2
	postModel["switchNumber"] = switchNumber
	postModel["called"] = called

	return postModel

def initPostHangup(ccNumber, callId, switchNumber, type, ccgeid, caller, isCaller, called, number, step):
	postModel = initPostModel()
	postModel["ccNumber"] = ccNumber
	postModel["callId"] = callId
	postModel["type"] = type
	postModel["ccgeid"] = ccgeid
	postModel["caller"] = caller
	postModel["isCaller"] = isCaller
	postModel["number"] = number
	postModel["step"] = step
	postModel["status"] = 3
	postModel["switchNumber"] = switchNumber
	postModel["called"] = called

	return postModel



post_dict97 = {
	"type" : 97,
	"ccgeid" : 243,
	"callId" : str(time()),
	"ccNumber" : str(time()),
	"callType" : 0,
	"ivrFlowId" : 245,
	"ivrQueryId" : 1,
	"caller" : "13260278209",
	"callerType" : 1,
	"switchNumber" : "02566687671",
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

post_dict95 = {
	"type" : 95,
	"ccgeid" : 243,
	"callId" : str(time()),
	"ccNumber" : str(time()),
	"callType" : 0,
	"ivrFlowId" : 245,
	"ivrQueryId" : 1,
	"caller" : "13260278209",
	"callerType" : 1,
	"switchNumber" : "02566687671",
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

post_dict99 = {
	"type" : 99,
	"callId" : str(time()),
	"ccNumber" : str(time()),
	"caller" : "15861800293",
	"switchNumber" : "02566699794",
	
	"timestamp" : "1519439787",
}

def callinCrIvrTest():
	ccgeid = 243
	ccNumber = callId = str(time())
	type = 5	#以前的推送的缺陷导致通话类型和中间的请求类型没有分开,所以这里type是通话类型
	caller = "15861800293"
	switchNumber = "02566687671"
	isCaller = 1
	number = "2133"
	called = "2133"
	step = 0

	global auth_info, domain, version, dataType
	p = ApiPost(auth_info, domain, version, dataType)
	postReqCaller = CommonPostModel(callId, ccNumber, ccgeid, switchNumber, True)
	pushup(p, postReqCaller.genQueryCalledInfo(step, 99, caller, isCaller, number, called, int(time())))

	callType = 5
	ivrFlowId = 1
	ivrQueryId = 1
	usrQueryId = 1
	inputKeys = 1
	callerType = 1
	calledType = 1
	variables = [{"725":""}, {"111":"10010"}]
	ivr = CommonIvrPostModel(callId, ccNumber, ccgeid, switchNumber, callType, ivrFlowId)
	dic = ivr.genCommonIvrInfo(step, caller, callerType, isCaller, number, called, calledType,
			int(time()), ivrQueryId, usrQueryId, inputKeys, variables)
	pushup(p, dic)

	sleep(2)

	ivr = CommonIvrPostModel(callId, ccNumber, ccgeid, switchNumber, callType, ivrFlowId)
	dic = ivr.genKeyBackInfo(step, caller, callerType, isCaller, number, called, calledType,
			int(time()), ivrQueryId, usrQueryId, inputKeys, variables)
	pushup(p, dic)
	sleep(2)

	pushup(p, postReqCaller.genCallingInfo(step, type, caller, isCaller, number, called, int(time())))
	sleep(2)

	pushup(p, postReqCaller.genEstablishedInfo(step, type, caller, isCaller, number, called, int(time())))
	sleep(2)

	pushup(p, postReqCaller.genHangupInfo(step, type, caller, isCaller, number, called, int(time())))

	'''
	pushup(p, post_dict98)

	postReqCaller = initPostCalling(ccNumber, callId, switchNumber, type, ccgeid, caller, isCaller, called, number, step) 
	pushup(p, postReqCaller)

	postReqCaller = initPostEstablish(ccNumber, callId, switchNumber, type, ccgeid, caller, isCaller, called, number, step)
	pushup(p, postReqCaller)

	postReqCaller = initPostHangup(ccNumber, callId, switchNumber, type, ccgeid, caller, isCaller, called, number, step)
	pushup(p, postReqCaller)
	'''
 
def callinCrCommonTest():
	ccgeid = 243
	ccNumber = callId = str(time())
	type = 5	#以前的推送的缺陷导致通话类型和中间的请求类型没有分开,所以这里type是通话类型
	caller = "15861800293"
	switchNumber = "02566687671"
	isCaller = 1
	number = "2133"
	called = "2133"
	step = 0

	global auth_info, domain, version, dataType
	p = ApiPost(auth_info, domain, version, dataType)

	postReqCaller = CommonPostModel(callId, ccNumber, ccgeid, switchNumber, True)
	pushup(p, postReqCaller.genCallingInfo(step, type, caller, isCaller, number, called, int(time())))
	pushup(p, postReqCaller.genEstablishedInfo(step, type, caller, isCaller, number, called, int(time())))
	pushup(p, postReqCaller.genHangupInfo(step, type, caller, isCaller, number, called, int(time())))

	'''
	postReqCaller = initPostCalling(ccNumber, callId, switchNumber, type, ccgeid, caller, isCaller, called, number, step) 
	pushup(p, postReqCaller)

	postReqCaller = initPostEstablish(ccNumber, callId, switchNumber, type, ccgeid, caller, isCaller, called, number, step)
	pushup(p, postReqCaller)

	postReqCaller = initPostHangup(ccNumber, callId, switchNumber, type, ccgeid, caller, isCaller, called, number, step)
	pushup(p, postReqCaller)
	'''
 
def callinPbxCommonTest():
	eid = 65761
	ccNumber = callId = str(time())
	type = 5	#以前的推送的缺陷导致通话类型和中间的请求类型没有分开,所以这里type是通话类型
	caller = "15861800293"
	switchNumber = "02566699734"
	isCaller = 1
	number = "1003"
	called = "1003"
	step = 0

	global auth_info, domain, version, dataType
	p = ApiPost(auth_info, domain, version, dataType)

	post_dict99["eid"] = 65761
	post_dict99["callId"] = callId
	post_dict99["ccNumber"] = ccNumber
	pushup(p, post_dict99)

	postReqCaller = CommonPostModel(callId, ccNumber, eid, switchNumber, False)
	pushup(p, postReqCaller.genCallingInfo(step, type, caller, isCaller, number, called, int(time())))
	pushup(p, postReqCaller.genEstablishedInfo(step, type, caller, isCaller, number, called, int(time())))
	pushup(p, postReqCaller.genHangupInfo(step, type, caller, isCaller, number, called, int(time())))

def callinTransferTest():
	eid = 65761
	ccNumber = callId = str(time())
	type = 5	#以前的推送的缺陷导致通话类型和中间的请求类型没有分开,所以这里type是通话类型
	caller = "15861800293"
	switchNumber = "02566699734"
	isCaller = 0
	number = "1003"
	called = "1003"
	step = 0

	global auth_info, domain, version, dataType
	p = ApiPost(auth_info, domain, version, dataType)

	postReqCaller = CommonPostModel(callId, ccNumber, eid, switchNumber, False)
	pushup(p, postReqCaller.genCallingInfo(step, type, caller, isCaller, number, called, int(time())))
	pushup(p, postReqCaller.genFaliedInfo(step, type, caller, isCaller, number, called, int(time())))
	number = "1004"
	called = "1004"
	step += 1
	pushup(p, postReqCaller.genCallingInfo(step, type, caller, isCaller, number, called, int(time())))
	pushup(p, postReqCaller.genFaliedInfo(step, type, caller, isCaller, number, called, int(time())))

	number = "1003"
	called = "1003"
	step += 1
	pushup(p, postReqCaller.genCallingInfo(step, type, caller, isCaller, number, called, int(time())))
	pushup(p, postReqCaller.genFaliedInfo(step, type, caller, isCaller, number, called, int(time())))

	number = "1004"
	called = "1004"
	step += 1
	pushup(p, postReqCaller.genCallingInfo(step, type, caller, isCaller, number, called, int(time())))
	pushup(p, postReqCaller.genEstablishedInfo(step, type, caller, isCaller, number, called, int(time())))
	pushup(p, postReqCaller.genHangupInfo(step, type, caller, isCaller, number, called, int(time())))

#callinCrIvrTest()
#callinCrCommonTest()
#callinPbxCommonTest()
callinTransferTest()