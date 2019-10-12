package definition

type EncryptedId struct {
	EncryptedId []byte `json:"encryptedMetaDataId,omitempty"`
}

type Extensions struct {
	MetaDataExtension   string   `json:"metaDataExtension,omitempty"`
	ProofDataExtensions []string `json:"proofDataExtensions,omitempty"`
}

type Arbitrate struct {
	ArbitrateResult bool `json:"arbitrateResult,omitempty"`
}
