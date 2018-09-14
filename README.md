# dockertest
this is a docker test basic golang

包管理工具 glide

###########################
#初始化包
glide init
#生成依赖
glide install
#生成docker镜像文件
docker build .
#运行docker 切配置文件挂载
#-v /path:/path 将宿主文件挂载在容器对应的目录
#这里我在开始做的时候闹了一个笑话，容器名称dockertest写在了-v 前边，结果导致挂载不上。
#/bin/bash表示接受shell
#如果向进入容器内部
#  先列出容器对应的ID
#   docker ps
#  然后exec进入容器内部
#   docker exec -it 容器id /bin/bash
#   如果执行上面命令报错则把/bin/bash缓存/bin/sh
#   docker exec -it 36476bbaff1d /bin/sh
docker run -p 8000:8000  -v /home/gosrc/src/dockertest/config:/config dockertest   /bin/bash

