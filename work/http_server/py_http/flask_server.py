from flask import Flask, request
import json
from handle_module import SyncHanleWithReq, OldSyncHanleWithReqForGroup, OldSyncHanleWithReqForNumber
import logging
import http_log

app = Flask(__name__)

@app.route('/', methods=['POST', 'GET'])
def ivr_arg_search():
	logging.info(request.headers)
	info = request.stream.read().decode(encoding='UTF-8')
	
	logging.info(info)

	rsp = SyncHanleWithReq(info)
	logging.info(rsp)
	return rsp

@app.route('/oldGroup', methods=['POST', 'GET'])
def old_ivr_group_arg_search():
	logging.info(request.headers)
	info = request.stream.read().decode(encoding='UTF-8')
	
	logging.info(info)

	rsp = OldSyncHanleWithReqForGroup(info)
#	rsp = OldSyncHanleWithReqForNumber(info)

	logging.info(rsp)
	return rsp

@app.route('/oldNumber', methods=['POST', 'GET'])
def old_ivr_number_arg_search():
	logging.info(request.headers)
	info = request.stream.read().decode(encoding='UTF-8')
	
	logging.info(info)

	rsp = OldSyncHanleWithReqForNumber(info)
	logging.info(rsp)
	return rsp

if __name__ == '__main__':
	app.run(host='172.27.0.3',  port=8777, debug=True)
