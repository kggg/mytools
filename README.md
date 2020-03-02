# mytools
mytools 一个工具集, 可以对文件进行相应的操作，也可以对远程主机进行远程SSH操作， 可以提供一个go-gin的web服务进行测试  <br>
* 文件操作: ./mytools files -p "/home/user/a.txt" -w "#" -a shift    //删除行首的注释 <br>
* 远程发送文件： ./mytools ssh -h webserver -a sendfile /etc/nginx/nginx.conf /etc/nginx <br>
     ./mytools  ssh -host webserver -action sendfile /etc/nginx/nginx.conf /etc/nginx/  <br>

# 主要功能模块  
*   file  对文件的操作 <br>
*   ssh  对远程主机的操作 <br>
*   web  运行web服务 <br>


###   files  

Usage of file:
  -a, --action string   操作行为, operation
  -p, --path string     文件路径, filepath
  -s, --save            操作文件时， 是否保存操作结果，默认不保存, true|false
  -w, --word string     需要操作的内容word


    append 追加到行尾
    pop 删除行尾的内容
    unshift 在行首添加内容
    shift  删除行首内容
    delete 删除内容
    search 查找文件中的相关内容
    rsearch 以正则匹配模式查找文件中的内容
    replace 替换文件中的内容
    save   对文件的操作进行保存

###   ssh    

Usage of ssh:
  -a, --action string   远程执行的模块
  -h, --host string     远程主机名， 在配置文件config/remote.ini中设置 (default "localhost")

    -action 为空时， 执行words中的SSH命令
    action为 sendfile发送文件到远程主机,
    getfile从远程主机复制文件到本地 