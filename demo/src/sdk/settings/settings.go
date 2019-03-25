package settings

import (
	"github.com/pkg/errors"
	ev "github.com/scryinfo/iscap/demo/src/sdk/core/ethereum/events"
	cf "github.com/scryinfo/iscap/demo/src/sdk/util/configuration"
)

const (
	SettingLocation = "D:/EnglishRoad/workspace/Go/src/github.com/scryinfo/iscap/demo/src/sdk/settings/definition.yaml"
)

func SaveLastEvent(event ev.Event) error {
	rv, err := LoadSettings()
	if err != nil {
		return errors.Wrap(err, "Load definition failed. ")
	}

	rv.Chain.LastEvent = event

	if err = cf.SaveChanges(SettingLocation, rv); err != nil {
		return errors.Wrap(err, "Save definition failed. ")
	}

	return nil
}

func LoadLastEvent() (ev.Event, error) {
	rv, err := LoadSettings()
	if err != nil {
		return ev.Event{}, errors.Wrap(err, "Load definition failed. ")
	}

	return rv.Chain.LastEvent, nil

}

func LoadLogPath() (*Log, error) {
	rv, err := LoadSettings()
	if err != nil {
		return nil, errors.Wrap(err, "Load definition failed. ")
	}

	return &rv.Log, nil
}

func LoadSettings() (*scryinfo, error) {
	rv, err := cf.GetYAMLStructure(SettingLocation, &scryinfo{})
	if err != nil {
		return nil, errors.Wrap(err, "Get YAML structure failed. ")
	}

	scryinfo, ok := rv.(*scryinfo)
	if !ok {
		return nil, errors.New("Convert data stream to YAML structure failed. ")
	}

	return scryinfo, nil
}
