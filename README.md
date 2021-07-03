# gewu_jxc_server

格物进销存服务端

## 常用命令

给 script.sh 添加权限

```bash
sudo chmod -R 777 script.sh
```

编译 linux

```
./script.sh linux
```

本地执行

```
./script.sh dev
```

单元测试

```
./script.sh test
```

单元测试并且显示测试覆盖率

```
./script.sh test
```

## 目录说明

```text
app // 程序源码文件夹
cmd // 项目执行文件
logs // 日志文件
models // 数据库表结构文件，自动生成
migrations // 数据库迁移代码
.air.toml // air 开发环境热更新配置
.env  // 环境参数配置（gitignore忽略）
database.json // sql-migrate 基本配置
dbconfig.yml // sql-migrate cli 配置（gitignore忽略）
go.mod
go.sum
README.md
sqlc.yaml // sqlc/sqlm 配置
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
migrations="migrations"
redisAddr="127.0.0.1:6379"
redisPassword="123456"
redisDB=0
```

## controllers

利用 fiber 进行路由管理和接口参数验证

## services

业务代码

## 启动服务/迁移数据库

migrate 会读取 `migrations` 进行迁移

启动并执行迁移，999 次

```bash
upMigrate=999 go run main.go
```

启动并执行回滚，1 次

```bash
downMigrate=1 go run main.go
```

执行完 Migrate 后自动退出

```bash
upMigrate=999 onlyMigrate=1 go run main.go
```

开发环境热更新启动

```bash
upMigrate=999 air
```

## sqlm

使用 sqlm 编译 表结构模型, sqlm 的安装和使用查看：https://github.com/ymzuiku/sqlc/tree/main-ex, sqlm 会读取 migrations 的历史生成 model，并且会忽略 migrate-down 的代码块

执行命令：

```bash
sqlm generate
```

## 关于测试

理论上，gewu_jxc 所有用户行为都有 100%的客户端自动化集成测试，从用户最终行为模拟上覆盖了所有业务。

每次提交之后，只需要确保当前的业务最终通过可前端的所有自动化测试即可。

所以后端仅需要编写部分组件的单元测试

### 单元测试

单元测试，安装 golang 的建议，直接写在业务代码中，具体参考 `kit/randomCode_test.go`

- -count=1 忽略测试缓存，有副作用的测试，如查询数据库，需要忽略测试缓存
- -cover 查看测试覆盖率
- -tags test 忽略文件首行包含 `//+build !test` 的文件进行覆盖测试,记得文件第二行需要换行
- -coverprofile=t.out 导出测试描述

测试并计算覆盖率

```bash
go test ./app/... -count=1 -cover -tags test -coverprofile=t.out
```

显示具体测试覆盖率 html

```
go tool cover -html=t.out
```

## 配置 dbconfig.yml(可选)

若需使用 cli 执行 sql-migrate,可以配置 dbconfig.yml, gitignore 已忽略 dbconfig.yml

以下是一个 dbconfig.yml 配置的例子

```yml
development:
  dialect: postgres
  datasource: host=localhost port=5432 user=postgres password=123456 dbname=dev_dog sslmode=disable TimeZone=Asia/Shanghai
  dir: migrations

production:
  dialect: postgres
  datasource: host=localhost port=5432 user=postgres password=123456 dbname=dev_fish sslmode=disable TimeZone=Asia/Shanghai
  dir: migrations
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

迁移所有

```bash
sql-migrate up
```

回滚所有

```bash
sql-migrate down -limit=999
```
