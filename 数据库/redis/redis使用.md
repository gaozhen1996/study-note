# 什么是NoSQL

## NoSQL

NoSQL = Not Only SQL （不仅仅是SQL）

关系型数据库：表格 ，行 ，列

泛指非关系型数据库的，随着web2.0互联网的诞生！传统的关系型数据库很难对付web2.0时代！尤其
是超大规模的高并发的社区！ 暴露出来很多难以克服的问题，NoSQL在当今大数据环境下发展的十分迅
速，Redis是发展最快的，而且是我们当下必须要掌握的一个技术！

很多的数据类型用户的个人信息，社交网络，地理位置。这些数据类型的存储不需要一个固定的格式！
不需要多月的操作就可以横向扩展的 ！ Map<String,Object> 使用键值对来控制！

## NoSQL 特点

1. 方便扩展（数据之间没有关系，很好扩展！）

2. 大数据量高性能（Redis 一秒写 8 万次，读取 11 万，NoSQL的缓存记录级，是一种细粒度的缓存，性
   能会比较高！）

3. 数据类型是多样型的！（不需要事先设计数据库！随取随用！如果是数据量十分大的表，很多人就无
   法设计了！）

4. 传统 RDBMS 和 NoSQL对比

RDBMS

- SQL
- 数据和关系都存在单独的表中 row col
- 操作操作，数据定义语言
- 严格的一致性
- 基础的事务

Nosql
- 不仅仅是数据
- 没有固定的查询语言
- 键值对存储，列存储，文档存储，图形数据库（社交关系）
- 最终一致性，
- CAP定理和BASE （异地多活） 初级架构师！（狂神理念：只要学不死，就往死里学！）
- 高性能，高可用，高可扩
- ....

## 不同的数据使用场景

- 商品的基本信息
  - 名称、价格、商家信息
  - 关系型数据库就可以解决了！ MySQL / Oracle 

- 商品的描述、评论（文字比较多）
  - 文档型数据库中，MongoDB

- 图片
  - 分布式文件系统 FastDFS
  - 淘宝自己的 TFS
  - Gooale的 GFS
  - Hadoop HDFS
  - 阿里云的 oss
- 商品的关键字 （搜索）
  - 搜索引擎 solr elasticsearch
  - ISerach
- 商品热门的波段信息
  - 内存数据库，Redis Tair、Memache..

# Redis入门

## 概述

Redis（Remote Dictionary Server )，即远程字典服务!

是一个开源的使用ANSI C语言编写、支持网络、可基于内存亦可持久化的日志型、Key-Value数据库，
并提供多种语言的API。

**Redis 作用**

1 、内存存储、持久化，内存中是断电即失、所以说持久化很重要（rdb、aof）

2 、效率高，可以用于高速缓存

3 、发布订阅系统

4 、地图信息分析

5 、计时器、计数器（浏览量！）

6 、........

## 测试性能

**redis-benchmark** 是一个压力测试工具！

- 测试： 100 个并发连接 100000 请求

```
redis-benchmark -h localhost -p 6379 -c 100 -n 100000
```

## 基础知识

- redis默认有 16 个数据库，默认使用的是第 0 个

- 可以使用 select 进行切换数据库！

```
127 .0.0.1:6379> select 3 # 切换数据库
OK
127 .0.0.1:6379[3]> DBSIZE  # 查看DB大小！
(integer) 0
127 .0.0.1:6379[3]> keys *  # 查看数据库所有的key
1 ) "name"
```

- 清除当前数据库 **flushdb**

- 清除全部数据库的内容 **FLUSHALL**

- **Redis 是单线程的**

明白Redis是很快的，官方表示，Redis是基于内存操作，CPU不是Redis性能瓶颈，Redis的瓶颈是根据
机器的内存和网络带宽，既然可以使用单线程来实现，就使用单线程了！所有就使用了单线程了！

Redis 是C 语言写的，官方提供的数据为 100000+ 的QPS，完全不比同样是使用 key-vale的
Memecache差！

**Redis 为什么单线程还这么快？**

核心：redis 是将所有的数据全部放在内存中的，所以说使用单线程去操作效率就是最高的，多线程
（CPU上下文会切换：耗时的操作！！！），对于内存系统来说，如果没有上下文切换效率就是最高
的！多次读写都是在一个CPU上的，在内存情况下，这个就是最佳的方案！

# 五大数据类型

Redis 是一个开源（BSD许可）的，内存中的数据结构存储系统，它可以用作数据库、缓存和消息中间
件MQ。 它支持多种类型的数据结构，如 字符串（strings）， 散列（hashes）， 列表（lists）， 集合
（sets）， 有序集合（sorted sets） 与范围查询， bitmaps， hyperloglogs 和 地理空间
（geospatial） 索引半径查询。 Redis 内置了 复制（replication），LUA脚本（Lua scripting）， LRU
驱动事件（LRU eviction），事务（transactions） 和不同级别的 磁盘持久化（persistence）， 并通过
Redis哨兵（Sentinel）和自动 分区（Cluster）提供高可用性（high availability）。

## Redis-Key

## String（字符串）

90% 的 java程序员使用 redis 只会使用一个String类型！

- 字符串范围 range

```redis
127.0.0.1:6379> keys *
(empty list or set)
127.0.0.1:6379> set key1 "hello,gaozhen"
OK
127.0.0.1:6379> get key1
"hello,gaozhen"
127.0.0.1:6379> getrange key1 0 3  #get key1 0-3的字符串
"hell"
127.0.0.1:6379> 
```

- 替换

```redis
127.0.0.1:6379> set key2 123abc
OK
127.0.0.1:6379> get key2
"123abc"
127.0.0.1:6379> setrange key2 1 gz # 替换指定位置开始的字符串
(integer) 6
127.0.0.1:6379> get key2
"1gzabc"
127.0.0.1:6379> 
```

- setex (set with expire) 设置过期时间

```
# 设置key3 的值为 hello,30秒后过期
127.0.0.1:6379> setex key3 30 "hello" 
OK
127.0.0.1:6379> ttl key3
(integer) 23
127.0.0.1:6379> ttl key3
(integer) 22
127.0.0.1:6379> ttl key3
(integer) 1
127.0.0.1:6379> ttl key3
(integer) -2
127.0.0.1:6379> 
# 使用expire设置过去时间
127.0.0.1:6379> set key3 123
OK
127.0.0.1:6379> expire key3 10
(integer) 1
127.0.0.1:6379> ttl key3
(integer) 7
127.0.0.1:6379> ttl key3
(integer) -2
127.0.0.1:6379> 
```

- setnx (set if not exist) # 不存在在设置 （在分布式锁中会常常使用！）

```redis
127.0.0.1:6379> set lock:user "user1"
OK
127.0.0.1:6379> setnx lock:user "user2"
(integer) 0
127.0.0.1:6379> get lock:user
"user1"
127.0.0.1:6379>
```

- 其他命令

```
127 .0.0.1:6379> msetnx k1 v1 k4 v4  # msetnx 是一个原子性的操作，要么一起成功，要么一起失败！
(integer) 0
127 .0.0.1:6379> get k4
(nil)
##########################################################################
getset # 先get然后在set
127 .0.0.1:6379> getset db redis # 如果不存在值，则返回 nil
(nil)
127 .0.0.1:6379> get db
"redis
127 .0.0.1:6379> getset db mongodb  # 如果存在值，获取原来的值，并设置新的值
"redis"
127 .0.0.1:6379> get db
"mongodb"
```



- String类似的使用场景：value除了字符串还可以是数字
  - 计数器
  - 统计多单位的数量
  - 粉丝数
  - 对象缓存存储

## List（列表）

在redis里面，我们可以把list玩成 ，栈、队列、阻塞队列！

所有的list命令都是用l开头的，Redis不区分大小命令

```reids
127 .0.0.1:6379> LPUSH list one  # 将一个值或者多个值，插入到列表头部 （左）
(integer) 1
127 .0.0.1:6379> LPUSH list two
(integer) 2
127 .0.0.1:6379> LPUSH list three
(integer) 3
127 .0.0.1:6379> LRANGE list 0 -1 # 获取list中值！
1 ) "three"
2 ) "two"
3 ) "one"
127 .0.0.1:6379> LRANGE list 0 1 # 通过区间获取具体的值！
1 ) "three"
2 ) "two"
127 .0.0.1:6379> Rpush list righr  # 将一个值或者多个值，插入到列表位部 （右）
(integer) 4
127 .0.0.1:6379> LRANGE list 0 -1
1 ) "three"
2 ) "two"
3 ) "one"
4 ) "righr"
##########################################################################
LPOP
RPOP
127 .0.0.1:6379> LRANGE list 0 -1
1 ) "three"
2 ) "two"
3 ) "one"
4 ) "righr"
127 .0.0.1:6379> Lpop list  # 移除list的第一个元素
"three"
127 .0.0.1:6379> Rpop list  # 移除list的最后一个元素
"righr"
127 .0.0.1:6379> LRANGE list 0 -1
1 ) "two"
2 ) "one"
##########################################################################
Lindex
127 .0.0.1:6379> LRANGE list 0 -1
1 ) "two"
2 ) "one"
127 .0.0.1:6379> lindex list 1 # 通过下标获得 list 中的某一个值！
"one"
127 .0.0.1:6379> lindex list 0
"two"
##########################################################################
Llen
127 .0.0.1:6379> Lpush list one
(integer) 1
127 .0.0.1:6379> Lpush list two
(integer) 2
127 .0.0.1:6379> Lpush list three
(integer) 3
127 .0.0.1:6379> Llen list # 返回列表的长度
(integer) 3
```



\##########################################################################
移除指定的值！
取关 uid

Lrem
127 .0.0.1:6379> LRANGE list 0 -1
1 ) "three"
2 ) "three"
3 ) "two"
4 ) "one"
127 .0.0.1:6379> lrem list 1 one # 移除list集合中指定个数的value，精确匹配
(integer) 1
127 .0.0.1:6379> LRANGE list 0 -1
1 ) "three"
2 ) "three"
3 ) "two"
127 .0.0.1:6379> lrem list 1 three
(integer) 1
127 .0.0.1:6379> LRANGE list 0 -1
1 ) "three"
2 ) "two"
127 .0.0.1:6379> Lpush list three
(integer) 3
127 .0.0.1:6379> lrem list 2 three
(integer) 2
127 .0.0.1:6379> LRANGE list 0 -1
1 ) "two"

\##########################################################################
trim 修剪。； list 截断!

127 .0.0.1:6379> keys *
(empty list or set)
127 .0.0.1:6379> Rpush mylist "hello"
(integer) 1
127 .0.0.1:6379> Rpush mylist "hello1"
(integer) 2
127 .0.0.1:6379> Rpush mylist "hello2"
(integer) 3
127 .0.0.1:6379> Rpush mylist "hello3"
(integer) 4
127 .0.0.1:6379> ltrim mylist 1 2 # 通过下标截取指定的长度，这个list已经被改变了，截断了
只剩下截取的元素！
OK
127 .0.0.1:6379> LRANGE mylist 0 -1
1 ) "hello1"
2 ) "hello2"

\##########################################################################
rpoplpush # 移除列表的最后一个元素，将他移动到新的列表中！

127 .0.0.1:6379> rpush mylist "hello"

```
他实际上是一个链表，before Node after ， left，right 都可以插入值
如果key 不存在，创建新的链表
如果key存在，新增内容
```

(integer) 1
127 .0.0.1:6379> rpush mylist "hello1"
(integer) 2
127 .0.0.1:6379> rpush mylist "hello2"
(integer) 3
127 .0.0.1:6379> rpoplpush mylist myotherlist # 移除列表的最后一个元素，将他移动到新的
列表中！
"hello2"
127 .0.0.1:6379> lrange mylist 0 -1 # 查看原来的列表
1 ) "hello"
2 ) "hello1"
127 .0.0.1:6379> lrange myotherlist 0 -1 # 查看目标列表中，确实存在改值！
1 ) "hello2"

\##########################################################################
lset 将列表中指定下标的值替换为另外一个值，更新操作
127 .0.0.1:6379> EXISTS list # 判断这个列表是否存在
(integer) 0
127 .0.0.1:6379> lset list 0 item # 如果不存在列表我们去更新就会报错
(error) ERR no such key
127 .0.0.1:6379> lpush list value1
(integer) 1
127 .0.0.1:6379> LRANGE list 0 0
1 ) "value1"
127 .0.0.1:6379> lset list 0 item # 如果存在，更新当前下标的值
OK
127 .0.0.1:6379> LRANGE list 0 0
1 ) "item"
127 .0.0.1:6379> lset list 1 other # 如果不存在，则会报错！
(error) ERR index out of range
\##########################################################################
linsert # 将某个具体的value插入到列把你中某个元素的前面或者后面！

127 .0.0.1:6379> Rpush mylist "hello"
(integer) 1
127 .0.0.1:6379> Rpush mylist "world"
(integer) 2
127 .0.0.1:6379> LINSERT mylist before "world" "other"
(integer) 3
127 .0.0.1:6379> LRANGE mylist 0 -1
1 ) "hello"
2 ) "other"
3 ) "world"
127 .0.0.1:6379> LINSERT mylist after world new
(integer) 4
127 .0.0.1:6379> LRANGE mylist 0 -1
1 ) "hello"
2 ) "other"
3 ) "world"
4 ) "new"

消息排队！消息队列 （Lpush Rpop）， 栈（ Lpush Lpop）！

## Set（集合）

set中的值是不能重读的！

```
##########################################################################
127 .0.0.1:6379> sadd myset "hello" # set集合中添加匀速
(integer) 1
127 .0.0.1:6379> sadd myset "kuangshen"
(integer) 1
127 .0.0.1:6379> sadd myset "lovekuangshen"
(integer) 1
127 .0.0.1:6379> SMEMBERS myset # 查看指定set的所有值
1 ) "hello"
2 ) "lovekuangshen"
3 ) "kuangshen"
127 .0.0.1:6379> SISMEMBER myset hello  # 判断某一个值是不是在set集合中！
(integer) 1
127 .0.0.1:6379> SISMEMBER myset world
(integer) 0
##########################################################################
127 .0.0.1:6379> scard myset  # 获取set集合中的内容元素个数！
(integer) 4
##########################################################################
rem
127 .0.0.1:6379> srem myset hello  # 移除set集合中的指定元素
(integer) 1
127 .0.0.1:6379> scard myset
(integer) 3
127 .0.0.1:6379> SMEMBERS myset
1 ) "lovekuangshen2"
2 ) "lovekuangshen"
3 ) "kuangshen"
##########################################################################
set 无序不重复集合。抽随机！
127 .0.0.1:6379> SMEMBERS myset
1 ) "lovekuangshen2"
2 ) "lovekuangshen"
3 ) "kuangshen"
127 .0.0.1:6379> SRANDMEMBER myset  # 随机抽选出一个元素
"kuangshen"
127 .0.0.1:6379> SRANDMEMBER myset
"kuangshen"
127 .0.0.1:6379> SRANDMEMBER myset
"kuangshen"
127 .0.0.1:6379> SRANDMEMBER myset
"kuangshen"
127 .0.0.1:6379> SRANDMEMBER myset 2 # 随机抽选出指定个数的元素
bilibili：狂神说Java
```

1 ) "lovekuangshen"
2 ) "lovekuangshen2"
127 .0.0.1:6379> SRANDMEMBER myset 2
1 ) "lovekuangshen"
2 ) "lovekuangshen2"
127 .0.0.1:6379> SRANDMEMBER myset # 随机抽选出一个元素
"lovekuangshen2"

\##########################################################################
删除定的key，随机删除key！

127 .0.0.1:6379> SMEMBERS myset
1 ) "lovekuangshen2"
2 ) "lovekuangshen"
3 ) "kuangshen"
127 .0.0.1:6379> spop myset # 随机删除一些set集合中的元素！
"lovekuangshen2"
127 .0.0.1:6379> spop myset
"lovekuangshen"
127 .0.0.1:6379> SMEMBERS myset
1 ) "kuangshen"

\##########################################################################
将一个指定的值，移动到另外一个set集合！
127 .0.0.1:6379> sadd myset "hello"
(integer) 1
127 .0.0.1:6379> sadd myset "world"
(integer) 1
127 .0.0.1:6379> sadd myset "kuangshen"
(integer) 1
127 .0.0.1:6379> sadd myset2 "set2"
(integer) 1
127 .0.0.1:6379> smove myset myset2 "kuangshen" # 将一个指定的值，移动到另外一个set集
合！
(integer) 1
127 .0.0.1:6379> SMEMBERS myset
1 ) "world"
2 ) "hello"
127 .0.0.1:6379> SMEMBERS myset2
1 ) "kuangshen"
2 ) "set2"

\##########################################################################
微博，B站，共同关注！(并集)
数字集合类：

- 差集 SDIFF
- 交集
- 并集
  127 .0.0.1:6379> SDIFF key1 key2 # 差集
  1 ) "b"
  2 ) "a"
  127 .0.0.1:6379> SINTER key1 key2 # 交集 共同好友就可以这样实现
  1 ) "c"
  127 .0.0.1:6379> SUNION key1 key2 # 并集
  1 ) "b"
  2 ) "c"
  3 ) "e"
  4 ) "a"

```
bilibili：狂神说Java
```

微博，A用户将所有关注的人放在一个set集合中！将它的粉丝也放在一个集合中！

共同关注，共同爱好，二度好友，推荐好友！（六度分割理论）

## Hash（哈希）

Map集合，key-map! 时候这个值是一个map集合！ 本质和String类型没有太大区别，还是一个简单的
key-vlaue！

set myhash field kuangshen

```
5 ) "d"
```

#### 

```
127 .0.0.1:6379> hset myhash field1 kuangshen  # set一个具体 key-vlaue
(integer) 1
127 .0.0.1:6379> hget myhash field1  # 获取一个字段值
"kuangshen"
127 .0.0.1:6379> hmset myhash field1 hello field2 world # set多个 key-vlaue
OK
127 .0.0.1:6379> hmget myhash field1 field2 # 获取多个字段值
1 ) "hello"
2 ) "world"
127 .0.0.1:6379> hgetall myhash # 获取全部的数据，
1 ) "field1"
2 ) "hello"
3 ) "field2"
4 ) "world"
127 .0.0.1:6379> hdel myhash field1  # 删除hash指定key字段！对应的value值也就消失了！
(integer) 1
127 .0.0.1:6379> hgetall myhash
1 ) "field2"
2 ) "world"
##########################################################################
hlen
127 .0.0.1:6379> hmset myhash field1 hello field2 world
OK
127 .0.0.1:6379> HGETALL myhash
1 ) "field2"
2 ) "world"
3 ) "field1"
4 ) "hello"
127 .0.0.1:6379> hlen myhash  # 获取hash表的字段数量！
(integer) 2
##########################################################################
127 .0.0.1:6379> HEXISTS myhash field1  # 判断hash中指定字段是否存在！
(integer) 1
127 .0.0.1:6379> HEXISTS myhash field3
(integer) 0
##########################################################################
# 只获得所有field
# 只获得所有value
127 .0.0.1:6379> hkeys myhash  # 只获得所有field
1 ) "field2"
2 ) "field1"
bilibili：狂神说Java
```

hash变更的数据 user name age,尤其是是用户信息之类的，经常变动的信息！ hash 更适合于对象的
存储，String更加适合字符串存储！

## Zset（有序集合）

在set的基础上，增加了一个值，set k1 v1 zset k1 score1 v1

```
127 .0.0.1:6379> hvals myhash  # 只获得所有value
1 ) "world"
2 ) "hello"
##########################################################################
incr decr
127 .0.0.1:6379> hset myhash field3 5 #指定增量！
(integer) 1
127 .0.0.1:6379> HINCRBY myhash field3 1
(integer) 6
127 .0.0.1:6379> HINCRBY myhash field3 -1
(integer) 5
127 .0.0.1:6379> hsetnx myhash field4 hello  # 如果不存在则可以设置
(integer) 1
127 .0.0.1:6379> hsetnx myhash field4 world  # 如果存在则不能设置
(integer) 0
127 .0.0.1:6379> zadd myset 1 one # 添加一个值
(integer) 1
127 .0.0.1:6379> zadd myset 2 two 3 three # 添加多个值
(integer) 2
127 .0.0.1:6379> ZRANGE myset 0 -1
1 ) "one"
2 ) "two"
3 ) "three"
##########################################################################
排序如何实现
127 .0.0.1:6379> zadd salary 2500 xiaohong  # 添加三个用户
(integer) 1
127 .0.0.1:6379> zadd salary 5000 zhangsan
(integer) 1
127 .0.0.1:6379> zadd salary 500 kaungshen
(integer) 1
# ZRANGEBYSCORE key min max
127 .0.0.1:6379> ZRANGEBYSCORE salary -inf +inf  # 显示全部的用户 从小到大！
1 ) "kaungshen"
2 ) "xiaohong"
3 ) "zhangsan"
127 .0.0.1:6379> ZREVRANGE salary 0 -1 # 从大到进行排序！
1 ) "zhangsan"
2 ) "kaungshen"
127 .0.0.1:6379> ZRANGEBYSCORE salary -inf +inf withscores # 显示全部的用户并且附带成
绩
1 ) "kaungshen"
2 ) "500"
3 ) "xiaohong"
bilibili：狂神说Java
```

#### 其与的一些API，通过我们的学习吗，你们剩下的如果工作中有需要，这个时候你可以去查查看官方文

#### 档！

案例思路：set 排序 存储班级成绩表，工资表排序！

普通消息， 1 ， 重要消息 2 ，带权重进行判断！

排行榜应用实现，取Top N 测试！

# 三种特殊数据类型

## Geospatial 地理位置

#### 朋友的定位，附近的人，打车距离计算？

Redis 的 Geo 在Redis3.2 版本就推出了！ 这个功能可以推算地理位置的信息，两地之间的距离，方圆
几里的人！

可以查询一些测试数据：http://www.jsons.cn/lngcodeinfo/0706D99C19A781A3/

只有 六个命令：

#### 4 ) "2500"

```
5 ) "zhangsan"
6 ) "5000"
127 .0.0.1:6379> ZRANGEBYSCORE salary -inf 2500 withscores # 显示工资小于 2500 员工的升
序排序！
1 ) "kaungshen"
2 ) "500"
3 ) "xiaohong"
4 ) "2500"
##########################################################################
# 移除rem中的元素
127 .0.0.1:6379> zrange salary 0 -1
1 ) "kaungshen"
2 ) "xiaohong"
3 ) "zhangsan"
127 .0.0.1:6379> zrem salary xiaohong # 移除有序集合中的指定元素
(integer) 1
127 .0.0.1:6379> zrange salary 0 -1
1 ) "kaungshen"
2 ) "zhangsan"
127 .0.0.1:6379> zcard salary  # 获取有序集合中的个数
(integer) 2
##########################################################################
127 .0.0.1:6379> zadd myset 1 hello
(integer) 1
127 .0.0.1:6379> zadd myset 2 world 3 kuangshen
(integer) 2
127 .0.0.1:6379> zcount myset 1 3 # 获取指定区间的成员数量！
(integer) 3
127 .0.0.1:6379> zcount myset 1 2
(integer) 2
bilibili：狂神说Java
```

#### 、

官方文档：https://www.redis.net.cn/order/3685.html

```
getadd
getpos
```

#### 获得当前定位：一定是一个坐标值！

#### GEODIST

```
# getadd 添加地理位置
# 规则：两级无法直接添加，我们一般会下载城市数据，直接通过java程序一次性导入！
# 有效的经度从-180度到 180 度。
# 有效的纬度从-85.05112878度到85.05112878度。
# 当坐标位置超出上述指定范围时，该命令将会返回一个错误。
# 127.0.0.1:6379> geoadd china:city 39.90 116.40 beijin
(error) ERR invalid longitude,latitude pair 39 .900000,116.400000
# 参数 key 值（）
127 .0.0.1:6379> geoadd china:city 116 .40 39 .90 beijing
(integer) 1
127 .0.0.1:6379> geoadd china:city 121 .47 31 .23 shanghai
(integer) 1
127 .0.0.1:6379> geoadd china:city 106 .50 29 .53 chongqi 114 .05 22 .52 shengzhen
(integer) 2
127 .0.0.1:6379> geoadd china:city 120 .16 30 .24 hangzhou 108 .96 34 .26 xian
(integer) 2
127 .0.0.1:6379> GEOPOS china:city beijing  # 获取指定的城市的经度和纬度！
1 ) 1 ) "116.39999896287918091"
2 ) "39.90000009167092543"
127 .0.0.1:6379> GEOPOS china:city beijing chongqi
1 ) 1 ) "116.39999896287918091"
2 ) "39.90000009167092543"
2 ) 1 ) "106.49999767541885376"
2 ) "29.52999957900659211"
bilibili：狂神说Java
```

#### 两人之间的距离！

#### 单位：

```
m 表示单位为米。
km 表示单位为千米。
mi 表示单位为英里。
ft 表示单位为英尺。
georadius 以给定的经纬度为中心， 找出某一半径内的元素
```

#### 我附近的人？ （获得所有附近的人的地址，定位！）通过半径来查询！

#### 获得指定数量的人， 200

所有数据应该都录入：china:city ，才会让结果更加请求！

```
127 .0.0.1:6379> GEODIST china:city beijing shanghai km  # 查看上海到北京的直线距离
"1067.3788"
127 .0.0.1:6379> GEODIST china:city beijing chongqi km # 查看重庆到北京的直线距离
"1464.0708"
127 .0.0.1:6379> GEORADIUS china:city 110 30 1000 km  # 以 110 ， 30 这个经纬度为中心，寻
找方圆1000km内的城市
1 ) "chongqi"
2 ) "xian"
3 ) "shengzhen"
4 ) "hangzhou"
127 .0.0.1:6379> GEORADIUS china:city 110 30 500 km
1 ) "chongqi"
2 ) "xian"
127 .0.0.1:6379> GEORADIUS china:city 110 30 500 km withdist  # 显示到中间距离的位置
1 ) 1 ) "chongqi"
2 ) "341.9374"
2 ) 1 ) "xian"
2 ) "483.8340"
127 .0.0.1:6379> GEORADIUS china:city 110 30 500 km withcoord  # 显示他人的定位信息
1 ) 1 ) "chongqi"
2 ) 1 ) "106.49999767541885376"
 2 ) "29.52999957900659211"
2 ) 1 ) "xian"
2 ) 1 ) "108.96000176668167114"
 2 ) "34.25999964418929977"
127 .0.0.1:6379> GEORADIUS china:city 110 30 500 km withdist withcoord count 1 #
筛选出指定的结果！
1 ) 1 ) "chongqi"
2 ) "341.9374"
3 ) 1 ) "106.49999767541885376"
 2 ) "29.52999957900659211"
127 .0.0.1:6379> GEORADIUS china:city 110 30 500 km withdist withcoord count 2
1 ) 1 ) "chongqi"
2 ) "341.9374"
3 ) 1 ) "106.49999767541885376"
 2 ) "29.52999957900659211"
2 ) 1 ) "xian"
2 ) "483.8340"
3 ) 1 ) "108.96000176668167114"
 2 ) "34.25999964418929977"
bilibili：狂神说Java
```

#### GEORADIUSBYMEMBER

```
GEOHASH 命令 - 返回一个或多个位置元素的 Geohash 表示
```

该命令将返回 11 个字符的Geohash字符串!

```
GEO 底层的实现原理其实就是 Zset！我们可以使用Zset命令来操作geo！
```

## Hyperloglog

#### 什么是基数？

#### A {1,3,5,7,8,7}

#### B{1，3,5,7,8}

#### 基数（不重复的元素） = 5，可以接受误差！

#### 简介

Redis 2.8.9 版本就更新了 Hyperloglog 数据结构！

Redis Hyperloglog 基数统计的算法！

#### # 找出位于指定元素周围的其他元素！

```
127 .0.0.1:6379> GEORADIUSBYMEMBER china:city beijing 1000 km
1 ) "beijing"
2 ) "xian"
127 .0.0.1:6379> GEORADIUSBYMEMBER china:city shanghai 400 km
1 ) "hangzhou"
2 ) "shanghai"
```

#### # 将二维的经纬度转换为一维的字符串，如果两个字符串越接近，那么则距离越近！

```
127 .0.0.1:6379> geohash china:city beijing chongqi
1 ) "wx4fbxxfke0"
2 ) "wm5xzrybty0"
127 .0.0.1:6379> ZRANGE china:city 0 -1 # 查看地图中全部的元素
1 ) "chongqi"
2 ) "xian"
3 ) "shengzhen"
4 ) "hangzhou"
5 ) "shanghai"
6 ) "beijing"
127 .0.0.1:6379> zrem china:city beijing  # 移除指定元素！
(integer) 1
127 .0.0.1:6379> ZRANGE china:city 0 -1
1 ) "chongqi"
2 ) "xian"
3 ) "shengzhen"
4 ) "hangzhou"
5 ) "shanghai"
bilibili：狂神说Java
```

#### 优点：占用的内存是固定，2^64 不同的元素的技术，只需要废 12KB内存！如果要从内存角度来比较的

话 Hyperloglog 首选！

**网页的 UV （一个人访问一个网站多次，但是还是算作一个人！）**

传统的方式， set 保存用户的id，然后就可以统计 set 中的元素数量作为标准判断!

这个方式如果保存大量的用户id，就会比较麻烦！我们的目的是为了计数，而不是保存用户id；

0.81% 错误率！ 统计UV任务，可以忽略不计的！

#### 测试使用

如果允许容错，那么一定可以使用 Hyperloglog ！

如果不允许容错，就使用 set 或者自己的数据类型即可！

## Bitmap

#### 为什么其他教程都不喜欢讲这些？这些在生活中或者开发中，都有十分多的应用场景，学习了，就是就

#### 是多一个思路！

#### 技多不压身！

#### 位存储

#### 统计用户信息，活跃，不活跃！ 登录 、 未登录！ 打卡， 365 打卡！ 两个状态的，都可以使用

Bitmaps！

Bitmap 位图，数据结构！ 都是操作二进制位来进行记录，就只有 0 和 1 两个状态！

365 天 = 365 bit 1字节 = 8bit 46 个字节左右！

#### 测试

```
127 .0.0.1:6379> PFadd mykey a b c d e f g h i j # 创建第一组元素 mykey
(integer) 1
127 .0.0.1:6379> PFCOUNT mykey  # 统计 mykey 元素的基数数量
(integer) 10
127 .0.0.1:6379> PFadd mykey2 i j z x c v b n m # 创建第二组元素 mykey2
(integer) 1
127 .0.0.1:6379> PFCOUNT mykey2
(integer) 9
127 .0.0.1:6379> PFMERGE mykey3 mykey mykey2  # 合并两组 mykey mykey2 => mykey3 并集
OK
127 .0.0.1:6379> PFCOUNT mykey3  # 看并集的数量！
(integer) 15
bilibili：狂神说Java
```

使用bitmap 来记录 周一到周日的打卡！

周一： 1 周二： 0 周三： 0 周四：1 ......

#### 查看某一天是否有打卡！

#### 统计操作，统计 打卡的天数！

# 事务

Redis 事务本质：一组命令的集合！ 一个事务中的所有命令都会被序列化，在事务执行过程的中，会按
照顺序执行！

一次性、顺序性、排他性！执行一些列的命令！

Redis事务没有没有隔离级别的概念！

所有的命令在事务中，并没有直接被执行！只有发起执行命令的时候才会执行！Exec

Redis单条命令式保存原子性的，但是事务不保证原子性！

redis的事务：

```
开启事务（multi）
命令入队（......）
执行事务（exec）
```

#### 正常执行事务！

```
127 .0.0.1:6379> getbit sign 3
(integer) 1
127 .0.0.1:6379> getbit sign 6
(integer) 0
127 .0.0.1:6379> bitcount sign  # 统计这周的打卡记录，就可以看到是否有全勤！
(integer) 3
------ 队列 set set set 执行------
127 .0.0.1:6379> multi  # 开启事务
OK
# 命令入队
127 .0.0.1:6379> set k1 v1
bilibili：狂神说Java
```

#### 放弃事务！

#### 编译型异常（代码有问题！ 命令有错！） ，事务中所有的命令都不会被执行！

#### 运行时异常（1/0）， 如果事务队列中存在语法性，那么执行命令的时候，其他命令是可以正常执行

#### 的，错误命令抛出异常！

#### QUEUED

127 .0.0.1:6379> set k2 v2
QUEUED
127 .0.0.1:6379> get k2
QUEUED
127 .0.0.1:6379> set k3 v3
QUEUED
127 .0.0.1:6379> exec # 执行事务
1 ) OK
2 ) OK
3 ) "v2"
4 ) OK

127 .0.0.1:6379> multi # 开启事务
OK
127 .0.0.1:6379> set k1 v1
QUEUED
127 .0.0.1:6379> set k2 v2
QUEUED
127 .0.0.1:6379> set k4 v4
QUEUED
127 .0.0.1:6379> DISCARD # 取消事务
OK
127 .0.0.1:6379> get k4 # 事务队列中命令都不会被执行！
(nil)

127 .0.0.1:6379> multi
OK
127 .0.0.1:6379> set k1 v1
QUEUED
127 .0.0.1:6379> set k2 v2
QUEUED
127 .0.0.1:6379> set k3 v3
QUEUED
127 .0.0.1:6379> getset k3 # 错误的命令
(error) ERR wrong number of arguments for 'getset' command
127 .0.0.1:6379> set k4 v4
QUEUED
127 .0.0.1:6379> set k5 v5
QUEUED
127 .0.0.1:6379> exec # 执行事务报错！
(error) EXECABORT Transaction discarded because of previous errors.
127 .0.0.1:6379> get k5 # 所有的命令都不会被执行！
(nil)

```
bilibili：狂神说Java
监控！ Watch （面试常问！）
```

#### 悲观锁：

#### 很悲观，认为什么时候都会出问题，无论做什么都会加锁！

#### 乐观锁：

#### 很乐观，认为什么时候都不会出问题，所以不会上锁！ 更新数据的时候去判断一下，在此期间是否

#### 有人修改过这个数据，

```
获取version
更新的时候比较 version
Redis测监视测试
```

#### 正常执行成功！

测试多线程修改值 , 使用watch 可以当做redis的乐观锁操作！

```
127 .0.0.1:6379> set k1 "v1"
OK
127 .0.0.1:6379> multi
OK
127 .0.0.1:6379> incr k1  # 会执行的时候失败！
QUEUED
127 .0.0.1:6379> set k2 v2
QUEUED
127 .0.0.1:6379> set k3 v3
QUEUED
127 .0.0.1:6379> get k3
QUEUED
127 .0.0.1:6379> exec
1 ) (error) ERR value is not an integer or out of range  # 虽然第一条命令报错了，但是
依旧正常执行成功了！
2 ) OK
3 ) OK
4 ) "v3"
127 .0.0.1:6379> get k2
"v2"
127 .0.0.1:6379> get k3
"v3"
127 .0.0.1:6379> set money 100
OK
127 .0.0.1:6379> set out 0
OK
127 .0.0.1:6379> watch money # 监视 money 对象
OK
127 .0.0.1:6379> multi # 事务正常结束，数据期间没有发生变动，这个时候就正常执行成功！
OK
127 .0.0.1:6379> DECRBY money 20
QUEUED
127 .0.0.1:6379> INCRBY out 20
QUEUED
127 .0.0.1:6379> exec
1 ) (integer) 80
2 ) (integer) 20
bilibili：狂神说Java
```

#### 如果修改失败，获取最新的值就好

# Jedis

我们要使用 Java 来操作 Redis，知其然并知其所以然，授人以渔！ 学习不能急躁，慢慢来会很快！

```
什么是Jedis 是 Redis 官方推荐的 java连接开发工具！ 使用Java 操作Redis 中间件！如果你要使用
java操作redis，那么一定要对Jedis 十分的熟悉！
```

#### 测试

#### 1 、导入对应的依赖

```
127 .0.0.1:6379> watch money # 监视 money
OK
127 .0.0.1:6379> multi
OK
127 .0.0.1:6379> DECRBY money 10
QUEUED
127 .0.0.1:6379> INCRBY out 10
QUEUED
127 .0.0.1:6379> exec  # 执行之前，另外一个线程，修改了我们的值，这个时候，就会导致事务执行失
败！
(nil)
<!--导入jedis的包-->
<dependencies>
<!-- https://mvnrepository.com/artifact/redis.clients/jedis -->
<dependency>
<groupId>redis.clients</groupId>
<artifactId>jedis</artifactId>
<version>3.2.0</version>
</dependency>
<!--fastjson-->
<dependency>
<groupId>com.alibaba</groupId>
<artifactId>fastjson</artifactId>
<version>1.2.62</version>
</dependency>
</dependencies>
bilibili：狂神说Java
```

#### 2 、编码测试：

#### 连接数据库

#### 操作命令

#### 断开连接！

#### 输出：

## 常用的API

String

List

Set

Hash

Zset

```
所有的api命令，就是我们对应的上面学习的指令，一个都没有变化！
```

#### 事务

```
package com.kuang;
import redis.clients.jedis.Jedis;
public class TestPing {
public static void main(String[] args) {
// 1、 new Jedis 对象即可
Jedis jedis = new Jedis("127.0.0.1", 6379 );
// jedis 所有的命令就是我们之前学习的所有指令！所以之前的指令学习很重要！
System.out.println(jedis.ping());
}
}
public class TestTX {
public static void main(String[] args) {
Jedis jedis = new Jedis("127.0.0.1", 6379 );
jedis.flushDB();
JSONObject jsonObject = new JSONObject();
jsonObject.put("hello","world");
jsonObject.put("name","kuangshen");
// 开启事务
Transaction multi = jedis.multi();
String result = jsonObject.toJSONString();
bilibili：狂神说Java
```

# SpringBoot整合

SpringBoot 操作数据：spring-data jpa jdbc mongodb redis！

SpringData 也是和 SpringBoot 齐名的项目！

说明： 在 SpringBoot2.x 之后，原来使用的jedis 被替换为了 lettuce?

jedis : 采用的直连，多个线程操作的话，是不安全的，如果想要避免不安全的，使用 jedis pool 连接
池！ 更像 BIO 模式

lettuce : 采用netty，实例可以再多个线程中进行共享，不存在线程不安全的情况！可以减少线程数据
了，更像 NIO 模式

源码分析：

```
// jedis.watch(result)
try {
multi.set("user1",result);
multi.set("user2",result);
int i = 1 / 0 ; // 代码抛出异常事务，执行失败！
multi.exec(); // 执行事务！
} catch (Exception e) {
multi.discard(); // 放弃事务
e.printStackTrace();
} finally {
System.out.println(jedis.get("user1"));
System.out.println(jedis.get("user2"));
jedis.close(); // 关闭连接
}
}
}
@Bean
@ConditionalOnMissingBean(name = "redisTemplate") // 我们可以自己定义一个
redisTemplate来替换这个默认的！
public RedisTemplate<Object, Object> redisTemplate(RedisConnectionFactory
redisConnectionFactory)
throws UnknownHostException {
// 默认的 RedisTemplate 没有过多的设置，redis 对象都是需要序列化！
// 两个泛型都是 Object, Object 的类型，我们后使用需要强制转换 <String, Object>
RedisTemplate<Object, Object> template = new RedisTemplate<>();
template.setConnectionFactory(redisConnectionFactory);
return template;
}
@Bean
@ConditionalOnMissingBean // 由于 String 是redis中最常使用的类型，所以说单独提出来了一
个bean！
public StringRedisTemplate stringRedisTemplate(RedisConnectionFactory
redisConnectionFactory)
throws UnknownHostException {
StringRedisTemplate template = new StringRedisTemplate();
template.setConnectionFactory(redisConnectionFactory);
return template;
bilibili：狂神说Java
```

#### 整合测试一下

#### 1 、导入依赖

#### 2 、配置连接

#### 3 、测试！

#### }

```
<!-- 操作redis -->
<dependency>
<groupId>org.springframework.boot</groupId>
<artifactId>spring-boot-starter-data-redis</artifactId>
</dependency>
# 配置redis
spring.redis.host=127.0.0.1
spring.redis.port= 6379
@SpringBootTest
class Redis02SpringbootApplicationTests {
@Autowired
private RedisTemplate redisTemplate;
@Test
void contextLoads() {
// redisTemplate 操作不同的数据类型，api和我们的指令是一样的
// opsForValue 操作字符串 类似String
// opsForList 操作List 类似List
// opsForSet
// opsForHash
// opsForZSet
// opsForGeo
// opsForHyperLogLog
// 除了进本的操作，我们常用的方法都可以直接通过redisTemplate操作，比如事务，和基本的
CRUD
// 获取redis的连接对象
// RedisConnection connection =
redisTemplate.getConnectionFactory().getConnection();
// connection.flushDb();
// connection.flushAll();
redisTemplate.opsForValue().set("mykey","关注狂神说公众号");
System.out.println(redisTemplate.opsForValue().get("mykey"));
}
}
bilibili：狂神说Java
```

#### 关于对象的保存：

我们来编写一个自己的 RedisTemplete

```
package com.kuang.config;
import com.fasterxml.jackson.annotation.JsonAutoDetect;
import com.fasterxml.jackson.annotation.PropertyAccessor;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.connection.RedisConnectionFactory;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.serializer.Jackson2JsonRedisSerializer;
import org.springframework.data.redis.serializer.StringRedisSerializer;
@Configuration
public class RedisConfig {
bilibili：狂神说Java
```

所有的redis操作，其实对于java开发人员来说，十分的简单，更重要是要去理解redis的思想和每一种数
据结构的用处和作用场景！

# Redis.conf详解

#### 启动的时候，就通过配置文件来启动！

#### 工作中，一些小小的配置，可以让你脱颖而出！

#### 行家有没有，出手就知道

#### 单位

#### // 这是我给大家写好的一个固定模板，大家在企业中，拿去就可以直接使用！

```
// 自己定义了一个 RedisTemplate
@Bean
@SuppressWarnings("all")
public RedisTemplate<String, Object> redisTemplate(RedisConnectionFactory
factory) {
// 我们为了自己开发方便，一般直接使用 <String, Object>
RedisTemplate<String, Object> template = new RedisTemplate<String,
Object>();
template.setConnectionFactory(factory);
// Json序列化配置
Jackson2JsonRedisSerializer jackson2JsonRedisSerializer = new
Jackson2JsonRedisSerializer(Object.class);
ObjectMapper om = new ObjectMapper();
om.setVisibility(PropertyAccessor.ALL, JsonAutoDetect.Visibility.ANY);
om.enableDefaultTyping(ObjectMapper.DefaultTyping.NON_FINAL);
jackson2JsonRedisSerializer.setObjectMapper(om);
// String 的序列化
StringRedisSerializer stringRedisSerializer = new
StringRedisSerializer();
// key采用String的序列化方式
template.setKeySerializer(stringRedisSerializer);
// hash的key也采用String的序列化方式
template.setHashKeySerializer(stringRedisSerializer);
// value序列化方式采用jackson
template.setValueSerializer(jackson2JsonRedisSerializer);
// hash的value序列化方式采用jackson
template.setHashValueSerializer(jackson2JsonRedisSerializer);
template.afterPropertiesSet();
return template;
}
}
bilibili：狂神说Java
```

1 、配置文件 unit单位 对大小写不敏感！

#### 包含

就是好比我们学习Spring、Improt， include

#### 网络

#### 通用 GENERAL

```
bind 127 .0.0.1  # 绑定的ip
protected-mode yes # 保护模式
port 6379 # 端口设置
daemonize yes # 以守护进程的方式运行，默认是 no，我们需要自己开启为yes！
pidfile /var/run/redis_6379.pid  # 如果以后台的方式运行，我们就需要指定一个 pid 文件！
# 日志
# Specify the server verbosity level.
# This can be one of:
bilibili：狂神说Java
```

#### 快照

持久化， 在规定的时间内，执行了多少次操作，则会持久化到文件 .rdb. aof

redis 是内存数据库，如果没有持久化，那么数据断电及失！

#### REPLICATION 复制，我们后面讲解主从复制的，时候再进行讲解

#### SECURITY 安全

可以在这里设置redis的密码，默认是没有密码！

```
# debug (a lot of information, useful for development/testing)
# verbose (many rarely useful info, but not a mess like the debug level)
# notice (moderately verbose, what you want in production probably) 生产环境
# warning (only very important / critical messages are logged)
loglevel notice
logfile "" # 日志的文件位置名
databases 16 # 数据库的数量，默认是 16 个数据库
always-show-logo yes # 是否总是显示LOGO
# 如果900s内，如果至少有一个1 key进行了修改，我们及进行持久化操作
save 900 1
# 如果300s内，如果至少10 key进行了修改，我们及进行持久化操作
save 300 10
# 如果60s内，如果至少10000 key进行了修改，我们及进行持久化操作
save 60 10000
# 我们之后学习持久化，会自己定义这个测试！
stop-writes-on-bgsave-error yes # 持久化如果出错，是否还需要继续工作！
rdbcompression yes # 是否压缩 rdb 文件，需要消耗一些cpu资源！
rdbchecksum yes # 保存rdb文件的时候，进行错误的检查校验！
dir ./  # rdb 文件保存的目录！
127 .0.0.1:6379> ping
PONG
127 .0.0.1:6379> config get requirepass # 获取redis的密码
1 ) "requirepass"
2 ) ""
127 .0.0.1:6379> config set requirepass "123456" # 设置redis的密码
OK
127 .0.0.1:6379> config get requirepass # 发现所有的命令都没有权限了
(error) NOAUTH Authentication required.
127 .0.0.1:6379> ping
(error) NOAUTH Authentication required.
127 .0.0.1:6379> auth 123456 # 使用密码进行登录！
OK
127 .0.0.1:6379> config get requirepass
1 ) "requirepass"
2 ) "123456"
bilibili：狂神说Java
```

#### 限制 CLIENTS

```
APPEND ONLY 模式 aof配置
```

具体的配置，我们在 Redis持久化 中去给大家详细详解！

# Redis持久化

#### 面试和工作，持久化都是重点！

Redis 是内存数据库，如果不将内存中的数据库状态保存到磁盘，那么一旦服务器进程退出，服务器中
的数据库状态也会消失。所以 Redis 提供了持久化功能！

## RDB（Redis DataBase）

#### 什么是RDB

在主从复制中，rdb就是备用了！从机上面！

```
maxclients 10000 # 设置能连接上redis的最大客户端的数量
maxmemory <bytes>  # redis 配置最大的内存容量
maxmemory-policy noeviction  # 内存到达上限之后的处理策略
 1 、volatile-lru：只对设置了过期时间的key进行LRU（默认值）
 2 、allkeys-lru ： 删除lru算法的key
 3 、volatile-random：随机删除即将过期key
 4 、allkeys-random：随机删除
 5 、volatile-ttl ： 删除即将过期的
 6 、noeviction ： 永不过期，返回错误
appendonly no  # 默认是不开启aof模式的，默认是使用rdb方式持久化的，在大部分所有的情况下，
rdb完全够用！
appendfilename "appendonly.aof" # 持久化的文件的名字
# appendfsync always # 每次修改都会 sync。消耗性能
appendfsync everysec # 每秒执行一次 sync，可能会丢失这1s的数据！
# appendfsync no # 不执行 sync，这个时候操作系统自己同步数据，速度最快！
bilibili：狂神说Java
```

在指定的时间间隔内将内存中的数据集快照写入磁盘，也就是行话讲的Snapshot快照，它恢复时是将快
照文件直接读到内存里。

Redis会单独创建（fork）一个子进程来进行持久化，会先将数据写入到一个临时文件中，待持久化过程
都结束了，再用这个临时文件替换上次持久化好的文件。整个过程中，主进程是不进行任何IO操作的。
这就确保了极高的性能。如果需要进行大规模数据的恢复，且对于数据恢复的完整性不是非常敏感，那
RDB方式要比AOF方式更加的高效。RDB的缺点是最后一次持久化后的数据可能丢失。我们默认的就是
RDB，一般情况下不需要修改这个配置！

有时候在生产环境我们会将这个文件进行备份！

rdb保存的文件是dump.rdb 都是在我们的配置文件中快照中进行配置的！

#### 触发机制

1 、save的规则满足的情况下，会自动触发rdb规则

2 、执行 flushall 命令，也会触发我们的rdb规则！

3 、退出redis，也会产生 rdb 文件！

备份就自动生成一个 dump.rdb

```
bilibili：狂神说Java
如果恢复rdb文件！
```

1 、只需要将rdb文件放在我们redis启动目录就可以，redis启动的时候会自动检查dump.rdb 恢复其中
的数据！

2 、查看需要存在的位置

#### 几乎就他自己默认的配置就够用了，但是我们还是需要去学习！

#### 优点：

#### 1 、适合大规模的数据恢复！

#### 2 、对数据的完整性要不高！

#### 缺点：

1 、需要一定的时间间隔进程操作！如果redis意外宕机了，这个最后一次修改数据就没有的了！

2 、fork进程的时候，会占用一定的内容空间！！

### AOF（Append Only File）

将我们的所有命令都记录下来，history，恢复的时候就把这个文件全部在执行一遍！

#### 是什么

```
127 .0.0.1:6379> config get dir
1 ) "dir"
2 ) "/usr/local/bin" # 如果在这个目录下存在 dump.rdb 文件，启动就会自动恢复其中的数据
bilibili：狂神说Java
```

以日志的形式来记录每个写操作，将Redis执行过的所有指令记录下来（读操作不记录），只许追加文件
但不可以改写文件，redis启动之初会读取该文件重新构建数据，换言之，redis重启的话就根据日志文件
的内容将写指令从前到后执行一次以完成数据的恢复工作

Aof保存的是 appendonly.aof 文件

```
append
```

默认是不开启的，我们需要手动进行配置！我们只需要将 appendonly 改为yes就开启了 aof！

重启，redis 就可以生效了！

如果这个 aof 文件有错位，这时候 redis 是启动不起来的吗，我们需要修复这个aof文件

redis 给我们提供了一个工具 redis-check-aof --fix

```
bilibili：狂神说Java
```

#### 如果文件正常，重启就可以直接恢复了！

#### 重写规则说明

aof 默认就是文件的无限追加，文件会越来越大！

如果 aof 文件大于 64m，太大了！ fork一个新的进程来将我们的文件进行重写！

#### 优点和缺点！

```
bilibili：狂神说Java
```

#### 优点：

#### 1 、每一次修改都同步，文件的完整会更加好！

#### 2 、每秒同步一次，可能会丢失一秒的数据

#### 3 、从不同步，效率最高的！

#### 缺点：

1 、相对于数据文件来说，aof远远大于 rdb，修复的速度也比 rdb慢！

2 、Aof 运行效率也要比 rdb 慢，所以我们redis默认的配置就是rdb持久化！

#### 扩展：

#### 1 、RDB 持久化方式能够在指定的时间间隔内对你的数据进行快照存储

#### 2 、AOF 持久化方式记录每次对服务器写的操作，当服务器重启的时候会重新执行这些命令来恢复原始

的数据，AOF命令以Redis 协议追加保存每次写的操作到文件末尾，Redis还能对AOF文件进行后台重
写，使得AOF文件的体积不至于过大。

3 、只做缓存，如果你只希望你的数据在服务器运行的时候存在，你也可以不使用任何持久化

4 、同时开启两种持久化方式

```
在这种情况下，当redis重启的时候会优先载入AOF文件来恢复原始的数据，因为在通常情况下AOF
文件保存的数据集要比RDB文件保存的数据集要完整。
RDB 的数据不实时，同时使用两者时服务器重启也只会找AOF文件，那要不要只使用AOF呢？作者
建议不要，因为RDB更适合用于备份数据库（AOF在不断变化不好备份），快速重启，而且不会有
AOF可能潜在的Bug，留着作为一个万一的手段。
```

5 、性能建议

```
因为RDB文件只用作后备用途，建议只在Slave上持久化RDB文件，而且只要 15 分钟备份一次就够
了，只保留 save 900 1 这条规则。
如果Enable AOF ，好处是在最恶劣情况下也只会丢失不超过两秒数据，启动脚本较简单只load自
己的AOF文件就可以了，代价一是带来了持续的IO，二是AOF rewrite 的最后将 rewrite 过程中产
生的新数据写到新文件造成的阻塞几乎是不可避免的。只要硬盘许可，应该尽量减少AOF rewrite
的频率，AOF重写的基础大小默认值64M太小了，可以设到5G以上，默认超过原大小100%大小重
写可以改到适当的数值。
如果不Enable AOF ，仅靠 Master-Slave Repllcation 实现高可用性也可以，能省掉一大笔IO，也
减少了rewrite时带来的系统波动。代价是如果Master/Slave 同时倒掉，会丢失十几分钟的数据，
启动脚本也要比较两个 Master/Slave 中的 RDB文件，载入较新的那个，微博就是这种架构。
```

# Redis发布订阅

```
appendonly no  # 默认是不开启aof模式的，默认是使用rdb方式持久化的，在大部分所有的情况下，
rdb完全够用！
appendfilename "appendonly.aof" # 持久化的文件的名字
# appendfsync always # 每次修改都会 sync。消耗性能
appendfsync everysec # 每秒执行一次 sync，可能会丢失这1s的数据！
# appendfsync no # 不执行 sync，这个时候操作系统自己同步数据，速度最快！
# rewrite 重写，
bilibili：狂神说Java
```

Redis 发布订阅(pub/sub)是一种消息通信模式：发送者(pub)发送消息，订阅者(sub)接收消息。微信、
微博、关注系统！

Redis 客户端可以订阅任意数量的频道。

订阅/发布消息图：

第一个：消息发送者， 第二个：频道 第三个：消息订阅者！

下图展示了频道 channel1 ， 以及订阅这个频道的三个客户端 —— client2 、 client5 和 client1 之间的
关系：

当有新消息通过 PUBLISH 命令发送给频道 channel1 时， 这个消息就会被发送给订阅它的三个客户
端：

```
bilibili：狂神说Java
```

#### 命令

这些命令被广泛用于构建即时通信应用，比如网络聊天室(chatroom)和实时广播、实时提醒等。

#### 测试

#### 订阅端：

#### 发送端：

#### 原理

Redis是使用C实现的，通过分析 Redis 源码里的 pubsub.c 文件，了解发布和订阅机制的底层实现，籍
此加深对 Redis 的理解。

```
127 .0.0.1:6379> SUBSCRIBE kuangshenshuo  # 订阅一个频道 kuangshenshuo
Reading messages... (press Ctrl-C to quit)
1 ) "subscribe"
2 ) "kuangshenshuo"
3 ) (integer) 1
# 等待读取推送的信息
1 ) "message" # 消息
2 ) "kuangshenshuo" # 那个频道的消息
3 ) "hello,kuangshen" # 消息的具体内容
1 ) "message"
2 ) "kuangshenshuo"
3 ) "hello,redis"
127 .0.0.1:6379> PUBLISH kuangshenshuo "hello,kuangshen" # 发布者发布消息到频道！
(integer) 1
127 .0.0.1:6379> PUBLISH kuangshenshuo "hello,redis" # 发布者发布消息到频道！
(integer) 1
127 .0.0.1:6379>
bilibili：狂神说Java
```

Redis 通过 PUBLISH 、SUBSCRIBE 和 PSUBSCRIBE 等命令实现发布和订阅功能。

微信：

通过 SUBSCRIBE 命令订阅某频道后，redis-server 里维护了一个字典，字典的键就是一个个 频道！，
而字典的值则是一个链表，链表中保存了所有订阅这个 channel 的客户端。SUBSCRIBE 命令的关键，
就是将客户端添加到给定 channel 的订阅链表中。

通过 PUBLISH 命令向订阅者发送消息，redis-server 会使用给定的频道作为键，在它所维护的 channel
字典中查找记录了订阅这个频道的所有客户端的链表，遍历这个链表，将消息发布给所有订阅者。

Pub/Sub 从字面上理解就是发布（Publish）与订阅（Subscribe），在Redis中，你可以设定对某一个
key值进行消息发布及消息订阅，当一个key值上进行了消息发布后，所有订阅它的客户端都会收到相应
的消息。这一功能最明显的用法就是用作实时消息系统，比如普通的即时聊天，群聊等功能。

**使用场景：**

1 、实时消息系统！

2 、事实聊天！（频道当做聊天室，将信息回显给所有人即可！）

3 、订阅，关注系统都是可以的！

稍微复杂的场景我们就会使用 消息中间件 MQ （）

# Redis主从复制

```
bilibili：狂神说Java
```

### 概念

主从复制，是指将一台Redis服务器的数据，复制到其他的Redis服务器。前者称为主节点
(master/leader)，后者称为从节点(slave/follower)；数据的复制是单向的，只能由主节点到从节点。
Master以写为主，Slave 以读为主。

默认情况下，每台Redis服务器都是主节点；

且一个主节点可以有多个从节点(或没有从节点)，但一个从节点只能有一个主节点。（）

**主从复制的作用主要包括：**

1 、数据冗余：主从复制实现了数据的热备份，是持久化之外的一种数据冗余方式。

2 、故障恢复：当主节点出现问题时，可以由从节点提供服务，实现快速的故障恢复；实际上是一种服务
的冗余。

3 、负载均衡：在主从复制的基础上，配合读写分离，可以由主节点提供写服务，由从节点提供读服务
（即写Redis数据时应用连接主节点，读Redis数据时应用连接从节点），分担服务器负载；尤其是在写
少读多的场景下，通过多个从节点分担读负载，可以大大提高Redis服务器的并发量。

4 、高可用（集群）基石：除了上述作用以外，主从复制还是哨兵和集群能够实施的基础，因此说主从复
制是Redis高可用的基础。

一般来说，要将Redis运用于工程项目中，只使用一台Redis是万万不能的（宕机），原因如下：

1 、从结构上，单个Redis服务器会发生单点故障，并且一台服务器需要处理所有的请求负载，压力较
大；

2 、从容量上，单个Redis服务器内存容量有限，就算一台Redis服务器内存容量为256G，也不能将所有
内存用作Redis存储内存，一般来说，单台Redis最大使用内存不应该超过20G。

电商网站上的商品，一般都是一次上传，无数次浏览的，说专业点也就是"多读少写"。

对于这种场景，我们可以使如下这种架构：

```
bilibili：狂神说Java
```

#### 主从复制，读写分离！ 80% 的情况下都是在进行读操作！减缓服务器的压力！架构中经常使用！ 一主

#### 二从！

只要在公司中，主从复制就是必须要使用的，因为在真实的项目中不可能单机使用Redis！

## 环境配置

#### 只配置从库，不用配置主库！

#### 复制 3 个配置文件，然后修改对应的信息

#### 1 、端口

2 、pid 名字

3 、log文件名字

4 、dump.rdb 名字

修改完毕之后，启动我们的 3 个redis服务器，可以通过进程信息查看！

## 一主二从

默认情况下，每台Redis服务器都是主节点； 我们一般情况下只用配置从机就好了！

认老大！ 一主 （ 79 ）二从（ 80 ， 81 ）

```
127 .0.0.1:6379> info replication # 查看当前库的信息
# Replication
role:master  # 角色 master
connected_slaves:0 # 没有从机
master_replid:b63c90e6c501143759cb0e7f450bd1eb0c70882a
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:0
second_repl_offset:-1
repl_backlog_active:0
repl_backlog_size:1048576
repl_backlog_first_byte_offset:0
repl_backlog_histlen:0
127 .0.0.1:6380> SLAVEOF 127 .0.0.1 6379 # SLAVEOF host 6379 找谁当自己的老大！
OK
127 .0.0.1:6380> info replication
# Replication
role:slave  # 当前角色是从机
master_host:127.0.0.1 # 可以的看到主机的信息
master_port:6379
bilibili：狂神说Java
```

#### 如果两个都配置完了，就是有两个从机的

#### 真实的从主配置应该在配置文件中配置，这样的话是永久的，我们这里使用的是命令，暂时的！

#### 细节

#### 主机可以写，从机不能写只能读！主机中的所有信息和数据，都会自动被从机保存！

#### 主机写：

```
master_link_status:up
master_last_io_seconds_ago:3
master_sync_in_progress:0
slave_repl_offset:14
slave_priority:100
slave_read_only:1
connected_slaves:0
master_replid:a81be8dd257636b2d3e7a9f595e69d73ff03774e
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:14
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:14
# 在主机中查看！
127 .0.0.1:6379> info replication
# Replication
role:master
connected_slaves:1  # 多了从机的配置
slave0:ip= 127 .0.0.1,port= 6380 ,state=online,offset= 42 ,lag= 1 # 多了从机的配置
master_replid:a81be8dd257636b2d3e7a9f595e69d73ff03774e
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:42
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:42
bilibili：狂神说Java
```

#### 从机只能读取内容！

#### 测试：主机断开连接，从机依旧连接到主机的，但是没有写操作，这个时候，主机如果回来了，从机依

#### 旧可以直接获取到主机写的信息！

#### 如果是使用命令行，来配置的主从，这个时候如果重启了，就会变回主机！只要变为从机，立马就会从

#### 主机中获取值！

#### 复制原理

Slave 启动成功连接到 master 后会发送一个sync同步命令

Master 接到命令，启动后台的存盘进程，同时收集所有接收到的用于修改数据集命令，在后台进程执行
完毕之后，master将传送整个数据文件到slave，并完成一次完全同步。

全量复制：而slave服务在接收到数据库文件数据后，将其存盘并加载到内存中。

增量复制：Master 继续将新的所有收集到的修改命令依次传给slave，完成同步

但是只要是重新连接master，一次完全同步（全量复制）将被自动执行！ 我们的数据一定可以在从机中
看到！

#### 层层链路

#### 上一个M链接下一个 S！

#### 这时候也可以完成我们的主从复制！

```
bilibili：狂神说Java
```

#### 如果没有老大了，这个时候能不能选择一个老大出来呢？ 手动！

#### 谋朝篡位

如果主机断开了连接，我们可以使用 SLAVEOF no one 让自己变成主机！其他的节点就可以手动连
接到最新的这个主节点（手动）！如果这个时候老大修复了，那就重新连接！

## 哨兵模式

#### （自动选举老大的模式）

#### 概述

#### 主从切换技术的方法是：当主服务器宕机后，需要手动把一台从服务器切换为主服务器，这就需要人工

#### 干预，费事费力，还会造成一段时间内服务不可用。这不是一种推荐的方式，更多时候，我们优先考虑

哨兵模式。Redis从2.8开始正式提供了Sentinel（哨兵） 架构来解决这个问题。

谋朝篡位的自动版，能够后台监控主机是否故障，如果故障了根据投票数自动将从库转换为主库。

哨兵模式是一种特殊的模式，首先Redis提供了哨兵的命令，哨兵是一个独立的进程，作为进程，它会独
立运行。其原理是 **哨兵通过发送命令，等待Redis服务器响应，从而监控运行的多个Redis实例。**

#### 这里的哨兵有两个作用

```
通过发送命令，让Redis服务器返回监控其运行状态，包括主服务器和从服务器。
当哨兵监测到master宕机，会自动将slave切换成master，然后通过 发布订阅模式 通知其他的从服
务器，修改配置文件，让它们切换主机。
```

然而一个哨兵进程对Redis服务器进行监控，可能会出现问题，为此，我们可以使用多个哨兵进行监控。
各个哨兵之间还会进行监控，这样就形成了多哨兵模式。

```
bilibili：狂神说Java
```

假设主服务器宕机，哨兵 1 先检测到这个结果，系统并不会马上进行failover过程，仅仅是哨兵 1 主观的认
为主服务器不可用，这个现象成为 **主观下线** 。当后面的哨兵也检测到主服务器不可用，并且数量达到一
定值时，那么哨兵之间就会进行一次投票，投票的结果由一个哨兵发起，进行failover[故障转移]操作。
切换成功后，就会通过发布订阅模式，让各个哨兵把自己监控的从服务器实现切换主机，这个过程称为
**客观下线** 。

#### 测试！

#### 我们目前的状态是 一主二从！

1 、配置哨兵配置文件 sentinel.conf

后面的这个数字 1 ，代表主机挂了，slave投票看让谁接替成为主机，票数最多的，就会成为主机！

2 、启动哨兵！

```
# sentinel monitor 被监控的名称 host port 1
sentinel monitor myredis 127 .0.0.1 6379 1
[root@kuangshen bin]# redis-sentinel kconfig/sentinel.conf
26607 :X 31 Mar 2020 21 :13:10.027 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
26607 :X 31 Mar 2020 21 :13:10.027 # Redis version=5.0.8, bits=64,
commit=00000000, modified=0, pid=26607, just started
26607 :X 31 Mar 2020 21 :13:10.027 # Configuration loaded
_._
_.-``__ ''-._
_.-`` `. `_.  ''-._ Redis 5 .0.8 (00000000/0) 64 bit
.-`` .-```. ```\/ _.,_ ''-._
(  ' , .-` | `, ) Running in sentinel mode
|`-._`-...-` __...-.``-._|'` _.-'| Port: 26379
|  `-._ `._ / _.-' | PID: 26607
`-._ `-._  `-./ _.-' _.-'
|`-._`-._  `-.__.-' _.-'_.-'|
|  `-._`-._ _.-'_.-' | http://redis.io
`-._ `-._`-.__.-'_.-' _.-'
|`-._`-._  `-.__.-' _.-'_.-'|
|  `-._`-._ _.-'_.-' |
`-._ `-._`-.__.-'_.-' _.-'
  `-._`-.__.-' _.-'
     `-._ _.-'
     `-.__.-'
bilibili：狂神说Java
```

如果Master 节点断开了，这个时候就会从从机中随机选择一个服务器！ （这里面有一个投票算法！）

#### 哨兵日志！

#### 如果主机此时回来了，只能归并到新的主机下，当做从机，这就是哨兵模式的规则！



```
26607 :X 31 Mar 2020 21 :13:10.029 # WARNING: The TCP backlog setting of 511
cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value
of 128.
26607 :X 31 Mar 2020 21 :13:10.031 # Sentinel ID is
4c780da7e22d2aebe3bc20c333746f202ce72996
26607 :X 31 Mar 2020 21 :13:10.031 # +monitor master myredis 127.0.0.1 6379 quorum
1
26607 :X 31 Mar 2020 21 :13:10.031 * +slave slave 127 .0.0.1:6380 127 .0.0.1 6380 @
myredis 127 .0.0.1 6379
26607 :X 31 Mar 2020 21 :13:10.033 * +slave slave 127 .0.0.1:6381 127 .0.0.1 6381 @
myredis 127 .0.0.1 6379
bilibili：狂神说Java
```

#### 哨兵模式

#### 优点：

#### 1 、哨兵集群，基于主从复制模式，所有的主从配置优点，它全有

#### 2 、 主从可以切换，故障可以转移，系统的可用性就会更好

#### 3 、哨兵模式就是主从模式的升级，手动到自动，更加健壮！

#### 缺点：

1 、Redis 不好啊在线扩容的，集群容量一旦到达上限，在线扩容就十分麻烦！

2 、实现哨兵模式的配置其实是很麻烦的，里面有很多选择！

#### 哨兵模式的全部配置！

```
# Example sentinel.conf
# 哨兵sentinel实例运行的端口 默认 26379
port 26379
# 哨兵sentinel的工作目录
dir /tmp
# 哨兵sentinel监控的redis主节点的 ip port
# master-name 可以自己命名的主节点名字 只能由字母A-z、数字0-9 、这三个字符".-_"组成。
# quorum 配置多少个sentinel哨兵统一认为master主节点失联 那么这时客观上认为主节点失联了
# sentinel monitor <master-name> <ip> <redis-port> <quorum>
sentinel monitor mymaster 127 .0.0.1 6379 2
# 当在Redis实例中开启了requirepass foobared 授权密码 这样所有连接Redis实例的客户端都要提供
密码
# 设置哨兵sentinel 连接主从的密码 注意必须为主从设置一样的验证密码
# sentinel auth-pass <master-name> <password>
sentinel auth-pass mymaster MySUPER--secret-0123passw0rd
# 指定多少毫秒之后 主节点没有应答哨兵sentinel 此时 哨兵主观上认为主节点下线 默认 30 秒
# sentinel down-after-milliseconds <master-name> <milliseconds>
sentinel down-after-milliseconds mymaster 30000
# 这个配置项指定了在发生failover主备切换时最多可以有多少个slave同时对新的master进行 同步，
这个数字越小，完成failover所需的时间就越长，
但是如果这个数字越大，就意味着越 多的slave因为replication而不可用。
可以通过将这个值设为 1 来保证每次只有一个slave 处于不能处理命令请求的状态。
# sentinel parallel-syncs <master-name> <numslaves>
sentinel parallel-syncs mymaster 1
# 故障转移的超时时间 failover-timeout 可以用在以下这些方面：
#1. 同一个sentinel对同一个master两次failover之间的间隔时间。
#2. 当一个slave从一个错误的master那里同步数据开始计算时间。直到slave被纠正为向正确的master那
里同步数据时。
#3.当想要取消一个正在进行的failover所需要的时间。
#4.当进行failover时，配置所有slaves指向新的master所需的最大时间。不过，即使过了这个超时，
slaves依然会被正确配置为指向master，但是就不按parallel-syncs所配置的规则来了
# 默认三分钟
# sentinel failover-timeout <master-name> <milliseconds>
bilibili：狂神说Java
```

#### 社会目前程序员饱和（初级和中级）、高级程序员重金难求！（提升自己！）

# Redis缓存穿透和雪崩

#### 服务的高可用问题！

#### 在这里我们不会详细的区分析解决方案的底层！

Redis缓存的使用，极大的提升了应用程序的性能和效率，特别是数据查询方面。但同时，它也带来了一
些问题。其中，最要害的问题，就是数据的一致性问题，从严格意义上讲，这个问题无解。如果对数据
的一致性要求很高，那么就不能使用缓存。

另外的一些典型问题就是，缓存穿透、缓存雪崩和缓存击穿。目前，业界也都有比较流行的解决方案。

```
sentinel failover-timeout mymaster 180000
# SCRIPTS EXECUTION
#配置当某一事件发生时所需要执行的脚本，可以通过脚本来通知管理员，例如当系统运行不正常时发邮件通知
相关人员。
#对于脚本的运行结果有以下规则：
#若脚本执行后返回 1 ，那么该脚本稍后将会被再次执行，重复次数目前默认为 10
#若脚本执行后返回 2 ，或者比 2 更高的一个返回值，脚本将不会重复执行。
#如果脚本在执行过程中由于收到系统中断信号被终止了，则同返回值为 1 时的行为相同。
#一个脚本的最大执行时间为60s，如果超过这个时间，脚本将会被一个SIGKILL信号终止，之后重新执行。
#通知型脚本:当sentinel有任何警告级别的事件发生时（比如说redis实例的主观失效和客观失效等等），
将会去调用这个脚本，这时这个脚本应该通过邮件，SMS等方式去通知系统管理员关于系统不正常运行的信
息。调用该脚本时，将传给脚本两个参数，一个是事件的类型，一个是事件的描述。如果sentinel.conf配
置文件中配置了这个脚本路径，那么必须保证这个脚本存在于这个路径，并且是可执行的，否则sentinel无
法正常启动成功。
#通知脚本
# shell编程
# sentinel notification-script <master-name> <script-path>
sentinel notification-script mymaster /var/redis/notify.sh
# 客户端重新配置主节点参数脚本
# 当一个master由于failover而发生改变时，这个脚本将会被调用，通知相关的客户端关于master地址已
经发生改变的信息。
# 以下参数将会在调用脚本时传给脚本:
# <master-name> <role> <state> <from-ip> <from-port> <to-ip> <to-port>
# 目前<state>总是“failover”,
# <role>是“leader”或者“observer”中的一个。
# 参数 from-ip, from-port, to-ip, to-port是用来和旧的master和新的master(即旧的slave)通
信的
# 这个脚本应该是通用的，能被多次调用，不是针对性的。
# sentinel client-reconfig-script <master-name> <script-path>
sentinel client-reconfig-script mymaster /var/redis/reconfig.sh # 一般都是由运维来配
置！
bilibili：狂神说Java
```

### 缓存穿透（查不到）

#### 概念

缓存穿透的概念很简单，用户想要查询一个数据，发现redis内存数据库没有，也就是缓存没有命中，于
是向持久层数据库查询。发现也没有，于是本次查询失败。当用户很多的时候，缓存都没有命中（秒
杀！），于是都去请求了持久层数据库。这会给持久层数据库造成很大的压力，这时候就相当于出现了
缓存穿透。

#### 解决方案

#### 布隆过滤器

布隆过滤器是一种数据结构，对所有可能查询的参数以hash形式存储，在控制层先进行校验，不符合则
丢弃，从而避免了对底层存储系统的查询压力；

#### 缓存空对象

```
bilibili：狂神说Java
```

#### 当存储层不命中后，即使返回的空对象也将其缓存起来，同时会设置一个过期时间，之后再访问这个数

#### 据将会从缓存中获取，保护了后端数据源；

#### 但是这种方法会存在两个问题：

#### 1 、如果空值能够被缓存起来，这就意味着缓存需要更多的空间存储更多的键，因为这当中可能会有很多

#### 的空值的键；

#### 2 、即使对空值设置了过期时间，还是会存在缓存层和存储层的数据会有一段时间窗口的不一致，这对于

#### 需要保持一致性的业务会有影响。

### 缓存击穿（量太大，缓存过期！）

#### 概述

这里需要注意和缓存击穿的区别，缓存击穿，是指一个key非常热点，在不停的扛着大并发，大并发集中
对这一个点进行访问，当这个key在失效的瞬间，持续的大并发就穿破缓存，直接请求数据库，就像在一
个屏障上凿开了一个洞。

当某个key在过期的瞬间，有大量的请求并发访问，这类数据一般是热点数据，由于缓存过期，会同时访
问数据库来查询最新数据，并且回写缓存，会导使数据库瞬间压力过大。

#### 解决方案

#### 设置热点数据永不过期

从缓存层面来看，没有设置过期时间，所以不会出现热点 key 过期后产生的问题。

**加互斥锁**

分布式锁：使用分布式锁，保证对于每个key同时只有一个线程去查询后端服务，其他线程没有获得分布
式锁的权限，因此只需要等待即可。这种方式将高并发的压力转移到了分布式锁，因此对分布式锁的考
验很大。

```
bilibili：狂神说Java
```

### 缓存雪崩

#### 概念

缓存雪崩，是指在某一个时间段，缓存集中过期失效。Redis 宕机！

产生雪崩的原因之一，比如在写本文的时候，马上就要到双十二零点，很快就会迎来一波抢购，这波商
品时间比较集中的放入了缓存，假设缓存一个小时。那么到了凌晨一点钟的时候，这批商品的缓存就都
过期了。而对这批商品的访问查询，都落到了数据库上，对于数据库而言，就会产生周期性的压力波

峰。于是所有的请求都会达到存储层，存储层的调用量会暴增，造成存储层也会挂掉的情况。

#### 其实集中过期，倒不是非常致命，比较致命的缓存雪崩，是缓存服务器某个节点宕机或断网。因为自然

#### 形成的缓存雪崩，一定是在某个时间段集中创建缓存，这个时候，数据库也是可以顶住压力的。无非就

#### 是对数据库产生周期性的压力而已。而缓存服务节点的宕机，对数据库服务器造成的压力是不可预知

#### 的，很有可能瞬间就把数据库压垮。

#### 解决方案

```
bilibili：狂神说Java
```

**redis高可用**

这个思想的含义是，既然redis有可能挂掉，那我多增设几台redis，这样一台挂掉之后其他的还可以继续
工作，其实就是搭建的集群。（异地多活！）

**限流降级（在SpringCloud讲解过！）**

这个解决方案的思想是，在缓存失效后，通过加锁或者队列来控制读数据库写缓存的线程数量。比如对
某个key只允许一个线程查询数据和写缓存，其他线程等待。

**数据预热**

数据加热的含义就是在正式部署之前，我先把可能的数据先预先访问一遍，这样部分可能大量访问的数
据就会加载到缓存中。在即将发生大并发访问前手动触发加载缓存不同的key，设置不同的过期时间，让
缓存失效的时间点尽量均匀。