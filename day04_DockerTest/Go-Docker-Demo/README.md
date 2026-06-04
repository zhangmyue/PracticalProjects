# Go-Docker-Demo 项目说明
## 项目简介
这是一个 Go 极简 Web 服务 + Docker 容器化 实战项目。
项目采用 Go 原生 net/http 搭建 Web 服务，使用 Docker 多阶段构建 打包镜像，极致精简、符合生产级规范，适合学习 Go 后端、Docker 容器化、云原生基础部署。
## 项目结构
```
├── Dockerfile      # 多阶段构建Docker镜像配置
├── go.mod          # Go模块依赖管理
├── main.go         # 主程序入口
└── README.md       # 项目说明文档
```
## 技术栈
- 后端：Golang 原生 Web
- 容器化：Docker（多阶段构建、Alpine 极简镜像）
- 部署方式：本地 Docker 手动部署
  Dockerfile 核心特性
- 多阶段构建：编译环境与运行环境分离，大幅缩小镜像体积
- Alpine 轻量镜像：无多余依赖，安全、启动快
- 分层缓存优化：优先加载依赖，提升重复构建速度
- 规范端口声明：EXPOSE 声明容器服务端口
## 快速运行
1. ### 构建镜像
```
# 在项目根目录执行
docker build -t go-docker-app .
```
2. ### 启动容器
```
# 端口映射：主机8080 -> 容器8080
docker run -d -p 8080:8080 --name my-go-app go-docker-app
```
3. ### 访问服务
   浏览器 / curl 访问： http://localhost:8080

   成功返回：✅ Go 程序在 Docker 中运行成功！
## 常用运维命令
```
# 查看运行容器
docker ps

# 查看容器日志
docker logs my-go-app

# 停止容器
docker stop my-go-app

# 删除容器
docker rm my-go-app

# 删除镜像
docker rmi go-docker-app
```
核心知识点总结
- WORKDIR：指定容器/镜像内部工作目录
- COPY：左侧为本地主机文件，右侧为镜像内部路径
- EXPOSE：仅声明容器内部端口，不生效端口映射
- -p 主机端口:容器端口：实现真实端口映射，外部可访问
- 多阶段构建：解决 Go 镜像体积过大问题，生产最优方案
  学习拓展方向
- 搭配 Docker Compose 实现 Go + MySQL + Redis 集群部署
- 适配 K8s 实现容器编排、自动扩缩容
- 加入单元测试、性能 Benchmark、CI/CD 自动构建
  许可证
  仅供个人学习、练手使用。