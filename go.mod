module github.com/scryinfo/dp

go 1.12

require (
	cloud.google.com/go v0.39.0 // indirect
	github.com/allegro/bigcache v1.2.0 // indirect
	github.com/aristanetworks/goarista v0.0.0-20190429220743-799535f6f364 // indirect
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/cespare/cp v1.1.1 // indirect
	github.com/chilts/sid v0.0.0-20180928232130-250d10e55bf4
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/edsrzf/mmap-go v1.0.0 // indirect
	github.com/ethereum/go-ethereum v1.8.27
	github.com/fjl/memsize v0.0.0-20180929194037-2a09253e352a // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/golang/mock v1.3.1 // indirect
	github.com/golang/protobuf v1.3.1
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/google/go-cmp v0.3.0 // indirect
	github.com/google/pprof v0.0.0-20190515194954-54271f7e092f // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/huin/goupnp v1.0.0 // indirect
	github.com/ipfs/go-ipfs-api v0.0.1
	github.com/jackpal/go-nat-pmp v1.0.1 // indirect
	github.com/karalabe/hid v0.0.0-20181128192157-d815e0c1a2e2 // indirect
	github.com/kr/pty v1.1.4 // indirect
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/mattn/go-isatty v0.0.8 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/rs/cors v1.6.0 // indirect
	github.com/scryinfo/dot v0.1.3-0.20190528055454-5ce7a05174dc
	github.com/scryinfo/scryg v0.1.3-0.20190523074957-3a6377ac45ea
	github.com/sirupsen/logrus v1.4.1
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20190513172903-22d7a77e9e5f // indirect
	golang.org/x/exp v0.0.0-20190510132918-efd6b22b2522 // indirect
	golang.org/x/image v0.0.0-20190523035834-f03afa92d3ff // indirect
	golang.org/x/lint v0.0.0-20190409202823-959b441ac422 // indirect
	golang.org/x/mobile v0.0.0-20190509164839-32b2708ab171 // indirect
	golang.org/x/net v0.0.0-20190522155817-f3200d17e092
	golang.org/x/oauth2 v0.0.0-20190523182746-aaccbc9213b0 // indirect
	golang.org/x/sys v0.0.0-20190528012530-adf421d2caf4 // indirect
	golang.org/x/text v0.3.2 // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	golang.org/x/tools v0.0.0-20190525145741-7be61e1b0e51 // indirect
	google.golang.org/appengine v1.6.0 // indirect
	google.golang.org/genproto v0.0.0-20190522204451-c2c4e71fbf69 // indirect
	google.golang.org/grpc v1.21.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0 // indirect
	gopkg.in/urfave/cli.v1 v1.20.0 // indirect
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc // indirect
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.39.0
	go.uber.org/atomic => github.com/uber-go/atomic v1.4.0
	go.uber.org/zap => github.com/uber-go/zap v1.10.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190513172903-22d7a77e9e5f
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190510132918-efd6b22b2522
	golang.org/x/image => github.com/golang/image v0.0.0-20190507092727-e4e5bf290fec
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190509164839-32b2708ab171
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
	google.golang.org/grpc => github.com/grpc/grpc-go v1.20.1
	gopkg.in/natefinch/npipe.v2 => github.com/natefinch/npipe v0.0.0-20160621034901-c1b8fa8bdcce
	gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.20.0
	honnef.co/go/tools => github.com/dominikh/go-tools v0.0.0-20190418001031-e561f6794a2a
)
