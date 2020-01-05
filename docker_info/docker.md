## 1 安装(环境为ubuntu)
1. 安装docker
sudo apt  install docker.io
2. 更改用户组
1) 创建docker用户组 sudo groupadd docker
2) 应用用户加入docker用户组 sudo usermod -aG docker (USER)
3) 重启docker服务 sudo systemctl restart docker
4) 切换当前账户再重新登入 su root su (USER)
3. 
