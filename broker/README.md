# broker
**消息队列** 对消息队列SDK做了封装, 提供统一接口, 方便快速使用和切换

## 使用方式
```golang
import (
   "github.com/joshqu1985/lego/broker"
)

func main() {
   // 生产者
   producer := broker.NewProducer(conf)

   _ = producer.Send(context.Background(), "topic1", &broker.Message{
      Payload: []byte("hello kitty 1"),
   })

   // 消费者
   consumer := broker.NewConsumer(conf)

   consumer.Register("topic1", func(ctx context.Context, msg *broker.Message) error {
      fmt.Println(time.Now().Unix(), string(msg.Payload))
      return nil
   })
   go consumer.Start()
}
```

## 内存
基于链表和优先级队列实现。支持延迟消息，优先级队列实现。
<br>每个topic有10000条消息限制，当topic队列数据量大于10000条时，调用Send会失败
```golang
  // 配置信息
  conf := broker.Config{
    Source:  "memory",
    GroupId: "g1",
    Topics:  map[string]string { "topic1": "real_topic1" },
  }
```

## redis
基于redis stream实现，需要redis版本>=5.0。支持延迟消息，使用zset和hash实现
<br>每个topic有10000条消息限制，当topic队列数据量大于10000条时，调用Send会失败
```golang
  // 配置信息
  conf := broker.Config{
    Source:    "redis",
    Endpoints: []string{"127.0.0.1:6379"},
    GroupId:   "g1",
    Topics:    map[string]string { "topic1": "real_topic1" },
  }
```

## kafka
基于 github.com/segmentio/kafka-go 封装，不支持延迟消息。
```golang
  // 配置信息
  conf := broker.Config{
    Source:    "kafka",
    Endpoints: []string{"127.0.0.1:9092"},
    GroupId:   "g1",
    Topics:    map[string]string { "topic1": "real_topic1" },
  }
```

## rocketmq
基于 github.com/apache/rocketmq-clients/golang/v5 封装
```golang
  // 配置信息
  conf := broker.Config{
    Source:    "rocketmq",
    Endpoints: []string{"127.0.0.1:8081"},
    GroupId:   "g1",
    Topics:    map[string]string { "topic1": "real_topic1" },
  }
```

## pulsar
基于 github.com/apache/pulsar-client-go/pulsar 封装
```golang
  // 配置信息
  conf := broker.Config{
    Source:    "pulsar", 
    Endpoints: []string{"pulsar://localhost:6650"},
    GroupId:   "g1",
    Topics:    map[string]string{"test": "persistent://public/default/test"},
  }
```
