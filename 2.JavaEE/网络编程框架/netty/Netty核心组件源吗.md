# NioEventLoopGroup源码剖析

- 创建一定数量的NioEventLoop线程组并初始化。 
  默认线程数是当前系统的CPU的2倍
- 创建线程选择器chooser。当获取线程时，通过选择器来获取。
- 创建线程工厂并构建线程执行器。

# NioEventLoop源码剖析

- 开启Selector并初始化。
- 把ServerSocketChannel注册到Selector上。
- 处理各种I/O事件，如OP_ACCEPT、OP_CONNECT、OP_READ、OP_WRITE事件。
- 执行定时调度任务。
- 解决JDK空轮询bug。

