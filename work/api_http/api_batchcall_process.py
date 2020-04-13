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

commonWorkNumber = "6677"
commonType = 1
commonDevice = '2104'
#commonType = 0
#commonDevice = ''

monitorWorkNumber = "hhhlll4"
monitorType = 1
monitorDevice = '2150'
#monitorType = 0
#monitorDevice = ''

transferWorkNumber = "1012"
transferType = 0
transferDevice = ''

to = "15861800293"
#to = "15551208601"
#gid = '1000002855'


f = ApiPost(auth_info, domain, version, dataType)
task_id = input("please input taskid:")
operation = "BatchCalls"
function = "addNewBatch"
params = {
    "taskId":task_id,
    "unique_task_ids":"",
#   "gid":gid,
    "tels":"15861800293",
}
authType = 1

f.set_func_mode(operation,function,authType)
rsp = f.post_info_to_api(params)
print(rsp)
rsp_dic = json.loads(rsp)
try:
    if(rsp_dic['resp']['respCode'] != 0):
        print("Error, please check and retry")
except Exception as e:
    print(e)
    print("Error, please check and retry")
