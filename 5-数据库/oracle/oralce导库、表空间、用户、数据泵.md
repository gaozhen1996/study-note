# 一、导入dmp流程
> 一般导入dmp文件需要先创建表空间、创建用户并指定表空间、导库。
# 1.使用sys登录

```
sqlplus / as sysdba
```

# 2.创建表空间
## 2.1创建临时表空间 

```
create temporary tablespace FD20170731B_temp 
tempfile 'F:\app\hspcadmin\oradata\orcl\FD20170731B_temp.dbf' 
size 64m 
autoextend on 
next 50m maxsize unlimited 
extent management local;
```
- test_temp：l临时表空间名称
- 'F:\app\hspcadmin\oradata\orcl\FD20180816B_temp.dbf'：数据文件位置

## 2.2创建数据表空间 

```
create tablespace FD20170731B_data 
logging  
datafile 'F:\app\hspcadmin\oradata\orcl\FD20170731B_data.dbf' 
size 256M 
autoextend on 
next 50M maxsize unlimited 
extent management local; 
```
- test_data：表空间名称
- 'F:\app\hspcadmin\oradata\orcl\FD20170731A_data.dbf'：数据文件位置

# 3.创建用户并指定表空间

```
create user FD20170731B identified by FD20170731B 
default tablespace FD20170731B_data 
temporary tablespace FD20170731B_temp;
```
- FD20180816B：用户名
- FD20180816B：密码
- FD20180816B_data ：默认用户表空间
- FD20180816B_temp：默认临时表空间

# 4.授权

```
GRANT CREATE USER,DROP USER,ALTER USER ,CREATE ANY VIEW ,
DROP ANY VIEW,EXP_FULL_DATABASE,IMP_FULL_DATABASE,
DBA,CONNECT,RESOURCE,CREATE SESSION TO FD20170731B;
```

# 5.导库
## 导出
```
exp FD20170731B/FD20170731B@192.168.225.162:1521/orcl   file=F:\FD20170731B.dmp  full=y
```
- exp 数据库用户名/密码@服务器ip:1521/实例   file=c:\devdb.dmp  full=y

## 导入
```
imp userid=FD20170731A/FD20170731A full=y file=F:\FD20170731A14.dmp
```
- userid=用户名/口令 
- full=y 导入一个完整的库
- file= dmp的文件路径
- log= 日志文件的路径

# 二、删除表空间和用户
> 当表空间没有用时，需要将表空间或者用户删除，删除用户有两种

## 1.删除用户，不删除数据
```
drop user  FD20170731A;
```
- username为用户名

## 2.删除用户，且删除数据
```
drop user FD20170731B cascade;
```
- username为用户名

## 3.删除删除非空表空间，包含物理文件。

```
DROP TABLESPACE FD20170731B_data INCLUDING CONTENTS AND DATAFILES;
```
- test_datas为表空间的名称


# 三、数据泵
## 1.导出数据
**创建转储文件和日志的目录--sqlplus 下执行**

```
create directory FD0816B as 'F:\hs_valuate\20180816B';
```

*注意：要保证在操作系统存在指定的目录*

**操作系统下执行：**

```
expdp system/root directory=FD0816B schemas=FD20180816B dumpfile=FD20180816B10.dmp  logfile=FD20180816B10.log
```
- system :用户名
- 123456：密码
- directory：导出目录
- schemas：用户级别导出
- dumpfile：导出文件
- logfile：日志文件

## 2.导入数据
**操作系统下执行：**

```
impdp system/root directory=dump_dir dumpfile=FD20180816D.dmp  logfile=dbin.log schemas=FD20180816D
```
- system :用户名
- 123456：密码
- directory：导出目录
- logfile：日志文件
- schemas：用户级别

