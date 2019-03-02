package sdkinterface

import "testing"

func TestInitialize(t *testing.T) {
    err := Initialize()
    if err != nil {
        t.Fail()
    }

    if scryInfo == nil {
        t.Fail()
    }
}
