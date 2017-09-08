# API文档

为了简洁明了,这里采用例子进行说明

假设服务器根地址为http://server.com/

## 注册

1. 方法: POST
2. 地址: http://server.com/register
3. 参数: username, password

httpie模拟请求如下

```
http -f POST server.com/register username="username" password="password"
```

注册成功,返回success

注册失败,返回wrong

## 登录

1. 方法: POST
2. 地址: http://server.com/login
3. 参数: username, password

httpie模拟请求如下

```
http -f POST server.com/login username="username" password="password"
```

登录成功,返回success

登录失败,返回wrong

## 连接WebSocket

这里的逻辑如下:在上一步调用登录接口,如果服务器返回success,代表服务器允许进行WebSocket连接,由于WebSocket连接只能由客户端主动发起,所以请手动发起连接

WebSocket连接地址

```
ws://server.com/ws
```

JS连接示例

```javascript
var ws = new WebSocket("ws://" + document.location.host + "/ws")
```

##WebSocket的事件绑定

上面提到,比较好的办法是在调用登录接口后如果服务器返回success就立即建立WebSocket连接,用JS表述为类似于下面代码的逻辑

```javascript
$.post(
  "/login",
  {
    username: sender, //sender为客户端采集的username
    password: password
  },
  function (data) {
    if (data == "success") {
      ws = new WebSocket("ws://" + document.location.host + "/ws") //这里马上建立了websocket连接
    } else {
      $("body").html("<h1>wrong password</h1>")
    }
  }
)
```

### WebSocket的onopen事件

在建立了WebSocket后,需要首先向服务器端WebSocket服务表明自己的身份,为了简便,**建议在onopen事件中首先发送一条消息表明身份**,用JS表述为类似于下面代码的逻辑

```javascript
ws.onopen = function (evt) {
  var message = {
    "Sender": sender, //sender为客户端采集的username
    "MessageType": "login", //注意这里的MessageType必须为login
    "Content": "null", //这两个乱填都可以,反正服务器不会看
    "Reciver": "null" //这两个乱填都可以,反正服务器不会看
  }
  ws.send(JSON.stringify(message))
}
```

### WebSocket的onclose事件

如果断开WebSocket,为了保险起见客户端**最好还是调用一下服务器的登出http接口**(不过不调用也行,服务器本来就在监测每个用户的WebSocket状态,会自动登出的,此外不用担心多次登出的问题,服务器已经做了处理)

用JS表述为类似于下面代码的逻辑

```javascript
ws.onclose = function (evt) {
  $.post(
    "/logout",
    {
      username: sender,
    }
  )
}
```

1. 方法: POST
2. 地址: http://server.com/logout
3. 参数: username

httpie模拟请求如下

```
http -f POST server.com/logout username="username"
```

登出成功,返回success

登出失败,返回wrong

客户端自动断开或是意外断开的话其实不用看返回值,登出的判断服务器端判断已经比较自动了

这个api主要用于主动登出(其实主动登出也可以暴力登出嘛,何必调用api)

## 通过WebSocket向服务器发送消息

message JSON格式

JS示例

```javascript
var message = {
  "Sender": sender,
  "MessageType": $("#MessageType").val(),
  "Content": $("#Content").val(),
  "Reciver": $("#Reciver").val()
}
```

分为四个字段

Sender实际是客户端的username

MessageType分为login, send_to_user

Content在MessageType是send_to_user的时候是消息内容

Reciver为目标用户的username

用JS示例一次消息发送

```javascript
$("#user-panel").submit(function () {
  if ($("#Content").val() == "" || $("#MessageType").val() == "" || $("#Reciver").val() == "") {
    return false
  }
  var message = {
    "Sender": sender,
    "MessageType": $("#MessageType").val(),
    "Content": $("#Content").val(),
    "Reciver": $("#Reciver").val()
  }

  ws.send(JSON.stringify(message))

  $("#Content").val("")
  $("#Reciver").val("")
  return false
})
```

发送了消息后,服务器会向自己回显这条消息,同时服务器会向目标用户推送这条消息

不用担心格式错误,或是目标用户不存在,或是目标用户不是自己的好友,服务器会过滤

### WebSocket的onmessage事件

假如有用户通过上面的方法发送了消息,并且消息的Reciver是自己的话,服务器就会把这条message推送到自己,如果自己不在线,自己下次登录的时候服务器会把消息推送过来

用JS示例一次消息接收

```javascript
ws.onmessage = function (evt) {
  var message = JSON.parse(evt.data)
  $("#messages").append($("<li>").text(JSON.stringify(message)))
}
```

## 获取好友列表

采用http轮询

1. 方法: POST
2. 地址: http://server.com/getfriends
3. 参数: username

httpie模拟请求如下

```
http -f POST server.com/addfriend username="username"
```

请求成功,返回好友列表,好友之间用","分隔

请求失败,返回wrong

用JS示例

```javascript
var getfriends = function () {
  $.post(
    "/getfriends",
    {
      username: sender
    },
    function (data) {
      if (data != "wrong") {
        friends = data.split(",")
        $("#friends").html("")
        for (var i = 0; i < friends.length; ++i) {
          $("#friends").append($("<li>").text(friends[i]))
        }
      }
    }
  )
}
setInterval(function () {
  getfriends()
}, 2000) //两秒查询一次
```

## 添加好友

用JS示例

```javascript
$("#add-friend").submit(function () {
  if ($("#add-friend-name").val() == "") {
    return false
  }
  $.post(
    "/addfriend",
    {
      username: sender,
      friend: $("#add-friend-name").val()
    },
    function (data) {
      $("#add-friend-name").val("")
      getfriends() //主动重新获取好友列表
    }
  )
  return false
})
```

1. 方法: POST
2. 地址: http://server.com/addfriend
3. 参数: username, friend

httpie模拟请求如下

```
http -f POST server.com/addfriend username="username" friend="friend-username"
```

添加成功,返回success

添加失败,返回wrong

添加好友是双向的,自己加对方为好友,对方也会自动添加自己为好友

### 删除好友

用JS示例

```javascript
$("#delete-friend").submit(function () {
  if ($("#delete-friend-name").val() == "") {
    return false
  }
  $.post(
    "/deletefriend",
    {
      username: sender,
      friend: $("#delete-friend-name").val()
    },
    function (data) {
      $("#delete-friend-name").val("")
      getfriends()
    }
  )
  return false
})
```

1. 方法: POST
2. 地址: http://server.com/deletefriend
3. 参数: username, friend

httpie模拟请求如下

```
http -f POST server.com/deletefriend username="username" friend="friend-username"
```

删除成功,返回success

删除失败,返回wrong

删除好友是双向的,自己删除对方为好友,对方也会在好友列表中自动删除你