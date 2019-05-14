package settings

import (
	"github.com/pkg/errors"
	"github.com/scryInfo/dp/dots/binary/sdk/util/config_yaml"
)

type ScryInfo struct {
	Chain    Chain    `yaml:"chain",json:"chain"`
	Services Services `yaml:"services",json:"services"`
	Config   Config   `yaml:"config",json:"config"`
}

type Chain struct {
	Contracts Contracts `yaml:"contracts",json:"contracts"`
	Ethereum  Ethereum  `yaml:"ethereum",json:"ethereum"`
}
type Contracts struct {
	TokenAddr        string `yaml:"token_contract_addr",json:"tokenAddr"`
	ProtocolAddr     string `yaml:"protocol_contract_addr",json:"protocolAddr"`
	DeployerKeyJson  string `yaml:"deployer_keyjson",json:"deployerKeyJson"`
	DeployerPassword string `yaml:"deployer_password",json:"deployerPassword"`
}
type Ethereum struct {
	EthNode string `yaml:"node",json:"ethNode"`
}

type Services struct {
	Ipfs     string `yaml:"ipfs",json:"ipfs"`
	Keystore string `yaml:"keystore",json:"keystore"`
}

type Config struct {
	WSPort         string `yaml:"websocket_port",json:"wsPort"`
	UIResourcesDir string `yaml:"ui_resources_dir",json:"uiResourcesDir"`
	LogPath        string `yaml:"log_path",json:"logPath"`
	AppId          string `yaml:"app_id",json:"appId"`
	IPFSOutDir     string `yaml:"ipfs_out_dir",json:"ipfsOutDir"`
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
