

# 1.redis安装

在 Ubuntu 系统安装 Redis 可以使用以下命令:

```shell
$sudo apt-get update
$sudo apt-get install redis-server
```

#  2.启动 Redis

```
$ redis-server
$ redis-cli
```

会出现下面的界面

> redis 127.0.0.1:6379>

(127.0.0.1 是本机 IP ，6379 是 redis 服务端口。现在我们输入 PING 命令)

> redis 127.0.0.1:6379> ping

> PONG

以上说明我们已经成功安装了redis。

#  3.redis远程

修改配置文件 /etc/redis/redis.conf

在redis的配置文件redis.conf中，找到bind localhost注释掉。

注释掉本机,局域网内的所有计算机都能访问。

band localhost  只能本机访问,局域网内计算机不能访问。

bind 局域网IP  只能局域网内IP的机器访问, 本地localhost都无法访问。

验证方法：

> [root@mch ~]# ps -ef | grep redis
>
> root   2175   1 0 08:15 ?    00:00:05 /usr/local/bin/redis-server *:6379

/usr/local/bin/redis-server *:6379 中通过"*"就可以看出此时是允许所有的ip连接登录到这台redis服务上。