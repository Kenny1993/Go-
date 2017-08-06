# 开发环境

## Vagrant 虚拟机管理工具

Vagrant是一个虚拟机管理工具，因此需要安装 VirtualBox 虚拟化系统软件，官网：https://www.virtualbox.org/ ,然后还需要安装Vagrant，官网： https://www.vagrantup.com ，安装完成后，使用 vagrant version 测试是否安装成功。

Vagrant 和 VirtualBox 都安装好后，下载想要在虚拟机上安装的操作系统，这个操作系统环境是以 .box 结尾的文件，实际上是一个 zip 压缩包，astaxie 为我们提供了一个 Ubuntu lucid 64 操作系统，这个版本过旧，有些软件会因为和依赖的软件版本不兼容而导致安装时会有bug，我从官网 （ https://vagrantcloud.com/boxes/search ）下载了最新的Ubuntu 14.04.5 LTS ，然后进入下载目录，使用命令行 vagrant box add box --name “ubuntu”。

成功添加 box 后，到想要建立虚拟机的目录下使用初始化命令 vagrant init，会生成一个 Vagrantfile 文件，然后启动虚拟机 vagrant up，启动成功后，可通过 ssh 连接虚拟机 vagrant ssh

下面讲解 Vagrantfile 配置文件

    /# -- mode: ruby --

    /# vi: set ft=ruby :

    Vagrant.configure("2") do |config|

    /# 虚拟机默认初始化 box ，可以在 vagrant init 命令中指定要初始化的 box

    config.vm.box = "ubuntu"

    /# 设置虚拟机内存大小

    config.vm.provider "virtualbox" do |vb|

    vb.memory = "1024"

    end

    /# 设置 host-only 主机模式，静态 IP

    config.vm.network "private_network", ip: "192.168.2.1"

    /# 设置局域网模式，外网可直接访问虚拟机IP

    /# config.vm.network "public_network",ip: “192.168.2.1”

    /# 设置虚拟机主机名 config.vm.hostname = "kenny"

    /# 设置同步目录，第一个参数是主机目录，第二个参数是虚拟机挂载目录，默认会把当前虚拟机所在目录同步到虚拟机操作系统中的/vagrant目录

    /#config.vm.synced_folder "F:/vagrant_data", "F:/vagrant/data"

    /# 设置转发端口，如主机的8080端口转发到虚拟机80端口

    config.vm.network :forwarded_port, guest: 80, host: 8080

    end

设置完成后，使用 vagrant reload 重启 VM，以使配置生效

vagrant 常用命令

vagrant box add	添加box的操作

vagrant box remove	删除box的操作

vagrant init 初始化box的操作

vagrant up	启动虚拟机的操作

vagrant ssh	登录虚拟机的操作

vagrant box list	显示所有box

vagrant destroy	停止虚拟机，销毁所有相关资源

vagrant halt	虚拟机关机

vagrant package	打包虚拟机

vagrant plugin	管理虚拟机插件

vagrant provision	配置管理工具（通常使用 chef 或 puppet）

vagrant status	查看虚拟机状态

vagrant resume	虚拟机恢复运行

vagrant suspend	虚拟机挂起

vagrant ssh-config	查看ssh连接信息

Vagrant 模拟分布式系统，通过修改配置文件就可以实现

如一台 WEB 服务器和一台 DB 服务器

    Vagrant.configure("2") do |config|

    config.vm.define :web do |web|

    web.vm.provider "virtualbox" do |v|

      v.customize ["modifyvm", :id, "--name", "web", "--memory", "512"]
      
    end

    web.vm.box = "base"

    web.vm.hostname = "web"

    web.vm.network :private_network, ip: "11.11.1.1"
    end

    config.vm.define :db do |db|

    db.vm.provider "virtualbox" do |v|

      v.customize ["modifyvm", :id, "--name", "db", "--memory", "512"]
      
    end

    db.vm.box = "base"

    db.vm.hostname = "db"

    db.vm.network :private_network, ip: "11.11.1.2"
    end

    end

使用 vagrant up 启动以上配置的两台虚拟机后，要使用 vagrant ssh web 或 vagrant ssh db 来登录虚拟机

使用 vagrant 模拟分布式系统，两台 web 服务器，一台 redis 服务器，一台 db 服务器

    /# -- mode: ruby --

    /# vi: set ft=ruby :

    Vagrant.configure("2") do |config|

    config.vm.define :web1 do |web1|

    web1.vm.provider "virtualbox" do |v|

      v.customize ["modifyvm", :id, "--name", "web1", "--memory", "512"]
          
    end
    web1.vm.box = "ubuntu"

    web1.vm.hostname = "web1"

    web1.vm.network :private_network, ip: "11.11.1.1"
    end

    config.vm.define :web2 do |web2|

    web2.vm.provider "virtualbox" do |v|

      v.customize ["modifyvm", :id, "--name", "web2", "--memory", "512"]
      
    end

    web2.vm.box = "ubuntu"

    web2.vm.hostname = "web2"

    web2.vm.network :private_network, ip: "11.11.1.2"
    end

    config.vm.define :db do |db|

    db.vm.provider "virtualbox" do |v|

      v.customize ["modifyvm", :id, "--name", "db", "--memory", "512"]
      
    end

    db.vm.box = "ubuntu"

    db.vm.hostname = "db"

    db.vm.network :private_network, ip: "11.11.1.3"
    end

    config.vm.define :redis do |redis|

    redis.vm.provider "virtualbox" do |v|

      v.customize ["modifyvm", :id, "--name", "redis", "--memory", "512"]
      
    end

    redis.vm.box = "ubuntu"

    redis.vm.hostname = "redis"

    redis.vm.network :private_network, ip: "11.11.1.4"
    end

    end

## 下载 Go

配置系统编码为 UTF-8，因为Go是基于UTF-8编码的

    sudo vi /etc/default/locale

    LANG="en_US.UTF-8"

    LANGUAGE="en_US:en"

wget 下载Go

    sudo wget --no-check-certificate -O /usr/local/go1.8.3.linux-amd64.tar.gz https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz

    sudo tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz

## 配置环境变量

    sudo vi /etc/profile

    export GOROOT=/usr/local/go

    export GOBIN=$GOROOT/bin

    export GOPATH=/vagrant/gopath

    export PATH=$GOBIN:$GOPATH/BIN:$PATH

### 查看 go 环境变量

    go env

### 安装 git

ubuntu 使用 apt-get install git

### 测试 git 是否安装成功

    git version

## Go安装目录

/bin：包含可执行文件，如：编译器，Go 工具

/doc：包含示例程序，代码工具，本地文档等

/lib：包含文档模版

/misc：包含与支持 Go 编辑器有关的配置文件以及 cgo 的示例

/os_arch：包含标准库的包的对象文件（.a）

/src：包含源代码构建脚本和标准库的包的完整源代码（Go 是一门开源语言）

/src/cmd：包含 Go 和 C 的编译器和命令行脚本

### Go命令

go install

go fix

go test

go build

go run

go doc

go fmt

## Go runtime

runtime 类似 Java 和 .NET 语言所用到的虚拟机，它负责管理包括内存分配、垃圾回收、栈处理、goroutine、channel、切片（slice）、map 和反射（reflection）等等。

垃圾回收器采取高效的标记-清除算法

## GOPATH

GOPATH 用于存放Go 语言源码文件、包文件和可执行文件的路径，而这些路径下又必须分别包含三个规定的目录：src、pkg 和 bin，这三个目录分别用于存放源码文件、包文件和可执行文件。

如：go get -u -v github.com/astaxie/beego

完成后，可在环境变量 GOPATH/src/ 目录下找到该项目源码

## GOX

交叉编译，实现GO跨平台

go get -u -v github.com/mitchellh/gox

gox -h，测试gox命令是否可用

gox -build-toolchain 编译跨平台库（go 1.5 之后不需要编译了）

进入 GOPATH，编辑 hello.go

    package main

    import "fmt"

    func main() { 
        fmt.Println("hello world")
    }

然后执行 gox 进行交叉编译，编译完成后，生成各平台的可执行文件

src_darwin_386*

src_darwin_amd64*

src_freebsd_386*

src_freebsd_amd64*

src_freebsd_arm*

src_linux_386*

src_linux_amd64*

src_linux_arm*

src_netbsd_386*

src_netbsd_amd64*

src_netbsd_arm*

src_openbsd_386*

src_openbsd_amd64*

src_windows_386.exe*

src_windows_amd64.exe*

我的操作系统是 linux 64位，因此可以执行 src_linux_amd64* 文件

gox 也可以指定平台进行编译，gox -osarch "windows/amd64 linux/amd64"

进阶学习 https://github.com/laher/goxc
