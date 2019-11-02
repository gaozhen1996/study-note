数据库服务器上运行：

```
mysql -u root -p -e "set global net_buffer_length=1000000; set global max_allowed_packet=1000000000;"
```

