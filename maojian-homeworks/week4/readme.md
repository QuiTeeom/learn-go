## 第四周作业

### Q
1. 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。
### A
> 时间问题，没能补充代码，暂时写下目录结构

实现一个基于etcd的分布式任务调度平台
#### 服务端 ./jobd-repo
* 提供web管理页面
* 提供api接口
* job调度
* ...

目录结构
```
|-- api 服务的api定义，比如pb swagger 等
|   |-- api-server
|   |   |-- v1 v1版本的接口
|   |-- schedule-server
|   |   |-- v1 v1版本的接口
|-- app 各个应用服务
|   |-- api-server  api服务
|   |   |-- cmd
|   |   |   |-- server
|   |   |   |   |-- main.go api-server的启动程序，负责服务生命周期的管理
|   |   |-- configs 存放服务使用的配置文件
|   |   |-- internal
|   |   |   |-- biz 核心业务逻辑代码
|   |   |   |   ｜-- entity 领域对象
|   |   |   |   ｜-- repos 定义repo接口
|   |   |   |   ｜-- service 领域服务
|   |   |   |-- conf 配置相关代码和工具
|   |   |   |-- infra 基础设施，我还是更倾向使用infra来表示，而不是data
|   |   |   |-- pkg 一些工具类型或者internal内公用的方法
|   |   |   |-- server http、rpc服务启动，注册等，相当于服务的 facade
|   |   |   |-- service 应用层服务
|   |-- schedule-server  调度服务
|   |   ...同api-server
|-- configs  工程的配置文件
|-- build 构建相关信息
|-- deployments 部署相关
|-- docs 文档
|-- tools  工具文件
...其他 license，readme等开源规范
```

#### 客户端 ./jobd-client-go-repo
* 提供golang sdk
```
|-- docs 文档
|-- examples demo案例
|-- pkg 一些公用方法
|-- v1 v1版本的sdk
|-- v2 v2版本的sdk
...其他 license，readme等开源规范
```