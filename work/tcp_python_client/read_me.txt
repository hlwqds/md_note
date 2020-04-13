此脚本是为了测试tts，目前只支持通过cli命令发送单条请求
此脚本应运行在tts_server所在服务器上
在脚本目录中执行:sudo python3 voice_action.py -a 1 -t weada -c xunfi -s 1 -v 1
如果想要了解参数含义，执行sudo python3 voice_action.py -h
如果未指定文件存放目录，默认生成在脚本目录中所在的voice_dir文件夹
