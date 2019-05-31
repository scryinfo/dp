package storage

import (
    shell "github.com/ipfs/go-ipfs-api"
    "github.com/pkg/errors"
    "github.com/scryinfo/dot/dot"
    "strings"
)

const (
    IpfsTypeId = "6763f97f-dfd2-40eb-8925-b8a031aab461"
)

type Ipfs struct {
    sh     *shell.Shell
    config ipfsConfig
}

type ipfsConfig struct {
    nodeAddr string
}

func (c *Ipfs) Create(l dot.Line) error {
    if c.sh == nil {
        c.sh = shell.NewShell(c.config.nodeAddr)
        if c.sh == nil {
            return errors.New("Failed to create ipfs shell.")
        }
    }

    return nil
}

//construct dot
func newIpfsDot(conf interface{}) (dot.Dot, error) {
    var err error
    var bs []byte
    if bt, ok := conf.([]byte); ok {
        bs = bt
    } else {
        return nil, dot.SError.Parameter
    }

    dConf := &ipfsConfig{}
    err = dot.UnMarshalConfig(bs, dConf)
    if err != nil {
        return nil, err
    }

    d := &Ipfs{config: *dConf}

    return d, err
}

//Data structure needed when generating newer component
func IpfsTypeLive() *dot.TypeLives {
    return &dot.TypeLives{
        Meta: dot.Metadata{TypeId: IpfsTypeId,
            NewDoter: func(conf interface{}) (dot dot.Dot, err error) {
                return newIpfsDot(conf)
            }},
    }
}

func (c *Ipfs) Save(value []byte) (string, error) {
    if c.sh == nil {
        return "", errors.New("Ipfs api shell is nil")
    }

    return c.sh.Add(strings.NewReader(string(value)))
}

func (c *Ipfs) Get(key string, outDir string) error {
    if c.sh == nil {
        return errors.New("Get from ipfs failed, ipfs api shell is nil. ")
    }

    return c.sh.Get(key, outDir)
}
