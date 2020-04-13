from urllib3 import *
import time
from my_hash import my_md5,my_base64
import json


class ApiPost():
    def __init__(self,auth_info,domain,softVersion,dataType):
        self.isCallPush = False
        try:
            self.accountSid = auth_info["accountSid"]
            self.accountToken = auth_info["accountToken"]
            if "subAccountSid" in auth_info:
                self.subAccountSid = auth_info["subAccountSid"]
                self.subAccountToken = auth_info["subAccountToken"]
            if "appId" in auth_info:
                self.appId = auth_info["appId"]
            else:
                self.appId = ""
                self.isCallPush = True
        except Exception as e:
            print("auth init error")
            print(e)
        self.domain = domain
        self.softVersion = softVersion
        self.dataType = dataType
        self.authType = 0

    def _gen_sig(self,unix_time_stamp):
        if(self.authType):
            tmp = self.subAccountSid + self.subAccountToken + str(unix_time_stamp)
        else:
	        tmp = self.accountSid + self.accountToken + str(unix_time_stamp)
        sig = my_md5(tmp)
        return sig

    def _gen_url(self,sig):
        tmp = ""
        if self.isCallPush == False:
            tmp = "http://" + self.domain + '/' + self.softVersion + '/'
            if(self.authType):
                tmp += 'SubAccounts/' + self.subAccountSid
            else:
                tmp += 'Accounts/' + self.accountSid

            tmp += '/' + self.operation

            if self.function:
                tmp += '/' + self.function
        else:
            tmp = "http://" + self.domain + '/' +self.operation
        tmp += '?sig=' + sig

        return tmp

    def _gen_authorization(self,unix_time_stamp):
        if(self.authType == 0):
	        tmp = self.accountSid + ':' + str(unix_time_stamp)
        else:
            tmp = self.subAccountSid + ':' + str(unix_time_stamp)
        authorization = my_base64(tmp)

        return authorization

    def set_func_mode(self,operation,function,authType):
        self.operation = operation
        self.function = function
        self.authType = authType

    def post_info_to_api(self,params):
        disable_warnings()
        http = PoolManager()
    
        unix_time_stamp = int(time.time())
        sig = self._gen_sig(unix_time_stamp)
    
        url = self._gen_url(sig)
    
        if self.dataType == "json":
            if self.appId:
                params["appId"] = self.appId
                request_dic = {self.function:params}
            else:
                request_dic = params
            request_string = json.dumps(request_dic)
            #request_string = gen_batch_request_json(json_task_file)
        elif self.dataType == "xml":
            #request_string = gen_batch_request_xml(xml_task_file)
            print("none")
        
        print(request_string)
        header = {}
        authorization = self._gen_authorization(unix_time_stamp)
        header['Accept'] = 'application/' + self.dataType
        header['Content-Type'] = 'application/' + self.dataType + ';charset=utf-8'
        header['Content-Length'] = str(len(request_string))
        header['Authorization'] = authorization
	    
        try:
            response = http.request('POST',url,body=request_string.encode("utf-8"),headers=header)
            data = response.data.decode('UTF-8')
            return data
        except Exception as e:
            print("post error")
            print(e)

