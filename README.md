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

|     姓名     |               分工                |            进度            |
| :----------: | :-------------------------------: | :------------------------: |
| 胡杨（队长） |        Github仓库管理，xxx        |          尚未开始          |
|    周恋程    | 项目架构搭建，user接口，jwt中间件 | 已完成user部分，未严格测试 |
|    陈博宇    |                xxx                |          尚未开始          |
|    董弘宇    |                xxx                |          尚未开始          |
|    王君宇    |                xxx                |          尚未开始          |

### 三. 开发环境配置

1. 代码运行环境（版本号为本机环境，不需要完全相同）：

   - Golang 1.19.5

   - Mysql 8.0

   - Redis 5.0.14

   以上环境的安装教程百度即可。

2. 根据自己的环境修改`DouyinOrigin/config/config.yaml`中的内容，一般只需要修改用户名和密码。

   完成以上两步后，在终端输入`go run main.go msgServer.go`即可自动下载依赖并运行。

3. 使用安卓模拟器或安卓手机进行测试，[可以参考这篇文章](https://juejin.cn/post/7192600701745233979)。

4. 个人习惯：使用Goland进行开发，蓝叠模拟器进行测试，Typora写文档，Navicat查看数据库信息。

### 四. Github协同开发tips

1. 加入organization: **DouyinSixwithone**，找到仓库**DouyinOrigin**，确认是否拥有dev分支的修改权限（理论上来说没有main分支的修改权限？）。

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

   * 推送到远程分支之前，请保证程序是**可运行**的状态，并**在commit中尽量详细地写明更新的内容**，并**简单更新README中的进度部分**。

   * 如果push失败，可能是因为其他队员的最新提交和你试图推送的提交有冲突，可以先用`git pull`把最新的提交从远程分支dev中抓下来，然后在本地合并，解决冲突。

     更多细节可参考：[廖雪峰的git教程-多人协作](https://www.liaoxuefeng.com/wiki/896043488029600/900375748016320)。
   
   * 如果不喜欢输命令行，可以使用Goland中的Git菜单进行操作，尤其建议使用Goland进行提交前的检查。
   
     <img src="https://raw.githubusercontent.com/Leng-Chu/picture/main/2023/02/upgit_20230204_1675506189.png" alt="image-20230203170910729" style="zoom:60%;" />

5. 根据开发进度，队长会不定时将dev分支中的内容合并到main分支。

### 五. 项目架构

1. 使用到的技术

   * 框架：gin、gorm

   * 数据库：Mysql

   * 其他
     * Redis：缓存
     * jwt：生成token、鉴权
     * bcrypt：对输入的password进行加密，数据库中存储加密后的密码
     * yaml：写配置文件

2. 采用 **repository → service → controller** 的分层结构：

   <img src="https://raw.githubusercontent.com/Leng-Chu/picture/main/2023/02/upgit_20230204_1675513814.png" alt="image-20230204203013698" style="zoom: 60%;" />

   * **controller层**
     * 解析得到参数，传递给service层。
     
     * 如果需要返回数据信息，则调用service层的逻辑得到数据；如果不需要返回数据信息，只需要执行特定动作修改数据库，那么调用service层的逻辑执行这个动作。
     
     * 将得到的数据（如果有）与状态码和状态描述打包，返回响应。
   * **service层**

     * 如果上层需要返回数据信息，则进行参数检查、数据准备、数据打包；如果上层不需要返回数据信息，则进行参数检查、动作的执行。

     * 进行数据准备或动作执行时，需要调用repository层的逻辑。
   * **repository层**

     * 面向数据库进行增删改查。

3. 文件目录说明

   其中controller和service文件夹中根据功能模块做了分包，如果不涉及对其他模块的调用，可以只专注于自己负责的部分。

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
   ├── msgServer.go  demo中提供的消息服务，不太清楚怎么用
   ├── main.go  程序入口
   └── README.md
   ```

4. 个人实现过程中参考了以下几个项目：

   * https://github.com/HammerCloth/tiktok
   * https://github.com/ACking-you/byte_douyin_project
   * https://github.com/Henrik-Yao/douyin

   如果不知道如何开始，也可参考已经实现好的user部分。

5. 其他注意事项

   * 数据库不需要手动建表，已在db_init.go文件中调用DB.AutoMigrate函数，根据结构体自动建表，示例如下：

     ```go
     type User struct {
     	gorm.Model
     	Name     string
     	Password string
     }
     if err := DB.AutoMigrate(User{}); err != nil {
     		return err
     }
     ```

     事实上表已经在repository的各个文件中定义好，应该没什么需要修改的（确信）。

   * 对token进行鉴权的操作已经在中间件中完成，也就是说写接口时不太需要管token参数。

     部分接口只传入了token参数，未传入user_id，此时需要使用以下代码，从token中获取user_id：
   
     ```
     //方法1：从token中解析出id，解析的步骤已经在中间件中写好，直接调用get方法即可
      idToken, ok := c.Get("user_id")
      // 根据ok判断是否合法
      id, err := idToken.(uint)
      // 判断err
     
     //方法2：如果传入了user_id参数，可以直接调用query方法得到id
     	idStr := c.Query("user_id")
     	id, err := strconv.ParseUint(idStr, 10, 64)
     	// 判断err
     	// 注意此时id的类型为uint64，可能需要强制转换为uint再使用
     ```
   
   * Redis我也还没用过，但先把配置文件和初始化函数写好了看你们写的时候用不用得到（比如评论和赞的部分可以调用缓存？）
   
   * 很多地方的接口需要返回用户信息，可以直接调用service/user中的GetUserInfo函数。
   
   * 有问题可以在群里问或者私戳我，同时我也是新手，欢迎大家对项目架构进行进一步优化。
