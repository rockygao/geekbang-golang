# 对当下自己项目中的业务，进行一个微服务改造，需要考虑如下技术点：

* 微服务架构（BFF、Service、Admin、Job、Task 分模块）
* API 设计（包括 API 定义、错误码规范、Error 的使用）
* gRPC 的使用
* Go 项目工程化（项目结构、DI、代码分层、ORM 框架）
* 并发的使用（errgroup 的并行链路请求）
* 微服务中间件的使用（ELK、Opentracing、Prometheus、Kafka）
* 缓存的使用优化（一致性处理、Pipeline 优化）

# 感受
时光匆匆，接近课程尾声，目前工作中用不到go，所以学习起来有些吃力，不管怎么样，还是有进步的

# 启动
~~~
//因读取配置文件相对路径问题，需要在项目目录直接执行下面命令
go run cmd/srv/main.go
~~~