from tcp_client import tcp_socket_set
import json

tts_req = {
	'text_file':'/home/huanglin/tools/tcp_python_client/voice_dir/test.txt',
	'dest_file':'/home/huanglin/tools/tcp_python_client/voice_dir/dest.wav',
	'ttsCompany':'xf',
	'tts_voice':'xiaoyan',
	'tts_speed':'1.0'
}

host = '127.0.0.1'
port = '8600'

json_string = json.dumps(tts_req)
print(json_string)
result = tcp_socket_set(host,int(port),json_string)

print(result)
