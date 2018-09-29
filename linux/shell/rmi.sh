#删除名称为ai-market的本地镜像
for img in `docker images |grep ai-market`;
do
        if test ${#img} -eq $[12]
        then
                echo $img
                docker rmi $img
        fi
done
