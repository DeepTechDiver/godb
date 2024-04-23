# 重学Golang

> 能学到什么？

一步一步从零开始实现框架的编写
代码量少，学习编程思路，提高代码质量


> 用Go从零实现Web框架 - Gee
Gee 是一个模仿 gin 实现的 Web 框架，Go Gin简明教程可以快速入门。

一：前置知识(http.Handler接口)

二：上下文设计(Context) 

三：Trie树路由(Router)

四：分组控制(Group)

五：中间件(Middleware)

六：HTML模板(Template)

七：错误恢复(Panic Recover)

> 用Go从零实现分布式缓存 GeeCache
GeeCache 是一个模仿 groupcache 实现的分布式缓存系统

一：LRU 缓存淘汰策略

二：单机并发缓存

三：HTTP 服务端

四：一致性哈希(Hash)

五：分布式节点

六：防止缓存击穿

七：使用 Protobuf 通信

> 用Go从零实现ORM框架 GeeORM

GeeORM 是一个模仿 gorm 和 xorm 的 ORM 框架

gorm 准备推出完全重写的 v2 版本(目前还在开发中)，相对 gorm-v1 来说，xorm 的设计更容易理解，所以 geeorm 接口设计上主要参考了 xorm，一些细节实现上参考了 gorm。

一：database/sql 基础

二：对象表结构映射

三：记录新增和查询

四：链式操作与更新删除

五：实现钩子(Hooks)

六：支持事务(Transaction)

七：数据库迁移(Migrate)

> 用Go从零实现RPC框架 GeeRPC

GeeRPC 是一个基于 net/rpc 开发的 RPC 框架 GeeRPC 是基于 Go 语言标准库 net/rpc 实现的，添加了协议交换、服务注册与发现、负载均衡等功能，代码约 1k。

一 - 服务端与消息编码

二 - 支持并发与异步的客户端

三 - 服务注册(service register)

四 - 超时处理(timeout)

五 - 支持HTTP协议

六 - 负载均衡(load balance)

七 - 服务发现与注册中心(registry)
