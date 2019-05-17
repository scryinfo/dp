## 编译环境：
- go (12.5)
- node.js (>=6, 10.15.0)
- ipfs客户端
- geth客户端
- 浏览器 （支持websocket协议，chrome 66）

## 编译步骤：
### 连接ipfs：

	我们假设你已经完成了ipfs的下载与安装。
	
app只使用ipfs的读写功能，如果你只在一台终端上使用，单节点（本机）启动ipfs守护进程即可（成功执行ipfs daemon命令），不需要接入一个节点网络。

然而，值得一提的是，app使用js执行ipfs上传，所以你至少需要允许ipfs跨域执行post请求。

### 搭建一条私链：

	我们假设你已经完成了geth的下载与安装。
	
使用geth搭建一条私链，[参考链接](https://github.com/ethereum/go-ethereum/wiki/Private-network)。

### 部署智能合约：

在上一步搭建好的链上部署dp/dots/binary/contracts目录下的智能合约。

在上述目录下执行npm.ps1文件，下载相关的依赖。

如果你的链使用了自定义的端口等内容，修改truffle配置文件与之匹配，然后使用truffle部署智能合约。

请查询并记录 智能合约各自的地址 与合约部署者的私钥，以供后续使用。

### 打包UI资源文件：

    我们假设你已经完成了node.js的下载与安装。

下载dp/app/app/ui/package.json中配置好的依赖，然后使用webpack打包UI资源文件。

你可以通过ui/config/index.js中的bundleAnalyzerReport控制是否显示webpack结果分析报告。

我们将上述命令执行过程写成了脚本，你也可以通过执行ui目录下的webpackUI.ps1文件完成这一步骤。

### 修改app配置文件：
| key | value |
|:------- |:------- |
app.chain.contracts.tokenAddr | 修改为前面记录的token合约地址 
app.chain.contracts.protocolAddr | 修改为前面记录的protocol合约地址
app.chain.contracts.deployerKeyjson | 使用contracts/tool目录下的genKeyJSON.js文件生成，生成前应替换合约部署者的私钥和密码，替换前应转义双引号。
app.chain.contracts.deployerPassword | 与.js文件中的password一致
app.chain. ethereum.ethNode | 按照格式修改为geth的节点地址
app.services.keystore | 修改为认证服务的地址
app.config.uiResourcesDir | 修改dp的目录即可
app.config.ipfsOutDir | 修改为你期望的ipfs下载路径

### 构建app可执行文件：

在dp/app/app/main目录下执行go build。