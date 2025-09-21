## 工具分享

utools

插件：

* 快捷输入
* 程序员手册
* 编码小助手

离线插件下载网址：https://api.u-tools.cn/Plugins/Developer/allPlugins



## 参考链接

https://tenqaz.github.io/pages/39f36e/
https://tenqaz.github.io/pages/f3cf17/



### 问题

1. 容器进程和docker run进程的父进程是？

容器的父进程是containerd-shim，而contained-shim是contained拉起来的，而docker run是我们执行的命令进程，所以其父进程是shell.
https://tenqaz.github.io/pages/39f36e/#docker%E6%9E%B6%E6%9E%84

2. proc一定要挂载吗？
proc目录是在内存的文件系统，需要挂载在目录中，才能访问。


3. k8s如何使用cgroup限制？
后面会讲

4. docker export

容器的导出：docker export
容器的导入变成镜像：docker import
镜像的导出：docker save
镜像的导入: docker load
