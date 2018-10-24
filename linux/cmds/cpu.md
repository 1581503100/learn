##### 查看物理cpu个数 
````
cat /proc/cpuinfo |grep "physical id" |sort |uniq|wc -l
````

##### 查看每个物理cpu核数
````
cat /proc/cpuinfo |grep "cpu cores" |uniq
````

##### 查看内存信息
````
cat /proc/meminfo
````

##### 查看磁盘信息
````
fdisk -l
````


