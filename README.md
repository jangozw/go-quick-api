# gin-api-common 
github: https://github.com/jangozw/gin-api-common

* 基于go web框架gin的简单封装, 因为gin没有项目目录等传统框架定义的东西, 封装的目的是扩展一些现代接口开发的必备功能, 提高开发效率, 你可以拿来即用。

* 1小时入门go web开发， 为go初学者展示如何使用gin框架快速web开发，这个项目有完整示范: 
启动server->路由定义->中间件过滤->用户身份验证->控制器处理->数据库操作->redis缓存->返回给客户端, 这是web接口开发基本流程

# 功能特色
* gin框架为基础。文档： https://github.com/gin-gonic/gin
* 数据库操作orm 用 gorm。 文档： http://gorm.book.jasperxu.com 
* 使用 JWT 生成token, 结合redis双重验证。 文档 http://jwt.io
* gin 使用的验证器文档: https://godoc.org/gopkg.in/go-playground/validator.v8
* 日志用logrus  文档: https://github.com/sirupsen/logrus
* docker 部署, docker-compose 编排容器一键启动整套服务
* fresh 本地监测编译工具:https://github.com/gravityblast/fresh

# 版本要求

 * go1.11+
 * 使用 go modules 包管理

# 启动运行

### docker 方式

1, 构建项目代码镜像

```shell script
docker build -t ginapicommon_main:latest -f deploy/docker-compose/project.build .
```
2, 构建mysql镜像(含初始化数据)
```shell script
docker build -t ginapicommon_mysql:latest -f deploy/docker-compose/mysql.build . 
```
3, 启动

```shell script
docker compose up 
```
或
```shell script
docker-compose up -d
```


成功后如图：

```text
ginapicommon_main |
ginapicommon_main | [GIN-debug] GET    /                         --> github.com/jangozw/gin-api-common/routes.registerNoLogin.func1 (3 handlers)
ginapicommon_main | [GIN-debug] POST   /v0/login                 --> github.com/jangozw/gin-api-common/apis/v0.Login (3 handlers)
ginapicommon_main | [GIN-debug] POST   /v0/logout                --> github.com/jangozw/gin-api-common/apis/v0.Logout (4 handlers)
ginapicommon_main | [GIN-debug] GET    /v0/user/list             --> github.com/jangozw/gin-api-common/apis/v0.UserList (4 handlers)
ginapicommon_main | [GIN-debug] GET    /v0/user/detail           --> github.com/jangozw/gin-api-common/apis/v0.UserDetail (4 handlers)
ginapicommon_main | [GIN-debug] POST   /v0/user/add              --> github.com/jangozw/gin-api-common/apis/v0.AddUser (4 handlers)
ginapicommon_main | [GIN-debug] Listening and serving HTTP on :8080
```


#### 停止和编译代码重启


1, 停止服务
```shell script
docker-compose down
```

2, 修改代码后重新编译启动

```shell script
docker-compose down
docker rmi ginapicommon_main:latest
docker build -t ginapicommon_main:latest -f deploy/docker-compose/project.build .
docker compose up 
```



### 本地启动

1, 创建配置文件的软连接
```shell script
sudo ln -s $(pwd)/config.ini /etc/ginapicommon_config.ini
````

2， 本地启动

```shell script
go run main.go
```


2，使用fresh实时检测代码变动并重新编译启动, 适合开发期间频繁修改代码



在项目根目录
```sh 
fresh
```



选择一种方式启动后, 打开浏览器访问： http://127.0.0.1:8080 看效果
```
⇒  curl http://127.0.0.1:8080/
{"code":200,"msg":"请求成功","timestamp":1559253308,"data":"Welcome!"}%
```


# 请求示例

* 添加用户
路由在/routes/api.go 中可以看到 

 
* 登陆


路由在/routes/api.go 中可以看到 ```/v0/login``` 




参数:
```json
{"mobile": "1500000000", "pwd": "1234546"}
```
响应:
```json
{
    "code": 200,
    "msg": "请求成功",
    "timestamp": 1576573879,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIsInVlbiI6ImVjNDc2ZDJkNGU3ODhkYzA3YzFkNDI3NGVkZjA1Y2Y1YmQyMGI4YWYwYTdlODcwYTAzMzRmYjZlZDg2MzNiZDQiLCJleHAiOjE1NzY2NjAyNzksImlzcyI6InRlc3QifQ.erealfYAsbxkvoyf3IxXvRSX46hZt4G6JxPQmYoNvNc"
    }
}
```

客户端要将 token 保存， 然后下次请求的HEADER 中 Authorization的值设为token


* 添加用户

请求：


POST  ```/v0/user/add```
header 中Authorization的值设为token

```json
{
    "mobile": "15000000000",
    "pwd": "123456",
    "name": "test"
    }
```
返回：


成功




* 用户列表


路由: ```/v0/user/list```


参数: 无 


响应:
```json
{
    "code": 200,
    "msg": "请求成功",
    "timestamp": 1576574102,
    "data": {
        "total": 2,
        "page_size": 20,
        "list": [
            {
                "id": 1,
                "mobile": "1500000000",
                "name": "test"
            },
            {
                "id": 2,
                "mobile": "1500000001",
                "name": "test"
            }
        ]
    }
}

```
* 中间件使用方法

见```/middleewares/api.go```
 



# 扩展

* go 是编译型的， 开发期间每次改动代码都要重新 go build, 实时监视代码改动可用 fresh 工具: https://github.com/gravityblast/fresh
* 将数据库中的表生成go struct 可用gormt 工具: https://github.com/xxjwxc/gormt
* 开发期间为提高效率，不用docker部署，直接本地开发，修改config.ini 数据库账号即可
