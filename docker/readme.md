
#### 通过正则获取容器
```text
for i in $(docker ps -a | grep "REGEXP_PATTERN" | cut -f1 -d" "); do echo $i; done
```

#### 删除停止容器
```text
docker rm -v `docker ps -a -q -f status=exited`
```