[中文](./README-cn.md)  
[EN](./README.md)  
# Introduction
Through SDK for data exchange with blockchain, developers can get easy access to DAPP development. It mainly includes data encryption and decryption, signature, smart contract, event notification, data storage API, data acquisition and search, digital currency payment, third party APP payment API and so on.The process is as following:  
Data provider inputs data and metadata through SDK (data includes static data and dynamic data which have the agreed format; metadata mainly includes data signature, data description and etc). Data demander can find the required data through SDK and obtain the data after paying digital currency. The data verifier can be qualified by pledging a certain amount of digital currency to the smart contract. In the process of data exchange, the data demander can initiate the compensable data verification request or transaction arbitration to the smart contract, and the verifier will be randomly selected by the smart contract. All participants in the data exchange can score each other in the transaction；The smart contract would record the transaction and scores of the participants thus generate the reputation evaluation of the participants which can be inquired through SDK
# Windows
##  Edit
###  Edit environment
> The following environment should be installed yourself. Excepted environment(like webpack, truffle) and optional environment(like python) is not listed here
> The following is suggested version that has been tested
- go (1.12.5)
- node.js (10.15.3)
- gcc (mingw-w64 v8.1.0)
### Package UI source files：
> We assume that you have finished node.js download and installation  

Run **webpackUI.ps1** script in dp/app/app/ui content to finish this process 
You can control whether to display webpack result analysis through *bundleAnalyzerReport* in ui/config/index.js  
### Build app executable file：
Run: go build in dp/app/app/main content，entrance file: **main.exe** will be generated if succeeded.
##  Operation
### Operating environment
- ipfs client (0.4.20)
- geth client (1.8.27)
- Browser (chrome 74)
### Start user service:
Run user service executable file in dp/dots/auth content，default API is 48080
### ipfs connection：
> We assume that you have finished ipfs download and installation
- Adjust config files: find config files in your ipfs download path like following, add following 3 "Access..." for config item "API"   Config：  
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
- Run ```ipfs daemon``` in command line，"Daemon is ready" will be displayed if succeeded，Keep command line open
> Since app use js to upload ipfs ，so "permit ipfs Cross-domain execute post request" config is added above  
### Build one private chain:
> We assume that you have finished geth download and installation

Run **geth_init.ps1** script in dp/dots/binary/contracts/geth_init content to finish private chain building   
Run **geth_acc_mine.ps1** script in the same content to create user and start mining
### Deploy smart contract:
Run **contract.ps1** script in dp/dots/binary/contracts content to finish this process 
Script will input part of result to migrate.log in the same content, *ScryToken*、*ScryProtocol* two  42-character address with "0x" in the beginning can be found in the file end  
### Adjust app config file：
| key | value |
|:------- |:------- |
app.chain.contracts.tokenAddr | Adjust to ScryToken address found in logfile 
app.chain.contracts.protocolAddr | Adjust to ScryProtocol address found in logfile 
app.chain.contracts.deployerKeyjson | Adjust to unique file contents under dp/dots/binary/contracks/geth_init/chain/keystore content，pay attention to double quotes
app.config.uiResourcesDir | Adjust the content of dp 
app.config.ipfsOutDir | Adjust to your selected ipfs download path
### Experience
After finishing all process above, you can experience it through dp/app/app/main/main.exe entrance file
## Exception handling：
- windows banned ps1 script operation：Use administrator privileges to open command line, run Set-ExecutionPolicy unrestricted
- npm install error，python exec is not found：install python2 or ignore this problem
- User service start failure, vcruntime140.dll is not found：[install vcre](https://www.microsoft.com/zh-cn/download/details.aspx?id=48145).
- Smart contract deployment failure, failed to get connected to ether client: Check whether customized API is used to build private chain, adjust truffle.js config file network.geth.port in contracts content to get consistent with it
- Smart contract deployment is not displayed: Check powershell opened by geth_init.ps1 is still mining or not(information will be refreshed constantly).
# [Code Style -- Go](https://github.com/scryinfo/scryg/blob/master/codestyle_go.md)
# [ScryInfo Protocol Layer SDK API Document v0.0.5](https://github.com/scryinfo/dp/blob/master/document/ScryInfo%20protocol%20layer%20SDK%20%20v0.0.5.md)
