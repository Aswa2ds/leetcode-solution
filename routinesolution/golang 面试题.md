# golang 面试题

1. new & make 区别？

   make：slice map channel

   new：struct 指针

2. 传参的时候，什么情况传指针、什么情况传值？

   传指针在go 中会可能发生逃逸，原本在栈上的内存逃逸到堆上，导致内存浪费

3. cpu 缓存加速（经典的矩阵横纵读取）、以及**多核cpu 缓存原理**

4. 无符号数 运算溢出补码

5. 有三个函数，分别可以打印"cat","dog","fish"，要求每个函数都起一个GoRoutine，按照"cat","dog","fish"的顺序打印到屏幕上

   锁、chnnel 两种实现

   讲一下channel 有缓存和无缓存（声明的时候）

   ```go
   ch := make(chan struct{}) \\ 无缓存，首先要有一个消费者，即向channel 中发送的数据会直接流出，不会等消费者来接收
   ch := make(chan struct{}, 1) \\有缓存，类似消息队列，生产者的数据会缓存在channel 中，等待消费者来接收
   ```

   chnnel 底层实现、是否线程安全、mutex 使用悲观锁的原因

6. mutex 锁机制

   乐观锁、悲观锁

   mutex 有几种模式？两种：互斥、饥饿

   mutex 是否可以作为自旋锁

   RWmutex（读写锁互斥），业务场景读多写少，读操作之间不互斥，提高并发效率

   其他共享内存 线程安全的方式

7. GoRoutine 

   协程的内存消耗比较小，相对与线程

   什么时候会发生阻塞，阻塞的时候调度器的行为

   PMG 模型（就是调度模型？）

   信号协作

   调度模型

   自旋 系统抢占

   线程 OOM、GoRoutine OOM 的情况

   **stop the world & 栈扫描** gc 的是一个安全点 

8. 错误处理

   error if/else

   eror gRPC 透传

   某一个goroute 发生panic 会发生什么，有什么影响

9. gRPC

   proto 管理：monorepo

   服务发现：注册中心

   负载均衡

10. linux 64种信号 了解

11. linux 中的线程模型

12. gc

    三色标记法 

    gc 触发有哪几种方式？手动触发、内存触发

13. go 的内存管理机制

    阈值触发、时间间隔触发

14. go 的内存分配

    堆栈 栈区逃逸

    channel 是分配在堆还是栈

    哪些变量会分配在堆区和栈区

    大对象、小对象：大小临界值

15. 开发中的坑

16. cpu 流水线 如何让A、B 指令在cpu中顺序执行？ 内存屏障

17. gorm 

18. [effectiv go](https://learnku.com/docs/effective-go/2020)