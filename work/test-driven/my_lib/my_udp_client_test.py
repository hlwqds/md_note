from my_udp import*

udp = UdpClient()
try:
    udp.send("lala".encode(encoding='utf-8'))
    data = udp.recv()
    print(data)
except Exception as e:
    print(e)

udp.close()
