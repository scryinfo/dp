package settings

import (
	"github.com/scryinfo/dot/dot"
)

const (
	ConfTypeId = "1241ebcb-4b9d-498d-9749-de0e0cd7d6a2"
	ConfLiveId = "1241ebcb-4b9d-498d-9749-de0e0cd7d6a2"
)

type Config struct {
	WSPort         string `json:"wsPort"`
	UIResourcesDir string `json:"uiResourcesDir"`
	IPFSOutDir     string `json:"IPFSOutDir"`
	AccsBackupFile string `json:"accsBackupFile"`
}

func (c *Config) Create(l dot.Line) error {
	return nil
}

func GetConfig() *Config {
	logger := dot.Logger()
	l := dot.GetDefaultLine()
	if l == nil {
		logger.Errorln("the line do not create, do not call it")
		return nil
	}
	d, err := l.ToInjecter().GetByLiveId(ConfLiveId)
	if err != nil {
		logger.Errorln(err.Error())
		return nil
	}

	if g, ok := d.(*Config); ok {
		return g
	}

	logger.Errorln("do not get the Config dot")
	return nil
}

//construct dot
func newConfDot(conf interface{}) (dot.Dot, error) {
	var err error
	var bs []byte
	if bt, ok := conf.([]byte); ok {
		bs = bt
	} else {
		return nil, dot.SError.Parameter
	}

	d := &Config{}
	err = dot.UnMarshalConfig(bs, d)

	if err != nil {
		return nil, err
	}

	return d, err
}

//Data structure needed when generating newer component
func ConfTypeLive() *dot.TypeLives {
	return &dot.TypeLives{
		Meta: dot.Metadata{TypeId: ConfTypeId,
			NewDoter: func(conf interface{}) (dot dot.Dot, err error) {
				return newConfDot(conf)
			}},
	}
}
