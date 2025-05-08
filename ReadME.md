todo:
+ error handling: fmt.Errof("%w")  pkg/errors库 包含堆栈信息

## Consul KV

StayEaseGo/user_srv
```json
{
  "name": "user_srv",
  "mysql": {
    "host": "127.0.0.1",
    "port": 3306,
    "user": "root",
    "password": "root",
    "db": "StayEaseGo",
    "salt": "123"
  },
  "redis": {
    "host": "127.0.0.1",
    "port": 6379
  }
}
```

StayEaseGo/user_web

```json
{
  "name":"user_web",
  "otel":{
    "endpoint":"localhost:4317"
  },
  "user_srv":{
    "name":"user_srv"
  },
  "jwt":{
    "key":"test"
  }
}
```

StayEaseGo/homestay_srv
```json
{
  "name":"homestay_srv",
  "otel":{
    "endpoint":"localhost:4317"
  },
  "mysql": {
    "host": "127.0.0.1",
    "port": 3306,
    "user": "root",
    "password": "root",
    "db": "StayEaseGo",
    "salt": "123"
  },
  "es":{
    "host":"127.0.0.1",
    "port":9201
  }
}
```

StayEaseGo/homestay_web
```json
{
  "name":"homestay_web",
  "otel":{
    "endpoint":"localhost:4317"
  },
  "user_srv":{
    "name":"user_srv"
  },
  "homestay_srv":{
    "name":"homestay_srv"
  }
}
```

StayEaseGo/order_srv
```json
{
  "name":"order_srv",
  "otel":{
    "endpoint":"localhost:4317"
  },
  "mysql": {
    "host": "127.0.0.1",
    "port": 3306,
    "user": "root",
    "password": "root",
    "db": "StayEaseGo"
  },
  "homestay_srv":{
    "name":"homestay_srv"
  },
  "kafka":{
    "brokers":["localhost:9092"]
  },
  "redis": {
    "host": "127.0.0.1",
    "port": 6379
  }
}
```

StayEaseGo/order_web
```json
{
  "name":"order_web",
  "otel":{
    "endpoint":"localhost:4317"
  },
  "order_srv":{
    "name":"order_srv"
  }
}
```

StayEaseGo/payment_srv
```json
{
  "name":"payment_srv",
  "otel":{
    "endpoint":"localhost:4317"
  },
  "mysql": {
    "host": "127.0.0.1",
    "port": 3306,
    "user": "root",
    "password": "root",
    "db": "StayEaseGo"
  },
  "kafka":{
    "brokers":["localhost:9092"]
  }
}
```

StayEaseGo/payment_web
```json
{
  "name":"order_web",
  "otel":{
    "endpoint":"localhost:4317"
  },

  "jwt":{
    "key":"test"
  },
  "wx_mini":{
    "AppId":"wxe1ea796246f1df67",
    "Secret":"7f6f688ca75e7794728a5a2acaf3361b"
  },
  "wx_pay":{
    "MchId":"1305638280",
    "SerialNon":"013467007045764",
    "APIv3Key":"013467007045764",
    "PrivateKey":"",
    "NotifyUrl":"http://localhost:8004/api/v1/pay/wxpay/callback"
  },
  "order_srv":{
    "name":"order_srv"
  },
  "payment_srv":{
    "name":"payment_srv"
  },
  "user_srv":{
    "name":"user_srv"
  }

}
```


StayEaseGo/mq
```json
{
  "name":"mq",
  "kafka":{
    "brokers":["localhost:9092"]
  },
    "order_srv":{
    "name":"order_srv"
  },
    "user_srv":{
    "name":"user_srv"
  }
}
```

StayEaseGo/asynq
```json
{
  "name":"asynq",
  "order_srv":{
    "name":"order_srv"
  },
  "redis": {
    "host": "127.0.0.1",
    "port": 6379
  }
}
```