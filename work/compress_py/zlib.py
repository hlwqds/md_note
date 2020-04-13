import json
import zlib

result = {
	"name":"huanglin",
	"age":18,
}

with open("zlib.txt", "wb") as fp1:
	data1 = bytes(json.dumps(result),"utf-8")
	fp1.write(zlib.compress(data1))

with open("zlib.txt","rb") as fp:
	data = zlib.decompress(fp.read()).decode("utf-8")
	read_result = json.loads(data)
	print(read_result)
