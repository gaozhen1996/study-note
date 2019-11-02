# 安装 Nginx

安装nginx的命令很简单

```
sudo apt-get install nginx 
```

在浏览器中输入，服务器地址，比如：

```
127.0.0.1
```

就会出现Welcome  to nginx！界面，代表nginx安装成功。

# nginx 的启动和关闭

启动 nginx：

```
nginx -c /etc/nginx/nginx.conf 
```

关闭 

```
nginx# nginx -s stop
```

重读配置文件

```
nginx -s reload
```