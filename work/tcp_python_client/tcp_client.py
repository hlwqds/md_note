import socket

def tcp_socket_set(host,port,request):
	bufferSize = 1024
	addr = (host, port)
	tcpClientSocket = socket.socket()
	tcpClientSocket.setblocking(False)
	tcpClientSocket.settimeout(3)
	print('connect to ip:%s,port:%d'%(host,port))
	tcpClientSocket.connect(addr)
	tcpClientSocket.send(request.encode('utf-8'))
	recvlen = 0
	data = ''
	while True:
		response = tcpClientSocket.recv(bufferSize)
		if not response:
			break
		data = data + response.decode()

	tcpClientSocket.close()
	return data
