# lego-certmgr 一款使用 lego 生成域名证书的代理服务

`lego-certmgr` 是一个基于 [lego - Github](https://github.com/go-acme/lego) Libiray 封装的证书申请 **代理** 。

其目的是

1. 为了快速方便的申请 **Let's Encrypt** 证书
2. 提供 RESTful API 接口， 方便下游系统 (ex `cmdb`) 调用并进行资源管理

因此

1. `certmgr` 为了方便快速返回已生成过的证书而缓存了一份结果。
2. 由于 `certmgr` 定位是 **代理** ， 所以并未考虑证书的 **持久化** 和 **过期重建** 操作。 

## 使用说明

### 下载 

访问 Github 下载最新版 lego-certmgr [GitHub Release - lego-certmgr](https://github.com/tangx/srv-lego-certmgr/releases/latest)


## Usage

```bash
export DNSPOD_API_KEY=123123123,123123
export ADMIN_EMAIL=xxxx@example.com

./certmgr --dnspod

# 
export ALICLOUD_ACCESS_KEY=ACCasdfasdfasdf
export ALICLOUD_SECRET_KEY=SECaasdf0sdfa02sdfa
export ADMIN_EMAIL=xxxx@example.com

./certmgr --alidns

```

**路由**

```
[GIN-debug] POST   /certmgr/gen/:provider/:domain   --> 创建证书
[GIN-debug] GET    /certmgr/gen/:provider/:domain   --> 查询证书， 303 redirect
[GIN-debug] GET    /certmgr/query/:domain --> 查询证书
[GIN-debug] GET    /certmgr/query/:domain/download --> 下载证书
[GIN-debug] GET    /certmgr/list        --> 查询缓存中生成的所有证书
```

> provider: `alidns` or `dnspod`

## todo

+ [x] 优化 `routes/qcloud` ， 使其完成多 provider 注册式功能， 以支持多 provider
+ [x] 优化 **初始化设置** 支持读取配置文件或环境变量， 实现多 provider 注册。 
+ [ ] 优化 `initial` 逻辑 ， 同一个 email 只向 `let's encrypt` 注册一次
