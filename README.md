# GIN-KIT
提供使用gin构建RESTful API,需要的一系列套件以及目录结构参考
## Structure
```
.
├── README.md
├── go.mod
├── sqlboiler.example.toml
├── config
│   ├── config.go
│   ├── development.yaml
│   ├── production.yaml
│   └── test.yaml
├── controller
│   └── user.go
├── repository
├── db
│   └── db.go
├── models
│   └── user.go
├── service
│   └── user.go
├── middleware
│   └── auth.go
├── entity
│   └── user.go
└── cmd
│   └── server
│   │   ├── router.go
│   │   └── server.go
│   └── main.go

```

## Kits

* [x] 数据库mapping工具(sqlboiler)
* [x] 环境变量(flag)
* [x] 配置管理(viper)
* [x] 日志
* [x] http(GIN)
* [x] 依赖管理(go mod)
* [x] 验证器(gin自带的验证器)
* [x] 部署(docker)
* [ ] error封装


## Usage
Init Entity
~~~
cd [root of the project]
//初始化数据库连接
cp sqlboiler.example.toml sqlboiler.toml

//生成entity层(https://github.com/volatiletech/sqlboiler)
sqlboiler mysql
~~~

Deployment
~~~
docker build . -t [docker image name]
~~~