# srv-cert-manager

使用 lego 创建 `let's encrypt` 证书

根据注册的 provider 添加路由
## Usage

```bash
export DNSPOD_API_KEY=123123123,123123
export DNSPOD_API_EMAIL=xxxx@example.com

./certmgr
```

**路由**

```
[GIN-debug] POST   /certmgr/gen/:provider/:domain   --> 创建证书
[GIN-debug] GET    /certmgr/gen/:provider/:domain   --> 查询证书， 303 redirect
[GIN-debug] GET    /certmgr/cert/query/:domain --> 查询证书
[GIN-debug] GET    /certmgr/cert/query/:domain/download --> 下载证书
[GIN-debug] GET    /certmgr/cert/list        --> 查询缓存中生成的所有证书
```


## todo

+ [x] 优化 `routes/qcloud` ， 使其完成多 provider 注册式功能， 以支持多 provider
+ [ ] 优化 **初始化设置** 支持读取配置文件或环境变量， 实现多 provider 注册。 
+ [ ] 优化 `initial` 逻辑 ， 同一个 email 只向 `let's encrypt` 注册一次