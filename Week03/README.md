# 本周作业

基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

```bash
# 1. http的启动和关闭，30s超时自动关闭服务

# 2. kill后关闭服务
$ go run .
signal: killed
```