import argparse
from voice import genfile_and_send_to_tts

def read_auto_args():
	parser = argparse.ArgumentParser()

	parser.add_argument(
		'-a','--action',
		help = '输入你的操作',
		metavar = 'action',
	)

	parser.add_argument(
		'-t','--text',
		help = '操作内容',
		default=None,
		metavar = 'text',
	)

	parser.add_argument(
		'-c','--company',
		help = '语音企业',
		default=None,
		metavar = 'company',
	)

	parser.add_argument(
		'-voice','--voice',
		help = '声音',
		default=None,
		metavar = 'voice',
	)

	parser.add_argument(
		'-s','--speed',
		help = '速度',
		default=None,
		metavar = 'speed',
	)
	
	parser.add_argument(
		'-i','--host',
		help = 'ip地址',
		default=None,
		metavar = 'host',
	)

	parser.add_argument(
		'-p','--port',
		help = '端口号',
		default=None,
		metavar = 'port',
	)

	parser.add_argument(
		'-d','--local_dir',
		help = '录音存放目录',
		default=None,
		metavar = 'local_dir'
	)

	return parser.parse_args()


if __name__ == '__main__':
	args = read_auto_args()
	#注册回调函数，或者称为自定义switch
	switcher = {
		'1':genfile_and_send_to_tts,
		'2':'',
		'3':'',
		'4':'',
	}

	print(args)
	
	#根据输入确定是哪种action
	func = switcher.get(args.action)

	result = func(args)
	print(result)
