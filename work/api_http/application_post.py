from flask_easy_post import *
from urllib3 import *

def application_post():
	json_task_file = 'app.txt'
	data_type = 'json'
	url = 'http://testcrm.hanhuioffice.com/admin/call/list'
	disable_warnings()
	http = PoolManager()
	if data_type == "json":
		request_string = gen_batch_request_json(json_task_file)
	elif data_type == "xml":
		request_string = gen_batch_request_xml(xml_task_file)
	
	header = {}
	header['Content-Type'] = 'application/' + data_type
	response = http.request('POST',url,body=request_string,headers=header)
	data = response.data.decode('UTF-8')
	print(url)	
	print(data)


application_post()
