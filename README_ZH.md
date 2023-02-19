# ZIM
Zim, 开源即时通讯软件

## System Design 

![](./doc/img/ZIM.jpg)

#### 系统分层
- 接入层  
  -  聊天服务接入层 
    - Websocket
    - GRPC
    - UDP
  - 管理端接入层
    - GRPC
    - GraphQL
    - RESTful
- 消息预处理
  - 消息ID生成器
  - NSQ 处理队列
- 用户在线状态
  - redis
- API服务
  - 用户关系
  - 群组关系
  - 登陆注册
  - pgsql
- im推送服务 Timeline
  - 一对一：单件箱，多写
  - 标准群：单件箱，多写
  - 千人群： 多件箱，多读

#### 数据库设计

##### 表设计


> super_admin 最高角色表

- 表名称：super_admin
- 表描述：最高角色表
- 存储方式： Pgsql

| 行名称      | 行描述     | 行类型    |
|----------|---------|--------|
| SupID    | 超管id    | uuid   |
| SupToken | 超管token | String |

> 

