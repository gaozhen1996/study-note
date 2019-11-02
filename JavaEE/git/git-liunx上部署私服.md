在liunx上配置git服务器,配置的时候出现了权限问题,希望能够解决和我遇到了相同问题的朋友们. 因为git仓库上的项目是以root权限创建的,pull的时候却是以git用户pull的,导致出现了权限问题


## 一、安装git服务器所需软件

打开终端输入以下命令：

```
 sudo apt-get install git-core openssh-server openssh-client
```

git-core是git版本控制核心软件

安装openssh-server和openssh-client是由于git需要通过ssh协议来在服务器与客户端之间传输文件

然后中间有个确认操作，输入Y后等待系统自动从镜像服务器中下载软件安装，安装完后会回到用户当前目录。如果

安装提示失败，可能是因为系统软件库的索引文件太旧了，先更新一下就可以了，更新命令如下：

```
 sudo apt-get update 
```

更新完软件库索引后继续执行上面的安装命令即可

---
## 二、新加git用户,管理git仓库

```
sudo useradd -m git
```

```
sudo passwd git
```

## 三、新建项目,初始化

到/home/git目录下新建项目Test
```
cd ~
mkdir Test
git init --bare
```

## 四、pull项目

```
ssh://git@192.168.31.208:22//home/git/Test
```
IP需要更换为git服务器的ip