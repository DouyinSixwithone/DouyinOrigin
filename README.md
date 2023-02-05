# 抖音项目服务端 -- 六带一队

作者：周恋程

（如果 README.md 中的图片无法显示，建议科学上网）

### 一. 重要资料

- [抖音项目方案说明](https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof)
- [接口说明文档](https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707523)
- [极简抖音App使用说明](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)
- [服务端Demo仓库地址](https://github.com/RaymondCode/simple-demo)
- [本项目仓库地址](https://github.com/DouyinSixwithone/DouyinOrigin)

### 二. 小组成员及分工

|     姓名     |          分工          |        进度        |
| :----------: | :--------------------: | :----------------: |
| 胡杨（队长） |  Github仓库管理，xxx   |      尚未开始      |
|    周恋程    | 项目架构搭建，user部分 | 完成注册、登录接口 |
|    陈博宇    |          xxx           |      尚未开始      |
|    董弘宇    |          xxx           |      尚未开始      |
|    王君宇    |          xxx           |      尚未开始      |
|    陈毅杰    |          xxx           |      尚未开始      |

### 三. 开发环境配置

1. 代码运行环境（版本号为本机环境，不需要完全相同）：

   - Golang 1.19.5

   - Mysql 8.0

   - Redis 5.0.14

   以上环境的安装教程百度即可。

2. 根据自己的环境修改`DouyinOrigin/config/config.yaml`中的内容，一般只需要修改用户名和密码。

   完成以上两步后，在终端输入`go run main.go`即可自动下载依赖并运行。

3. 使用安卓模拟器或安卓手机进行测试，[可以参考这篇文章](https://juejin.cn/post/7192600701745233979)。

4. 个人习惯：使用Goland进行开发，蓝叠模拟器进行测试，Typora写文档，Navicat查看数据库信息。

### 四. Github协同开发tips

1. 加入organization: **DouyinSixwithone**，找到仓库**DouyinOrigin**，确认是否拥有dev分支的修改权限（理论上来说没有main分支的修改权限）。

2. 配置 Github 的 SSH key，[可以参考这篇文章](https://blog.csdn.net/zhouzhiwengang/article/details/122247683)。

3. 打开 git bash，使用SSH链接克隆项目到本地。

   ```shell
   git clone git@github.com:DouyinSixwithone/DouyinOrigin.git
   ```

   若使用 Goland，也可以从新建项目处选择**Get from VCS**，复制项目SSH到URL处，确认本地工作目录，创建项目。

   <img src="https://raw.githubusercontent.com/Leng-Chu/picture/main/2023/02/upgit_20230203_1675430178.png" alt="image-20230203170910729" style="zoom:50%;" />

4. 进入dev分支进行add commit push操作，**不要直接在main上修改**。

    ```shell
   git checkout -b dev origin/dev  // 在本地创建dev分支
   
   git add . //将所有的更改都拉到commit上
   
   git commit -m "备注信息" //提交
   
   git push origin dev //将更改推送到远程分支dev
   
   git pull //从远程分支中得到最新的提交并合并到本地
   ```

   * 推送到远程分支之前，请保证程序是可运行的状态，并在commit中尽量详细地写明更新的内容。

   * 如果push失败，可能是因为其他队员的最新提交和你试图推送的提交有冲突，可以先用`git pull`把最新的提交从远程分支dev中抓下来，然后在本地合并，解决冲突。

     更多细节可参考：[廖雪峰的git教程-多人协作](https://www.liaoxuefeng.com/wiki/896043488029600/900375748016320)。
   
   * 如果不喜欢输命令行，可以使用Goland中的Git菜单进行操作，尤其建议使用Goland进行提交前的检查。
   
     <img src="https://raw.githubusercontent.com/Leng-Chu/picture/main/2023/02/upgit_20230204_1675506189.png" alt="image-20230203170910729" style="zoom:45%;" />

5. 根据开发进度，队长会不定时将dev分支中的内容合并到main分支。

### 五. 项目架构

1. 使用到的技术

   * 框架：gin、gorm

   * 数据库：Mysql

   * 其他
     * Redis：缓存
     * jwt：生成token、鉴权
     * bcrypt：对输入的password进行加密
     * yaml：写配置文件

2. 采用 **repository → service → controller** 的分层结构：

   <img src="https://raw.githubusercontent.com/Leng-Chu/picture/main/2023/02/upgit_20230204_1675513814.png" alt="image-20230204203013698" style="zoom: 43%;" />

   * **controller层**
     * 解析得到参数，传递给service层。
     
     * 如果需要返回数据信息，则调用service层的逻辑得到数据；如果不需要返回数据信息，只需要执行特定动作修改数据库，那么调用service层的逻辑执行这个动作。
     
     * 将得到的数据（如果有）与状态码和状态描述打包，返回响应。
     
   * **service层**

     * 如果上层需要返回数据信息，则进行参数检查、数据准备、数据打包；如果上层不需要返回数据信息，则进行参数检查、动作的执行。

     * 进行数据准备或动作执行时，需要调用repository层的逻辑。

   * **repository层**

     * 面向数据库进行增删改查。

     * 根据结构体自动建表，需要在db_init文件中调用DB.AutoMigrate函数。

3. 文件目录说明

   ```
   Douyin 
   ├── /config/ 配置文件
   ├── /common/ 通用结构体
   ├── /controller/ 视图层
   ├── /service/ 逻辑层
   ├── /repository/ 数据层
   ├── /middleware/ 中间件
   │   ├── jwt/ 鉴权
   │   └── redis/ 缓存
   ├── /router/ 路由配置
   ├── /data/ 上传的视频文件存储在本地
   ├── /go.mod/
   ├── main.go
   └── README.md
   ```

4. 个人实现过程中参考了以下几个项目：

   * https://github.com/HammerCloth/tiktok
   * https://github.com/ACking-you/byte_douyin_project
   * https://github.com/Henrik-Yao/douyin
