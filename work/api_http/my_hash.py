import hashlib
import base64
#string = 'huanglin'

#m = hashlib.md5()

#b = string.encode(encoding='utf-8')
#m.update(b)
#str_md5 = m.hexdigest()

#print(str_md5.upper())

#n = hashlib.md5()
#n.update(b"huanglin")
#str_md5 = n.hexdigest()
#print(str_md5.upper())

def my_md5(string):
	m = hashlib.md5()
	string_utf8 = string.encode(encoding='utf-8')
	m.update(string_utf8)
	return m.hexdigest().upper()

def my_base64(string):
	tmp = string.encode('utf-8')
	result = base64.b64encode(tmp)
	return result.decode('utf-8')

#s = 'huanglin'
#tmp = s.encode('utf-8')
#result = base64.b64encode(tmp)
#string = base64.b64decode(result)
#print(result.decode('utf-8'))
#print(string.decode('utf-8'))
