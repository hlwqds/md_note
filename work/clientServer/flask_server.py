from flask import Flask, request
import json
from handle_module import SyncHanleWithReq
app = Flask(__name__)

@app.route('/', methods=['POST', 'GET'])
def ivr_arg_search():
	print(request.headers)
	info = request.stream.read().decode(encoding='UTF-8')
	
	print(info)

	rsp = SyncHanleWithReq(info)
	print(rsp)
	return rsp

if __name__ == '__main__':
	app.run(host='172.27.0.3',  port=8777, debug=True)
