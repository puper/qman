# qman
- docker run --name=kafka -p 2181:2181 -p 9092:9092 --env ADVERTISED_HOST=127.0.0.1 --env ADVERTISED_PORT=9092 spotify/kafka
- docker run -d --name rabbitmq -p 5672:5672 rabbitmq

- etcd维护订阅关系
- 简单的http handler
- 数据持久化

# topic 数量


# message in mysql
- topic
- tag
- key
- partition
- value
- message_time

# message
- topic
- tag
- key
- value


- 消息

# subscription
    - topic
    - tag
    - handler
    - timeout
    - type 并发 / 局部有序
    - retry
    - max_flight_message

# message
- topic
- tag
- key
- data

# consumer
- 顺序消息
    - 阻塞，同时只有1个消费者
    - 单独topic
- 事务消息 X
- 普通消息

订阅
    - topic
    - tag

- 普通消息
    - 持久化pending消息

- 广播
    - 每个应用一个组

- 顺序
    - 固定topic， 并且分区唯一

- 并发
    - 随意分区

- 根据应用监听

- push
    - 消费者不在线问题

- pull
    - 

- agent模式

- worker模式
    - 


- agent
    - topics



- topic
- tag
    

- topic
- tag
- app
- handler
- max.in.flight.requests.per.connection


- topic
    - 顺序消费


- kv storage
- 每个订阅一个组
- 每个订阅一个flight list


# mesh模式


# 消费者分布式

- 一分代码 两种部署方式，异步任务和对外服务隔离
- 分topic
- 消费者组


- 同key的顺序


flight messages



- 存储
    - offset
    - 失败记录
    - pending的消息
    - 


- message
    - topic
    - partition
    - tag
    - key
    - offset
    - message_time
    - value

- failed_message
    - topic
    - partition
    - tag
    - key
    - offset
    - message_time
    - data
    - response

- etcd
    - subscriptions
        - name
        - topic
        - tag
        - handler_config

没有性别数据的病例
图片编辑
推荐标签

SELECT
 id ID, 
 code "药品编码", 
 standardCode "药品本位码", 
 (CASE WHEN category = 1 THEN '西药' WHEN category = 2 THEN '中成药' WHEN category = 3 THEN '中草药' END) '药品分类', 
 name '药品名称', 
 genericName '药品通用名称', 
 alias '药品别名',
 pinyin '药品拼音',
 images '图片json数组',
 (CASE WHEN type=1 THEN '非处方药' WHEN type=2 THEN '处方药' END) '类型',
 
 (CASE WHEN originType=1 THEN '国产' WHEN originType=2 THEN '进口' END) '产地类型',
 dosageForm '剂型',
 composition '药品成分',
 `character` '药品性状',
 packing '包装',
 storage '贮藏',
 approvalNum '批准文号/注册证号',
 note '说明/注意',
 manufacturer '生产厂商',
 manufacturerEn '生产厂商（英文）',
 brand '品牌',
 specification '规格',
 attending '功能主治',
 adverseReaction '不良反应',
 taboo '禁忌',
 drugInteract '药物相互作用',
 `usage` '用法',
 dosage '用量',
 unit '单位'




 FROM drug LIMIT 15000,5000;