from socket import *
class UdpInitException(Exception):
    pass

class UdpConnectException(Exception):
    pass

class UdpSendException(Exception):
    pass

class UdpRecvException(Exception):
    pass

class UdpCloseException(Exception):
    pass

class UdpBindException(Exception):
    pass

class Udp():
    def __init__(self,host='localhost',port=9876,bufferSize=1024):
        self._seraddr = (host,port)
        self._bufferSize = bufferSize
        try:
            self._socket = socket(AF_INET, SOCK_DGRAM)
        except Exception as e:
            print(e)
            raise UdpInitException()

    def close(self):
        try:
            self._socket.close()
            print("OK")
        except Exception as e:
            raise UdpCloseException()

class UdpClient(Udp):
    def send(self,request):
        try:
            self._socket.sendto(request,self._seraddr)
        except Exception as e:
            raise UdpSendException()
    def recv(self):
        try:
            response, seraddr = self._socket.recvfrom(self._bufferSize)
            return response
        except Exception as e:
            print(e)
            raise UdpRecvException()

class UdpServer(Udp):
    def bind(self):
        try:
            self._socket.bind(self._seraddr)
        except Exception as e:
            raise UdpBindException()

    def recv(self):
        try:
            request, cliaddr = self._socket.recvfrom(self._bufferSize)
            self._cliaddr = cliaddr
            return request
        except Exception as e:
            raise UdpRecvException()

    def send(self, response):
        try:
            self._socket.sendto(response, self._cliaddr)
        except Exception as e:
            print(e)
            raise UdpSendException()
#example
#udp = UdpClient()
#try:
#    udp.connect()
#    udp.send("lala")
#    data = udp.recv()
#    print(data)
#except Exception as e:
#    print(e)
#udp.close()

#udp = UdpServer()
#try:
#    udp.bind()
#    data = udp.recv()
#    print(data)
#    udp.send(data)
#except Exception as e:
#    print(e)

#udp.close()
