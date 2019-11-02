### 1.现象
MySql 执行 DELETE FROM Table 时，报 Error Code: 1175.

### 2.原因
因为 MySql 运行在 safe-updates模式下，该模式会导致非主键条件下无法执行update或者delete命令

### 3.解决方法
```
SET SQL_SAFE_UPDATES = 0;
```
