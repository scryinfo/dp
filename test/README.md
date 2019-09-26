# 自动化测试  
> 这是一个自动化测试的工具（未完成）
---  
### 使用  
#### 准备
初始化一条链、部署智能合约，启动用户服务（参考app）  
将智能合约地址，复制到dp/test/binary/tools/config.txt文件中的对应位置  
运行相同目录下的**prepare.ps1**脚本，完成**生成可执行文件**以及**修改配置文件中的智能合约地址**步骤  
按照你的终端路径，修改dp/test/binary/utils/utils.go中的地址  
#### 运行  
在test/binary目录，启动上面的脚本生成的**main.exe**文件，或执行```go run main.go```命令即可。  
下面会介绍本工具的具体实现，如果你只是想要使用它，请忽略以下内容。     
---  
### 流程介绍    
1. 注册一批验证者  
1. 执行cases  
1. 记录执行结果  
  
在每一个case中，我们为每一个身份都单独开启了一个进程，这些进程是阻塞的、与主程序通过pipe通信；主程序的pipe接收到约定的终止信号(string "[exit]")时，会结束对应的进程：  
 - 当一个case的全部进程都正常结束时，记录结果：对应case测试通过；  
 - 当一定时间(200s)内，一个case的某个进程仍没有发送终止信号时，记录结果：对应case测试不通过。（未实现）  
   
最后，将所有case的执行结果保存到文件：dp/test/binary/result.txt  
---  
### cases  
> 描述每一个case覆盖的功能流程  

| cases | processes | inVerification |  
| --- | --- | --- |  
| case1 | publish(x) -> pre-buy -> cancel | false |  
