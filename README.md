# 抖音项目服务端 -- 六带一队

作者：周恋程

### 一. 重要资料

- [抖音项目方案说明](https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof)
- [接口说明文档](https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707523)
- [极简抖音App使用说明](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)
- [服务端Demo仓库地址](https://github.com/RaymondCode/simple-demo)
- [本项目仓库地址](https://github.com/DouyinSixwithone/DouyinOrigin)

### 二. 开发环境配置

1. 代码运行环境（版本号为本机环境，应该不需要完全相同）：

   - Golang 1.19.5

   - Mysql 8.0

   - Redis 5.0.14.1

   以上环境的安装教程百度即可。

2. 运行程序需要更改配置：进入config目录更改对应的mysql、redis、server、path信息。

3. 使用安卓模拟器或安卓手机进行测试，[可以参考这篇文章](https://juejin.cn/post/7192600701745233979)。

4. 个人习惯：使用Goland进行开发，Typora写文档。

### 三. github协同开发tips

1. 进入organization，确认能否可见仓库以及是否拥有修改权限。

2. 配置 github 的 SSH key，[可以参考这篇文章](https://blog.csdn.net/zhouzhiwengang/article/details/122247683)。

3. 打开 git bash，使用SSH链接克隆项目到本地。

   ```shell
   git clone git@github.com:DouyinSixwithone/DouyinOrigin.git
   ```

   若使用 Goland，也可以从新建项目处选择Get from VCS，复制项目SSH到URL处，确认本地工作目录，创建项目。

   <img src="C:\Users\lc\AppData\Roaming\Typora\typora-user-images\image-20230203170910729.png" alt="image-20230203170910729" style="zoom:40%;float:left;" />

4. 进入dev分支进行add commit push操作，不要直接在main上修改。

   如果push失败，可能是因为其他队员的最新提交和你试图推送的提交有冲突，可以先用`git pull`把最新的提交从远程分支dev中抓下来，然后在本地合并，解决冲突。

   更多细节可参考：[廖雪峰的git教程-多人协作](https://www.liaoxuefeng.com/wiki/896043488029600/900375748016320)。

   ```shell
   git checkout -b dev origin/dev  // 在本地创建dev分支
   
   git add . //将所有的更改都拉到commit上
   
   git commit -m "备注信息" //提交
   
   git push origin dev //将更改推送到远程分支dev
   ```

5. 阶段性开发完毕后，队长会负责将dev分支中的内容合并到main分支。

