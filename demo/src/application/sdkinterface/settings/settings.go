package settings

import (
    "github.com/pkg/errors"
    cf "github.com/scryinfo/iscap/demo/src/sdk/util/configuration"
)

const (
    SettingLocation = "D:/EnglishRoad/workspace/Go/src/github.com/scryinfo/iscap/demo/src/application/sdkinterface/settings/definition.yaml"
)

func LoadSettings() (*ScryInfo, error) {
    rv, err := cf.GetYAMLStructure(SettingLocation, &ScryInfo{})
    if err != nil {
        return nil, errors.Wrap(err, "Get YAML structure failed. ")
    }

    scryinfo, ok := rv.(*ScryInfo)
    if !ok {
        return nil, errors.New("Convert data stream to YAML structure failed. ")
    }

    return scryinfo, nil
}
