# gewu_jxc_server

格物进销存服务端

## 目录说明

```js
-- app  // 程序源码文件夹
  |- controllers  // 路由
      |- user //每个业务一个文件
          |- userControllers  // 该业务的路由，做参数解析、约束，调用 servers
          |- userServices  // 该业务的业务逻辑，调用 DAO 层，业务逻辑处理
          |- user.go
  |- tools  // 所有相关工具
  |- app.go  // 程序具体实例入口
-- logs  // 日志文件
-- sql  // sql相关
  |- db  // DAO层，自动生成代码，无修改
  |- migrations  // 数据库迁移代码
  |- query  // 数据库操作方法，每个表一个文件
main.go  // 仅仅执行 app.go
.air.toml  // air 开发环境热更新配置
.env  // 环境参数配置（gitignore忽略）
database.json  // sql-migrate 基本配置
dbconfig.yml  // sql-migrate cli 配置（gitignore忽略）
go.mod
go.sum
README.md
sqlc.yaml  // sqlc 配置
```


## 配置 .env (必须)

启动的前提是配置好 .env 文件，.env 请确保仅在每台设备环境中单独配置，切勿同步至 git

以下是一个 .env 配置的例子

```bash
DEV=no
DB_CONNECT_URL="host=localhost port=5432 user=postgres password=123456 dbname=dev_dog sslmode=disable TimeZone=Asia/Shanghai"
maxOpenConns=20
maxIdleConns=20
maxLifetime=5
redisAddr="127.0.0.1:6379"
redisPassword="123456"
redisDB=0
```
## controllers

利用 fiber 进行路由管理和接口参数验证

## services

业务代码

## 启动服务/迁移数据库

migrate 会读取 `sql/migrations` 进行迁移

启动并执行迁移，999次

```bash
upMigrate=999 go run main.go
```

启动并执行回滚，1次

```bash
downMigrate=1 go run main.go
```

执行完Migrate后自动退出

```bash
upMigrate=999 onlyMigrate=1 go run main.go
```

开发环境热更新启动

```bash
upMigrate=999 air
```

## sqlc 

使用 sqlc 代替 orm 和 DAO 层，代码更简洁，并且利用 migrate 的文件内容生成 model，减少了重复工作
使用 sqlc 的另一个好处是相对于 gorm 它的性能更高（节省了orm反射的开销），其次是每一个 sql 都是手工编写的，更容易调优。

执行命令：

```bash
sqlc generate
```

1. sqlc 会读取 sql/migrations 的历史生成 model，并且会忽略 migrate-down 的代码块
2. sqlc 会读取 sql/query 生成 数据操作（DAO）和 Models


## 关于测试

理论上，gewu_jxc 所有用户行为都有100%的客户端自动化集成测试，从用户最终行为模拟上覆盖了所有业务。

每次提交之后，只需要确保当前的业务最终通过可前端的所有自动化测试即可。

所以后端仅需要编写部分组件的单元测试


## 配置dbconfig.yml(可选)

若需使用 cli 执行 sql-migrate,可以配置 dbconfig.yml, gitignore 已忽略 dbconfig.yml

以下是一个 dbconfig.yml 配置的例子

```yml
development:
    dialect: postgres
    datasource: host=localhost port=5432 user=postgres password=123456 dbname=dev_dog sslmode=disable TimeZone=Asia/Shanghai
    dir: sql/migrations

production:
    dialect: postgres
    datasource: host=localhost port=5432 user=postgres password=123456 dbname=dev_fish sslmode=disable TimeZone=Asia/Shanghai
    dir: sql/migrations
```

### 配置完可以执行 sql-migrate cli 命令：

创建一个迁移文件

```bash
sql-migrate new file-name
```

跳过当前所有迁移文件

```bash
sql-migrate skip
```