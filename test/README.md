# 自动化测试  
> 这是一个自动化测试的工具
---  
### 使用  
#### 准备
初始化一条链、部署智能合约，启动用户服务，开启ipfs守护进程（参考app）  
将智能合约地址，复制到test/binary/tools/contract.txt文件中的对应位置  
运行相同目录下的[placeholder]()脚本，生成可执行文件  
运行相同目录下的[placeholder]()脚本，修改配置文件中的智能合约地址  
按照你的终端路径，修改test/binary/utils/utils.go中的地址  
#### 运行  
在test/binary目录，执行```go run main.go```即可运行这个工具（计划把这一部也生成可执行文件）  
---  
### cases  
> 描述每一个case执行的功能流程  

| cases | processes |  
| --- | --- |  
| case1 | publish -> pre-buy -> cancel |  
