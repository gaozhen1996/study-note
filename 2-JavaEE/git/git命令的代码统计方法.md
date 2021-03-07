### 1.统计所有人代码量

统计所有人代码增删量，拷贝如下命令，直接在git bash等终端，git项目某分支下执行

```
git log --format='%aN' | sort -u | while read name; do echo -en "$name\t"; git log --author="$name" --pretty=tformat: --numstat | awk '{ add += $1; subs += $2; loc += $1 - $2 } END { printf "added lines: %s, removed lines: %s, total lines: %s\n", add, subs, loc }' -; done
```
### 2.统计指定提交者代码量
统计单个提交者代码量，将下面的--author="username" 中的 username 替换成具体的提交者，然后执行


```
git log --author="username" --pretty=tformat: --numstat | awk '{ add += $1; subs += $2; loc += $1 - $2 } END { printf "added lines: %s, removed lines: %s, total lines: %s\n", add, subs, loc }' -
```
