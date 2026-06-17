# 1. 执行部署
## 1.1 创建配置、密钥
kubectl apply -f k8s/configmap.yaml
kubectl apply -f k8s/secret.yaml

## 1.2 构建镜像
docker build -t order-service:v1 .

## 1.3 导入kind集群
kind load docker-image order-service:v1 --name demo

## 1.4 发布deployment与service
kubectl apply -f k8s/order.yaml

## 1.5 滚动重启加载新配置
kubectl rollout restart deployment order-service

## 1.6 查看日志确认mysql连接成功
kubectl logs -f deployment/order-service

# 2. 测试接口
## 2.1 前置 sql
kubectl exec -it \$(kubectl get pod -l app=mysql -o name) -- mysql -uroot -proot test
```
CREATE TABLE users (
    id INT PRIMARY KEY,
    name VARCHAR(32) NOT NULL
);
INSERT INTO users(id,name) VALUES(1,"Tom");
INSERT INTO users(id,name) VALUES(2,"Lili");

CREATE TABLE orders (
id INT AUTO_INCREMENT PRIMARY KEY COMMENT '订单ID',
user_id INT NOT NULL COMMENT '用户ID',
create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
);
```

## 2.2 进入 order pod
kubectl exec -it \$(kubectl get pods -l app=order | grep Running | head -1 | awk '{print $1}') -- sh

## 2.3 创建订单
wget -qO- http://127.0.0.1:8080/order/create?uid=1

# 2.3 查询订单
wget -qO- http://127.0.0.1:8080/order/get?id=1