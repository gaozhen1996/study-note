# 一、 TCP粘包，拆包

## 1.粘包的主要原因

- 发送方每次写入数据 < 套接字缓冲区大小
- 接收方读取套接字缓冲区数据不够及时

## 2.拆包的主要原因

- 发送方写入数据 > 套接字缓冲区大小
- 发送的数据大于协议的最大传输单元，必须拆包

## 3.根本原因

- TCP 是流式协议，消息无边界。

## 4.Netty解决手段

- 固定长度
  - 解码：FixedLengthFrameDecoder
- 分割符
  - 解码：DelimiterBasedFrameDecoder
- 固定长度字段存个内容的长度信息 
  - 解码：LengthFieldBasedFrameDecoder

