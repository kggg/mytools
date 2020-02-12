# mytools
mytools 一个工具集, 可以对文件进行相应的操作，也可以对远程主机进行远程SSH操作， 可以提供一个go-gin的web服务进行测试  <br>
* 文件操作: ./mytools -m files -t "/home/user/a.txt" -w "#" -a shift    //删除行首的注释 <br>
* 远程发送文件： ./mytools -m ssh -t webserver -w "/etc/nginx/nginx.conf /etc/nginx" -a sendfile<br>
     ./mytools -module ssh -target webserver -words "/etc/nginx/nginx.conf /etc/nginx/" -action sendfile <br>

## -m | --module  使用的模块  
*   files  对文件的操作 <br>
*   ssh  对远程主机的操作 <br>
*   web  运行web服务 <br>

## -t | --target  目标
*  files  文件路径<br>
*   ssh  目的主机  <br>
*   web  监听的主机和端口, 如: 10.68.2.10:8080  <br>

## -w | --words  内容  
*   files  需要操作文件的内容  <br>
*   ssh   远程执行的命令, 或者复制  <br>
         远程文件复制 -words "src dst" src本地文件， dst远程文件 <br> 

## -a | --action  操作行为  
###   files  
    append 追加到行尾
    pop 删除行尾的内容
    unshift 在行首添加内容
    shift  删除行首内容
    delete 删除内容
    search 查找文件中的相关内容
    rsearch 以正则匹配模式查找文件中的内容
    save   对文件的操作进行保存

###   ssh    
    -action 为空时， 执行words中的SSH命令
    action为 sendfile发送文件到远程主机,
    getfile从远程主机复制文件到本地 