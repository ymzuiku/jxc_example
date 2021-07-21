# gewu_jxc_server

格物进销存服务端

## 参考

目录结构: https://github.com/golang-standards/project-layout

代码规范: https://github.com/xxjwxc/uber_go_guide_cn

## 常用命令

编译 linux

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o jxc cmd/jxc/main.go
```

本地执行

```bash
up_igrate=all go run cmd/jxc/main.go
```

本地 watch 模式（需要安装 [air](https://github.com/cosmtrek/air)）

```bash
up_igrate=all air
```

单元测试

```bash
go test ./... | { grep -v 'no test files'; true; }
```

单元测试并且显示测试覆盖率

```
go test ./... -count=1 -cover -tags test -coverprofile=t.out
go tool cover -html=t.out
```

## 迁移数据库

migrate 会读取 `migrations` 进行迁移

启动并执行所有迁移

```bash
up_migrate=all go run cmd/jxc/main.go
```

执行回滚，1 次, down_migrate 结束会自动退出

```bash
down_migrate=1 go run cmd/jxc/main.go
```

执行跳过，2 次, skip_migrate 结束会自动退出

```bash
skip_migrate=2 go run cmd/jxc/main.go
```

使用其他 migrations 路径

```bash
up_migrate=all dir_migrate=sql/migrations go run cmd/jxc/main.go
```

## 配置本地 .env

默认情况工程读取 .base.env, 若要覆盖其中的选项，可以创建一个 .env 文件，在文件中写编写需要覆盖的属性

以下是一个 .env 配置的例子

```bash
DEV=yes
DB_CONNECT_URL="host=localhost port=5432 user=postgres password=qwe123jkl dbname=dog sslmode=disable TimeZone=Asia/Shanghai"
redisAddr="localhost:6379"
redisPassword="qwe123jkl"
```

## sqlm 编译表结构

使用 sqlm 编译 表结构模型, sqlm (https://github.com/ymzuiku/sqlc/tree/main-ex), sqlm 会读取 migrations 的历史生成 model，并且会忽略 migrate-down 的代码块。

### 安装 sqlm

下载 main-ex 分支:

```
git clone -b main-ex github.com/ymzuiku/sqlc
```

进入到 ymzuiku/sqlc 工程，确保本地有 go 环境, 编译 sqlm 全局执行程序：

```bash
go install ./cmd/sqlm
```

### 编译表结构

安装完毕，回到本工程，执行 sqlm 命令：

```bash
sqlm generate
```

## 测试

前端，gewu_jxc 所有用户行为都有 100%的客户端自动化集成测试，从用户最终行为模拟上覆盖了所有业务。

服务端，gewu_jxc 要求对每个 `app/apis/account` 下的业务代码以 TDD 的方式开发

每次提交之后，只需要确保当前的业务最终通过可前端的所有自动化测试即可

### 单元测试

单元测试，安装 golang 的建议，直接写在业务代码中，具体参考 `kit/randomCode_test.go`

- -count=1 忽略测试缓存，有副作用的测试，如查询数据库，需要忽略测试缓存
- -cover 查看测试覆盖率
- -tags test 忽略文件首行包含 `//+build !test` 的文件进行覆盖测试,记得文件第二行需要换行
- -coverprofile=t.out 导出测试描述
- -v 测试即便成功，也显示过程和日志

测试并计算覆盖率

```bash
go test ./app/... -count=1 -cover -tags test -coverprofile=t.out
```

显示具体测试覆盖率 html

```bash
go tool cover -html=t.out
```

仅执行某个函数

```bash
go test ./app/... -count=1 -test.run TestHas
```

## 性能测试

首先性能测试需要进入到相应的目录,如 xxx/benchamark 然后执行:

- -bench='Generate' 表示匹配 func BenchmarkXXX 中包含 `Generate` 字样的名称:
- -test.benchmen 表示显示内存用量和内存分配次数
- -benchtime=5s -benchtime=50x 表示执行 5 秒或者默认的 50 倍

```bash
go test -bench='Generate' .  -test.benchmem
```
