# RabbitMQ Note
<br>

### RabbitMQ cluster
```
# scale down mq hosts

# 在mq pod 中, 先停用 host 後再將他移除
rabbitmqctl -n rabbit@xxx-rabbitmq-2.xxx-rabbitmq-headless.{namespace}.svc.cluster.local stop_app

rabbitmqctl forget_cluster_node rabbit@xxx-rabbitmq-2.xxx-rabbitmq-headless.{namespace}.svc.cluster.local

# 移除 pvc
kubectl delete pvc XXXX

```

