delete from 表名;

truncate table 表名;

不带where参数的delete语句可以删除mysql表中所有内容，使用truncate table也可以清空mysql表中所有内容。

效率上truncate比delete快，但truncate删除后不记录mysql日志，不可以恢复数据。

delete的效果有点像将mysql表中所有记录一条一条删除到删完，

而truncate相当于保留mysql表的结构，重新创建了这个表，所有的状态都相当于新表。

