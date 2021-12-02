# filestore-server




# V1.0   云存储”系统原型(实现一个超精简版云盘)


## 环境配置

```s

➜  ~ multipass launch -n master -c 2 -m 4G -d 40G
Launched: main
➜  ~ multipass shell master

multipass launch -n node1 -c 2 -m 4G -d 40G

multipass shell node1


ubuntu@main:~$ sudo mv /etc/apt/sources.list /etc/apt/sources.list.bak
ubuntu@main:~$ sudo vim /etc/apt/sources.list
ubuntu@main:~$ cat /etc/apt/sources.list
# ubuntu 20.04(focal) 配置如下
deb http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse

ubuntu@main:~$ sudo apt-get update && sudo apt-get upgrade -y



➜  share multipass mount /Users/x/share master:/home/ubuntu/share
➜  share multipass info main
Name:           main
State:          Running
IPv4:           192.168.105.9
Release:        Ubuntu 20.04.3 LTS
Image hash:     939be728cbc7 (Ubuntu 20.04 LTS)
Load:           0.34 0.13 0.03
Disk usage:     1.6G out of 19.2G
Memory usage:   217.2M out of 3.8G
Mounts:         /Users/x/share => /home/ubuntu/share
                    UID map: 501:default
                    GID map: 20:default


➜  ~ multipass ls
Name                    State             IPv4             Image
main                    Running           192.168.105.9    Ubuntu 20.04 LTS
                                          172.17.0.1
                
```


## Ubuntu 安装 Docker

```go

# 使用官方安装脚本自动安装
ubuntu@master:~$ curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun

ubuntu@master:~$ docker images
Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Get "http://%2Fvar%2Frun%2Fdocker.sock/v1.24/images/json": dial unix /var/run/docker.sock: connect: permission denied

ubuntu@master:~$ sudo groupadd docker
groupadd: group 'docker' already exists
ubuntu@master:~$ sudo gpasswd -a ubuntu docker
Adding user ubuntu to group docker
ubuntu@master:~$ sudo service docker restart

ubuntu@master:~$ docker info

Cgroup Driver: cgroupfs

ubuntu@master:~$ sudo vim /etc/docker/daemon.json

{ "registry-mirrors": ["https://hkaofvr0.mirror.aliyuncs.com"],
        "exec-opts": ["native.cgroupdriver=systemd"]
}

systemctl restart docker
systemctl status docker

ubuntu@master:~$ sudo systemctl daemon-reload
ubuntu@master:~$ sudo systemctl restart docker





# 重启 iTerm2
ubuntu@node1:~$ exit
logout
➜  ~ multipass shell node1

ubuntu@master:~$ docker info

 Registry Mirrors:
  https://hkaofvr0.mirror.aliyuncs.com/



# Install Compose on Linux systems

# 方法一：
sudo apt install docker-compose -y

# 方法二：
ubuntu@master:~$ sudo curl -L "https://github.com/docker/compose/releases/download/v2.0.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

ubuntu@master:~$ sudo chmod +x /usr/local/bin/docker-compose
ubuntu@master:~$ sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
ubuntu@master:~$ docker-compose --version

Docker Compose version v2.0.1




```

```go
ubuntu@main:~$ sudo snap install go --classic
go 1.17.3 from Michael Hudson-Doyle (mwhudson) installed

```

## golang的goproxy配置

```go
1.首先开启go module

go env -w GO111MODULE=on           // macOS 或 Linux
2.配置goproxy:


七牛云配置：
export GOPROXY=https://goproxy.cn         // macOS 或 Linux
注意：

Go 1.13设置了默认的GOSUMDB=sum.golang.org，是用来验证包的有效性。这个网址由于墙的原因可能无法访问，所以可以使用下面命令来关闭：

export GOSUMDB=off // macOS 或 Linux

```


## 测试

```go


http://192.168.105.9:8080/file/upload



ubuntu@main:/tmp$ sha1sum /tmp/noface.png
e87999a1ac4defe6f25153d2dd41091fdd89c884  /tmp/noface.png




http://192.168.105.9:8080/file/meta/?filehash=e87999a1ac4defe6f25153d2dd41091fdd89c884

{
    "FileSha1": "e87999a1ac4defe6f25153d2dd41091fdd89c884",
    "FileName": "noface.png",
    "FileSize": 582157,
    "Location": "/tmp/noface.png",
    "UploadAt": "2021-11-18 15:52:41"
}




ubuntu@main:/tmp$ sha1sum /tmp/01.png
e87999a1ac4defe6f25153d2dd41091fdd89c884  /tmp/01.png

http://192.168.105.9:8080/file/meta/?filehash=e87999a1ac4defe6f25153d2dd41091fdd89c884

http://192.168.105.9:8080/file/download?filehash=e87999a1ac4defe6f25153d2dd41091fdd89c884


POST 
http://192.168.105.9:8080/file/update?op=0&filehash=e87999a1ac4defe6f25153d2dd41091fdd89c884&filename=111.png

{
    "FileSha1": "e87999a1ac4defe6f25153d2dd41091fdd89c884",
    "FileName": "111.png",
    "FileSize": 582157,
    "Location": "/tmp/01.png",
    "UploadAt": "2021-11-18 17:57:31"
}

POST
http://192.168.105.9:8080/file/delete?filehash=e87999a1ac4defe6f25153d2dd41091fdd89c884

```


# V2.0 “云存储”系统之基于MySQL实现的文件数据库(持久化云文件信息)



```
docker volume create portainer_data

docker run -d -p 8000:8000 -p 9443:9443 --name portainer \
    --restart=always \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v portainer_data:/data \
    portainer/portainer-ce:latest


https://192.168.105.9:9443



mac:

localhost:9443
```


# 查看服务器内存
```go
sudo apt install htop
```

## mysql 环境配置

```sql
ubuntu@main:~$ sudo docker pull mysql

docker run -itd --name master -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql

docker run -itd --name slave -p 3307:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql


配置Master

使用如下命令进入到Master容器内部，使用容器ID或者名称均可：

ubuntu@main:~$ docker exec -it master /bin/bash
root@87dc3b1876ac:/# cd /etc/mysql
root@87dc3b1876ac:/etc/mysql# apt-get update && apt-get install vim -y

root@87dc3b1876ac:/etc/mysql# ls
conf.d	my.cnf	my.cnf.fallback
root@87dc3b1876ac:/etc/mysql# vim my.cnf


[mysqld]
pid-file        = /var/run/mysqld/mysqld.pid
socket          = /var/run/mysqld/mysqld.sock
datadir         = /var/lib/mysql
secure-file-priv= NULL
# 添加
server-id=100
log-bin=master-bin
binlog-format=ROW



## 解释
[mysqld]
## 同一局域网内注意要唯一
server-id=100  
## 开启二进制日志功能，可以随便取（关键）
log-bin=master-bin
binlog-format=ROW     // 二级制日志格式，有三种 row，statement，mixed
binlog-do-db=数据库名  //同步的数据库名称,如果不配置，表示同步所有的库



root@87dc3b1876ac:/etc/mysql# exit
exit
ubuntu@main:~$ sudo docker restart master


# 配置Slave






ubuntu@main:~$ sudo docker exec -it slave /bin/bash

root@e762b72982bc:/# apt-get update && apt-get install vim -y

root@e762b72982bc:/# vim /etc/mysql/my.cnf

[mysqld]
pid-file        = /var/run/mysqld/mysqld.pid
socket          = /var/run/mysqld/mysqld.sock
datadir         = /var/lib/mysql
secure-file-priv= NULL
# 新增配置
server-id=101
log-bin=mysql-slave-bin
relay_log=mysql-relay-bin
read_only=1


## 解释
[mysqld]
## 设置server_id,注意要唯一
server-id=101  
## 开启二进制日志功能，以备Slave作为其它Slave的Master时使用
log-bin=mysql-slave-bin   
## relay_log配置中继日志
relay_log=mysql-relay-bin  
read_only=1  ## 设置为只读,该项如果不设置，表示slave可读可写





root@e762b72982bc:/# exit
exit
ubuntu@main:~$ sudo docker restart slave



# 开启Master-Slave主从复制






ubuntu@main:~$ sudo apt install mysql-client-core-8.0


ubuntu@main:~$ mysql -uroot -h127.0.0.1 -P3306 -p123456


ubuntu@main:~$ mysql -uroot -h127.0.0.1 -P3307 -p123456






## 3306 主

mysql> show master status;
+-------------------+----------+--------------+------------------+-------------------+
| File              | Position | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set |
+-------------------+----------+--------------+------------------+-------------------+
| master-bin.000001 |      156 |              |                  |                   |
+-------------------+----------+--------------+------------------+-------------------+
1 row in set (0.05 sec)





ubuntu@main:~$ sudo docker inspect --format='{{.NetworkSettings.IPAddress}}' master
172.17.0.3
ubuntu@main:~$ sudo docker inspect --format='{{.NetworkSettings.IPAddress}}' slave
172.17.0.4






## 3307 从

CHANGE MASTER TO
  MASTER_HOST='172.17.0.3',
  MASTER_USER='root',
  MASTER_PASSWORD='123456',
  MASTER_PORT=3306,
  MASTER_LOG_FILE='master-bin.000001',
  MASTER_LOG_POS=156;








mysql> start slave;

mysql> show slave status\G;
*************************** 1. row ***************************
               Slave_IO_State: Waiting for source to send event
                  Master_Host: 172.17.0.3
                  Master_User: root
                  Master_Port: 3306
                Connect_Retry: 60
              Master_Log_File: master-bin.000001
          Read_Master_Log_Pos: 156
               Relay_Log_File: mysql-relay-bin.000002
                Relay_Log_Pos: 325
        Relay_Master_Log_File: master-bin.000001
             Slave_IO_Running: Yes
            Slave_SQL_Running: Yes






## 测试 主

mysql> create database test1 default character set utf8;
Query OK, 1 row affected, 1 warning (0.11 sec)

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
| test1              |
+--------------------+
5 rows in set (0.10 sec)

## 测试 从

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
| test1              |
+--------------------+
5 rows in set (0.18 sec)



```



## 1. 创建数据库和表结构

```
-- 创建数据库
create database fileserver default character set utf8;

-- 切换数据库
use fileserver;

-- 创建文件表
CREATE TABLE `tbl_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `file_sha1` char(40) NOT NULL DEFAULT '' COMMENT '文件hash',
  `file_name` varchar(256) NOT NULL DEFAULT '' COMMENT '文件名',
  `file_size` bigint(20) DEFAULT '0' COMMENT '文件大小',
  `file_addr` varchar(1024) NOT NULL DEFAULT '' COMMENT '文件存储位置',
  `create_at` datetime default NOW() COMMENT '创建日期',
  `update_at` datetime default NOW() on update current_timestamp() COMMENT '更新日期',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '状态(可用/禁用/已删除等状态)',
  `ext1` int(11) DEFAULT '0' COMMENT '备用字段1',
  `ext2` text COMMENT '备用字段2',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_file_hash` (`file_sha1`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 查看数据表
show create table tbl_file\G;


mysql> show tables;
+----------------------+
| Tables_in_fileserver |
+----------------------+
| tbl_file             |
+----------------------+
1 row in set (0.07 sec)


```









```git 
git switch -c v2.0

git push origin HEAD:v2.0
```



# V3.0 “云存储”系统之基于用户系统实现的资源隔离及鉴权 (账号和应用收入息息相关)


mysql用户表设计
```sql
-- 创建用户表
CREATE TABLE `tbl_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `user_pwd` varchar(256) NOT NULL DEFAULT '' COMMENT '用户encoded密码',
  `email` varchar(64) DEFAULT '' COMMENT '邮箱',
  `phone` varchar(128) DEFAULT '' COMMENT '手机号',
  `email_validated` tinyint(1) DEFAULT 0 COMMENT '邮箱是否已验证',
  `phone_validated` tinyint(1) DEFAULT 0 COMMENT '手机号是否已验证',
  `signup_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '注册日期',
  `last_active` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后活跃时间戳',
  `profile` text COMMENT '用户属性',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '账户状态(启用/禁用/锁定/标记删除等)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`user_name`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

```


## 404 page not found


main.go 添加

```go

	// 静态资源处理
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	pwd, _ := os.Getwd()
	fmt.Println(pwd + " " + os.Args[0])
	http.Handle("/static/", http.FileServer(http.Dir(filepath.Join(pwd, "./"))))

  // 用户相关接口
	http.HandleFunc("/", handler.SignInHandler)
```


修改 signin.html

```js
      success: function (body, textStatus, jqXHR) {
        var resp = JSON.parse(body);
        localStorage.setItem("token", resp.data.Token)
        localStorage.setItem("username", resp.data.Username)
        window.location.href = resp.data.Location;
      }
```



```git 
git switch -c v3.0

git push origin HEAD:v3.0
```




## V4.0 基于Hash计算实现秒传 （大幅提升拥有海量文件的云盘性能）


```go
handler/handler.go 实现：

FileQueryHandler : 查询批量的文件元信息
TryFastUploadHandler : 尝试秒传接口
DownloadHandler : 文件下载接口


db/userfile.go 实现：

QueryUserFileMetas : 批量获取用户文件信息

meta/filemeta.go 实现：

GetLastFileMetas : 获取批量的文件元信息列表
GetLastFileMetasDB : 批量从mysql获取文件元信息



```




```go

# POST 请求

http://192.168.105.9:8080/user/signin?username=admin&password=admin

{"code":0,"msg":"OK","data":{"Location":"http://192.168.105.9:8080/static/view/home.html","Username":"admin","Token":"d5fbd378f2a64aa19ecc51204ec74831619c3fb5"}}

# POST 请求

http://192.168.105.9:8080/file/fastupload?username=admin&token=d5fbd378f2a64aa19ecc51204ec74831619c3fb5&filehash=e87999a1ac4defe6f25153d2dd41091fdd89c884&filename=mojo.png&filesize=582157


{"code":0,"msg":"秒传成功","data":null}



mysql> SHOW CREATE TABLE tbl_user_file;

| tbl_user_file | CREATE TABLE `tbl_user_file` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_name` varchar(64) NOT NULL,
  `file_sha1` varchar(64) NOT NULL DEFAULT '' COMMENT '??hash',
  `file_size` bigint DEFAULT '0' COMMENT '????',
  `file_name` varchar(256) NOT NULL DEFAULT '' COMMENT '???',
  `upload_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '????',
  `last_update` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '??????',
  `status` int NOT NULL DEFAULT '0' COMMENT '????(0??1???2??)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_file` (`user_name`,`file_sha1`),
  KEY `idx_status` (`status`),
  KEY `idx_user_id` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci



mysql> alter table tbl_user_file drop index `idx_user_file`;

Query OK, 0 rows affected (0.06 sec)
Records: 0  Duplicates: 0  Warnings: 0



mysql> SELECT * FROM tbl_user_file;
+----+-----------+------------------------------------------+-----------+---------------+---------------------+---------------------+--------+
| id | user_name | file_sha1                                | file_size | file_name     | upload_at           | last_update         | status |
+----+-----------+------------------------------------------+-----------+---------------+---------------------+---------------------+--------+
|  3 | admin     | e87999a1ac4defe6f25153d2dd41091fdd89c884 |    582157 | 02.png        | 2021-11-22 09:05:43 | 2021-11-22 09:05:42 |      0 |
|  4 | admin     | 7ccc908cad1e9204a530c5cac8fc27d3211fc736 |    410735 | cv.pdf        | 2021-11-22 09:12:15 | 2021-11-22 09:12:14 |      0 |
|  6 | admin     | 60b82b6e25de5e86c797f24e8bd57a8b350fbdcd |   2211509 |  LaTeX-cn.pdf | 2021-11-22 14:20:33 | 2021-11-22 14:20:32 |      1 |
|  8 | admin     | e87999a1ac4defe6f25153d2dd41091fdd89c884 |    582157 | mojo.png      | 2021-11-23 01:35:41 | 2021-11-23 01:35:40 |      1 |
+----+-----------+------------------------------------------+-----------+---------------+---------------------+---------------------+--------+
4 rows in set (0.00 sec)



```




```git 
git switch -c v4.0

git push origin HEAD:v4.0
```




## V5.0 基于Redis实现分块上传及断点续传（可大幅提升传输速度）



```go



➜  ~ docker pull redis

➜  ~ docker run -itd --name redis -p 6379:6379 redis

➜  ~ docker exec -it redis /bin/bash

root@03ac20e1c401:/data# redis-cli
127.0.0.1:6379>




# redis

version: '2'

services:
  redis:
    image: redis
    restart: always
    volumes:
      - data:/data
    ports:
      - 6379:6379
      
volumes:
  data: {}


ubuntu@main:~$ docker exec -it redis_redis_1 /bin/bash
root@8d5040d0ee09:/data# redis-cli
127.0.0.1:6379>

1,查询默认密码
127.0.0.1:6379>  config get requirepass
1) "requirepass"
2) ""

2.修改密码
127.0.0.1:6379> config set requirepass 123456
OK

3.查询修改后的密码
127.0.0.1:6379> config get requirepass
1) "requirepass"
2) "123456"

 4.登录redis

 redis-cli -p 6379 -a 123456



 127.0.0.1:6379> auth 123456
OK
127.0.0.1:6379> keys *
(empty array)




```







```git 
git switch -c v5.0

git push origin HEAD:v5.0
```



## V6.0 基于Ceph实现私有云存储服务 (自主管理集群，技术可控)


```go
# 1. 创建Ceph专用网络
docker network create --driver bridge --subnet 172.20.0.0/16 ceph-network
docker network inspect ceph-network

# 2. 删除旧的ceph相关容器
docker rm -f $(docker ps -a | grep ceph | awk '{print $1}')

# 3. 清理旧的ceph相关目录文件，加入有的话
rm -rf /www/ceph /var/lib/ceph/  /www/osd/

# 4. 创建相关目录及修改权限，用于挂载volume
mkdir -p /www/ceph /var/lib/ceph/osd /www/osd/
chown -R 64045:64045 /var/lib/ceph/osd/
chown -R 64045:64045 /www/osd/



# 5. 创建monitor节点



docker run -itd --name monnode --network ceph-network --ip 172.20.0.10 -e MON_NAME=monnode -e MON_IP=172.20.0.10 -v /www/ceph:/etc/ceph ceph/mon

# 6. 在monitor节点上标识3个osd节点
docker exec monnode ceph osd create
docker exec monnode ceph osd create
docker exec monnode ceph osd create

# 7. 创建OSD节点
docker run -itd --name osdnode0 --network ceph-network -e CLUSTER=ceph -e WEIGHT=1.0 -e MON_NAME=monnode -e MON_IP=172.20.0.10 -v /www/ceph:/etc/ceph -v /www/osd/0:/var/lib/ceph/osd/ceph-0 ceph/osd 
docker run -itd --name osdnode1 --network ceph-network -e CLUSTER=ceph -e WEIGHT=1.0 -e MON_NAME=monnode -e MON_IP=172.20.0.10 -v /www/ceph:/etc/ceph -v /www/osd/1:/var/lib/ceph/osd/ceph-1 ceph/osd
docker run -itd --name osdnode2 --network ceph-network -e CLUSTER=ceph -e WEIGHT=1.0 -e MON_NAME=monnode -e MON_IP=172.20.0.10 -v /www/ceph:/etc/ceph -v /www/osd/2:/var/lib/ceph/osd/ceph-2 ceph/osd

# 8. 增加monitor节点，组件成集群
docker run -itd --name monnode_1 --network ceph-network --ip 172.20.0.11 -e MON_NAME=monnode_1 -e MON_IP=172.20.0.11 -v /www/ceph:/etc/ceph ceph/mon
docker run -itd --name monnode_2 --network ceph-network --ip 172.20.0.12 -e MON_NAME=monnode_2 -e MON_IP=172.20.0.12 -v /www/ceph:/etc/ceph ceph/mon

# 9. 创建gateway节点
docker run -itd --name gwnode --network ceph-network --ip 172.20.0.9 -p 9080:80 -e RGW_NAME=gwnode -v /www/ceph:/etc/ceph ceph/radosgw

# 10. 查看ceph集群状态
sleep 10 && docker exec monnode ceph -s


```



```go
ubuntu@x:~$ docker exec -it gwnode radosgw-admin user create --uid=user1 --display-name=user1
{
    "user_id": "user1",
    "display_name": "user1",
    "email": "",
    "suspended": 0,
    "max_buckets": 1000,
    "auid": 0,
    "subusers": [],
    "keys": [
        {
            "user": "user1",
            "access_key": "35WISHNZFS6ZGCNIOV6Z",
            "secret_key": "DTD9a6DwymN5CSEaTmls3HHVPSyn81EkqX9JRqPm"
        }
    ],
    "swift_keys": [],
    "caps": [],
    "op_mask": "read, write, delete",
    "default_placement": "",
    "placement_tags": [],
    "bucket_quota": {
        "enabled": false,
        "max_size_kb": -1,
        "max_objects": -1
    },
    "user_quota": {
        "enabled": false,
        "max_size_kb": -1,
        "max_objects": -1
    },
    "temp_url_keys": []
}

```












```git 
git switch -c v6.0

git push origin HEAD:v6.0
```


## V7.0 阿里云OSS实现海量数据上云



```git 
git switch -c v7.0

git push origin HEAD:v7.0
```







## V8.0 RabbitMQ实现异步存储 (上传与转移解耦，提高稳定性)

```go
ubuntu@x:~$ sudo mkdir /data/rabbitmq -p

ubuntu@x:~$ docker run -d --hostname rabbit-svr --name rabbit -p 5672:5672 -p 15672:15672 -p 25672:25672 -v /data/rabbitmq:/var/lib/rabbitmq rabbitmq:management

# for RabbitMQ 3.9, the latest series
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.9-management




ubuntu@x:~/share/filestore-server$ go run service/transfer/main.go
2021/11/26 11:44:33 文件转移服务启动中，开始监听异步转移任务队列...
2021/11/26 11:45:34 {"FileHash":"66a173004cb8d07745d04c05eac69a7d2ffca1da","CurLocation":"/tmp/redis-trib.rb","DestLocation":"oss/66a173004cb8d07745d04c05eac69a7d2ffca1da","DestStoreType":3}

```





```git 
git switch -c v8.0

git push origin HEAD:v8.0
```



## V9.0 架构微服务化 (全面优化系统架构及性能)

新增轻量级Gin web框架及改造演示； 
新增gRPC网络通讯框架； 
新增API网关, 实现限流及熔断等功能； 
新增服务注册功能，实现服务自发现和负载均衡；



```go
sudo apt install protobuf-compiler -y

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

go install github.com/golang/protobuf/protoc-gen-go@latest

go get github.com/micro/protoc-gen-micro/v2
go install github.com/micro/protoc-gen-micro@latest



## Ubuntu

ubuntu@x:~$ cd ~/go/bin
ubuntu@x:~/go/bin$ ls
protoc-gen-go  protoc-gen-go-grpc  protoc-gen-micro

ubuntu@x:~/go/bin$ sudo cp protoc-gen-go /usr/local/bin/
ubuntu@x:~/go/bin$ sudo cp protoc-gen-micro /usr/local/bin/


# Mac 
brew install protobuf

# Protoc环境变量
export PROTOBUF=/usr/local/share/emacs/site-lisp/protobuf
export PATH=$PROTOBUF/bin:$PATH
export PATH=$PATH:$GOPATH/bin



ubuntu@x:~/share/filestore-server$ 
protoc --proto_path=service/account/proto --go_out=service/account/proto --micro_out=service/account/proto service/account/proto/user.proto



docker run -d -p 8400:8400 -p 8500:8500 -p 8600:53/udp --restart=always -v /data/consul:/consul/data -e CONSUL_BIND_INTERFACE='eth0' --name=consul consul:latest agent -server -bootstrap -ui -node=1 -client='0.0.0.0'


http://192.168.105.4:8500


提示：/consul/data 是 Consul 持久化地方，如果需要持久化那 Dooker 启动时候需要给它指定一个数据卷 -v /data/consul:/consul/data。

Consul 命令简单介绍
agent : 表示启动 Agent 进程。
-server：表示启动 Consul Server 模式。
-client：表示启动 Consul Cilent 模式。
-bootstrap：表示这个节点是 Server-Leader ，每个数据中心只能运行一台服务器。技术角度上讲 Leader 是通过 Raft 算法选举的，但是集群第一次启动时需要一个引导 Leader，在引导群集后，建议不要使用此标志。
-ui：表示启动 Web UI 管理器，默认开放端口 8500，所以上面使用 Docker 命令把 8500 端口对外开放。
-node：节点的名称，集群中必须是唯一的。
-client：表示 Consul 将绑定客户端接口的地址，0.0.0.0 表示所有地址都可以访问。
-join：表示加入到某一个集群中去。 如：-json=192.168.1.23
Web 管理器
上面命令已经启动了 Consul 和 Web 管理器，我们现在打开 Web 管理器来看一下是否启动成功。通过浏览器浏览 Http://{serverIp}:8500 。


go run service/account/main.go --registry=consul








protoc --proto_path=service/account/proto --go_out=service/account/proto --micro_out=service/account/proto service/account/proto/user.proto


protoc --proto_path=service/dbproxy/proto --go_out=service/dbproxy/proto --micro_out=service/dbproxy/proto service/dbproxy/proto/proxy.proto


protoc --proto_path=service/download/proto --go_out=service/download/proto --micro_out=service/download/proto service/download/proto/download.proto


protoc --proto_path=service/upload/proto --go_out=service/upload/proto --micro_out=service/upload/proto service/upload/proto/upload.proto


```


```git 
git switch -c v9.0

git push origin HEAD:v9.0
```







## V10.0 k8s&Docker容器化实战 (容器化：应用上云部署的基石)


```go
multipass launch -n master -c 2 -m 4G -d 40G
multipass shell master

multipass launch -n node1 -c 2 -m 4G -d 40G
multipass shell node1

```



```go

docker network create fileserver

# 启动微服务
sudo docker-compose up --scale upload=2 --scale download=2 -d


# 启动 traefik 反向代理
sudo docker-compose up -d


vim /etc/hosts

192.168.105.4 upload.fileserver.com

192.168.105.4 download.fileserver.com

192.168.105.4 apigw.fileserver.com





```


## Ubuntu 安装 k8s
```go


# 在master添加hosts
sudo vim /etc/hosts

192.168.105.5 master
192.168.105.6 node1


# 关闭防火墙
systemctl stop firewalld
systemctl disable firewalld

# 关闭防火墙
ubuntu@master:~$ sudo apt-get install ufw

ubuntu@master:~$ sudo ufw disable
Firewall stopped and disabled on system startup
ubuntu@master:~$ sudo ufw status
Status: inactive


# 关闭selinux

setenforce 0

cat /etc/selinux/config

# 关闭selinux

ubuntu@master:~$ sudo vim /etc/selinux/config

SELINUX=disabled


# 关闭swap
ubuntu@master:~$ sudo sed -ri 's/.*swap.*/#&/' /etc/fstab


# 开启路由转发

vim /etc/sysctl.d/k8s.conf


net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1


# 生效
sysctl -p /etc/sysctl.d/k8s.conf

```






## 允许 iptables 检查桥接流量

确保 br_netfilter 模块被加载。这一操作可以通过运行 lsmod | grep br_netfilter 来完成。若要显式加载该模块，可执行 sudo modprobe br_netfilter。

为了让你的 Linux 节点上的 iptables 能够正确地查看桥接流量，你需要确保在你的 sysctl 配置中将 net.bridge.bridge-nf-call-iptables 设置为 1。例如：
```go


cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF

cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
sudo sysctl --system
```


1. 更新 apt 包索引并安装使用 Kubernetes apt 仓库所需要的包：
```go
sudo apt-get update
sudo apt-get install -y apt-transport-https ca-certificates curl
```

2. 下载 Google Cloud 公开签名秘钥：
```go
sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg
```

3. 添加 Kubernetes apt 仓库：
```go
echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
```


4. 更新 apt 包索引，安装 kubelet、kubeadm 和 kubectl，并锁定其版本：
```go

sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl

```






```go


ubuntu@master:~$ docker info

Cgroup Driver: cgroupfs

ubuntu@master:~$ sudo vim /etc/docker/daemon.json

{ "registry-mirrors": ["https://hkaofvr0.mirror.aliyuncs.com"],
        "exec-opts": ["native.cgroupdriver=systemd"]
}

systemctl restart docker && systemctl status docker



vim /etc/systemd/system/kubelet.service.d/10-kubeadm.conf

Environment="KUBELET_EXTAR_ARGS=--fail-swap-on=false"



systemctl daemon-reload

重新启动kubelet.service
		1.systemctl daemon-reload
		2.systemctl restart kubelet.service



ubuntu@master:~$ kubeadm config images list
k8s.gcr.io/kube-apiserver:v1.22.4
k8s.gcr.io/kube-controller-manager:v1.22.4
k8s.gcr.io/kube-scheduler:v1.22.4
k8s.gcr.io/kube-proxy:v1.22.4
k8s.gcr.io/pause:3.5
k8s.gcr.io/etcd:3.5.0-0
k8s.gcr.io/coredns/coredns:v1.8.4


node1:



kubeadm config images pull --image-repository=registry.aliyuncs.com/google_containers

ubuntu@master:~$ kubeadm config images pull --image-repository=registry.aliyuncs.com/google_containers
[config/images] Pulled registry.aliyuncs.com/google_containers/kube-apiserver:v1.22.4
[config/images] Pulled registry.aliyuncs.com/google_containers/kube-controller-manager:v1.22.4
[config/images] Pulled registry.aliyuncs.com/google_containers/kube-scheduler:v1.22.4
[config/images] Pulled registry.aliyuncs.com/google_containers/kube-proxy:v1.22.4
[config/images] Pulled registry.aliyuncs.com/google_containers/pause:3.5
[config/images] Pulled registry.aliyuncs.com/google_containers/etcd:3.5.0-0
[config/images] Pulled registry.aliyuncs.com/google_containers/coredns:v1.8.4



ubuntu@master:~$ docker images
REPOSITORY                                                        TAG       IMAGE ID       CREATED        SIZE
registry.aliyuncs.com/google_containers/kube-apiserver            v1.22.4   8a5cc299272d   13 days ago    128MB
registry.aliyuncs.com/google_containers/kube-controller-manager   v1.22.4   0ce02f92d3e4   13 days ago    122MB
registry.aliyuncs.com/google_containers/kube-scheduler            v1.22.4   721ba97f54a6   13 days ago    52.7MB
registry.aliyuncs.com/google_containers/kube-proxy                v1.22.4   edeff87e4802   13 days ago    104MB
registry.aliyuncs.com/google_containers/etcd                      3.5.0-0   004811815584   5 months ago   295MB
registry.aliyuncs.com/google_containers/coredns                   v1.8.4    8d147537fb7d   6 months ago   47.6MB
registry.aliyuncs.com/google_containers/pause                     3.5       ed210e3e4a5b   8 months ago   683kB






images=(  # 下面的镜像应该去除"k8s.gcr.io/"的前缀，版本换成上面获取到的版本
    kube-apiserver:v1.22.4
    kube-controller-manager:v1.22.4
    kube-scheduler:v1.22.4
    kube-proxy:v1.22.4
    pause:3.5
    etcd:3.5.0-0
    coredns:v1.8.4
)


for imageName in ${images[@]} ; do
    docker tag registry.aliyuncs.com/google_containers/$imageName k8s.gcr.io/$imageName
    docker rmi registry.aliyuncs.com/google_containers/$imageName
done

docker tag k8s.gcr.io/coredns:v1.8.4 k8s.gcr.io/coredns/coredns:v1.8.4
docker rmi k8s.gcr.io/coredns:v1.8.4






ubuntu@master:~$ docker images
REPOSITORY                           TAG       IMAGE ID       CREATED        SIZE
k8s.gcr.io/kube-apiserver            v1.22.4   8a5cc299272d   13 days ago    128MB
k8s.gcr.io/kube-controller-manager   v1.22.4   0ce02f92d3e4   13 days ago    122MB
k8s.gcr.io/kube-scheduler            v1.22.4   721ba97f54a6   13 days ago    52.7MB
k8s.gcr.io/kube-proxy                v1.22.4   edeff87e4802   13 days ago    104MB
k8s.gcr.io/etcd                      3.5.0-0   004811815584   5 months ago   295MB
k8s.gcr.io/coredns/coredns           v1.8.4    8d147537fb7d   6 months ago   47.6MB
k8s.gcr.io/pause                     3.5       ed210e3e4a5b   8 months ago   683kB






sudo kubeadm init --kubernetes-version=v1.22.4 --pod-network-cidr=10.244.0.0/16 --apiserver-advertise-address=192.168.105.5

或：
sudo kubeadm init --kubernetes-version=v1.22.4 --pod-network-cidr=10.244.0.0/16 --apiserver-advertise-address=192.168.105.5 --image-repository registry.aliyuncs.com/google_containers



Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 192.168.105.5:6443 --token 8dtoxq.mc8y47svyn1qbs3o \
	--discovery-token-ca-cert-hash sha256:276f67ebefc068beeb31005935889d6874a36cc26c12fd3663d6d2aea1c15e0d



echo "export KUBECONFIG=/etc/kubernetes/admin.conf" >> /etc/profile

source /etc/profile




For Kubernetes v1.17+ 

https://github.com/flannel-io/flannel


kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml

kubectl apply -f kube-flannel.yml


root@master:/home/ubuntu# kubectl apply -f kube-flannel.yml
Warning: policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
podsecuritypolicy.policy/psp.flannel.unprivileged created
clusterrole.rbac.authorization.k8s.io/flannel created
clusterrolebinding.rbac.authorization.k8s.io/flannel created
serviceaccount/flannel created
configmap/kube-flannel-cfg created
daemonset.apps/kube-flannel-ds created

kubectl get pods --all-namespaces



下载到本地
https://github.com/flannel-io/flannel/releases/download/v0.15.1/flanneld-v0.15.1-amd64.docker


docker load < flanneld-v0.15.1-amd64.docker

root@master:~# docker images
REPOSITORY                                       TAG             IMAGE ID       CREATED        SIZE
k8s.gcr.io/kube-apiserver                        v1.22.4         8a5cc299272d   2 weeks ago    128MB
k8s.gcr.io/kube-scheduler                        v1.22.4         721ba97f54a6   2 weeks ago    52.7MB
k8s.gcr.io/kube-controller-manager               v1.22.4         0ce02f92d3e4   2 weeks ago    122MB
k8s.gcr.io/kube-proxy                            v1.22.4         edeff87e4802   2 weeks ago    104MB
quay.io/coreos/flannel                           v0.15.1         e6ea68648f0c   2 weeks ago    69.5MB
quay.io/coreos/flannel                           v0.15.1-amd64   e6ea68648f0c   2 weeks ago    69.5MB
rancher/mirrored-flannelcni-flannel-cni-plugin   v1.0.0          cd5235cd7dc2   4 weeks ago    9.03MB
k8s.gcr.io/etcd                                  3.5.0-0         004811815584   5 months ago   295MB
k8s.gcr.io/coredns/coredns                       v1.8.4          8d147537fb7d   6 months ago   47.6MB
k8s.gcr.io/pause                                 3.5             ed210e3e4a5b   8 months ago   683kB
root@master:~# kubectl get pods --all-namespaces
NAMESPACE     NAME                             READY   STATUS    RESTARTS      AGE
kube-system   coredns-78fcd69978-hqq4v         1/1     Running   0             97m
kube-system   coredns-78fcd69978-qm6qq         1/1     Running   0             97m
kube-system   etcd-master                      1/1     Running   0             98m
kube-system   kube-apiserver-master            1/1     Running   0             98m
kube-system   kube-controller-manager-master   1/1     Running   0             98m
kube-system   kube-flannel-ds-kkqf4            1/1     Running   0             26m
kube-system   kube-proxy-h5p5z                 1/1     Running   0             97m
kube-system   kube-scheduler-master            1/1     Running   1 (86m ago)   98m

root@master:~# kubectl get nodes
NAME     STATUS   ROLES                  AGE    VERSION
master   Ready    control-plane,master   102m   v1.22.4


```

## flannel
```go
https://github.com/flannel-io/flannel


chrome下载到本地：
https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml

vim kube-flannel.yml

kubectl apply -f kube-flannel.yml
```





## node1 Kubernetes Dashboard
```go
kubectl apply -f recommended.yaml

vim recommended.yaml
chrome下载到本地：
https://raw.githubusercontent.com/kubernetes/dashboard/v2.4.0/aio/deploy/recommended.yaml

```

## ## node1 加入集群
```go

kubeadm join 192.168.105.5:6443 --token 8dtoxq.mc8y47svyn1qbs3o \
	--discovery-token-ca-cert-hash sha256:276f67ebefc068beeb31005935889d6874a36cc26c12fd3663d6d2aea1c15e0d

This node has joined the cluster:
* Certificate signing request was sent to apiserver and a response was received.
* The Kubelet was informed of the new secure connection details.

Run 'kubectl get nodes' on the control-plane to see this node join the cluster.
```





## 错误❌解决：

```go

[wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
[kubelet-check] Initial timeout of 40s passed.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10248/healthz' failed with error: Get "http://localhost:10248/healthz": dial tcp [::1]:10248: connect: connection refused.


	Unfortunately, an error has occurred:
		timed out waiting for the condition

	This error is likely caused by:
		- The kubelet is not running
		- The kubelet is unhealthy due to a misconfiguration of the node in some way (required cgroups disabled)

	If you are on a systemd-powered system, you can try to troubleshoot the error with the following commands:
		- 'systemctl status kubelet'
		- 'journalctl -xeu kubelet'

	Additionally, a control plane component may have crashed or exited when started by the container runtime.
	To troubleshoot, list all containers using your preferred container runtimes CLI.



Failed to run kubelet" err="failed to run Kubelet: misconfiguration: kubelet cgroup driver: \"systemd\" is different from docker cgroup driver: \"cgroupfs\""


sudo vim /etc/docker/daemon.json

{
  "exec-opts": ["native.cgroupdriver=systemd"]
}

systemctl restart docker
systemctl status docker


sudo kubeadm reset

sudo rm -rf /etc/kubernetes/manifests/kube-apiserver.yaml
sudo rm -rf /etc/kubernetes/manifests/kube-controller-manager.yaml
sudo rm -rf /etc/kubernetes/manifests/kube-scheduler.yaml
sudo rm -rf /etc/kubernetes/manifests/etcd.yaml
sudo rm -rf /var/lib/etcd

sudo rm -rf $HOME/.kube/config

sudo kubeadm init --kubernetes-version=v1.22.4 --pod-network-cidr=10.244.0.0/16 --apiserver-advertise-address=192.168.105.5



```







## 在Docker中下载并运行Jenkins （在macOS和Linux上）
```go

docker run \
  -u root \
  --rm \
  -d \
  -p 8080:8080 \
  -p 50000:50000 \
  -v jenkins-data:/var/jenkins_home \
  -v /var/run/docker.sock:/var/run/docker.sock \
  jenkinsci/blueocean
```







```go

docker pull coredns/coredns:1.8.6
docker tag coredns/coredns:1.8.6 registry.aliyuncs.com/google_containers/coredns:v1.8.6
```



```git 
git switch -c v10.0

git push origin HEAD:v10.0
```