# 一.ubuntu上安装mysql

```
sudo apt-get update 

sudo apt-get install mysql-server 
```

# 二、修改权限

1.1进入mysql

```
mysql -u root -proot 
```

1.2修改数据权限

```
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'root' WITH GRANT OPTION; 
```

操作完后切记执行以下命令刷新权限 

```
FLUSH PRIVILEGES ; 
```

 

# 三、修改配置文件

**修改/etc/mysql/my.conf**

找到bind-address = 127.0.0.1这一行

改为bind-address = 0.0.0.0即可

