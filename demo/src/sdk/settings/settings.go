package settings


import (
    cf "../util/configuration"
    ev "../core/ethereum/events"
    "errors"
    "fmt"
    "../core/ethereum/events"
)

const (
    SETTING_LOCATION = "../sdk/settings/definition.yaml"
)

func SaveLastEvent(event ev.Event) (error) {
    rv, err := LoadSettings()
    if err != nil {
        fmt.Println("failed to load settings, error:", err)
        return err
    }

    rv.Chain.LastEvent = event

    err = cf.SaveChanges(SETTING_LOCATION, rv)
    if err != nil {
        fmt.Println("failed to save settings, error:", err)
        return err
    }

    return nil
}

func LoadLastEvent() (events.Event, error) {
    rv, err := LoadSettings()
    if err != nil {
        fmt.Println("failed to load settings, error:", err)
        return events.Event{}, err
    }

    return rv.Chain.LastEvent, nil

}

func LoadSettings() (*ScryInfo, error) {
    rv, err := cf.GetYAMLStructure(SETTING_LOCATION, &ScryInfo{})
    if err != nil {
        fmt.Println("failed to get yaml structure, error:", err)
        return nil, err
    }

    scryInfo, ok := rv.(*ScryInfo)
    if !ok {
        emsg := "failed to convert stream to yaml structure"
        fmt.Println(emsg)
        return nil, errors.New(emsg)
    }

    return scryInfo, nil
}