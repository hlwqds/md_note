# TURN服务器搭建和测试

通过对janus服务器的深入学习，对janus使用stun和turn服务器的作为NAT穿透方法的背景有了一定的了解。现在我希望能通过实践来进一步验证这些知识。

## 1 创建应用

```bash
 sudo apt-get install nodejs npm
 sudo git clone https://github.com/LingyuCoder/SkyRTC-demo
 cd SkyRTC-demo
 sudo npm install
```

以skyRtc为例

执行sudo nodejs server.js

## 2 搭建turn/stun

### 2.1 安装

sudo apt-get install libssl-dev
sudo apt-get install libevent-dev
sudo apt-get install libpq-dev
sudo apt-get install mysql-client
sudo apt-get install libmysqlclient-dev
sudo apt-get install libhiredis-dev
sudo apt-get install git

sudo git clone https://github.com/coturn/coturn
cd coturn
sudo ./configure
sudo make
sudo make install

### 2.2 配置

sudo cp /usr/local/etc/turnserver.conf.default  /usr/local/etc/turnserver.conf

注意：/usr/local/etc/turnserver.conf才是配置文件！ 
 /usr/local/etc/turnserver.conf.default并不是配置文件！

下面修改/usr/local/etc/turnserver.conf

1. 查看网卡，ifconfig
2. 