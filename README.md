# lego-certmgr 一款使用 lego 生成域名证书的代理服务

`lego-certmgr` 是一个基于 [lego - Github](https://github.com/go-acme/lego) Libiray 封装的证书申请 **代理** 。

其目的是

1. 为了快速方便的申请 **Let's Encrypt** 证书
2. 提供 RESTful API 接口， 方便下游系统 (ex `cmdb`) 调用并进行资源管理

因此

1. `certmgr` 为了方便快速返回已生成过的证书而缓存了一份结果。
2. 由于 `certmgr` 定位是 **代理** ， 所以并未考虑证书的 **持久化** 和 **过期重建** 操作。 

## 使用说明

访问 `http(s)://yourdomain.com` 可以进入图形化界面

![index.png](./docs/img/index.png)

### 下载 

访问 Github 下载最新版 lego-certmgr [GitHub Release - lego-certmgr](https://github.com/tangx/srv-lego-certmgr/releases/latest)


## Usage

使用 `viper` 进行配置管理， 可以通过 `环境变量` 或 `配置文件` 进行参数传递

### 通过环境变量

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

### 使用配置文件

+ 路径为 `$HOME/certmgr` 或 `程序当前目录`
+ 文件名为 `config.yml / config.yaml`

```yaml
# dnspod
DNSPOD_API_KEY : 123123123,123123

# alidns
ALICLOUD_ACCESS_KEY : ACCasdfasdfasdf
ALICLOUD_SECRET_KEY : SECaasdf0sdfa02sdfa

ADMIN_EMAIL : xxxx@example.com
```

**路由**

```
[GIN-debug] GET    /index/*filepath          --> 首页
[GIN-debug] GET    /lego-certmgr/query/:domain --> 查询域名证书信息
[GIN-debug] GET    /lego-certmgr/query/:domain/download --> 下载域名证书
[GIN-debug] GET    /lego-certmgr/list        --> 获取域名证书列表
[GIN-debug] GET    /lego-certmgr/healthy     --> 健康检查
[GIN-debug] POST   /certmgr/gen/:provider/:domain   --> 创建证书
[GIN-debug] GET    /certmgr/gen/:provider/:domain   --> 查询证书， 303 redirect
```

> provider: `alidns` or `dnspod`

## todo

+ [x] 优化 `routes/qcloud` ， 使其完成多 provider 注册式功能， 以支持多 provider
+ [x] 优化 **初始化设置** 支持读取配置文件或环境变量， 实现多 provider 注册。 
+ [ ] 优化 `initial` 逻辑 ， 同一个 email 只向 `let's encrypt` 注册一次
