# container

## batcher
使用场景：写入ES等, 需要通过批量写入优化的场景
```golang
import (
   "github.com/joshqu1985/lego/container"
)

func main() {
   // 500条 或者 1s 触发批量发送
   batcher := container.NewBatcher(container.WithBulkSize(500), container.WithInterval(time.Second))
   go BatchSend(batcher.Queue())

   for i := 0; i < 1000; i++ {
      batcher.Put("hello")
   }
}

func BatchSend(queue chan []any) {
   for {
      vals := <-queue
      b, _ := json.Marshal(vals)
      fmt.Println("len: ", len(vals), " data: ", string(b))
   }
}
```

## bloom filter
使用场景：减少数据库查询
```golang
import (
   "github.com/joshqu1985/lego/container"
)

func main() {
   // 设置bit大小 10k
   bloom := container.NewBloomFilter(1024 * 10)
   
   // 添加
   bloom.Add([]byte("hello"))
   
   // 判断是否存在
   ok := bloom.Exist([]byte("world"))
}
```
