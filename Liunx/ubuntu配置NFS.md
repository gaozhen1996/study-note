# 一、NFS简介
>　NFS:是Network FileSystem。最大的作用就是通过网络，让不同的机器、不同的作业系统、可以分享档案。
通过将共享目录挂接到本地，就可以像操作本地目录一样去操作共享的目录。

# 二、服务端安装
## 2.1安装命令
我的服务端环境是Ubuntu，因此nfs的服务端是安装在服务端上。
```
sudo apt-get install nfs-kernel-server
```
安装nfs-kernel-server时，apt会自动安装nfs-common和portmap，新的版本portmap已经被rpcbind代替了

## 2.2配置
配置共享的路径，只需要修改exports文件就可以了，而exports一般位于/etc/exports下

```
sudo vim /etc/exports
```
在文件的末尾添加

```
/mnt *(rw,sync,no_root_squash,insecure)
```
* mnt代表的是共享的目录
* 代表是允许所有的网络访问
* ro 该主机对该共享目录有只读权限
* rw 该主机对该共享目录有读写权限
* root_squash 客户机用root用户访问该共享文件夹时，将root用户映射成匿名用户
* no_root_squash 客户机用root访问该共享文件夹时，不映射root用户
* all_squash 客户机上的任何用户访问该共享目录时都映射成匿名用户anonuid 将客户机上的用户映射成指定的本地用户ID的用户
* anongid 将客户机上的用户映射成属于指定的本地用户组ID
* sync 资料同步写入到内存与硬盘中
* async 资料会先暂存于内存中，而非直接写入硬盘insecure 允许从这台机器过来的非授权访问
* 允许客户端从大于1024的tcp/ip端口连接服务器

## 2.3重启NFS服务

```
/etc/init.d/rpcbind restart
/etc/init.d/nfs-kernel-server restart
```

## 2.4客户端验证

```
showmount -e 47.94.131.201
```
输入上面命令，出现下面结果，代表nfs服务端配置完成

```
gaozhendeMacBook-Pro:~ gaozhen$ showmount -e 47.94.131.201
Exports list on 47.94.131.201:
/mnt                                *
gaozhendeMacBook-Pro:~ gaozhen$
```