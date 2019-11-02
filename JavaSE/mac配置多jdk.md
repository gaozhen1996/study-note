# 一、需要下载安装mac版的jdk

# 二、编辑.bash_profile
## 2.1 打开.bash_profile
输入命令：open .bash_profile，打开编辑器
## 2.2配置jdk环境
```language
export JAVA_7_HOME=/Library/Java/JavaVirtualMachines/jdk1.7.0_79.jdk/Contents/Home
export JAVA_8_HOME=/Library/Java/JavaVirtualMachines/jdk1.8.0_131.jdk/Contents/Home
```

## 2.3 创建默认的jdk版本
```language
export JAVA_HOME=$JAVA_8_HOME
```

## 2.4 创建alias命令，实现动态切换
```language
alias jdk8='export JAVA_HOME=$JAVA_8_HOME'
alias jdk7='export JAVA_HOME=$JAVA_7_HOME'
```

## 2.5保存并退出编辑器
## 2.6 bash_profile 使配置生效
在terminal中输入命令：
```language
source .bash_profile 
```
使配置生效

# 三、验证配置与切换
在terminal中 输入  jdk7，再输入java -version 查看当前版本即可实现动态切换,jdk8同样。