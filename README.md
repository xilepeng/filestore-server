# LeiliNetdisk




# V1.0   云存储”系统原型(实现一个超精简版云盘)


## 环境配置

```s

➜  ~ multipass launch -n main -c 1 -m 4G -d 20G
Launched: main
➜  ~ multipass shell main


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



➜  share multipass mount /Users/x/share main:/home/ubuntu/share
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


```s
ubuntu@main:~$ sudo snap install docker
docker 20.10.8 from Canonical✓ installed

ubuntu@main:~$ sudo snap install go --classic
go 1.17.3 from Michael Hudson-Doyle (mwhudson) installed


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
```



## mysql 环境配置

```sql
ubuntu@main:~$ sudo docker pull mysql

sudo docker run -itd --name master -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql

sudo docker run -itd --name slave -p 3307:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql


配置Master

使用如下命令进入到Master容器内部，使用容器ID或者名称均可：

ubuntu@main:~$ sudo docker exec -it master /bin/bash
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