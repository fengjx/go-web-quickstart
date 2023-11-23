# go-web-quickstart

go web 工程示例，可以作为一些项目工程结构参考，对一些常用功能做了简单封装。支持构建docker镜像。

示例中已完成了部分功能实现，方便做一些参考
- 用户登录、注册
- 博客查看、添加、删除、修改

## 环境依赖

- MySQL
  - 初始化sql在`deployments/db`
- Redis

对应配置在`configs`目录，可以自行修改

## make 指令

```bash
$ make help
Makefile cmd:

    build:                              项目打包
    build-go:                           构建 golang 包
    fmt-go:                             格式化 golang 代码
    tidy:                               去掉未使用的项目依赖
    clean:                              清理临时文件
    gen:                                代码生成，具体查看 gen.yml 配置
    help:                               Makefile 帮助
```

## 启动项目

```bash
go run cmd/main.go configs/app-local.yaml
```

## 打包

```bash
make build
```

打包后的可执行文件生成在`.dist`目录

## docker

打包
```bash
# 镜像名和版本可以自己定义
docker build . -t web-app:1.0.0
```

启动
```bash
# APP_ENV 启用不同环境配置
docker run  -p 8080:8080 -e APP_ENV=test --name webapp web-app:1.0.0
```

## 技术选型

- [gin](https://github.com/gin-gonic/gin) - web 框架
- [json-iterator](http://jsoniter.com/go-tips.cn.html) - 高效 json 类库
- [go-yaml](https://github.com/go-yaml/yaml) - yaml 文件加载
- [daox](https://github.com/fengjx/daox) - 基于 sqlx + go-redis 的数据访问工具库 
- [go-redis](https://github.com/redis/go-redis) redis 客户端
- [lo](https://github.com/samber/lo) 一个类似 lodash 的集合工具类库

## 工程结构
- build: 工程构建相关
- cmd: 应用启动入口
- configs: 项目配置
- deployments: 应用依赖
- init: 应用启动配置
- internal: 应用业务逻辑代码
  - endpoint：功能分包，每个功能点放一个目录，包括 dao、service、controller 等，其他功能有依赖则暴露相关api
- pkg: 放到应用外部依然能使用的代码库、工具类
- test: 测试相关
- tools: 项目工具，如代码生成脚本

## 参考

- [go 项目分层结构](https://github.com/golang-standards/project-layout/blob/master/README_zh.md)
