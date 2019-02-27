package settings

import (
	"errors"
	"github.com/scryinfo/iscap/demo/src/sdk/core/ethereum/events"
	ev "github.com/scryinfo/iscap/demo/src/sdk/core/ethereum/events"
	cf "github.com/scryinfo/iscap/demo/src/sdk/util/configuration"
	rlog "github.com/sirupsen/logrus"
)

const (
	SETTING_LOCATION = "../sdk/settings/definition.yaml"
)

func SaveLastEvent(event ev.Event) error {
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

func LoadSettings() (*scryinfo, error) {
	rv, err := cf.GetYAMLStructure(SETTING_LOCATION, &scryinfo{})
	if err != nil {
		rlog.Error("failed to get yaml structure, error:", err)
		return nil, err
	}

	scryinfo, ok := rv.(*scryinfo)
	if !ok {
		emsg := "failed to convert stream to yaml structure"
		rlog.Error(emsg)
		return nil, errors.New(emsg)
	}

	return scryinfo, nil
}
