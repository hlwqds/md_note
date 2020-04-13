import csv
from zipfile import ZipFile

f = ZipFile("test.zip", 'r')

for f_name in f.namelist():
	print(f_name)
	#f.extract(f_name, '../')
	data = f.read(f_name)
	with open('haha.txt', 'wb') as f:
		f.write(data)

f.close()
