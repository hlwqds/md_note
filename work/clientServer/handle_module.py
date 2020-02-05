import json
import xmltodict

format_json = False

def PlayBackKey(req):
	response_dict = { 
		"retcode": 0,
		"usrQueryId": "id_0002",
		"variables":[
		{"id_number": "110108198703127621"},
		{"name": "huanglin"},
		{"address": "nanjing"}
		],  
		"virtualKey":"1111",
		"nextAction" : { 
			"action" : 2,
			"params" : { 
				"voiceId" :"1",
				"voiceName" : "dwad",
				"allowBreak" : "1" 
			}   
		},  
		"userData":"FE87D3"
	}

	return response_dict

def PlayBack(req):
	response_dict = { 
		"retcode": 0,
		"usrQueryId": "id_0002",
		"variables":[
		{"id_number": "110108198703127621"},
		{"name": "huanglin"},
		{"address": "nanjing"}
		],  
		"virtualKey":"1111",
		"nextAction" : { 
			"action" : 1,
			"params" : { 
				"voiceId" :"1",
				"voiceName" : "dwad",
				"allowBreak" : "1",
				"getKeyNumber":"3",
				"getKeyTimeout":"60",
				"endWithHashKey":"1",
			}
		},  
		"userData":"FE87D3"
	}

	return response_dict

def ToGroup(req):
	response_dict = { 
		"retcode": 0,
		"usrQueryId": "id_0002",
		"variables":[
		{"id_number": "110108198703127621"},
		{"name": "huanglin"},
		{"address": "nanjing"}
		],  
		"virtualKey":"1111",
		"nextAction" : { 
			"action" : 3,
			"params" : { 
				"acdId" :"1",
				"acdName" : "dwad",
				"useAcdValue" : "1",
				"queueTime":40,
				"switchTimes": 32,
				"ringTimeout":10,
				"customerMemory":0,
			}   
		},  
		"userData":"FE87D3"
	}

	return response_dict

def ToSeat(req):
	response_dict = { 
		"retcode": 0,
		"usrQueryId": "id_0002",
		"variables":[
		{"id_number": "110108198703127621"},
		{"name": "huanglin"},
		{"address": "nanjing"}
		],  
		"virtualKey":"1111",
		"nextAction" : { 
			"action" : 4,
			"params" : { 
				"workNumber" :"dwa,da",
				"number" : "112",
				"queueTime" : 4,
				"ringTimeout": 10,
			}   
		},  
		"userData":"FE87D3"
	}

	return response_dict

def ToOutline(req):
	response_dict = { 
		"retcode": 0,
		"usrQueryId": "id_0002",
		"variables":[
		{"id_number": "110108198703127621"},
		{"name": "huanglin"},
		{"address": "nanjing"}
		],  
		"virtualKey":"1111",
		"nextAction" : { 
			"action" : 5,
			"params" : { 
				"called" :"15861800293",
				"outNumber" : "02566687671" 
			}   
		},  
		"userData":"FE87D3"
	}

	return response_dict

def ToOtherIVR(req):
	response_dict = { 
		"retcode": 0,
		"usrQueryId": "id_0002",
		"variables":[
		{"id_number": "110108198703127621"},
		{"name": "huanglin"},
		{"address": "nanjing"}
		],  
		"virtualKey":"1111",
		"nextAction" : { 
			"action" : 6,
			"params" : { 
				"ivrFlowId" :"15861800293",
				"ivrFlowName" : "02566687671" 
			}   
		},  
		"userData":"FE87D3"
	}

	return response_dict

def EndIVR(req):
	response_dict = { 
		"retcode": 0,
		"usrQueryId": "id_0002",
		"variables":[
		{"id_number": "110108198703127621"},
		{"name": "huanglin"},
		{"address": "nanjing"}
		],  
		"virtualKey":"1111",
		"nextAction" : { 
			"action" : 7,
		},  
		"userData":"FE87D3"
	}

	return response_dict

def HandleCalledQuery(req):
	response_dict = { 
		"retcode": 0,
		"action": 0,
		"called": "15861800293",
#		"number":"1111",
#		"workNumber":"1002",
		"waitTime": 20,
		"outNumber": "02566687671",
		"reason": "000",
		"userData":"FE87D3"
	}

	return response_dict



SyncHandleSet = {
	"1": PlayBack,
	"2": PlayBackKey,
	"3": ToGroup,
	"4": ToSeat,
	"5": ToOutline,
	"6": ToOtherIVR,
	"7": EndIVR,
}

a = 0

def SyncHanleWithReq(req):

	global a
	a += 1
	action = str((a % 7) + 1)

	if format_json == True:
		req_dict = json.loads(req)
		typeInt = int(req_dict["request"]["callType"])
		print(typeInt)
		if typeInt == 98 or typeInt == 99:
			response_dict = HandleCalledQuery(req_dict)
		else:
			response_dict = SyncHandleSet[action](req_dict)
		return json.dumps(response_dict)
	else:
		req_dict = xmltodict.parse(req)
		typeInt = int(req_dict["request"]["callType"])
		print(typeInt)
		if typeInt == 98 or typeInt == 99:
			response_dict = HandleCalledQuery(req_dict)
		else:
			response_dict = SyncHandleSet[action](req_dict)
		rsp = {
			"response":response_dict
		}
		return xmltodict.unparse(rsp)

