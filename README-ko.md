[中文](./README-cn.md)  
[EN](./README.md)  
[한국어](./README-ko.md)  
[日本語](./README-ja.md)  
# 소개
블록체인 데이터 교환 SDK를 통해 개발자는 아주 편리하고 신속한 속도로 DAPP 프로그램을 개발할 수 있다. 주로 데이터 암호화, 사인, 스마트 컨트랙트, 사건 통지, 데이터 저장 인터페이스, 데이터 수집 및 검색, 암호화폐 결제 그리고 서드 파티 페이 인터페이스 등 기능이 포함되여 있다. 그에 해당된 프로세스는 아래와 같다:  
데이터 지원자는 SDK를 통해 데이터 및 메타 데이터(데이터에는 정적 데이터, 동적 데이터가 있고 메타 데이터에는 데이터 사인, 데이터 설명 등 정보가 포함됨)를 작성할 수 있고 반면 데이터 수요자는 역시 SDK를 통해 필요한 데이터를 검색할 수 있을 뿐만 아니라 소량의 암호화폐를 결제하면 해당 데이터를 호출하여 사용할 수 있다. 또한 데이터 검증 신청자는 스마트 컨트랙트를 통해 일정한 비례의 암호화폐를 보증하면 진정한 데이터 검증자로 담임된다. 데이터 교환 프로세스에서 데이터 수요자는 유료 데이터 검증에 대해 요청 또는 컨트랙트에 대한 트랜잭션의 중재를 개시하게 되면 검증자는 스마트 컨트랙트에 따라 랜덤으로 투표하는 형식으로 검증을 진행한다. 데이터 교환에 참여한 모든 참여자들은 자체의 트랜잭션에서 서로 평점할 수 있으며 스마트 컨트랙트를 통해 참여자의 거래 및 평가한 정보를 기록하여 모든 참여자의 신뢰에 대한 평가가 생성된다. 여기에서 신뢰된 정보는 SDK를 통해 검색될 수 있다.
# Windows
## 컴파일
### 컴파일 환경
> 아래 환경은 스스로 설치해야 한다. 여기에 나열되지 않은 환경 (예: webpack, truffle)및 선택 가능한 환경(예: python)은 스스로 설치할 필요가 없다. 
> 아래에 괄호가 있는 버전은 이미 테스트한 버전이다.
- go (1.12.5)
- node.js (10.15.3)
- gcc (mingw-w64 v8.1.0)
### 패키지 UI 소스 파일：
> node.js의 다운로드 및 설치를 완료했다고 가정할 경우
첫 단계는 dp / app / app / ui 디렉토리에서 ** webpackUI.ps1 ** 스크립트 파일을 실행해야 한다.
webpack결과 분석 보고 관련 디스플레이의 여부는 ui/config/index.js에 있는 *bundleAnalyzerReport*로 제어할 수 있다.  
### app 실행 파일 작성
dp/app/app/main의 디렉토리에서 go build를 성공적으로 실행하게 되면 엔트런스 파일”**main.exe**”이 생성된다.    
## 오퍼레이션
### 오퍼레이션 환경
- ipfs 클라이언트 (0.4.20)ㄴ
- geth 클라이언트 (1.8.27)
- 웹브라우저(chrome 74)
### 유저 서비스의 스타트：
dp/dots/auth의 디렉토리에서 유저 서비스의 실행 파일을 작동할 수 있고 기본 API는 48080이다.
### ipfs에 연결：
> ipfs의 다운로드와 설치를 완료했다고 가정할 경우
- 구성 파일 조정: ipfs의 다운로드 경로에서 config파일을 찾을 수 있으며 설정 항목 ”API”에 아래와 같은 3” Access ...”를 추가할 수 있다.
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
- 커맨드 라인에 ipfs daemon을 실행하여 성공되면 "Daemon is ready"가 표시되고 커맨드 라인의 오픈을 유지한다.
> App는 js를 사용하여 ipfs를 업로드하기에 위에 “ipfs 크로스 실행 요청을 허가” 설정을 추가해야 한다.
###프라이빗 체인 구축：
> Geth의 다운로드 및 설치를 완료했다고 가정할 경우
dp / dots / binary / contracts / geth_init의 디렉토리에서 **geth_init.ps1** 스크립트를 실행해야 프라이빗 체인을 구축할 수 있다.
동일한 디렉토리에서 ** geth_acc_mine.ps1 ** 스크립트를 실행해야 유저를 생성하여 마이닝을 진행할 수 있다.
### 스마트 컨트랙트를 배치：
첫 단계 프로세스를 완료하려면 dp/dots/binary/contracts의 디렉토리에서 **contract.ps1** 스크립트를 실행해야 한다.
스크립트는 일부의 결과를 동일한 디렉토리에 있는 migrate.log 파일에 아웃풋할 수 있고 파일 끝부분에서 *ScryToken*、*ScryProtocol* 두개의 "0x" 시작된 42자리 주소를 찾을 수 있다.
### app 구성 파일을 수정：
| key | value |
|:------- |:------- |
app.chain.contracts.tokenAddr |를 로그 파일에서 발견된 ScryToken 주소로 조정
app.chain.contracts.protocolAddr |를 로그 파일에서 발견된 ScryToken 주소로 조정
app.chain.contracts.deployerKeyjson |를 dp / dots / binary / contracks / geth_init / chain / keystore 콘텐츠에 있는 고유한 파일 콘텐츠로 조정하고 큰 따옴표를 사용해야 한다.
app.config.uiResourcesDir |를 dp의 콘텐츠로 조정
app.config.ipfsOutDir |를 ipfs의 다운로드 경로로 조정
### 체험
위의 모든 절차를 완료해야만 dp / app / app / main / main.exe 엔트런스 파일을 통해 체험할 수 있다.
##예외 처리：
- Windows에서 ps1 스크립트 실행을 금지: 관리자 권한으로 커맨드 라인을 오픈하고 Set-ExecutionPolicy를 제한없이 실행할 수 있다.
- npm install error, python exec를 찾을 수 없음: python2를 설치하거나 문제를 무시한다.
- 유저 서비스의 스타트가 실패, vcruntime140.dll를 찾을 수 없음: [설치 vcre] (https://www.microsoft.com/zh-cn/download/details.aspx?id=48145)
-스마트 컨트랙트의 배치가 실패되여 이더리움 클라이언트에 연결되지 않을 경우: 자체 정의의 포트로 프라이빗 체인을 구축하였는지 확인, contracts 디렉토리의 truffle.js 구성 파일 network.geth.port과 맞게 조정
- 스마트 컨트랙트를 배치하여 그에 대한 디스플레이가 되지 않을 경우: geth_init.ps1에서 열린 powershell 창이 여전히 마아닝(정보가 지속적으로 업데이트됨)상태인지 확인
# [Code Style -- Go](https://github.com/scryinfo/scryg/blob/master/codestyle_go-ko.md)
# [ScryInfo Protocol Layer SDK API Document v0.0.5](https://github.com/scryinfo/dp/blob/master/document/ScryInfo%20protocol%20layer%20SDK%20%20v0.0.5.md)
