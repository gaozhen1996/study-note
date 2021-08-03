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



# 二、ByteBuf 内存泄漏

## 1.现象

```java
    public void channelRead(ChannelHandlerContext ctx, Object msg) {
        ByteBuf reqMsg = (ByteBuf)msg;
        byte [] body = new byte[reqMsg.readableBytes()];
//        ReferenceCountUtil.release(reqMsg);
        executorService.execute(()-> {
            //解析请求消息，做路由转发，代码省略...
            //转发成功，返回响应给客户端
            ByteBuf respMsg = allocator.heapBuffer(2);
            respMsg.writeBytes(new byte[]{0,1,1,1});//作为示例，简化处理，将请求返回
            ctx.writeAndFlush(respMsg);
        });
    }
```

报错：

> ERROR pool-1-thread-1 (io.netty.util.ResourceLeakDetector.reportTracedLeak:317) - LEAK: ByteBuf.release() was not called before it's garbage-collected.

## 2.原因
**调用堆栈：**

- NioEventLoop#processSelectedKey()

  - unsafe.read()

  - NioByteUnsafe#read()

    - byteBuf = allocHandle.allocate(allocator)

    - MaxMessageHandle#allocate

      - alloc.ioBuffer(this.guess())

      - AbstractByteBufAllocator#ioBuffer(int)

      - ...

        - PooledByteBufAllocator#newDirectBuffer



 **说明**   

msg 对象从堆外内存申请后，没有释放。导致内存泄漏

## 3.理解误区

有一种说法认为 Netty框架分配的 ByteBuf框架会自动释放，业务不需要释放；业务创建的ByteBuf则需要自己释放，Netty框架不会释放。

通过前面的案例分析和验证，我们可以看出这个观点是错误的。为了在实际项目中更好地管理ByteBuf，下面我们分4种场景进行说明。

**1.基于内存池的请求ByteBuf**

业务ChannelInboundHandler继承自SimpleChannelInboundHandler，实现它的抽象方法channelRead0（ChannelHandlerContext ctx，I msg），ByteBuf的释放业务不用关心，由SimpleChannelInboundHandler负责释放，相关代码如下（SimpleChannelInboundHandler）

```java
public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception {
    boolean release = true;

    try {
        if (this.acceptInboundMessage(msg)) {
            this.channelRead0(ctx, msg);
        } else {
            release = false;
            ctx.fireChannelRead(msg);
        }
    } finally {
        if (this.autoRelease && release) {
            ReferenceCountUtil.release(msg);
        }

    }

}
```

**2.基于非内存池的请求ByteBuf**

如果业务使用非内存池模式覆盖 Netty 默认的内存池模式创建请求 ByteBuf，例如通过如下代码修改内存申请策略为Unpooled。同样也是需要释放内存的

```java
.childHandler(new ChannelInitializer<SocketChannel>() {
    @Override
    public void initChannel(SocketChannel ch) throws Exception {
        ch.config().setAllocator(UnpooledByteBufAllocator.DEFAULT);
    }
});
```

**3.基于内存池的响应ByteBuf**

只要调用了writeAndFlush或者flush方法，在消息发送完成后都会由Netty框架进行内存释放，业务不需要主动释放内存。

**4.基于非内存池的响应ByteBuf**

无论是基于内存池还是非内存池分配的 ByteBuf，如果是堆内存，则将堆内存转换成堆外内存，然后释放 HeapByteBuffer，待消息发送完成，再释放转换后的DirectByteBuf；如果是 DirectByteBuffer，则不需要转换，待消息发送完成之后释放。因此对于需要发送的响应ByteBuf，由业务创建，但是不需要由业务来释放。
