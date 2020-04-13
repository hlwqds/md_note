from urllib import parse,request
import time
from my_hash import my_md5,my_base64
import json
import ssl

token_database = ""

def https_client(url, data):
    context = ssl._create_unverified_context()
    req = request.Request(url, data=data)

    with  request.urlopen(req, context=context) as response:
        rsp = response.read()
    return rsp.decode("utf-8")

def send_message():
    timestamp = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
    template_param = {"param1":"黄琳"}
    url = "http://api.189.cn/v2/emp/templateSms/sendSms"
    data = 'app_id=371756040000276471&access_token={0}&acceptor_tel=15861800293&template_id=91555741&template_param={1}&timestamp={2}'.format(token_database,template_param,timestamp)
    rsp = https_client(url,data.encode('utf-8'))
    rsp_json = json.loads(rsp)
    print(data)
    print(rsp_json)
    return rsp_json

def get_token():
    url = 'https://oauth.api.189.cn/emp/oauth2/v3/access_token'
    #data = 'grant_type=client_credentials&app_id=131769420000276439&app_secret=e2ac4d7e6f6af56439b64b253b0e95ac'.encode("utf-8")
    data = 'grant_type=client_credentials&app_id=371756040000276471&app_secret=14140b4bb123bdd77b3b5915f2dad595'.encode("utf-8")
    rsp = https_client(url,data)
    rsp_json = json.loads(rsp)
    token = rsp_json["access_token"]
    print(rsp)
    return token

rsp_json = send_message()
if token_database is None or rsp_json["res_code"] == 1:
    token_database = get_token()
    send_message()
