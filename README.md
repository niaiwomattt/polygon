# 多边形架构

全新的多边形架构，多监听，多驱动!

分层设计，无多层依赖，高内聚，低耦合，方便扩展和变换

基于接口实现，注册式，没有强依赖，只有需要时注册即可。


## 依赖层次
listener->handler->storer->driver

每层只依赖下一层，所有层中的配置，通过注册制，注册到 config 层中，只有存在当前层时，才有此配置，否则不存在配置。

每层只对上一层的接口负责，返回接口类型数据和对象，不对具体对象做依赖。

实际中，listener 和 driver 可能使用同一个包的俩种方式，比如 nats 的 订阅 和发布，这样的时候，listener 中包含自身的订阅器，不要和 driver 产生依赖，否则会造成循环依赖问题。

## 指导思想

1. [日志的艺术](https://www.cnblogs.com/xybaby/p/7954610.html)
2. [给应用加上日志](https://www.imzjy.com/blog/2018-07-06-writing-good-log)
3. 日志记录的时候可以设置记录级别，输出的分发器，也可以设置自己要的日志级别，等等
4. [优秀日志实践准则](https://cloud.tencent.com/developer/article/1116081)
5. [优秀日志实践准则2](https://dbaplus.cn/news-134-1658-1.html)
6. [兢兢业业的日志](https://cloud.tencent.com/developer/article/1625404)
