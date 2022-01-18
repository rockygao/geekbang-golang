## 对当下自己项目中的业务，进行一个微服务改造，需要考虑如下技术点：

* 微服务架构（BFF、Service、Admin、Job、Task 分模块）
* API 设计（包括 API 定义、错误码规范、Error 的使用）
* gRPC 的使用
* Go 项目工程化（项目结构、DI、代码分层、ORM 框架）
* 并发的使用（errgroup 的并行链路请求）
* 微服务中间件的使用（ELK、Opentracing、Prometheus、Kafka）
* 缓存的使用优化（一致性处理、Pipeline 优化）

## 感受

时光匆匆，接近课程尾声，目前工作中用不到go，所以学习起来有些吃力，不管怎么样，还是有进步的，从第四周作业一脸懵，无从下手，到现在照猫画虎大概能写个大概，把大致的依赖注入等用上。

## 整理

### 目录结构

- /api 
  - APi协议定义目录 使用proto文件定义具体服务的接口形状
- /cmd
  - 整个项目启动的入口文件
- /configs
  - 配置文件模板或默认配置
- /pkg
  - app    存放服务的具体实现代码，项目启动入口
  - boot   初始化
  - database 数据库连接
  - domain 领域对象
  - ginex 
  - repo 操作数据库
  - ut 服务公用的工具方法

### api定义

使用proto文件定义具体服务的接口形状

~~~
// api目录下 创建了user.proto 在对应目录执行
protoc -I ./ --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative user.proto
~~~

### 依赖注入

通过服务中的每一层使用wire.NewSet方法暴露出来的构造器函数，然后在pkg/app/srv的wire.go文件中使用暴露的构造器函数中的依赖定义，对各层的服务进行组装

### 使用说明

启动服务

~~~
//因读取配置文件相对路径问题，需要在项目目录直接执行下面命令
go run cmd/srv/main.go
~~~
