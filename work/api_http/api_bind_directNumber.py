from xml.dom.minidom import parse
import xml.dom.minidom
from api_post_class import *
from sys import exit
import json

def updateUser(f, mobile, directNumber):
	except_num = 0
	while(True):
		operation = "Enterprises"
		function = "updateUser"
		authType = 1
		params = {"phone":mobile, "directNumber":directNumber}
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
					return False
					exit(0)
				continue
		except Exception as e:
			print(e)
			return False
		break
	return True

auth_info = {}
auth_info["accountSid"] = "e9bb5de08a9f69fa19bc111e8d9283cf"
auth_info["accountToken"] = "0eb50bae2ce05b1efee0509af29af056"
auth_info["subAccountSid"] = "78228dbd3d0e8a6eedf3501355126386"
auth_info["subAccountToken"] = "5005a587cb398a4a1b89a08fda155fbf"
auth_info["appId"] = "9e5de7e655ee6067537d89d9a655ca17"
domain = "apiusertest.emic.com.cn"
version = "20161021"
dataType = "json"


mobile_list = []
directNumber_list = []

DOMTree = xml.dom.minidom.parse("yonghu.xml")
collection = DOMTree.documentElement
users = collection.getElementsByTagName("mobile")
for user in users:
	mobile = user.childNodes[0].data
	print(mobile)
	mobile_list.append(mobile)

f = open("directnumber.txt", "w")
#欠费返回对应号码集

DOMTree = xml.dom.minidom.parse("transferNumber.xml")
collection = DOMTree.documentElement
directNumbers = collection.getElementsByTagName("transferNumber")
for directNumber in directNumbers:
	directNumber_t = directNumber.childNodes[0].data
	print(directNumber_t)
	directNumber_list.append(directNumber_t)
	f.write(directNumber_t + "\n")

f.close()

#f = ApiPost(auth_info, domain, version, dataType)

#for mobile in mobile_list:
#	directNumber = directNumber_list.pop()
#	updateUser(f, str(mobile),str(directNumber))
	
