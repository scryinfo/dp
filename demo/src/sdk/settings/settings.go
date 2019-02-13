package settings

import (
    "../core/ethereum/events"
    ev "../core/ethereum/events"
    cf "../util/configuration"
    rlog "github.com/sirupsen/logrus"
    "errors"
)

const (
    SETTING_LOCATION = "../sdk/settings/definition.yaml"
)

func SaveLastEvent(event ev.Event) (error) {
    rv, err := LoadSettings()
    if err != nil {
        rlog.Error("failed to load settings, error:", err)
        return err
    }

    rv.Chain.LastEvent = event

    err = cf.SaveChanges(SETTING_LOCATION, rv)
    if err != nil {
        rlog.Error("failed to save settings, error:", err)
        return err
    }

    return nil
}

func LoadLastEvent() (events.Event, error) {
    rv, err := LoadSettings()
    if err != nil {
        rlog.Error("failed to load settings, error:", err)
        return events.Event{}, err
    }

    return rv.Chain.LastEvent, nil

}

func LoadLogPath() (*Log, error) {
    rv, err := LoadSettings()
    if err != nil {
        rlog.Error("failed to load settings, error:", err)
        return nil, err
    }

    return &rv.Log, nil
}

func LoadSettings() (*ScryInfo, error) {
    rv, err := cf.GetYAMLStructure(SETTING_LOCATION, &ScryInfo{})
    if err != nil {
        rlog.Error("failed to get yaml structure, error:", err)
        return nil, err
    }

    scryInfo, ok := rv.(*ScryInfo)
    if !ok {
        emsg := "failed to convert stream to yaml structure"
        rlog.Error(emsg)
        return nil, errors.New(emsg)
    }

    return scryInfo, nil
}