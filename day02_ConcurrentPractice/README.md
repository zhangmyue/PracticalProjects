# Go 并发编程实战：学生成绩提交系统

## 项目介绍
使用 Goroutine + WaitGroup + Channel + Mutex 实现并发成绩提交系统

## 实现版本
1. **Mutex 版本**：共享内存 + 互斥锁保护切片
2. **Channel 版本**：Go 风格无锁设计，通过通道传递数据

## 核心知识点
- 并发控制与限流（令牌桶）
- sync.WaitGroup 等待所有协程完成
- sync.Mutex 共享内存安全
- Channel 通道通信与数据传递
- 高并发、无锁、安全实践