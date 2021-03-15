http://www.cnblogs.com/xiuyangleiasp/p/5010311.html

http://blog.csdn.net/ly52352148/article/details/53943979

```language
JAVA_HOME=/usr/local/java
export JAVA_HOME
HADOOP_HOME=/usr/local/hadoop
export HADOOP_HOME
PATH=$JAVA_HOME/bin:$JAVA_HOME/jre/bin:$HADOOP_HOME/bin:$HADOOP_HOME/sbin:$PATH
```
保存文件后
```language
source /etc/profile
```
立即生效环境变量配置。
```language
java -version
```
测试java是否配置成功。
```language
hadoop version
```
测试hadoop是否配置成功。