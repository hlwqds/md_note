import json
from pymysql import *
import csv

def read_info_from_dic_format_txt(filename):
	f = open(filename,'r')
	return json.load(f)

def get_info_from_mysql(db,info_list,table,conditions):
	cursor = db.cursor()
	first_object = True
	info_string = ""
	for list_object in info_list:
		if first_object:
			info_string += list_object
			first_object = False
		else:
			info_string += "," + list_object
	
	sql = "select " + info_string + " from " + table + " where " + conditions
	print(sql)
	cursor.execute(sql)
	results = cursor.fetchall()
	fields = info_list
	records = []
	for row in results:
		records.append(dict(zip(fields,row)))
	return records

def get_appid_by_appid(db,appid):
	info_list = ["id"]
	table = "EmicallDev_system.applications"
	conditions = "appId='%s'"%appid
	results = get_info_from_mysql(db,info_list,table,conditions)
	return results[0]["id"]

def get_notify_count(db,app_id,stime,etime,call_type,statis_type):
	info_list = ["count(*)"]
	table = "EmicallDev_system.call_details"
	conditions = "callId IN(SELECT callId FROM EmicallDev_system.call_records WHERE app_id = %d AND createTime > FROM_UNIXTIME(%d) AND createTime < FROM_UNIXTIME(%d) AND TYPE = %d)"%(app_id,stime,etime,call_type)
	if statis_type == 1:
		conditions += " AND status = 3"
	elif statis_type == 2:
		conditions += " AND status = 1"
	elif statis_type == 3:
		conditions += " AND status != 1 AND status != 3"
	results = get_info_from_mysql(db,info_list,table,conditions)
	return results[0]["count(*)"]

def getnotify_count_list(db,appname,app_id,stime,etime,call_type):
	app_total = get_notify_count(db,app_id,stime,etime,call_type,0)
	app_success = get_notify_count(db,app_id,stime,etime,call_type,1)
	app_failure = get_notify_count(db,app_id,stime,etime,call_type,2)
	app_exception = get_notify_count(db,app_id,stime,etime,call_type,3)
	notify_info_list = [appname,call_type,app_total,app_success,app_failure,app_exception]
	return notify_info_list

applications_dic = read_info_from_dic_format_txt('test.txt')
application_list = applications_dic["name"]
stime = applications_dic["stime"]
etime = applications_dic["etime"]

db = connect("127.0.0.1","root","123456",charset="utf8")

with open('csvfile.csv','w',newline='') as csvf:
	spanwriter=csv.writer(csvf,dialect='excel')
	spanwriter.writerow(['名称','类型','总数','成功数','失败数','异常数'])
	while application_list:
		application_dic = application_list.pop()
		appname = application_dic["name"]
		appid = application_dic["appid"]
		app_id = get_appid_by_appid(db,appid)
		#语音验证码
		call_type = 2
		notify_info_list = getnotify_count_list(db,appname,app_id,stime,etime,call_type)
		spanwriter.writerow(notify_info_list)
		#语音通知
		call_type = 3
		notify_info_list = getnotify_count_list(db,appname,app_id,stime,etime,call_type)
		spanwriter.writerow(notify_info_list)
csvf.close()
db.close()
