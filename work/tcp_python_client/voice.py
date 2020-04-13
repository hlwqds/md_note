import uuid,json,os,stat
from tcp_client import tcp_socket_set

def genfile_and_send_to_tts(args):
	if args.text is None:
		result = 'Please input the text'
		return result
	
	if args.local_dir is None:
		dir_path = os.getcwd() + '/' + 'voice_dir'
	else:
		dir_path = args.local_dir

	if args.host is None:
		args.host = '127.0.0.1'
	if args.port is None:
		args.port = '8600'
	
	os.chmod(dir_path, stat.S_IRWXU+stat.S_IRWXG+stat.S_IRWXO)

	random_string = str(uuid.uuid1())
	textfile_name = dir_path + "/" + random_string + '.txt'
	voicefile_name = dir_path + "/" + random_string + '.wav'
	
	f = open(textfile_name, 'w+')
	f.write(args.text)
	f.close()
	
	data = {}
	data['text_file'] = textfile_name
	data['dest_file'] = voicefile_name
	if args.company is not None:
		data['ttsCompany'] = args.company
	if args.voice is not None:
		data['tts_voice'] = args.voice
	if args.speed is not None:
		data['tts_speed'] = args.speed
	
	json_request = json.dumps(data)
	
	print('request info:')
	print('textfile_name:' + textfile_name)
	print('voicefile_name:' + voicefile_name)
	if args.company is not None:
		print('company:' + args.company)
	if args.voice is not None:
		print('voice:' + args.voice)
	if args.speed is not None:
		print('speed:' + args.speed)
	
	try:
		result = tcp_socket_set(args.host,int(args.port),json_request)
		resp_data = {}
		resp_data = result
		print('resp_info:\n' + resp_data)
		return 'OK'
	except:
		os.remove(textfile_name)
		return 'Fail to connect'
