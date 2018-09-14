# dockertest
this is a docker test basic golang<br>

包管理工具 glide\<br>
dockerfile是用的vscode的插件生成的<br>
初始化包\<br>
glide init<br>
生成依赖<br>

glide install<br>

生成docker镜像文件<br>

docker build .<br>

运行docker 切配置文件挂载<br>
-v /path:/path 将宿主文件挂载在容器对应的目录<br>
这里我在开始做的时候闹了一个笑话，容器名称dockertest写在了-v 前边，结果导致挂载不上。<br>
/bin/bash表示接受shell<br>
如果向进入容器内部<br>
  先列出容器对应的ID<br>
   
   docker ps<br>
   
  然后exec进入容器内部<br>
   docker exec -it 容器id /bin/bash<br>
   如果执行上面命令报错则把/bin/bash缓存/bin/sh<br>
   
   docker exec -it 36476bbaff1d /bin/sh<br>
   

docker run -p 8000:8000  -v /home/gosrc/src/dockertest/config:/config dockertest   /bin/bash

