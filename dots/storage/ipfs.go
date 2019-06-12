package storage

import (
    "github.com/ipfs/go-ipfs-api"
    "github.com/pkg/errors"
    "github.com/scryinfo/dot/dot"
    "strings"
)

const (
    IpfsTypeId = "6763f97f-dfd2-40eb-8925-b8a031aab461"
    IPFSLiveId = "6763f97f-dfd2-40eb-8925-b8a031aab461"
)

type Ipfs struct {
    sh     *shell.Shell
    config ipfsConfig
}

type ipfsConfig struct {
    NodeAddr string `json:"nodeAddr"`
    OutDir string `json:"outDir"`
}

func (c *Ipfs) Create(l dot.Line) error {
    if c.sh == nil {
        c.sh = shell.NewShell(c.config.NodeAddr)
        if c.sh == nil {
            return errors.New("Failed to create ipfs shell.")
        }
    }

    return nil
}

func GetIPFSIns() *Ipfs {
    logger := dot.Logger()
    l := dot.GetDefaultLine()
    if l == nil {
        logger.Errorln("the line do not create, do not call it")
        return nil
    }
    d, err := l.ToInjecter().GetByLiveId(IPFSLiveId)
    if err != nil {
        logger.Errorln(err.Error())
        return nil
    }

    if g, ok := d.(*Ipfs); ok {
        return g
    }

    logger.Errorln("do not get the IPFS dot")
    return nil
}

func GetIPFSConfig() *ipfsConfig {
    logger := dot.Logger()
    l := dot.GetDefaultLine()
    if l == nil {
        logger.Errorln("the line do not create, do not call it")
        return nil
    }
    d, err := l.ToInjecter().GetByLiveId(IPFSLiveId)
    if err != nil {
        logger.Errorln(err.Error())
        return nil
    }

    if g, ok := d.(*Ipfs); ok {
        return &g.config
    }

    logger.Errorln("do not get the IPFS dot")
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
