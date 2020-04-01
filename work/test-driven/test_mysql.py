from my_lib.my_mysql import *
import time

def test_create_new_customer(db, info_dict={}):
    return insert_info_into_mysql(db, info_dict, "emicall_cc_man.customers")

def test_create_new_customertels(db, info_dict={}):
    return insert_info_into_mysql(db, info_dict, "emicall_cc_man.customer_telephones")

def test_create_new_customer_othertels(db, info_dict={}):
    return insert_info_into_mysql(db, info_dict, "emicall_cc_man.customer_other_telephones")

def test_init_one_thousand_customer():
    ret = 0
    ccgeid = 234
    seid = 10
    telNum = 5
    timemark = time.time()
    db = connect("127.0.0.1","root","123456",charset="utf8")
    try:
        cus_db_info = {}
        custels_db_info = {}
        cusothertels_db_info = {}

        for i in range(0, 1000):
            cus_db_info["seid"] = seid
            cus_db_info["ccgeid"] = ccgeid
            cus_db_info["cm_name"] = timemark
            cus_db_info["cm_gender"] = 1
            cus_db_info["cm_age"] = i
            cus_db_info["cm_addr"] = "taixing" + str(i)
            cus_db_info["cm_company"] = "yimi" + str(i)
            cus_db_info["cm_email"] = str(i) + ".qq.com"
            cus_db_info["cm_web_site"] = str(i) + ".kaisa.com"
            cus_db_info["cm_remark"] = str(i) + "remark"
            cus_db_info["cm_detail"] = str(i) + "cm_detail"
            cid = test_create_new_customer(db, cus_db_info)
            custels_db_info["cid"] = cid
            custels_db_info["seid"] = seid
            custels_db_info["ccgeid"] = ccgeid
            for j in range(0, telNum):
                custels_db_info["telephone"] = timemark + str(j)
                test_create_new_customertels(db, custels_db_info)
            cusothertels_db_info["cid"] = cid
            cusothertels_db_info["seid"] = seid
            cusothertels_db_info["ccgeid"] = ccgeid
            for j in range(0, telNum):
                cusothertels_db_info["telephone"] = str(1) + timemark + str(j)
                cusothertels_db_info["type"] = 1
                test_create_new_customer_othertels(db, cusothertels_db_info)
            for j in range(0, telNum):
                cusothertels_db_info["telephone"] = str(2) + timemark + j
                cusothertels_db_info["type"] = 2
                test_create_new_customer_othertels(db, cusothertels_db_info)

        del cus_db_info
        del custels_db_info
        del cusothertels_db_info

    except Exception as e:
        print(e)
    finally:
        print("exit")
        db.close()
        return ret