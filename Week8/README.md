# 作业内容
* 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
* 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。


# 1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
本机redis已安装redis服务
### 命令
~~~
redis-benchmark -d 10 -t get,set
~~~
* -t 表示需要执行的命令，如set、get
* -c 表示客户端连接数
* -d 表示单条数据大小，单位 byte
* -n 表示测试包数量
* -r 表示使用随机key数量

### SET
|-|执行次数和耗时|每秒请求次数|
|----|----|----|
|10|100000 requests completed in 0.85 seconds|117370.89 requests per second|
|20|100000 requests completed in 0.92 seconds|109289.62 requests per second|
|50|100000 requests completed in 0.87 seconds|115207.38 requests per second|
|100|100000 requests completed in 0.89 seconds|112359.55 requests per second|
|200|100000 requests completed in 0.86 seconds|116959.06 requests per second|
|1k|100000 requests completed in 0.87 seconds|114547.53 requests per second|
|5k|100000 requests completed in 0.92 seconds|109289.62 requests per second|
### GET
|-|执行次数和耗时|每秒请求次数|
|----|----|----|
|10|100000 requests completed in 0.85 seconds|117785.63 requests per second|
|20|100000 requests completed in 0.85 seconds|118343.19 requests per second|
|50|100000 requests completed in 0.85 seconds|118063.76 requests per second|
|100|100000 requests completed in 0.85 seconds|118343.19 requests per second|
|200|100000 requests completed in 0.85 seconds|117647.05 requests per second|
|1k|100000 requests completed in 0.86 seconds|116686.12 requests per second|
|5k|100000 requests completed in 0.94 seconds|106609.80 requests per second|


# 2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

* 代码工作原理，写入不同数量不同长度的value, 分析内存占用, 导出结果 reports 目录下的 csv 文件
* 结论 相同长度的value在写入数量越多情况下，平均每个value占用内存更多