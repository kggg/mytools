# mytools
mytools 一个工具集, 方便维护  
例： ./mytools -module ssh -target webserver -words "/etc/nginx/nginx.conf /etc/nginx/" -action sendfile  

## -module  使用的模块  
   files  
   ssh  
   web  

## -target  目标  
   files 文件路径  
   ssh  目的主机  
   web  监听的主机和端口, 如: 10.68.2.10:8080  

## -words  内容  
   files  需要操作文件的内容  
   ssh   远程执行的命令, 或者复制  
         远程文件复制 -words "src dst" src本地文件， dst远程文件  

## -action  操作行为  
   files  append,pop,unshift,shift  
   ssh    sendfile发送文件到远程主机, getfile从远程主机复制文件到本地  