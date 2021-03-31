> 前言：
>
> sqlload是一种常用的将外部数据导入到oralce中方法

# 一、导入txt文本文件

## 1.准备源数据文件

ttt,SCOTT,192.168.1.111,yes,1,
qq,JACK,192.168.1.20,no,1,
YY,TOM,192.168.1.20,no,1,
WEB1,HAHA,192.168.1.1,no,1,
XXX,ROBIN,111.111.111.111,no,1,20210331,
DB2,LUCY,192.168.10.10,no,1,
ORACLE,LILY,222.222.222.222,no,1,
WORKGROUP,DENNIS,133.133.133.133,no,0,20210331,
DCR,CANDY,192.168.100.10,no,1,
T3,FLY,192.168.10.33,no,1,
T1,LINDA,192.168.10.200,no,1,20210331,
T2,LILEI,192.168.100.31,no,1,20210331,

## 2.创建目标表

```sql
create table test_txt
(
 host          VARCHAR2(30),
 user_name VARCHAR2(30),
 ip_address      VARCHAR2(15),
 pass            VARCHAR2(4) default 'no' not null,
 judge           NUMBER default 0 not null,
 endtime         DATE
 );
```

## 3.写控制文件

控制文件有很多的参数可以设置，此处以基本的为例。测试将控制文件保存为txt文件即可

- 有部分数据最后字段为空，所以控制文件中需要加trailing nullcols 

```sql
LOAD DATA
INFILE *
APPEND INTO TABLE test_txt
fields terminated by ','
trailing nullcols 
(HOST,
USER_NAME,
IP_AddrESS,
PASS,
JUDge,
endTIME "to_date(:endTIME ,'yyyy-mm-dd')")
```

## 4.导入数据

sqlldr FD20180816C/FD20180816C control=C:\Users\hspcadmin\Desktop\ctrl_txt.txt  log=C:\Users\hspcadmin\Desktop\log.txt data=C:\Users\hspcadmin\Desktop\data.txt



# 二、导入excel文件

## 1.准备源数据文件

todo

## 2.创建目标表

```sql
create table test_excel
(
  effectivedate        DATE,
  indexcode     VARCHAR2(20),
  indexname     VARCHAR2(50),
  constituentcode VARCHAR2(100),
  constituentname VARCHAR2(100)
)
```

## 3.写控制文件

```sql
LOAD DATA
INFILE *
append INTO TABLE test_excel
fields terminated by ','
TRAILING NULLCOLS
(
  effectivedate  "to_date(:effectivedate,'yyyy-mm-dd')",
  indexcode,
  indexname,
  constituentcode,
  constituentname
)
```

## 4.导入数据

sqlldr FD20180816C/FD20180816C control=C:\Users\hspcadmin\Desktop\ctrl_excel.txt  log=C:\Users\hspcadmin\Desktop\sqlload.log data=C:\Users\hspcadmin\Desktop\sqlload源数据文件.xls

