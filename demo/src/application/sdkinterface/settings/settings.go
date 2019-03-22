package settings

import (
    "errors"
    cf "github.com/scryinfo/iscap/demo/src/sdk/util/configuration"
    rlog "github.com/sirupsen/logrus"
)

const (
    SettingLocation = "D:/EnglishRoad/workspace/Go/src/github.com/scryinfo/iscap/demo/src/application/sdkinterface/settings/definition.yaml"
)

func LoadSettings() (*ScryInfo, error) {
    rv, err := cf.GetYAMLStructure(SettingLocation, &ScryInfo{})
    if err != nil {
        rlog.Error("failed to get yaml structure, error:", err)
        return nil, err
    }

    scryinfo, ok := rv.(*ScryInfo)
    if !ok {
        emsg := "failed to convert stream to yaml structure"
        rlog.Error(emsg)
        return nil, errors.New(emsg)
    }

    return scryinfo, nil
}

func LoadServiceSettings() (*ScryInfoAS, error) {
    rv, err := cf.GetYAMLStructure(SettingLocation, &ScryInfoAS{})
    if err != nil {
        rlog.Error("failed to get yaml structure, error:", err)
        return nil, err
    }

    scryinfoas, ok := rv.(*ScryInfoAS)
    if !ok {
        emsg := "failed to convert stream to yaml structure"
        rlog.Error(emsg)
        return nil, errors.New(emsg)
    }

    return scryinfoas, nil
}