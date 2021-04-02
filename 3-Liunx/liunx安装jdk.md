# liunx 安装JDK
liunx安装jdk分为两步，第一是解压缩jdk，第二是配置环境变量
# 1.解压缩jdk
jdk在官网下载即可
# 配置环境变量
配置环境变量 环境变量分为用户变量和系统变量。  

- 用户变量配置文件：
位置 ：
~/.bashrc（在当前用户主目录下的隐藏文件，可以通过ls ­a查看到）  


- 系统环境配置文件：
位置 ：
/etc/profile 
```language
#set jdk environment #注意 JAVA_HOME变量为Java安装目录 export JAVA_HOME=/usr/lib/jvm/jdk1.8.0_131     export JRE_HOME=${JAVA_HOME}/jre export CLASSPATH=.:$JAVA_HOME/lib/dt.jar:$JRE_HOME/lib/tools.jar export PATH=${JAVA_HOME}/bin:$PATH
```
**最后一步：使环境变量生效 source /etc/profile，jdk配置完成**
