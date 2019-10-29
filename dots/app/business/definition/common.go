package definition

type EncryptedId struct {
	EncryptedId []byte `json:"encryptedMetaDataId,omitempty"`
}

type Arbitrate struct {
	ArbitrateResult bool `json:"arbitrateResult,omitempty"`
}

type User struct {
	Address    string   `json:"address,omitempty"`
	Nickname   string   `json:"nickname,omitempty"`
	FromBlock  int64    `json:"from_block,omitempty"`
	IsVerifier bool     `json:"is_verifier,omitempty"`
	Verify     []string `json:"verify,omitempty"`
}
