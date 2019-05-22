[中文](./README-cn.md)  
[EN](./README.md)  
# 介绍
通过区块链提供数据交换的SDK，让开发者可以快速方便的开发DAPP应用。主要包含这些内容：数据加解密、签名、智能合约、事件通知、数据存储接口、数据获取与查询、数字货币支付、第三方App支付接口等。过程如下：  
数据提供者通过SDK写入数据及元数据（数据包含静态数据、动态数据，数据有约定的格式，元数据主要包含数据签名、数据描述等信息）；数据需求者通过SDK查找需要的数据，支付数字货币后，可以获得数据；数据验证者，通过向智能合约抵押一定的数字货币来成为验证者。在数据交换过程中，数据需求者可以向合约发起有偿数据验证请求或交易仲裁，验证者由智能合约随机选取。数据交换的所有参与者可以在参与的交易中相互评分；智能合约会记录参与者交易及评分信息，生成参与者的信誉评价，信誉评价信息可以通过SDK查询。
# Windows
##  编译
###  编译环境
> 下列环境需要自行安装，未列出不需要自行安装的环境(如webpack、truffle)与可选的环境(如python)。    
括号内为经过测试的推荐版本。
- go (1.12.5)
- node.js (10.15.3)
- gcc (mingw-w64 v8.1.0)
### 打包UI资源文件：
> 我们假设你已经完成了node.js的下载与安装。

执行dp/app/app/ui目录下的**webpackUI.ps1**脚本文件完成这一步骤。  
你可以通过ui/config/index.js中的*bundleAnalyzerReport*控制是否显示webpack结果分析报告。  
### 构建app可执行文件：
在dp/app/app/main目录下执行go build，成功执行后，会生成入口文件：**main.exe**。
##  运行
### 依赖
- ipfs客户端 (0.4.20)
- geth客户端 (1.8.27)
- 浏览器 (chrome 74)
### 启动用户服务：
运行dp/dots/auth目录下的，用户服务的可执行文件，默认使用48080端口。
### 连接ipfs：
> 我们假设你已经完成了ipfs的下载与安装。
- 修改配置文件，在你的ipfs下载路径中，找到config文件。如下所示，为其一级配置项"API"添加下面三条"Access..."配置：  
```json
"API": {
  "HTTPHeaders": {
    "Server": [
      "go-ipfs/0.4.14"
    ],
    "Access-Control-Allow-Origin": [
      "*"
    ],
    "Access-Control-Allow-Credentials": [
      "true"
    ],
    "Access-Control-Allow-Methods": [
      "POST"
    ]
  }
},
```
- 在命令行执行 ipfs daemon 命令，执行成功时会显示"Daemon is ready"，保持命令行窗口开启。
> 因为app使用js进行ipfs上传，所以上面添加了"允许ipfs跨域执行post请求"的配置。
### 搭建一条私链：
> 我们假设你已经完成了geth的下载与安装。

执行dp/dots/binary/contracts/geth_init目录下的**geth_init.ps1**脚本文件完成私链搭建。  
执行相同目录下的**geth_acc_mine.ps1**脚本文件创建用户并开始挖矿。
### 部署智能合约：
执行dp/dots/binary/contracts目录下的**contract.ps1**脚本文件完成这一步骤。  
脚本会将部分结果输出到相同目录下的migrate.log文件，在文件末尾可以找到*ScryToken*、*ScryProtocol*两个"0x"开头的42个字符的地址。
### 修改app配置文件：
| key | value |
|:------- |:------- |
app.chain.contracts.tokenAddr | 修改为日志文件中找到的ScryToken地址 
app.chain.contracts.protocolAddr | 修改为日志文件中找到的ScryProtocol地址
app.chain.contracts.deployerKeyjson | 修改为dp/dots/binary/contracks/geth_init/chain/keystore目录下，唯一文件的内容，注意转义双引号
app.config.uiResourcesDir | 修改dp的目录即可
app.config.ipfsOutDir | 修改为你期望的ipfs下载路径
### 体验
完成上述所有步骤后，即可通过dp/app/app/main/main.exe入口文件进行体验。
## 异常处理：
- windows禁止ps1脚本执行：使用管理员权限打开命令行，执行Set-ExecutionPolicy unrestricted命令。
- npm install error，找不到python exec：安装python2或忽略该问题。
- 用户服务启动失败，找不到vcruntime140.dll：[安装vcre](https://www.microsoft.com/zh-cn/download/details.aspx?id=48145)。
- 智能合约部署失败，连接不到以太坊客户端：检查是否使用了自定义的端口搭建私链，修改contracts目录下的truffle.js配置文件network.geth.port与之一致。
- 智能合约部署无显示：查看geth_init.ps1打开的powershell窗口是否仍在挖矿（不断有消息刷新）。
# [Code Style -- Go](https://github.com/scryinfo/scryg/blob/master/codestyle_go-cn.md)
# [ScryInfo协议层SDK接口文档v0.0.5](https://github.com/scryinfo/dp/blob/master/document/ScryInfo%E5%8D%8F%E8%AE%AE%E5%B1%82SDK%E6%8E%A5%E5%8F%A3%E6%96%87%E6%A1%A3v0.0.5.md)
