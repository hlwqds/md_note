import json
import xmltodict
from time import time
import logging
import http_log

a = 0

def HandleCalledQuery(req):
	response_dict = { 
		"retcode": 0,
		"action": 0,
#		"called": "15861800293",
		"number":"2133",
#		"workNumber":"1002",
		"waitTime": 20,
		"outNumber": "02566687671",
		"reason": "000",
		"userData":"FE87D3"
	}

	return response_dict


def HandleGroupQuery(req):
	response_dict = { 
		"retcode": 0,
		"action": 1,
		"transferGroup": "496",
		"reason": "000",
		"userData":"FE87D3"
	}

	rsp = {
		"response":response_dict
	}
	return rsp

def HandleNumberQuery(req):
	response_dict = { 
		"retcode": 0,
		"action": 0,
		"number": "6666",
		"workNumber": "",
		"called": "",
		"waitTime": 20,
		"outNumber": "02566699794",
		"reason": "000",
		"userData":"FE87D3"
	}

	rsp = {
		"response":response_dict
	}
	return rsp

def HandleCommonReq(req):
	response_dict = { 
		"retcode": 0,
	}

	return response_dict

def NextActionPlayBack():
	nextActionDict = { 
		"action" : 1,

#		"params" : { 
#			"voiceId" :"3",
#			"voiceName" : "放音指定名称",
#			"allowBreak" : "1",
#		}
		"params" : {
			#"voiceTempId" :"56",
			"voiceTempName" : "语音模板测试",
			"voiceTempParams":"test1,test2",
			"allowBreak" : "1",
		}
	}
	return nextActionDict

def NextActionPlayBackKey():
	nextActionDict = { 
		"action" : 2,
		"params" : { 
#			"voiceId" :"3",
			"voiceName" : "放音指定名称",
			"allowBreak" : "1",
			"getKeyNumber":"3",
			"getKeyTimeout":"60",
			"endWithHashKey":"1",
		}   
	}

	return nextActionDict

def NextActionToGroup():
	nextActionDict = { 
		"action" : 3,
		"params" : { 
			"acdId" :"1000002855",
			"acdName" : "dwad",
			"useAcdValue" : "1",
			"queueTime":40,
			"switchTimes": 32,
			"ringTimeout":10,
			"customerMemory":0,
		}   
	}

	return nextActionDict

def NextActionToSeat():
	nextActionDict = { 
		"action" : 4,
		"params" : { 
			"workNumber" :"dwa,da",
			"number" : "2133",
			"queueTime" : 4,
			"ringTimeout": 10,
		}   
	}

	return nextActionDict

def NextActionToOutline():
	nextActionDict = {
		"action" : 5,
		"params" : { 
			"called" :"15861800293",
			"outNumber" : "02566687671" 
		}   
	}

	return nextActionDict

def NextActionToOtherIVR():
	nextActionDict = { 
		"action" : 6,
		"params" : { 
#			"ivrFlowId" :"6640",
			"ivrFlowName" : "转其他IVR流程" 
		}   
	}

	return nextActionDict

def NextActionEndIVR():
	nextActionDict = {
		"action" : 7,
	}

	return nextActionDict

NextActionSet = {
	"1": NextActionPlayBack,
	"2": NextActionPlayBackKey,
	"3": NextActionToGroup,
	"4": NextActionToSeat,
	"5": NextActionToOutline,
	"6": NextActionToOtherIVR,
	"7": NextActionEndIVR,
}

def GenVariablesJSON(req, response_dict):
	try:
		response_dict["variables"] = []
		for var in req["request"]["variables"]:
				for k, v in var.items():
					logging.info(k + ": " + v)
					response_dict["variables"].append({k:int(time())})
	except Exception as e:
		logging.info(e)
		response_dict["variables"] = []
	return response_dict

def GenVariablesXMLOld(req, response_dict):
	try:
		response_dict["variables"] = {}
		if hasattr(req["request"]["variables"], "items"):
			variableList = []
			for k, v in req["request"]["variables"].items():
				response_dict["variables"][k] = int(time())
#		else:
#			logging.info(req["request"]["variables"])
#			for var in req["request"]["variables"]:
#					for k, v in var.items():
#						variableList.append({k:int(time())})
#			response_dict["variables"] = variableList
	except Exception as e:
		logging.info(e)
		response_dict["variables"] = []
	return response_dict

def GenVariablesXML(req, response_dict):
	try:
		response_dict["variable"] = []
		logging.info(req["request"]["variable"])
		if hasattr(req["request"]["variable"], "items"):
			#只有一个元素，被解释为了字典
			variable_dic = {}
			variable_dic["name"] = req["request"]["variable"]["name"]
			variable_dic["value"] = int(time())
			response_dict["variable"].append(variable_dic)
		else:
			#有多个元素，被解释为列表
			for var in req["request"]["variable"]:
				for k, v in var.items():
					variable_dic = {}
					logging.info(k + ": " + str(v))
					if k == "name":
						variable_dic["name"] = v
					elif k == "value":
						variable_dic["value"] = int(time())
					else:
						continue
				response_dict["variable"].append(variable_dic)
	except Exception as e:
		logging.info(e)
		response_dict["variable"] = []
	return response_dict

def GenVariables(req, response_dict, ifJson):
	if ifJson:
		response_dict = GenVariablesJSON(req, response_dict)
	else:
		response_dict = GenVariablesXMLOld(req, response_dict)
	return response_dict

def HandleKeyBackQuery(req, ifJson):
	response_dict = { 
		"retcode": 0,
		"usrQueryId": req["request"]["usrQueryId"],
		"virtualKey":"1",

		"userData":"FE87D3"
	}
	response_dict = GenVariables(req, response_dict, ifJson)
	global a
	a += 1
	action = str((a % 7) + 1)
	nextActionDict = NextActionSet[action]()
	response_dict["nextAction"] = nextActionDict
	return response_dict

def HandleCommonIVRQuery(req, ifJson):
	response_dict = { 
		"retcode": 0,
		"usrQueryId": req["request"]["usrQueryId"],
		"userData":"FE87D3"
	}
	response_dict = GenVariables(req, response_dict, ifJson)
	global a
	a += 1
	action = "2"
	nextActionDict = NextActionSet[action]()
	response_dict["nextAction"] = nextActionDict
	return response_dict

def genResponseByRequestDic(req_dict, ifJson):
	callTypeInt = int(req_dict["request"]["callType"])
	typeInt = int(req_dict["request"]["type"])

	if callTypeInt == 98 or callTypeInt == 99:
		response_dict = HandleCalledQuery(req_dict)
	elif typeInt == 95:
		response_dict = HandleKeyBackQuery(req_dict, ifJson)
	elif typeInt == 96 or typeInt == 97:
		response_dict = HandleCommonIVRQuery(req_dict, ifJson)
	else:
		response_dict = HandleCommonReq(req_dict)
	rsp = {
		"response":response_dict
	}
	return rsp

def SyncHanleWithReq(req):
	format_json = False

	if format_json == True:
		req_dict = json.loads(req)

		response_dict = genResponseByRequestDic(req_dict, True)

		return json.dumps(response_dict) 
	else:
		req_dict = xmltodict.parse(req)
		response_dict = genResponseByRequestDic(req_dict, False)

		return xmltodict.unparse(response_dict)

def OldSyncHanleWithReqForGroup(req):
	format_json = True
	if format_json == True:
		req_dict = json.loads(req)

		response_dict = HandleGroupQuery(req_dict)

		return json.dumps(response_dict) 
	else:
		req_dict = xmltodict.parse(req)
		response_dict = HandleGroupQuery(req_dict)

		return xmltodict.unparse(response_dict)

def OldSyncHanleWithReqForNumber(req):
	format_json = True
	if format_json == True:
		req_dict = json.loads(req)

		response_dict = HandleNumberQuery(req_dict)

		return json.dumps(response_dict) 
	else:
		req_dict = xmltodict.parse(req)
		response_dict = HandleNumberQuery(req_dict)

		return xmltodict.unparse(response_dict)
