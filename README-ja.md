[![GoDoc](https://godoc.org/github.com/scryinfo/dp?status.svg)](https://godoc.org/github.com/scryinfo/dp)
[![Go Report Card](https://goreportcard.com/badge/github.com/scryinfo/dp)](https://goreportcard.com/report/github.com/scryinfo/dp)
[![Build Status](https://travis-ci.org/scryinfo/dp.svg?branch=master)](https://travis-ci.org/scryinfo/dp)
[![codecov](https://codecov.io/gh/scryinfo/dp/branch/master/graph/badge.svg)](https://codecov.io/gh/scryinfo/dp)

[中文](./README-cn.md)  
[EN](./README.md)  
[한국어](./README-ko.md)  
[日本語](./README-ja.md)  
# 紹介
ブロックチェーンでデータ交換のSDKを提供することで、開発者が迅速かつ容易にDAPPを開発することができます。主要な内容：データの暗号化と復号化、サイン、スマートコントラクト、イベントの知らせ、データ保存インターフェース、データの獲得と検索、仮想通貨の支払い、第三者Appの支払いインタフェースなどがあります。そのプロセスが以下の通りです。  
データの提供者がSDKを通じてデータとメタデータを書き込みます。（データは静的データ、動的データがあります。データの形式はルールがあります。メタデータは主にデジタル署名、データ記述などの情報があります。）；データ需要者はSDKで需要なデータを検索し、仮想通貨で支払ってデータを獲得します。データ検証者はスマートコントラクトに一定の仮想通貨を引き当てて検証者になります。データ交換の過程に、データ需要者はスマートコントラクトに有償なデータ検証や取引の仲裁を要請することができます。検証者はスマートコントラクトでランダムに選出されます。データ交換の全ての参加者が参加する取引を通じて互いに採点することができます。スマートコントラクトは参加者の取引と得点の情報を記録して参加者の信用評価を生成します。信用評価情報はSDKで検索することができます。  
# Windows
##  コンパイル
###  コンパイル環境
> 以下の環境は自らでインストールします。リストアップしていないのはインストールに必要がありません。（例えば、 webpack、truffle）及びオプショナルな環境（python）です。  
括弧内のはテストしたおすすめのバージョンです。
- go (1.12.5)
- node.js (10.15.3)
- gcc (mingw-w64 v8.1.0)
### UI資源ファイルをバックします：
> あなたがnode.jsのダウンロードとインストールを完成したと仮定します。  

コンテンツdp/app/app/uiの中の**webpackUI.ps1**スクリプトファイルを実行することでその手順を完成させます。  
webpackの結果分析レポートが現れるかどうかはui/config/index.jsの中の*bundleAnalyzerReport*を通じてコントロールします。
### 実行可能なappファイルを構築します：
コンテンツdp/app/app/mainの上にgo buildを実行します。成功的に実行した後に、エントリーファイル**main.exe**が生成します。
##  操作
### 操作環境
- ipfsクライアント(0.4.20)
- gethクライアント (1.8.27)
- ウェブブラウザ(chrome 74)
### ユーザーサービスの起動：
コンテンツdp/dots/authの中のユーザーサービスの実行可能なファイルを起動します。48080ポートを使用するとみなします。
### ipfsにアクセスします：  
> あなたがIpfsのダウンロードとインストールを完成したと仮定します。
- 配置ファイルを修正します。Ipfsのダウンロードルートの中に、configファイルを見つけます。そして、それの一次構成アイテムAPIに以下のように下記の三つの"Access..."構成を加えます。  
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
- コマンドラインにipfs daemonコマンドを実行します。成功した後に、"Daemon is ready"が現れます。コマンドラインのインターフェイスを開いておきます。  
> Appはjsを使ってipfsアップロードするため、ipfsはpost要求をクロスドメインで実行することが許可する構成を付け加えます。
### プライベートチェーンを構築します。
> あなたがgethのダウンロードとインストールを完成したと仮定します。

コンテンツdp/dots/binary/contracts/geth_initの中の**geth_init.ps1**スクリプトファイルを実行してプライベートチェーンを構築します。  
同じコンテンツの中の**geth_acc_mine.ps1**スクリプトファイルを実行し、ユーザー登録をしてマイニングを始めます。
### スマートコントラクトの配置：
コンテンツdp/dots/binary/contractsの中の**contract.ps1**スクリプトファイルを実行してこの手順を完成します。  
スクリプトは部分の結果を同じコンテンツのmigrate.logファイルへ送ります。ファイルの末に*ScryToken*、*ScryProtocol二つの"0x"で始まる42キャラクターのアドレスが見つかります。
### appの構成ファイルを修正します：
| key | value |
|:------- |:------- |
app.chain.contracts.tokenAddr | 日誌ファイルにScryTokenアドレスを見つけることに修正します。
app.chain.contracts.protocolAddr | 日誌ファイルにScryProtocolアドレスを見つけることに修正します。
app.chain.contracts.deployerKeyjson | コンテンツdp/dots/binary/contracks/geth_init/chain/keystoreの中にあるに修正します。唯一のファイルの内容、二重引用符の転換に注意してください。
app.config.uiResourcesDir | dpのコンテンツを修正して良いです。
app.config.ipfsOutDir | 希望するipfsのダウンロードルートに修正します。
### 体験
以上の全ての手順を完成した後に、エントリーファイルdp/app/app/main/main.exeを通じて体験できます。
## 異常の処理：
- windowsがスクリプトファイルps1の実行を禁止します。管理者権限でコマンドラインを起動し、Set-ExecutionPolicy unrestrictedコマンドラインを実行します。  
- npm install error，python execが見つかりません。安装python2をインストールします。もしくはこのエラーをみおとします。  
- ユーザーサービスの起動が失敗します。vcruntime140.dll：[インストールしvcre](https://www.microsoft.com/zh-cn/download/details.aspx?id=48145)が見つかりません。  
- スマートコントラクトの配置が失敗します。イーサリアムクライアントにアクセスできません。自己定義インタフェースを使ってプライベートチェインを構築するかどうかをチェックします。コンテンツcontractsの中のtruffle.jsの配置ファイルnetwork.geth.portを修正して一致させます。  
- スマートコントラクトの配置が現れません。geth_init.ps1で起動するインタフェースpowershellがマイニングしているかどうかをチェックします（メッセージが更新し続けています）。  
# [Code Style -- Go](https://github.com/scryinfo/scryg/blob/master/codestyle_go-ja.md)
# [ScryInfo Protocol Layer SDK API Document v0.0.5](https://github.com/scryinfo/dp/blob/master/document/ScryInfo%20protocol%20layer%20SDK%20%20v0.0.5.md)
