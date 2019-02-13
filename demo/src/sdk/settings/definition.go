package settings

import (
    "../core/ethereum/events"
)

type ScryInfo struct {
    Chain Chain  `yaml:"Chain"`
}

type Chain struct {
    LastEvent[] events.Event  `yaml:"LastEvent"`
}
