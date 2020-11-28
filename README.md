# 聊天室

## feature
- 登录/注册(防止重复登录)
- 群聊(支持文字、emoji、文件(图片)上传、多房间)
- 私聊(消息提醒)
- 历史消息查看(点击加载更多)
- 心跳检测，来自 https://github.com/zimv/websocket-heartbeat-js
- go mod 包管理
- 静态资源嵌入，运行只依赖编译好的可执行文件与mysql
- 支持 http/ws 、 https/wss


## todo
- [x] 心跳机制
- [x] 多频道聊天
- [x] 私聊
- [x] 在线用户列表
- [x] https支持