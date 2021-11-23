module github.com/scryinfo/dp/dots/eth/client

go 1.13

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.39.0
	go.uber.org/atomic => github.com/uber-go/atomic v1.4.0
	go.uber.org/zap => github.com/uber-go/zap v1.10.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190513172903-22d7a77e9e5f
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190510132918-efd6b22b2522
	golang.org/x/image => github.com/golang/image v0.0.0-20190507092727-e4e5bf290fec
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190509164839-32b2708ab171
	golang.org/x/mock => github.com/golang/mock v1.3.1
	golang.org/x/mod => github.com/golang/mod v0.1.0
	golang.org/x/net => github.com/golang/net v0.0.0-20190509222800-a4d6f7feada5
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190509141414-a5b02f93d862
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190513233021-7d589f28aaf4
	google.golang.org/api => github.com/google/google-api-go-client v0.5.0
	google.golang.org/appengine => github.com/golang/appengine v1.5.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190513181449-d00d292a067c
	google.golang.org/grpc => github.com/grpc/grpc-go v1.21.0
	gopkg.in/check.v1 => github.com/go-check/check v0.0.0-20180628173108-788fd7840127
	gopkg.in/errgo.v2 => github.com/go-errgo/errgo v0.0.0-20180823071505-f768c5ab0476
	gopkg.in/natefinch/npipe.v2 => github.com/natefinch/npipe v0.0.0-20160621034901-c1b8fa8bdcce
	gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.20.0
	honnef.co/go/tools => github.com/dominikh/go-tools v0.0.0-20190418001031-e561f6794a2a
)

require (
	github.com/ethereum/go-ethereum v1.10.10
	github.com/scryinfo/dot v0.1.3
	go.uber.org/zap v1.10.0
)
