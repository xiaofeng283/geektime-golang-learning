
# 本周作业

按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

项目目录：
```
|____cmd                
| |____server              # 编译入口
| | |____wire_gen.go
| | |____wire.go
| | |____main.go
|____test                  # 测试服务API
| |____user.go
|____bin
| |____server              # 编译好的二进制文件
|____Makefile              # 项目编译、运行、测试快捷命令
|____internal
| |____biz                 # 业务组装层
| | |____user.go
| |____service             # API实现层
| | |____user.go
| |____data                # 数据层
| | |____user.go
| |____pkg                 # 公共库
| | |____grpc              # gRPE Server封装
| | | |____service.go
|____README.md
|____api                   # rpc文件
| |____user
| | |____v1
| | | |____user_grpc.pb.go
| | | |____user.pb.go
| | | |____user.proto
```
