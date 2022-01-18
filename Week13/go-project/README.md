## 毕业总结

时光匆匆，接近课程尾声，经过13周的学习，从一开始对go一知半解，到补习基础语法，从第四周作业一脸懵无从下手，到现在照猫画虎大概能写个大概，把大致的依赖注入等用上。发现学习没有捷径，只能脚踏实地多学多练。目前工作中用不到go，所以学习起来有些吃力，不管怎么样，还是有进步的。

学习过程中，感谢班班领教助教们辛苦付出和有趣的学员们相互帮助，都太卷了，不学不行，课程也还需要自己逐步实践，学习要专注，不能东一耙子西一扫帚，按部就班学习，焦虑无用，就像毛老师说的活到老学到老。还有杯子上的英文：Stay hungry. Stay foolish. 看到这个解释很赞同，简单地说 Stay hungry 就是 hungry for success，是说要永不满足，Stay foolish 是说埋头做自己的事，不要理会前行路上的各种嘲讽声音。

2022继续加油前行！

## 毕业作业：进行一个微服务改造，需要考虑如下技术点：

* 微服务架构（BFF、Service、Admin、Job、Task 分模块）
* API 设计（包括 API 定义、错误码规范、Error 的使用）
* gRPC 的使用
* Go 项目工程化（项目结构、DI、代码分层、ORM 框架）
* 并发的使用（errgroup 的并行链路请求）
* 微服务中间件的使用（ELK、Opentracing、Prometheus、Kafka）
* 缓存的使用优化（一致性处理、Pipeline 优化）

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
