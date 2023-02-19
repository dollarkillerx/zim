# ZIM
Zim, open source instant messaging software

[中文文档](README_ZH.md)

## System Design

![](./doc/img/ZIM.jpg)

#### System Layering
- access layer
    - Chat service access layer
        - Websockets
        - GRPC
        - UDP
    - Management terminal access layer
      - GRPC
      - GraphQL
      - RESTful
- Message preprocessing
    - Message ID generator
    - NSQ processing queue
- User online status
  - redis
- API service
    - User relationship
    - Group relationship
    - sign up
    - pgsql
- im push service Timeline
    - One to one: single box, multiple write
    - Standard group: single box, multi-write
    - Thousands of Crowds: Multibox, Read More

