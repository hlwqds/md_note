from api_post_class import *
import json
from sys import exit

auth_info = {}
auth_info["accountSid"] = "42f7de84ff2ea9b4d71c2aa667455249"
auth_info["accountToken"] = "f6233e72475a6d72a90d624d051b6955"
auth_info["subAccountSid"] = "51d17d870e0c1da85869580195b31f1d"
auth_info["subAccountToken"] = "f36aca8093470104d3fb5a88958c1875"
auth_info["appId"] = "68a87b1250011b35cfb244f04a27ca9d"
domain = "58.240.254.106:4080"
version = "20161021"
dataType = "json"

commonWorkNumber = "8899"
#commonType = 1
#commonDevice = '2104'
commonType = 0
#commonDevice = '18955394329'
commonDevice = ''

monitorWorkNumber = "hhhlll4"
#monitorType = 1
#monitorDevice = '2150'
monitorType = 0
#monitorDevice = '15861800293'
monitorDevice = ''

transferWorkNumber = "zt"
transferType = 0
transferDevice = ''

#to = "15861800293"
to = "15551208601"
#to = "15861800293"
#gid = '1000002855'

callout = True
testMonitor = True
testTransfer = False

print("test monitor\n")
#相同操作应该封装成一个函数，但是先这样吧
def seatSignIn(f, worNumber, device, deviceType):
    except_num = 0
    while(True):
        operation = "CallCenter"
        function = "signIn"
        authType = 1
        params = {"workNumber":worNumber}
        params["deviceNumber"] = device
        params["type"] = deviceType
        f.set_func_mode(operation,function,authType)
    
        rsp = f.post_info_to_api(params)
        print(rsp)
        rsp_dic = json.loads(rsp)
        try:
            if(rsp_dic['resp']['respCode'] != 0):
                print("Error, please check and retry")
                except_num += 1
                if(except_num >= 3):
                    print("Max except num,break")
                    exit(0)
                continue
        except Exception as e:
            print(e)
            print("Error, please check and retry")

        break
    return 0

def seatCallout(f, workNumber, to, outNumber):
    except_num = 0
    while(callout):
#    gid = input("please input gid, R to retry:")
#    to = input("please input customer mobile:. R to retry:")
        operation = "CallCenter"
        function = "callOut"
        params = {
            "workNumber":commonWorkNumber,
            "to":to,
#           "gid":gid,
		    "outNumber":outNumber,
        }
        authType = 1

        f.set_func_mode(operation,function,authType)
        rsp = f.post_info_to_api(params)
        print(rsp)
        rsp_dic = json.loads(rsp)
        try:
            if(rsp_dic['resp']['respCode'] != 0):
                print("Error, please check and retry")
                except_num += 1
                if(except_num >= 3):
                    print("Max except num,break")
                    exit(0)
                continue
            callId = rsp_dic['resp']['callOut']['callId']
        except Exception as e:
            print(e)
            print("Error, please check and retry")
        break
    return callId

def seatNextOperate(f):
    while(True):
        func_dic = {
	        "0":"callMonitor",
            "1":"callInsert",
            "2":"callBreak",
            "3":"callIntercept"
        }
        print(func_dic)
        
        operation = "CallCenter"
        authType = 1
        print("please input your operation")
        func_type = input()
        try:
            function = func_dic[str(func_type)]
            print(function)
		
            if(int(func_type) in [0, 1, 2, 3]):
                params = {
                    "toWorkNumber":commonWorkNumber,
                    "workNumber":monitorWorkNumber
                }
#        elif(func_type == 4):
        
            f.set_func_mode(operation,function,authType)
            rsp = f.post_info_to_api(params)
            print(rsp)
        except Exception as e:
            print("Please input the correct type:")



f = ApiPost(auth_info, domain, version, dataType)

print("common seat sign in:")
seatSignIn(f, commonWorkNumber, commonDevice, commonType)

if(testMonitor):
    print("monitor seat sign in:")
    seatSignIn(f, monitorWorkNumber, monitorDevice, monitorType)

if(testTransfer):
    print("transfer seat sign in:")
    seatSingIn(f, transferWorkNumber, transferDevice, transferType)

if(callout):
    print("seat callout:")
    callId = seatCallout(f, commonWorkNumber, to, "02566687671")

print("wait for calling establish")
print("monitorseat function:")

while(True):
    func_dic = {
	    "0":"callMonitor",
        "1":"callInsert",
        "2":"callBreak",
        "3":"callIntercept"
    }
    print(func_dic)
    operation = "CallCenter"
    authType = 1

    print("please input your operation")
    func_type = input()
    try:
        function = func_dic[str(func_type)]
        print(function)
		
        if(int(func_type) in [0, 1, 2, 3]):
            params = {
                "toWorkNumber":commonWorkNumber,
                "workNumber":monitorWorkNumber
            }
#        elif(func_type == 4):
        
        f.set_func_mode(operation,function,authType)
        rsp = f.post_info_to_api(params)
        print(rsp)
    except Exception as e:
        print("Please input the correct type:")


