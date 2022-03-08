---
作业管理系统

---

## 作业管理系统 



### 需求: (我要的功能)

- :package:  **能上传并存储作业文件**
- :baby_chick:  **所有人能查看哪个同学交了，哪个同学没交，这样可以互相提醒**
- :biking_woman: **以学期、课程、和任务名字来管理文件**
- :balloon: **操作简单，不需要同学们登录，能随时查看他们上传了什么东西上来。**(PS: 不许上传涩图啊喂！！)
- :watch: **有显示作业截止时间，并且要有距离倒计时**
- :bowing_woman: **能随时创建任务，随时修改删除任务**
- :rage:   **还要有展示我伟岸形象的彩蛋**



### 前端页面展示

> [前端项目地址](https://github.com/yzcyayaya/task_upload_vue)

#####  同学们使用界面

> **主页**
>
> [![b2Sk7j.png](https://s1.ax1x.com/2022/03/08/b2Sk7j.png)](https://imgtu.com/i/b2Sk7j)
>
> [![b2FsUg.gif](https://s1.ax1x.com/2022/03/08/b2FsUg.gif)](https://imgtu.com/i/b2FsUg)

> 上传文件界面（**上传文件名必须包含学号，因为学号是我来判断他们交没交的唯一标识**）
>
> [![b2SwuD.png](https://s1.ax1x.com/2022/03/08/b2SwuD.png)](https://imgtu.com/i/b2SwuD)

> 查看提交人数
>
> [![b2Sy4I.png](https://s1.ax1x.com/2022/03/08/b2Sy4I.png)](https://imgtu.com/i/b2Sy4I)

> 彩蛋
>
> [![b2pP8x.gif](https://s1.ax1x.com/2022/03/08/b2pP8x.gif)](https://imgtu.com/i/b2pP8x)



------



##### 我使用的界面

- 主页

  > 任务管理
  >
  > [![b2pMGt.png](https://s1.ax1x.com/2022/03/08/b2pMGt.png)](https://imgtu.com/i/b2pMGt)

  > 任务编辑
  >
  > [![b2pGqg.png](https://s1.ax1x.com/2022/03/08/b2pGqg.png)](https://imgtu.com/i/b2pGqg)

  > 添加任务
  >
  > [![b2prsU.png](https://s1.ax1x.com/2022/03/08/b2prsU.png)](https://imgtu.com/i/b2prsU)

- 课程页面

  > 课程首页

  > [![b2pIsO.png](https://s1.ax1x.com/2022/03/08/b2pIsO.png)](https://imgtu.com/i/b2pIsO)

  > 课程添加
  >
  > [![b2pHdH.png](https://s1.ax1x.com/2022/03/08/b2pHdH.png)](https://imgtu.com/i/b2pHdH)

- 管理文件（minio）

  [![b2FlDK.png](https://s1.ax1x.com/2022/03/08/b2FlDK.png)](https://imgtu.com/i/b2FlDK)



### 技术栈以及设计细节

| 后端 |  go  |  gin  |   gorm2    | minio | MySQL8 |
| :--: | :--: | :---: | :--------: | ----- | ------ |
| 前端 | Vue2 | Axios | element-UI |       |        |

> 任务表
>
> [![b2AHj1.png](https://s1.ax1x.com/2022/03/08/b2AHj1.png)](https://imgtu.com/i/b2AHj1)

> 课程表
>
> [![b2EFDP.png](https://s1.ax1x.com/2022/03/08/b2EFDP.png)](https://imgtu.com/i/b2EFDP)

> 学生表
>
> [![b2EVUS.png](https://s1.ax1x.com/2022/03/08/b2EVUS.png)](https://imgtu.com/i/b2EVUS)

> 文件表 (只做了设计，并没有去实现)
>
> [![b2ElD0.png](https://s1.ax1x.com/2022/03/08/b2ElD0.png)](https://imgtu.com/i/b2ElD0)

### 运行

>  表结构不需要你过问，因为做了数据迁移，只要你数据库名字存在即可，配置文件在config/config.yaml
>
>  students.yaml为自己班级的名字和学号，按照模板来填写即可

```
# 安装包
go mod tidy
# 连接配置好mysql和minio，再修改students.yaml为自己班级同学
go run main.go
# 或者
go build task -o main.go
./task
```

> 推荐goland和webstorm打开运行即可

### 部署

- docker 

  >  二进制部署（由于我系统为Ubuntu, 该二进制部署方法不能使用）

  ```
  ## dockerfile  以下为正文
  ---------------------------------------------------------------------------------
  ## linux 直接采用二进制文件运行
  FROM  ubuntu:latest
  
  ## 设置时区为上海
  RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
  RUN echo 'Asia/Shanghai' >/etc/timezone
  ## 设置编码
  ENV LANG C.UTF-8
  ## 添加当前目录二进制文件进容器当中
  ADD ./task ./task
  CMD ["./task"]
  ```

  > 再配置好docker，可以远程服务器或者自己本地，**--network host不设置会设计容器通讯问题**

  [![b2C1Hg.png](https://s1.ax1x.com/2022/03/08/b2C1Hg.png)](https://imgtu.com/i/b2C1Hg)

  > 没有报错则成功了

  [![b2kPGd.png](https://s1.ax1x.com/2022/03/08/b2kPGd.png)](https://imgtu.com/i/b2kPGd)

  ```
  ##此时可以docker ps查看
  dockers ps
  ## 如果没有看见自己设置容器名字则查看日志 xx为容器id或者名字
  dockers ps -a
  docker logs xx
  ```

- 编译部署同理

  > 略..

### 简单粗暴部署法：

先装好mysql和minio。

##### 后端

> 直接go build 将二进制copy进服务器
>
> ```shell
> go build task -o main.go
> ```
>
> 在服务器直接nohup  日志文件输入out.txt  后面的与符合为后台运行
>
> ```
> nohup ./task >out.txt&
> ```

前端

> 将build后的所有文件放入Apache服务区的渲染文件夹里面
>
> 默认地址为/var/www/html

