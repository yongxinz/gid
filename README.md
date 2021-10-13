# gid

Golang 开发的一款分布式唯一 ID 生成系统

## 使用

有两种方式来调用接口：

1. HTTP 方式
2. gRPC 方式

### HTTP 方式

1、健康检查：

```
curl http://127.0.0.1:8080/ping
```

2、获取 ID：

获取 tag 是 test 的 ID：

```
curl http://127.0.0.1:8080/id/test
```

3、获取雪花 ID：

```
curl http://127.0.0.1:8080/snowid
```

### gRPC 方式

1、获取 ID：

```
grpcurl -plaintext -d '{"tag":"test"}' -import-path $HOME/src/gid/app/modules/home/homerpc/homepb -proto segment.proto localhost:50051 app.homepb.Gid/GetId
```

2、获取雪花 ID：

```
grpcurl -plaintext -import-path $HOME/src/gid/app/modules/home/homerpc/homepb -proto segment.proto localhost:50051 app.homepb.Gid/GetSnowId
```

## 开发

开发环境对应配置文件：conf/app.gid.toml

1、初始化 MySQL：

```
create database dev;
use dev;
```

```
create table segments
(
    biz_tag     varchar(32) not null,
    max_id      bigint       null,
    step        int          null,
    remark      varchar(200) null,
    create_time datetime       null,
    update_time datetime       null,
    constraint segments_pk
        primary key (biz_tag)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;
```

```
INSERT INTO segments(`biz_tag`, `max_id`, `step`, `remark`, `create_time`, `update_time`)
VALUES ('test', 0, 100, 'test', NOW(), NOW());
```

2、启服务

```
yago run main.go
```

## 部署

生产环境对应配置文件：conf/app.toml，并且提供了 Docker 的部署方式：

```
docker-compose build
docker-compose up -d
```

## 推荐阅读

- [hwholiday/gid](https://github.com/hwholiday/gid)
- [Leaf——美团点评分布式ID生成系统](https://tech.meituan.com/2017/04/21/mt-leaf.html)