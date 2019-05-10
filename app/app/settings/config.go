package settings

import (
	"github.com/pkg/errors"
	"github.com/scryInfo/dp/dots/binary/sdk/util/config_yaml"
)

type ScryInfo struct {
	Chain    Chain    `yaml:"chain"`
	Services Services `yaml:"services"`
	Config   Config   `yaml:"config"`
}

type Chain struct {
	Contracts Contracts `yaml:"contracts"`
	Ethereum  Ethereum  `yaml:"ethereum"`
}
type Contracts struct {
	TokenAddr        string `yaml:"token_contract_addr"`
	ProtocolAddr     string `yaml:"protocol_contract_addr"`
	DeployerKeyJson  string `yaml:"deployer_keyjson"`
	DeployerPassword string `yaml:"deployer_password"`
}
type Ethereum struct {
	EthNode string `yaml:"node"`
}

type Services struct {
	Ipfs     string `yaml:"ipfs"`
	Keystore string `yaml:"keystore"`
}

type Config struct {
	WSPort         string `yaml:"websocket_port"`
	UIResourcesDir string `yaml:"ui_resources_dir"`
	LogPath        string `yaml:"log_path"`
	AppId          string `yaml:"app_id"`
	IPFSOutDir     string `yaml:"ipfs_out_dir"`
}

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
