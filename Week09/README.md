
# 第九周作业

1. 总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用 

什么是socket 粘包现象？
- 发送方已经发送了消息给接收方，但是因为网络等原因没有接收到消息，或者不能进行回应。那发送方过一段时间会再次发送相同的消息，就会导致发送多次相同的消息。
- 但是，接收方最终接收到的消息可能会和发送的消息内容不一致。

socket 粘包的主要原因：
   - 发送方每次写入数据 < 套接字（Socket）缓冲区大小；
   - 接收方读取套接字（Socket）缓冲区数据不够及时。

解决方案：
   - fix length：固定缓冲区大小，只需要控制服务器端和客户端发送和接收字节的（数组）长度相同即可。
   - delimiter based：发送方封包时以特殊字符结尾，接收方就知道流的边界在哪里，实现按行读取。
   - length field based frame decoder：通过封装请求协议的方式解决粘包问题。将请求的数据封装为两部分：数据头+数据正文，在数据头中存储数据正文的大小，当读取的数据小于数据头中的大小时，继续读取数据，直到读取的数据长度等于数据头中的长度时才停止。


2. 实现一个从 socket connection 中解码出 goim 协议的解码器。

