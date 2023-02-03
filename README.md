# 抖音项目服务端 -- 六带一队

作者：周恋程

### 一. 重要资料

- [抖音项目方案说明](https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof)
- [接口说明文档](https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707523)
- [极简抖音App使用说明](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)
- [服务端Demo仓库地址](https://github.com/RaymondCode/simple-demo)
- [本项目仓库地址](https://github.com/DouyinSixwithone/DouyinOrigin)

### 二. 小组成员及分工

|     姓名     |              分工               |   进度   |
| :----------: | :-----------------------------: | :------: |
| 胡杨（队长） |       Github仓库管理，xxx       | 尚未开始 |
|    周恋程    | 项目架构搭建，开发文档编写，xxx | 尚未开始 |
|    陈博宇    |               xxx               | 尚未开始 |
|    董弘宇    |               xxx               | 尚未开始 |
|    王君宇    |               xxx               | 尚未开始 |
|    陈毅杰    |               xxx               | 尚未开始 |

### 三. 开发环境配置

1. 代码运行环境（版本号为本机环境，不需要完全相同）：

   - Golang 1.19.5

   - Mysql 8.0

   - Redis 5.0.14

   以上环境的安装教程百度即可。

3. 使用安卓模拟器或安卓手机进行测试，[可以参考这篇文章](https://juejin.cn/post/7192600701745233979)。

4. 个人习惯：使用Goland进行开发，蓝叠模拟器进行测试，Typora写文档。

### 四. Github协同开发tips

1. 加入organization: **DouyinSixwithone**，找到仓库**DouyinOrigin**，确认是否拥有dev分支的修改权限（理论上来说没有main分支的修改权限）。

2. 配置 Github 的 SSH key，[可以参考这篇文章](https://blog.csdn.net/zhouzhiwengang/article/details/122247683)。

3. 打开 git bash，使用SSH链接克隆项目到本地。

   ```shell
   git clone git@github.com:DouyinSixwithone/DouyinOrigin.git
   ```

   若使用 Goland，也可以从新建项目处选择Get from VCS，复制项目SSH到URL处，确认本地工作目录，创建项目。

   <img src="https://raw.githubusercontent.com/Leng-Chu/picture/main/2023/02/upgit_20230203_1675430178.png" alt="image-20230203170910729" style="zoom:40%;float:left;" />

4. 进入dev分支进行add commit push操作，**不要直接在main上修改**。

   如果push失败，可能是因为其他队员的最新提交和你试图推送的提交有冲突，可以先用`git pull`把最新的提交从远程分支dev中抓下来，然后在本地合并，解决冲突。

   更多细节可参考：[廖雪峰的git教程-多人协作](https://www.liaoxuefeng.com/wiki/896043488029600/900375748016320)。

   ```shell
   git checkout -b dev origin/dev  // 在本地创建dev分支
   
   git add . //将所有的更改都拉到commit上
   
   git commit -m "备注信息" //提交
   
   git push origin dev //将更改推送到远程分支dev
   
   git pull //从远程分支中得到最新的提交并合并到本地
   ```

5. 根据开发进度，队长会不定时将dev分支中的内容合并到main分支。

### 五. 项目架构

待完善
