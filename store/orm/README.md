# 数据库ORM模块

本模块是一个基于GORM封装的数据库操作模块，提供了统一的接口来操作MySQL和PostgreSQL数据库，并集成了性能监控功能。

## 核心特性

- **统一接口**：提供统一的API操作不同类型的数据库，避免业务代码直接依赖GORM
- **多数据库支持**：支持MySQL和PostgreSQL两种主流数据库
- **连接池管理**：内置连接池配置，优化数据库连接资源利用
- **性能监控**：集成metrics模块，自动记录查询耗时和慢查询统计
- **事务支持**：提供事务操作接口，确保数据一致性
- **灵活配置**：支持丰富的配置选项，适应不同的部署环境

## 接口定义

### 主要结构体

#### Store

`Store`是ORM模块的核心结构体，封装了数据库连接和操作方法：

```go
type Store struct {
    db     *DB
    source string
    durationHandle  metrics.HistogramVec  // 查询耗时监控
    slowCountHandle metrics.CounterVec    // 慢查询计数监控
}
```

#### Config

`Config`结构体用于配置数据库连接参数：

```go
type Config struct {
    Source       string `toml:"source" yaml:"source" json:"source"`                         // 数据库类型: mysql 或 postgre
    Host         string `toml:"host" yaml:"host" json:"host"`                               // 数据库地址
    Port         int    `toml:"port" yaml:"port" json:"port"`                               // 数据库端口
    User         string `toml:"user" yaml:"user" json:"user"`                               // 用户名
    Pass         string `toml:"pass" yaml:"pass" json:"pass"`                               // 密码
    Opts         string `toml:"opts" yaml:"opts" json:"opts"`                               // 额外连接参数
    Database     string `toml:"database" yaml:"database" json:"database"`                   // 数据库名
    MaxOpenConns int    `toml:"max_open_conns" yaml:"max_open_conns" json:"max_open_conns"` // 最大打开连接数
    MaxIdleConns int    `toml:"max_idle_conns" yaml:"max_idle_conns" json:"max_idle_conns"` // 最大空闲连接数
    MaxLife      int    `toml:"max_life" yaml:"max_life" json:"max_life"`                   // 连接最大存活时间(毫秒)
}
```

### 主要方法

#### New

初始化ORM连接池：

```go
func New(conf Config) (*Store, error)
```

#### Exec

执行数据库操作并记录性能指标：

```go
func (this *Store) Exec(fn func(*DB) error, labels ...string) error
```

#### Transaction

执行事务操作并记录性能指标：

```go
func (this *Store) Transaction(fn func(*DB) error, labels ...string) error
```

#### 事务控制

```go
func (this *Store) Begin() *DB
func (this *Store) Rollback() *DB
func (this *Store) Commit() *DB
```

## 使用示例

### 初始化数据库连接

```go
import (
    "github.com/joshqu1985/lego/store/orm"
)

// 初始化MySQL连接
mysqlConfig := orm.Config{
    Source:       "mysql",
    Host:         "localhost",
    Port:         3306,
    User:         "root",
    Pass:         "password",
    Database:     "test_db",
    Opts:         "charset=utf8mb4&parseTime=True",
    MaxOpenConns: 100,
    MaxIdleConns: 10,
    MaxLife:      3600000,
}

mysqlStore, err := orm.New(mysqlConfig)
if err != nil {
    panic(err)
}

// 初始化PostgreSQL连接
pgConfig := orm.Config{
    Source:       "postgre",
    Host:         "localhost",
    Port:         5432,
    User:         "postgres",
    Pass:         "password",
    Database:     "test_db",
    Opts:         "sslmode=disable TimeZone=Asia/Shanghai",
    MaxOpenConns: 100,
    MaxIdleConns: 10,
    MaxLife:      3600000,
}

pgStore, err := orm.New(pgConfig)
if err != nil {
    panic(err)
}
```

### 基本数据库操作

```go
// 定义模型
type User struct {
    ID    uint
    Name  string
    Email string
}

// 查询操作
var user User
err := mysqlStore.Exec(func(db *orm.DB) error {
    return db.First(&user, 1).Error
}, "find_user")

// 插入操作
newUser := User{Name: "张三", Email: "zhangsan@example.com"}
err = mysqlStore.Exec(func(db *orm.DB) error {
    return db.Create(&newUser).Error
}, "create_user")

// 更新操作
err = mysqlStore.Exec(func(db *orm.DB) error {
    return db.Model(&User{}).Where("id = ?", 1).Update("name", "李四").Error
}, "update_user")

// 删除操作
err = mysqlStore.Exec(func(db *orm.DB) error {
    return db.Delete(&User{}, 1).Error
}, "delete_user")
```

### 事务操作

```go
// 执行事务
err = mysqlStore.Transaction(func(db *orm.DB) error {
    // 插入用户
    user := User{Name: "张三", Email: "zhangsan@example.com"}
    if err := db.Create(&user).Error; err != nil {
        return err // 自动回滚
    }
    
    // 可以在事务中执行多个操作
    // ...
    
    return nil // 自动提交
}, "user_transaction")
```

### 使用条件表达式

```go
import (
    "github.com/joshqu1985/lego/store/orm"
    "time"
)

// 使用Expr
err = mysqlStore.Exec(func(db *orm.DB) error {
    return db.Model(&User{}).Update("last_login", orm.Expr("?", time.Now())).Error
}, "update_login_time")
```

### 检查记录是否存在

```go
var user User
err = mysqlStore.Exec(func(db *orm.DB) error {
    return db.First(&user, 1).Error
}, "find_user")

if err == orm.ErrRecordNotFound {
    // 记录不存在
} else if err != nil {
    // 发生其他错误
} else {
    // 记录存在，user包含查询结果
}
```

## 性能监控

模块自动集成了性能监控功能，主要记录以下指标：

1. **查询耗时**：`{namespace}_exec_duration` 直方图，记录所有数据库操作的耗时分布
   - MySQL: `mysql_exec_duration`
   - PostgreSQL: `postgre_exec_duration`

2. **慢查询计数**：`{namespace}_slow_count` 计数器，记录执行时间超过500ms的慢查询数量
   - MySQL: `mysql_slow_count`
   - PostgreSQL: `postgre_slow_total`

监控指标包含`command`标签，可用于区分不同类型的操作。

## 最佳实践

1. **使用Exec方法进行所有数据库操作**：确保性能指标被正确记录
2. **为每个操作提供有意义的labels**：方便在监控中区分不同类型的操作
3. **适当配置连接池参数**：根据应用负载调整MaxOpenConns、MaxIdleConns和MaxLife
4. **使用事务保证数据一致性**：对于涉及多表操作的业务逻辑，使用Transaction方法
5. **错误处理**：正确处理数据库操作可能返回的错误，特别是ErrRecordNotFound
6. **避免直接使用GORM类型**：使用模块提供的类型别名，如DB、Expr等

## 注意事项

1. **数据库配置安全**：避免在代码中硬编码数据库密码等敏感信息
2. **性能优化**：对于频繁执行的查询，考虑添加适当的索引
3. **连接池管理**：合理设置连接池参数，避免连接泄漏
4. **事务使用**：事务应该尽可能简短，减少锁定时间
5. **慢查询监控**：定期检查慢查询日志，优化性能问题

## 依赖说明

- **gorm.io/gorm**: ORM核心库
- **gorm.io/driver/mysql**: MySQL驱动
- **gorm.io/driver/postgres**: PostgreSQL驱动
- **github.com/joshqu1985/lego/metrics**: 性能监控模块
