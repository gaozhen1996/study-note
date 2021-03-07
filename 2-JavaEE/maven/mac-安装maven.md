# 1、下载Maven
打开Maven官网下载页面：http://maven.apache.org/download.cgi
下载:apache-maven-3.5.0-bin.tar.gz
解压下载的安装包到某一目录，比如：/Users/xxx/Documents/maven 

# 2、配置环境变量 
打开terminel输入以下命令： 
``` shell
open ~/.bash_profile 
```
 打开.bash_profile文件，在次文件中添加设置环境变量的命令 

```language
export M2_HOME=~/userApp/apache-maven-3.5.3  export PATH=$PATH:$M2_HOME/bin 
```


添加之后保存并推出，执行以下命令使配置生效：
```language
source ~/.bash_profile 
```

# 3、查看配置是否生效 
输入命令：
```language
mvn -v 
```
 