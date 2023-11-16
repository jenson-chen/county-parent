## 项目简介
项目整体架构为微服务架构，各服务之间采用gRPC方式通信，服务器与用户之间采用HTTP协议进行通信。应用场景：用户登录后查询所采集的区县数据。

## 项目结构
`county-parent`目录：顶级父目录；

`core`目录：核心子目录，其中包含了实体类的定义、配置项的加载、常量、工具类的定义等；

`county`目录：区县微服务(端口信息见`core/constant`)；

`user`目录：用户微服务(端口信息见`core/constant`)；

`web`目录：对外访问接口，使用`gin`框架做路由和访问。

流程：用户在前端通过发送一个HTTP请求至`web`应用，`web`中完成对用户的鉴权等操作，之后根据具体请求的服务决定调用哪一个微服务(以`county`服务为例)，`web`通过gRPC客户端接口远程调用`county`服务，`county`服务根据需求查询`es`或`redis`，两者都失败的情况下，查询`PostgreSQL`。

## 启动方式
1. 根据本地PostgreSQL、ElasticSearch、Redis的配置信息更改`county-parent/application.toml`中的相关信息。
2. 启动`user`和`county`微服务(两个文件夹下的`main.go`)。
3. 启动`web`服务(当前文件夹下的`main.go`)