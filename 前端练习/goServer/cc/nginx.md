## 1 shell配置

```shell
useradd -s /sbin/nologin -M nginx
```

## 2 下载源码

```shell
git clone https://github.com/hlwqds/nginx.git
cd nginx
mv ./auto/configure .
./configure --prefix=/usr/local/lnmp/nginx --with-http_ssl_module --with-http_stub_status_module --with-file-aio --with-threads --user=nginx --group=ngnix
```

## 3 安装依赖库

在配置过程中肯定会有一些动态库未安装

### ubuntu

```shell
#opssl
apt-get install libssl-dev
apt-get install openssl
#libpcre3
apt-get install libpcre3 libpcre3-dev
#zlib1g
apt-get install zlib1g zlib1g-dev
```

## 4 编译

```shell
cd nginx
make & make install
ln  -s  /usr/local/lnmp/nginx/sbin/nginx  /usr/local/sbin/
```

## 5 修改nginx配置

```shell
#安装时会有配置文件地址的打印
#进入文件修改第一行
user root root;
```

## 6 运行nginx

