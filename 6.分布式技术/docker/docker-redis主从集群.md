# 一、redis主节点

- 启动主节点

  ```shell
  docker run -p 6379:6379 --name redis-master ^
  -v /d/gaozhen/code-source/study-note/8.工具配置/redis/master-slave/master/redis.conf:/etc/redis/redis.conf ^
  -v /d/gaozhen/code-source/study-note/8.工具配置/redis/master-slave/master/:/data ^
  -d redis:5.0.5 redis-server /etc/redis/redis.conf
  ```

  **配置文件的路径需要修改**

- 查看主节点IP

  ```
  docker inspect redis-master
  ```

# 二、redis从节点

- 配置文件修改

  ```
  slaveof 172.17.0.2 6379
  ```

  172.17.0.2 是主节点的IP

  6379是主节点的端口

- redis-cluster1

  ```shell
  docker run -p 6380:6379 --name redis-cluster1 ^
  -v /d/gaozhen/code-source/study-note/8.工具配置/redis/master-slave/slave1/redis.conf:/etc/redis/redis.conf ^
  -v /d/gaozhen/code-source/study-note/8.工具配置/redis/master-slave/slave1/:/data ^
  -d redis:5.0.5 redis-server /etc/redis/redis.conf 
  ```
  
- redis-cluster2

  ```shell
  docker run -p 6381:6379 --name redis-cluster2 ^
  -v /d/gaozhen/code-source/study-note/8.工具配置/redis/master-slave/slave2/redis.conf:/etc/redis/redis.conf ^
  -v /d/gaozhen/code-source/study-note/8.工具配置/redis/master-slave/slave2/:/data ^
  -d redis:5.0.5 redis-server /etc/redis/redis.conf 
  ```

# 三、测试主从数据同步

