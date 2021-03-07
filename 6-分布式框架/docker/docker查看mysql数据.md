# 1.查看当前docker正在使用的容器

![img](img/docker查看mysql数据1.png)

**2.进入容器**

```
docker exec -it 8af579a08ead  /bin/bash
```

![img](img/docker查看mysql数据2.png)

**3.进入mysql**

```
mysql -u root -p
```

![img](img/docker查看mysql数据3.png)