from api_post_class import *
from sys import exit
from time import time
import json

auth_info = {}
auth_info["accountSid"] = "420d7b0d024b981da73c744ea887f3d2"
auth_info["accountToken"] = "9855e41e2ef2305c720c451540840a7c"
domain = "58.240.254.106:4080"
version = "20161021"
dataType = "json"

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

post_dict98 = {
	"type" : 98,
	"ccgeid" : 243,
	"callId" : str(time()),
	"ccNumber" : str(time()),
	"callType" : 98,
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


p = ApiPost(auth_info, domain, version, dataType)
pushup(p, post_dict98)
