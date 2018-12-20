package chainevents

import "github.com/qjpcpu/ethereum/events"

type EventCallback func(event events.Event) bool

