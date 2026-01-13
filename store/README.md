# 存储模块集合

本目录包含lego框架中各种存储相关的组件，每个子目录都是一个独立的存储模块，提供了统一的接口和易用的API。

## 模块列表

| 模块名称 | 描述 | 文件路径 |
|---------|------|----------|
| coss | 云对象存储服务，支持阿里云OSS、腾讯云COS、华为云OBS、亚马逊S3等 | [store/coss](coss) |
| mongo | MongoDB存储模块，基于MongoDB驱动封装，提供统一的接口和性能监控 | [store/mongo](mongo) |
| orm | 数据库ORM模块，基于GORM封装，支持MySQL和PostgreSQL | [store/orm](orm) |
| redis | Redis存储模块，基于go-redis/v9封装，支持单机和集群模式 | [store/redis](redis) |

## 统一特性

所有存储模块都具有以下共同特性：

- 统一的接口设计，便于集成和使用
- 丰富的配置选项，适应不同的部署环境
- 内置连接池管理，优化资源利用
- 与metrics模块集成，提供性能监控
- 错误处理和重试机制，提高稳定性

## 使用说明

每个存储模块都有自己的详细README文档，请参考相应子目录下的README文件获取具体使用方法。

```go
// 示例：导入并使用存储模块
import (
    "github.com/joshqu1985/lego/store/redis"
    "github.com/joshqu1985/lego/store/orm"
    "github.com/joshqu1985/lego/store/mongo"
    "github.com/joshqu1985/lego/store/coss"
)
```

## 注意事项

- 所有存储模块的配置都应该妥善保管，特别是密码等敏感信息
- 根据具体业务需求选择合适的存储模块
- 使用时请注意处理可能出现的网络异常和超时情况
- 对于高并发场景，建议合理配置连接池参数