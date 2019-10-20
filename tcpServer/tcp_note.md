# 备忘
## 代码
`
    // for和go程同时执行，for比go程先执行完，name永远是最后迭代的值
    for _,name :=j range names { // 如果names是个chan的话，会阻塞
        go func() { fmt.Println(name) }
    }
`

## 通道
elm, ok := <-IntChan （元素副本，map是引用）
1，通道没有元素时 -》 阻塞
2，关闭通道 -》 相当于chan<-nil，select-case结束（ok==false，说明chan被赋值nil，而不是因为通道关闭）
3，永远不要在接收端关闭通道；发送端，向一个已关闭的通道发送数据会引发panic
4，单向通道用在参数和返回值上，传递给它一个双向通道，起约束作用

# 高性能Go并发建议 
1. 管道chan传递业务数据会耗时，避免分割到多个go程-》减少chan的使用。
2. 互斥锁Mutex在高并发下性能不佳-》将数据分区，先定位再保护。
3. select随管道数增加性能线性下降。
4. 避免动态创建go程。
5. 读操作技术方案优先级：CDN -》 server内存 -》 数据库。

# go socket api
## net.Dial
1，网络不可达；2，服务没有启动 -》立即返回错误；
3，server backlog满了 -》阻塞；
4，网络延迟 -》阻塞并超时

## conn.Read
socket无数据 -》 “阻塞”，go程监听socket
socket有数据 -》 read读取并返回
socket关闭 -》 io.EOF
读取超时 -》 不会出现“读出部分且返回超时”的情况

## conn.Write
TCP通信两端的OS都会有数据缓冲区，写满了就会阻塞，所以Write在对方socket关闭后并不一定返回错误，因为它成功写到缓冲区了。

## socket属性设置
SetKeepAlive | SetKeepAlivePeriod | SetRead/WriteBuffer |
tcpConn, ok := conn.(*TCPConn); tcpConn.SetNoDelay(true)

## runtime包
提供控制goroutine运行的api。

# 连接池


# teleport
teleport 一个通用、高效、灵活的Socket框架。

## 特性
高性能、高效开发、DIY应用层协议、Body编码协商、RPC范式、插件、推送、连接管理、（Socket文件描述符/会话管理/上下文等）
兼容HTTP协议、平滑关闭/升级、Log接口、非阻塞异步IO、断线重连、对等通信、对等API、反向代理、慢响应报警
可用于Peer-Peer对等通信、RPC、长连接网关、微服务、推送服务，游戏服务等领域。

## 框架名词
Peer： 通信端点，可以是服务端或客户端
Socket： 对net.Conn的封装，增加自定义包协议、传输管道等功能
Message：数据包内容元素对应的结构体
Proto： 数据包封包／解包的协议接口
Codec： 用于Body的序列化工具
XferPipe： 数据包字节流的编码处理管道，如压缩、加密、校验等
XferFilter： 一个在数据包传输前，对数据进行加工的接口
Plugin： 贯穿于通信各个环节的插件
Session： 基于Socket封装的连接会话，提供的推、拉、回复、关闭等会话操作
Context： 连接会话中一次通信（如PULL-REPLY, PUSH）的上下文对象
Call-Launch： 从对端Peer拉数据
Call-Handle： 处理和回复对端Peer的拉请求
Push-Launch： 将数据推送到对端Peer
Push-Handle： 处理同伴的推送
Router： 通过请求信息（如URI）索引响应函数（Handler）的路由器
