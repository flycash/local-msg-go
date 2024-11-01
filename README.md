# local-msg-go
通用本地消息表-Go语言实现

## 存储设计
在设计上，我们是支持分表，但是不支持分库的。如果要支持分库，需要用户自己手动进行分库。

之所以不支持分库，是因为受制于本地事务这个概念。也就是说本地消息必须要和业务混在一起，所以本地消息没有办法分库，它如果要分库，就只能跟随业务方来分库。

而业务方怎么分库，这个是我们没有办法决定的。
