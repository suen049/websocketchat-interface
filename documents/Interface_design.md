# 接口设计

## 基本交互

1. 用户可以将消息发送给指定的对象
2. 用户可以将消息发送到特定的频道

## 基于基本交互功能的扩展

1. 可以创建用户
2. 用户可以登录登出
3. 用户可以通过id添加好友
4. 用户可以创建频道
5. 创建频道的用户为频道管理员
6. 频道管理员可以添加好友至频道
7. 用户可以退出频道
8. 系统预置部分公共频道

## 服务器端数据结构

### 用户结构(user)

#### 字段

1. id(string)
2. friends([]string),用于存储好友
3. channels([]string),用于存储频道

### 消息结构(message)

#### 字段

1. 发送者
2. 发送类别(用户->用户,用户->频道)
3. 内容
4. 接受者(用户,频道)

### 用户长链接集合结构(usermap)

#### 字段

1. map[string]*websocket.Conn,每个用户登陆后加入此map
2. sync.RWMutex锁,由于用户使用http方式登录,必定涉及到多线程访问上述map,所以在更改上述map时加写锁(不排读),而不干扰消息推送时读取该map

## 服务器端消息推送

1. 用户登录后建立长链接,并且加入usermap中的map[string]*websocket.Conn(需对usermap加锁)
2. 用户通过websocket向服务器推送消息,服务器接收消息后将消息推送至make(chan message)的chan实例
3. 服务器端维护一个无限循环,该无限循环不断从make(chan message)的chan实例中读取message,并且通过message中的字段在usermap中查询特定的对象所对应的*websocket.Conn,并推送消息
4. 如果在上一步的查询过程中没有查询到对应的*websocket.Conn,说明对方未上线,那么将此消息通过chan发送至另一个服务器维护的无限循环,该无限循环收到消息后将消息加入一个map[string][]message
5. 每个用户登录的时候,在建立websocket之初,未加入usermap之前,先读取map[string][]message,将推送给自己的消息取出,将自己从map[string][]message中删除,并且将消息推动给自己,推送完成后再加入usermap
