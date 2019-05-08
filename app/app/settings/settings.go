package settings

import (
	"github.com/pkg/errors"
	"github.com/scryInfo/dp/dots/binary/sdk/util/config_yaml"
)

func LoadSettings(setfile string) (*ScryInfo, error) {
	rv, err := config_yaml.GetYAMLStructure(setfile, &ScryInfo{})
	if err != nil {
		return nil, errors.Wrap(err, "Get YAML structure failed. ")
	}

	scryinfo, ok := rv.(*ScryInfo)
	if !ok {
		return nil, errors.New("Convert data stream to YAML structure failed. ")
	}

	return scryinfo, nil
}
