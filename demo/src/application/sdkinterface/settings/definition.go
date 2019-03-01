package settings

type ScryInfo struct {
    Chain Chain  `yaml:"chain"`
    Services Services      `yaml:"services"`
}

type Chain struct {
    Contracts Contracts `yaml:"contracts"`
    Ethereum Ethereum `yaml:"ethereum"`
}

type Contracts struct {
    ProtocolAddr string `yaml:"protocol_contract_addr"`
    TokenAddr string `yaml:"token_contract_addr"`
    ProtocolAbiPath string `yaml:"protocol_contract_abi_path"`
    TokenAbiPath string `yaml:"token_contract_abi_path"`
    ProtocolEvents string `yaml:"protocol_contract_events"`
    TokenEvents string `yaml:"token_contract_events"`
}

type Ethereum struct {
    EthNode string `yaml:"node"`
}

type Services struct {
    Ipfs string `yaml:"ipfs"`
    Keystore string `yaml:"keystore"`
}