// Scry Info.  All rights reserved.
// license that can be found in the license file.

package settings

// Scryinfo scryinfo
type Scryinfo struct {
	Chain    Chain    `yaml:"chain"`
	Services Services `yaml:"services"`
}

// Chain chain
type Chain struct {
	Contracts Contracts `yaml:"contracts"`
	Ethereum  Ethereum  `yaml:"ethereum"`
}

// Contracts contracts
type Contracts struct {
	ProtocolAddr     string `yaml:"protocol_contract_addr"`
	TokenAddr        string `yaml:"token_contract_addr"`
	ProtocolAbiPath  string `yaml:"protocol_contract_abi_path"`
	TokenAbiPath     string `yaml:"token_contract_abi_path"`
	ProtocolEvents   string `yaml:"protocol_contract_events"`
	TokenEvents      string `yaml:"token_contract_events"`
	DeployerKeyJson  string `yaml:"deployer_keyjson"`
	DeployerPassword string `yaml:"deployer_password"`
}

// Ethereum ethereum
type Ethereum struct {
	EthNode string `yaml:"node"`
}

// Services services
type Services struct {
	Ipfs     string `yaml:"ipfs"`
	Keystore string `yaml:"keystore"`
}
