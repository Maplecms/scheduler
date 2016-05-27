###  统一任务调度平台 for golang

###### 1、pre
* install golang env

  https://github.com/golang/go

#### 2、install  



*        go get github.com/astaxie/beego

*        go get github.com/shotdog/scheduler

*       go get github.com/shotdog/quartz

*       go get  github.com/go-sql-driver/mysql
*       init db  scheduler.sql
*       modify conf/app.conf -->database config


#### 3、run

*       cd $GOPATH
*       cd src
*       cd scheduler
*       go build main.go
*       ./main


#### 4、Screenshot

![image](https://github.com/shotdog/scheduler/raw/master/screenshot/1.png)

![image](https://github.com/shotdog/scheduler/raw/master/screenshot/2.png)

![image](https://github.com/shotdog/scheduler/raw/master/screenshot/3.png)

![image](https://github.com/shotdog/scheduler/raw/master/screenshot/4.png)

![image](https://github.com/shotdog/scheduler/raw/master/screenshot/5.png)

![image](https://github.com/shotdog/scheduler/raw/master/screenshot/6.png)

![image](https://github.com/shotdog/scheduler/raw/master/screenshot/7.png)

#### 5、Protocol
* see [invoker.go](https://github.com/shotdog/scheduler/blob/master/invoker/invoker.go)

#### 6、Client Test

* see [scheduler-client](https://github.com/shotdog/scheduler-client)

   * cd scheduler-client
   * go build main.go
   * ./main
