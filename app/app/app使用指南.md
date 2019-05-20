# app

> 一个基于dp的Dapp demo。

本例实现了基础交易功能与验证功能。

## 使用：

安装下面列出的[环境](#env)，并按照文档中的指引完成[编译过程](#compile)与[运行准备](#run_pre)步骤后，即可通过生成的可执行文件开始体验。

## <span id="env">环境</span>：

### 编译环境：

- go (1.12.5)
- node.js (10.15.3)
- gcc (mingw-w64 v8.1.0)

> 上述环境需要自行安装，未列出不需要自行安装的环境(如webpack、truffle)与可选的环境(如python)。

### 运行环境：

- ipfs客户端 (0.4.20)
- geth客户端 (1.8.27)
- 浏览器 (chrome 74)

## <span id="compile">编译过程</span>：

### 打包UI资源文件：

> 我们假设你已经完成了node.js的下载与安装。

下载dp/app/app/ui/package.json中配置好的依赖，然后使用webpack打包UI资源文件。

你可以通过ui/config/index.js中的bundleAnalyzerReport控制是否显示webpack结果分析报告。

我们将上述命令执行过程写成了脚本，你也可以通过执行ui目录下的webpackUI.ps1文件完成这一步骤。

### 构建app可执行文件：

在dp/app/app/main目录下执行go build。

## <span id="run_pre">运行准备</span>：

### 启动用户服务：

运行用户服务的可执行文件，默认使用48080端口。

### 连接ipfs：

> 我们假设你已经完成了ipfs的下载与安装。
	
app只使用ipfs的读写功能，如果你只在一台终端上使用，单节点（本机）启动ipfs守护进程即可（成功执行ipfs daemon命令），不需要接入一个节点网络。

然而，值得一提的是，app使用js执行ipfs上传，所以你至少需要允许ipfs跨域执行post请求。

### 搭建一条私链：

> 我们假设你已经完成了geth的下载与安装。
	
使用geth搭建一条私链，[参考链接](https://github.com/ethereum/go-ethereum/wiki/Private-network)。

### 部署智能合约：
    
在上一步搭建好的链上，部署dp/dots/binary/contracts目录下的智能合约。

如果你的链使用了自定义的端口等内容，请修改truffle配置文件与之匹配。

我们将上述命令执行过程写成了脚本，你也可以通过执行contracts目录下的contract.ps1文件完成这一步骤。

请查询并记录**所有智能合约的地址**，以供后续使用。

### 修改app配置文件：

| key | value |
|:------- |:------- |
app.chain.contracts.tokenAddr | 修改为前面记录的token合约地址 
app.chain.contracts.protocolAddr | 修改为前面记录的protocol合约地址
app.chain.contracts.deployerKeyjson | 修改为keystore目录下，对应用户的文件内容，注意转义双引号
app.chain.contracts.deployerPassword | 修改为创建账户时使用的密码
app.chain. ethereum.ethNode | 按照格式修改为geth的节点地址
app.services.keystore | 修改为认证服务的地址
app.config.uiResourcesDir | 修改dp的目录即可
app.config.ipfsOutDir | 修改为你期望的ipfs下载路径

## 异常处理：

- npm install error， 找不到python exec：安装python2或忽略该问题。

- 用户服务启动失败，找不到vcruntime140.dll：[安装vcre](https://www.microsoft.com/zh-cn/download/details.aspx?id=48145)。
