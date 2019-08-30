## docker proxy

### 1、启动分部代理端程序
* -l 分部docker代理端监听地址
* -m 接收端监听地址，根据需求可以改成http、tcp、mq等协议传输

```shell
[root@vm-master-240 ~]# ./docker-proxy slave docker -l 192.168.1.240:6000 -m 192.168.1.241:7000
```

### 2、配置各个节点docker为代理模式(Environment)
```shell
[root@vm-node-242 ~]# vi /usr/lib/systemd/system/docker.service
[Service]
Environment="HTTP_PROXY=http://192.168.1.240:6000/"
[root@vm-node-242 ~]# systemctl daemon-reload
[root@vm-node-242 ~]# systemctl restart docker
```

### 3、启动总部代理端程序
* -l 总部docker代理端监听地址, 如果换成mq等协议传输则不需要监听地址，接受数据后直接转发即可

```shell
[root@vm-master-241 ~]# ./docker-proxy master docker -l 192.168.1.241:7000
```
