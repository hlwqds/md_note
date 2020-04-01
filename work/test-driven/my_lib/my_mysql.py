from pymysql import *

def get_info_from_mysql(db,info_string,table,conditions):
    cursor = db.cursor()

    sql = "select " + info_string + " from " + table + " where " + conditions
    print(sql)
    cursor.execute(sql)
    results = cursor.fetchall()
    fields = info_list
    records = []

    for row in results:
        records.append(dict(zip(fields,row)))
    return records

def insert_info_into_mysql(db, info_dict, table):
    cursor = db.cursor()
    first_object = True
    mysql_insert_key = ""
    mysql_insert_value = ""
    for key, value in info_dict.items():
        if first_object:
            mysql_insert_key += "`{}`".format(key)
            mysql_insert_value += '"{}"'.format(str(value))
            first_object = False
        else:
            mysql_insert_key += "," + "`{}`".format(key)
            mysql_insert_value += "," + '"{}"'.format(str(value))

    sql = "insert into " + table + " ({})".format(mysql_insert_key) + \
            " values" + " ({})".format(mysql_insert_value)

    print(sql)
    cursor.execute(sql)
    this_id = cursor.lastrowid
    db.commit()
    return this_id

#moudle
#get
#def get_appid_by_appid(db,appid):
    #info_list = ["id"]
    #table = "EmicallDev_system.applications"
    #conditions = "appId='%s'"%appid
    #results = get_info_from_mysql(db,info_list,table,conditions)
    #return results[0]["id"]
#connect
#db = connect("127.0.0.1","root","123456",charset="utf8")
#db.close()

