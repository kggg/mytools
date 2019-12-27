# mytools
mytools 一个工具集, 方便维护  <br>
例： __./mytools -module ssh -target webserver -words "/etc/nginx/nginx.conf /etc/nginx/" -action sendfile __<br>

## -module  使用的模块  
   files  对文件的操作 <br>
   ssh  对远程主机的操作 <br>
   web  运行web服务 <br>

## -target  目标  
   ssh  目的主机  <br>
   web  监听的主机和端口, 如: 10.68.2.10:8080  <br>

## -words  内容  
   files  需要操作文件的内容  <br>
   ssh   远程执行的命令, 或者复制  <br>
         远程文件复制 -words "src dst" src本地文件， dst远程文件 <br> 

## -action  操作行为  
###   files  
    append 追加到行尾  <br>
    pop 删除行尾的内容  <br>
    unshift 在行首添加内容,  <br>
    shift  删除行首内容<br>

###   ssh    
      -action 为空时， 执行words中的SSH命令 <br> 
      action为 sendfile发送文件到远程主机, <br>
      getfile从远程主机复制文件到本地  <br>