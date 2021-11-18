


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