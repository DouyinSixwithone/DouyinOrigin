# 抖音项目服务端 -- 六带一队

### 一. 重要资料

- [抖音项目方案说明](https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof)
- [接口说明文档](https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707523)
- [极简抖音App使用说明](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)
- [服务端Demo仓库地址](https://github.com/RaymondCode/simple-demo)
- [github仓库地址](https://github.com/DouyinSixwithone/DouyinOrigin)

### 二. 开发环境配置

1. 代码运行环境（版本号为本机环境，应该不需要完全相同）：

   - Golang 1.19.5

   - Mysql 8.0

   - Redis 5.0.14.1

   以上环境的安装教程百度即可。

2. 运行程序需要更改配置：进入config目录更改对应的mysql、redis、server、path信息。

3. 使用安卓模拟器或安卓手机进行测试，[可以参考这篇文章](https://juejin.cn/post/7192600701745233979)。

4. 个人习惯：使用Goland进行开发，Typora写文档。