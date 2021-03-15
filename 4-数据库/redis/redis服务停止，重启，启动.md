如果是用apt-get或者yum install安装的redis，可以直接通过下面的命令停止/启动/重启redis

/etc/init.d/redis-server stop

/etc/init.d/redis-server start

/etc/init.d/redis-server restart

如果是通过源码安装的redis，则可以通过redis的客户端程序redis-cli的shutdown命令来重启redis



```shell
redis-cli -h 127.0.0.1 -p 6379 shutdown
```

如果上述方式都没有成功停止redis，则可以使用终极武器 kill -9

