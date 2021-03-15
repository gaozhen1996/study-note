
### 查询各个数据库占用空间大小
```
select TABLE_SCHEMA, concat(truncate(sum(data_length)/1024/1024,2),' MB') as data_size,
    concat(truncate(sum(index_length)/1024/1024,2),'MB') as index_size
    from information_schema.tables
    group by TABLE_SCHEMA
   order by data_length desc;
```

### 查询某个数据库内各个表的容量大小

```
select TABLE_NAME, concat(truncate(data_length/1024/1024,2),' MB') as data_size,
    concat(truncate(index_length/1024/1024,2),' MB') as index_size
    rom information_schema.tables where TABLE_SCHEMA = '数据库名字'
    group by TABLE_NAME
    order by data_length desc;
```
