package settings

import (
    "github.com/scryinfo/iscap/demo/src/sdk/core/ethereum/events"
)

type scryinfo struct {
    Chain Chain  `yaml:"Chain"`
    Log Log      `yaml:"Log"`
}

type Chain struct {
    LastEvent events.Event  `yaml:"LastEvent"`

}

type Log struct {
    Dir string
    File string
}

