from my_lib.my_mysql import *
from my_lib.my_udp import *
import time

def test_udp(taskdict):
    udp = UdpClient(host='10.0.1.49',port=1505)
    message = str(taskdict)
    try:
        udp.send(message.encode(encoding='utf-8'))
        #data = udp.recv().decode(encoding='utf-8')
        #print(data)
    except Exception as e:
        print(e)
    finally:    
        udp.close()
    return 0

def test_batchcall_create_new_job(db,info_dict={}):
    return insert_info_into_mysql(db,info_dict,"emicall_cc_man.batch_call_task_job")

def test_batchcall_new_asynctask(db,info_dict={}):
    return insert_info_into_mysql(db,info_dict,"emicall_cc_man.async_task_new")

def test_create_new_group_statis(db, info_dict={}):
    return insert_info_into_mysql(db, info_dict, "emicall_cc_man.preview_task_group_statistics")

def test_create_new_group_cm_statis(db, info_dict={}):
    return insert_info_into_mysql(db, info_dict, "emicall_cc_man.preview_task_cm_group_statistics")

def test_create_new_customer(db, info_dict={}):
    return insert_info_into_mysql(db, infor_dict, "emicall_cc_man.customer")

def test_batchcall():
    ret = 0
    ccgeid = 234
    seid = 10
    meduleType = 1
    action = 1
    task_id = 1234
    async_task_id = 0
    job_id = 0
    db = connect("127.0.0.1","root","123456",charset="utf8")

    try:
        job_db_info = {}
        job_db_info["seid"] = seid
        job_db_info["ccgeid"] = ccgeid
        job_db_info["task_id"] = task_id
        job_db_info["type"] = 1
        job_db_info["job_name"] = int(time.time())
        job_db_info["unique_id"] = "job_id2"
        job_id = test_batchcall_create_new_job(db,job_db_info)
        del job_db_info
        
        job_id = 3
        task_db_info = {}
        task_db_info["seid"] = seid
        task_db_info["ccgeid"] = ccgeid
        task_db_info["task_name"] = int(time.time())
        task_db_info["module"] = 1
        task_db_info["module_id"] = 1
        task_db_info["sub_module_id"] = 1
        
        task_db_condition = {}
        task_db_condition["task_id"] = task_id
        task_db_condition["job_id"] = job_id
        task_db_info["condition"] = task_db_condition
        async_task_id = test_batchcall_new_asynctask(db,task_db_info)
        del task_db_condition

        async_task_id = 198
        print(async_task_id)
        udpdict = {}
        udpdict["ccgeid"] = ccgeid
        udpdict["seid"] = seid
        udpdict["moduleType"] = meduleType
        udpdict["task_id"] = async_task_id
        udpdict["action"] = action
        test_udp(udpdict)
        del udpdict
    except Exception as e:
        print(e)
    finally:
        print("exit")
        db.close()
        return ret

def test_batchcall_recall():
    ret = 0
    ccgeid = 45
    seid = 10
    meduleType = 1
    action = 1
    task_id = 1

    async_task_id = 0
    job_id = 0
    db = connect("127.0.0.1","root","123456",charset="utf8")
    
    try:
        job_db_info = {}
        job_db_info["seid"] = seid
        job_db_info["ccgeid"] = ccgeid
        job_db_info["task_id"] = task_id
        job_db_info["type"] = 1
        job_db_info["job_name"] = int(time.time())
        job_db_info["unique_id"] = "job_id2"
        job_id = test_batchcall_create_new_job(db,job_db_info)
        del job_db_info
	
        task_db_info = {}
        task_db_info["seid"] = seid
        task_db_info["ccgeid"] = ccgeid
        task_db_info["task_name"] = int(time.time())
        task_db_info["module"] = 1
        task_db_info["module_id"] = 1
        task_db_info["sub_module_id"] = 3

        task_db_condition = {}
        screen_condition = {}
        #screen_condition["ids"] = "1,2,3"
        #screen_condition["is_called"] = 0
        #screen_condition["call_result"] = "0,1,2,3,4"
        #screen_condition["cm_result"] = "0,1,2,3,4"
        #screen_condition["call_count_condition"] = ">="
        #screen_condition["call_count"] = 0
        
        task_db_condition["task_id"] = task_id
        task_db_condition["job_id"] = job_id
        task_db_condition["condition"] = screen_condition
        task_db_info["condition"] = task_db_condition
        async_task_id = test_batchcall_new_asynctask(db,task_db_info)
        del screen_condition
        del task_db_condition
	
        print(async_task_id)
        udpdict = {}
        udpdict["ccgeid"] = ccgeid
        udpdict["seid"] = seid
        udpdict["moduleType"] = meduleType     #batch_call
        udpdict["task_id"] = async_task_id
        udpdict["action"] = action
        test_udp(udpdict)
        del udpdict
    except Exception as e:
        print(e)
    finally:
        print("exit")
        db.close()
        return ret

def test_customer_export():
    ret = 0
    ccgeid = 234
    seid = 10
    moduleType = 11
    action = 1

    async_task_id = 0
    job_id = 0
    db = connect("127.0.0.1","root","123456",charset="utf8")
    
    try:
        csv_name = str(int(time.time())) + ".csv"
	
        task_db_info = {}
        task_db_info["seid"] = seid
        task_db_info["ccgeid"] = ccgeid
        task_db_info["task_name"] = int(time.time())
        task_db_info["module"] = 11
        task_db_info["module_id"] = 11
        task_db_info["sub_module_id"] = 2
        task_db_info["file_path"] = csv_name

        task_db_condition = {}
        task_db_condition["column"] = "cm_name,cm_tels,home_tels,enterprise_tels,cm_gender,cm_email,cm_addr,cm_company,cm_web_site,cm_detail,modify_time,defineField_49,defineField_50,defineField_51"
        task_db_condition["fixed_column_mask"] = 1
        task_db_condition["hide_customer_number"] = "xx****x"
        task_db_condition["ids"] = "42144"
        task_db_info["condition"] = task_db_condition
        async_task_id = test_batchcall_new_asynctask(db,task_db_info)
        del task_db_condition
	
        print(async_task_id)
        udpdict = {}
        udpdict["ccgeid"] = ccgeid
        udpdict["seid"] = seid
        udpdict["moduleType"] = moduleType     #customer
        udpdict["task_id"] = async_task_id
        udpdict["action"] = action
        test_udp(udpdict)
        del udpdict
    except Exception as e:
        print(e)
    finally:
        print("exit")
        db.close()
        return ret

def test_preview_call_export():
    ret = 0
    ccgeid = 45
    seid = 10
    moduleType = 12
    action = 1

    async_task_id = 0
    job_id = 0
    db = connect("127.0.0.1","root","123456",charset="utf8")
    
    try:
        csv_name = str(int(time.time())) + ".csv"
	
        task_db_info = {}
        task_db_info["seid"] = seid
        task_db_info["ccgeid"] = ccgeid
        task_db_info["task_name"] = int(time.time())
        task_db_info["module"] = 12
        task_db_info["module_id"] = 12
        task_db_info["sub_module_id"] = 3
        task_db_info["file_path"] = csv_name

        task_db_condition = {}
        task_db_condition["enterprise_ids"] = "1,2,3,45"
        task_db_info["condition"] = task_db_condition
        async_task_id = test_batchcall_new_asynctask(db,task_db_info)
        del task_db_condition
	
        print(async_task_id)
        udpdict = {}
        udpdict["ccgeid"] = ccgeid
        udpdict["seid"] = seid
        udpdict["moduleType"] = moduleType     #customer
        udpdict["task_id"] = async_task_id
        udpdict["action"] = action
        test_udp(udpdict)
        del udpdict
    except Exception as e:
        print(e)
    finally:
        print("exit")
        db.close()
        return ret

def test_preview_call_detail_export():
    ret = 0
    ccgeid = 45
    seid = 10
    moduleType = 12
    action = 1

    async_task_id = 0
    job_id = 0
    db = connect("127.0.0.1","root","123456",charset="utf8")
    
    try:
        csv_name = str(int(time.time())) + ".csv"
	
        task_db_info = {}
        task_db_info["seid"] = seid
        task_db_info["ccgeid"] = ccgeid
        task_db_info["task_name"] = int(time.time())
        task_db_info["module"] = 12
        task_db_info["module_id"] = 12
        task_db_info["sub_module_id"] = 4
        task_db_info["file_path"] = csv_name	

        task_db_condition = {}
        task_db_condition["fixed_column_mask"] = 1
        task_db_condition["hide_customer_number"] = "XX**XX"
        task_db_condition["task_id"] = 1
        task_db_condition["batch_id"] = 1
        task_db_info["condition"] = task_db_condition
        async_task_id = test_batchcall_new_asynctask(db,task_db_info)
        del task_db_condition
	
        print(async_task_id)
        udpdict = {}
        udpdict["ccgeid"] = ccgeid
        udpdict["seid"] = seid
        udpdict["moduleType"] = moduleType     #customer
        udpdict["task_id"] = async_task_id
        udpdict["action"] = action
        test_udp(udpdict)
        del udpdict
    except Exception as e:
        print(e)
    finally:
        print("exit")
        db.close()
        return ret

def test_preview_call_cm_result_export():
    ret = 0
    ccgeid = 45
    seid = 10
    moduleType = 12
    action = 1

    async_task_id = 0
    job_id = 0
    db = connect("127.0.0.1","root","123456",charset="utf8")
    
    try:
        csv_name = str(int(time.time())) + ".csv"


#        for i in range(51):
#            db_info = {}
#            db_info["seid"] = 10
#            db_info["ccgeid"] = 45
#            db_info["task_id"] = 1
#            db_info["signed_count"] = 51
#            db_info["result_id"] = i + 1
#            db_info["result_name"] = "parent" + str(i+1)
#            test_create_new_group_cm_statis(db, db_info)
#            for j in range(51):
#                sub_db_info = {}
#                sub_db_info["seid"] = 10
#                sub_db_info["ccgeid"] = 45
#                sub_db_info["task_id"] = 1
#                sub_db_info["signed_count"] = 1
#                sub_db_info["result_id"] = j + 1
#                sub_db_info["parent_result_id"] = i + 1
#                sub_db_info["result_name"] = "sub" + str(j+1)
#                sub_db_info["parent_result_name"] = "parent" + str(i+1)
#                test_create_new_group_cm_statis(db, sub_db_info)


        task_db_info = {}
        task_db_info["seid"] = seid
        task_db_info["ccgeid"] = ccgeid
        task_db_info["task_name"] = int(time.time())
        task_db_info["module"] = 12
        task_db_info["module_id"] = 12
        task_db_info["sub_module_id"] = 5
        task_db_info["file_path"] = csv_name	

        task_db_condition = {}
        task_db_condition["fixed_column_mask"] = 1
        task_db_condition["hide_customer_number"] = 1
        task_db_condition["searchStrColumn"] = "cm_tels"
        task_db_condition["searchStr"] = "15"
        task_db_condition['seid'] = 10
        task_db_condition['ccgeid'] = 45
        task_db_info["condition"] = task_db_condition
        async_task_id = test_batchcall_new_asynctask(db,task_db_info)
        del task_db_condition
	
        print(async_task_id)
        udpdict = {}
        udpdict["ccgeid"] = ccgeid
        udpdict["seid"] = seid
        udpdict["moduleType"] = moduleType     #customer
        udpdict["task_id"] = async_task_id
        udpdict["action"] = action
        test_udp(udpdict)
        del udpdict
    except Exception as e:
        print(e)
    finally:
        print("exit")
        db.close()
        return ret

def test_preview_call_seat_statis_export():
    ret = 0
    ccgeid = 45
    seid = 10
    moduleType = 12
    action = 1

    async_task_id = 0
    job_id = 0
    db = connect("127.0.0.1","root","123456",charset="utf8")
    
    try:
        csv_name = str(int(time.time())) + ".csv"
	
        task_db_info = {}
        task_db_info["seid"] = seid
        task_db_info["ccgeid"] = ccgeid
        task_db_info["task_name"] = int(time.time())
        task_db_info["module"] = 12
        task_db_info["module_id"] = 12
        task_db_info["sub_module_id"] = 6
        task_db_info["file_path"] = csv_name	

        task_db_condition = {}
        task_db_condition["fixed_column_mask"] = 1
        task_db_condition["hide_customer_number"] = 1
        task_db_condition["searchStrColumn"] = "cm_tels"
        task_db_condition["searchStr"] = "15"
        #task_db_condition["batch_id"] = 1
        task_db_condition["task_id"] = 1
        task_db_condition["gid"] = 1
        task_db_condition["ccgeid"] = 45
        task_db_condition["seid"] = 10

        task_db_info["condition"] = task_db_condition
        async_task_id = test_batchcall_new_asynctask(db,task_db_info)
        del task_db_condition
	
        print(async_task_id)
        udpdict = {}
        udpdict["ccgeid"] = ccgeid
        udpdict["seid"] = seid
        udpdict["moduleType"] = moduleType     #customer
        udpdict["task_id"] = async_task_id
        udpdict["action"] = action
        test_udp(udpdict)
        del udpdict
    except Exception as e:
        print(e)
    finally:
        print("exit")
        db.close()
        return ret

def test_preview_call_group_statis_export():
    ret = 0
    ccgeid = 45
    seid = 10
    moduleType = 12
    action = 1

    async_task_id = 0
    job_id = 0
    db = connect("127.0.0.1","root","123456",charset="utf8")
    
    try:
        csv_name = str(int(time.time())) + ".csv"
	
        task_db_info = {}
        task_db_info["seid"] = seid
        task_db_info["ccgeid"] = ccgeid
        task_db_info["task_name"] = int(time.time())
        task_db_info["module"] = 12
        task_db_info["module_id"] = 12
        task_db_info["sub_module_id"] = 7
        task_db_info["file_path"] = csv_name	

        task_db_condition = {}
        task_db_condition["fixed_column_mask"] = 1
        task_db_condition["hide_customer_number"] = 1
        task_db_condition["searchStrColumn"] = "cm_tels"
        task_db_condition["searchStr"] = "15"
        task_db_condition["batch_id"] = 1
        task_db_condition["task_id"] = 1
        task_db_condition["ccgeid"] = 45
        task_db_condition["seid"] = 10

        task_db_info["condition"] = task_db_condition
        async_task_id = test_batchcall_new_asynctask(db,task_db_info)
        del task_db_condition
	
        print(async_task_id)
        udpdict = {}
        udpdict["ccgeid"] = ccgeid
        udpdict["seid"] = seid
        udpdict["moduleType"] = moduleType     #customer
        udpdict["task_id"] = async_task_id
        udpdict["action"] = action
        test_udp(udpdict)
        del udpdict
    except Exception as e:
        print(e)
    finally:
        print("exit")
        db.close()
        return ret

def test_preview_call_statis_export():
    ret = 0
    ccgeid = 45
    seid = 10
    moduleType = 12
    action = 1

    async_task_id = 0
    job_id = 0
    db = connect("127.0.0.1","root","123456",charset="utf8")
    
    try:
        csv_name = str(int(time.time())) + ".csv"
	
        task_db_info = {}
        task_db_info["seid"] = seid
        task_db_info["ccgeid"] = ccgeid
        task_db_info["task_name"] = int(time.time())
        task_db_info["module"] = 12
        task_db_info["module_id"] = 12
        task_db_info["sub_module_id"] = 8
        task_db_info["file_path"] = csv_name	

        task_db_condition = {}
        task_db_condition["fixed_column_mask"] = 1
        task_db_condition["hide_customer_number"] = 1
        task_db_condition["searchStrColumn"] = "cm_tels"
        task_db_condition["searchStr"] = "15"
        task_db_condition["task_id"] = 1
        task_db_condition["ccgeid"] = 45
        task_db_condition["seid"] = 10

        task_db_info["condition"] = task_db_condition
        async_task_id = test_batchcall_new_asynctask(db,task_db_info)
        del task_db_condition
	
        print(async_task_id)
        udpdict = {}
        udpdict["ccgeid"] = ccgeid
        udpdict["seid"] = seid
        udpdict["moduleType"] = moduleType     #customer
        udpdict["task_id"] = async_task_id
        udpdict["action"] = action
        test_udp(udpdict)
        del udpdict
    except Exception as e:
        print(e)
    finally:
        print("exit")
        db.close()
        return ret

def test_preview_call_batch_export():
    ret = 0
    ccgeid = 45
    seid = 10
    moduleType = 12
    action = 1

    async_task_id = 0
    job_id = 0
    db = connect("127.0.0.1","root","123456",charset="utf8")
    
    try:
        csv_name = str(int(time.time())) + ".csv"
	
#        for i in range(51):
#            db_info = {}
#            db_info["seid"] = 10
#            db_info["ccgeid"] = 45
#            db_info["task_id"] = 1
#            db_info["batch_id"] = 1
#            db_info["gid"] = i + 10
#            db_info["group_name"] = "group" + str(i)
#            test_create_new_group_statis(db, db_info)
            

        task_db_info = {}
        task_db_info["seid"] = seid
        task_db_info["ccgeid"] = ccgeid
        task_db_info["task_name"] = int(time.time())
        task_db_info["module"] = 12
        task_db_info["module_id"] = 12
        task_db_info["sub_module_id"] = 9
        task_db_info["file_path"] = csv_name

        task_db_condition = {}

        task_db_info["condition"] = task_db_condition
        async_task_id = test_batchcall_new_asynctask(db,task_db_info)
        del task_db_condition
	
        print(async_task_id)
        udpdict = {}
        udpdict["ccgeid"] = ccgeid
        udpdict["seid"] = seid
        udpdict["moduleType"] = moduleType     #customer
        udpdict["task_id"] = async_task_id
        udpdict["action"] = action
        test_udp(udpdict)
        del udpdict
    except Exception as e:
        print(e)
    finally:
        print("exit")
        db.close()
        return ret



if __name__ == "__main__":
    test_customer_export()
    #time.sleep(1)
    #test_preview_call_batch_export()
    #time.sleep(1)
    #test_preview_call_export()
    #time.sleep(1)
    #test_preview_call_detail_export()
    #time.sleep(1)
    #test_preview_call_cm_result_export()
    #time.sleep(1)
    #test_preview_call_seat_statis_export()
    #time.sleep(1)
    #test_preview_call_group_statis_export()
    #time.sleep(1)
    #test_preview_call_batch_export()
    #time.sleep(1)
    #test_preview_call_statis_export()
