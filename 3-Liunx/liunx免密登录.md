# 1.本地生成秘钥文件
本地系统执行 
```language
ssh-keygen -t rsa
```
生成密钥文件

# 2.查看秘钥文件
2.在相应的目录下查看生成的密钥文件，其中：id_rsa为私钥文件，
id_rsa.pub为公钥文件

# 3.上传秘钥文件
本地机器执行命令如：
```language
ssh-copy-id -i ~/.ssh/id_rsa.pub root@47.94.131.201
```
注意IP需要更新为自己服务器的IP
