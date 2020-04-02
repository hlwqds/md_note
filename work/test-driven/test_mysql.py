from my_lib.my_mysql import *
import time

def test_create_new_customer(db, info_dict={}):
    return insert_info_into_mysql(db, info_dict, "emicall_cc_man.customers")

def test_create_new_customertels(db, info_dict={}):
    return insert_info_into_mysql(db, info_dict, "emicall_cc_man.customer_telephones")

def test_create_new_customer_othertels(db, info_dict={}):
    return insert_info_into_mysql(db, info_dict, "emicall_cc_man.customer_other_telephones")

def test_create_new_defined_field(db, info_dict={}):
    return insert_info_into_mysql(db, info_dict, "emicall_cc_man.customer_self_defined_fields")

def test_create_new_defined_value(db, info_dict={}):
    return insert_info_into_mysql(db, info_dict, "emicall_cc_man.customer_self_defined_values")

def test_create_new_defined_option(db, info_dict={}):
    return insert_info_into_mysql(db, info_dict, "emicall_cc_man.customer_self_defined_options")

def test_init_one_thousand_customer_with_self_defined_field():
    ret = 0
    ccgeid = 234
    seid = 10
    telNum = 5
    timemark = str(int(time.time()))
    db = connect("127.0.0.1","root","123456",charset="utf8")
    try:
        self_defined_list = []
        for i in range(0,3):
            self_defined_field = {}
            self_defined_field["seid"] = seid
            self_defined_field["ccgeid"] = ccgeid
            self_defined_field["field_name"] = "one" + str(i) + timemark
            self_defined_field["field_type"] = 0
            field_id = test_create_new_defined_field(db, self_defined_field)
            self_defined_field["id"] = field_id
            self_defined_list.append(self_defined_field)
            del self_defined_field

        
        for i in range(0,3):
            self_defined_field = {}
            self_defined_field["seid"] = seid
            self_defined_field["ccgeid"] = ccgeid
            self_defined_field["field_name"] = "selection" + str(i) + timemark
            self_defined_field["field_type"] = 4
            field_id = test_create_new_defined_field(db, self_defined_field)
            self_defined_field["id"] = field_id
            
            self_defined_optionid_list = []
            self_defined_option = {}
            self_defined_option["seid"] = seid
            self_defined_option["ccgeid"] = ccgeid
            self_defined_option["field_id"] = field_id
            for j in range(0, 3):
                self_defined_option["field_option"] = "selection" + str(j) + timemark
                option_id = test_create_new_defined_option(db, self_defined_option)
                self_defined_optionid_list.append(option_id)
            self_defined_field["self_defined_option"] = self_defined_optionid_list
            self_defined_list.append(self_defined_field)
            del self_defined_option
            del self_defined_optionid_list
            del self_defined_field

            
        cus_db_info = {}
        custels_db_info = {}
        cusothertels_db_info = {}

        for i in range(0, 10):
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
                custels_db_info["telephone"] = str(i) + str(j) + timemark
                test_create_new_customertels(db, custels_db_info)
            cusothertels_db_info["cid"] = cid
            cusothertels_db_info["seid"] = seid
            cusothertels_db_info["ccgeid"] = ccgeid
            for j in range(0, telNum):
                cusothertels_db_info["telephone"] = str(1) + str(j) + timemark
                cusothertels_db_info["type"] = 1
                test_create_new_customer_othertels(db, cusothertels_db_info)
            for j in range(0, telNum):
                cusothertels_db_info["telephone"] = str(2) + str(j) + timemark
                cusothertels_db_info["type"] = 2
                test_create_new_customer_othertels(db, cusothertels_db_info)

            print(self_defined_list)
            for field in self_defined_list:
                self_defined_value = {}
                self_defined_value["seid"] = seid
                self_defined_value["ccgeid"] = ccgeid
                self_defined_value["cid"] = cid
                self_defined_value["field_id"] = field["id"]
                if field["type"] == 0:
                    self_defined_value["field_value"] = "onevalue" + str(cid) + timemark
                    test_create_new_defined_value(db, self_defined_value)
                elif field["type"] == 4:
                    self_defined_value["field_value"] = ""
                    first = True
                    for optionId in field["self_defined_option"]:
                        self_defined_value["field_value"] = str(optionId)
                        test_create_new_defined_value(db, self_defined_value)

        del cus_db_info
        del custels_db_info
        del cusothertels_db_info
        del self_defined_list

    except Exception as e:
        print(e)
    finally:
        print("exit")
        db.close()
        return ret
