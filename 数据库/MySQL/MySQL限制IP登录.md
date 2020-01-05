# 前言

一般开发库可能会允许所有的IP访问，但是生产环境这样是极其不安全的。

# 1.查看当前数据库允许登录的IP

```sql
SELECT DISTINCT CONCAT('User: ''',user,'''@''',host,''';') AS query FROM mysql.user;
```

如果结果集出现了User: 'root'@'%'; 表示所有IP都可以用root账号访问，需要将其删除了。

# 2.删除允许所有的IP登录的权限

```sql
drop user 用户名@'%';
```

# 3.赋予指定的IP访问

```sql
GRANT ALL PRIVILEGES ON *.* TO 'root'@'访问IP' IDENTIFIED BY '数据库密码';
flush privileges;
```

