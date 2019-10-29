// Scry Info.  All rights reserved.
// license that can be found in the license file.

package definition

type Callbacks struct {
	PublishId     string `json:",omitempty"`
	TransactionId string `json:",omitempty"`

	Title               string   `json:"Title,omitempty"`
	Keys                string   `json:"Keys,omitempty"`
	Description         string   `json:"Description,omitempty"`
	Seller              string   `json:"Seller,omitempty"`
	Buyer               string   `json:",omitempty,omitempty"`
	Price               string   `json:",omitempty,omitempty"`
	MetaDataExtension   string   `json:"MetaDataExtension,omitempty"`
	ProofDataExtensions []string `json:"ProofDataExtensions,omitempty"`

	SupportVerify    bool   `json:",omitempty"`
	StartVerify      bool   `json:",omitempty"`
	TransactionState string `json:",omitempty"`
	Block            uint64 `json:",omitempty"`

	VerifyResult     VerifyResult     `json:",omitempty"`
	VerifierDisabled VerifierDisabled `json:",omitempty"`
	Arbitrate        Arbitrate        `json:",omitempty"`
	EncryptedId      EncryptedId      `json:",omitempty"`

	CurrentUser User `json:",omitempty"`
}

type VerifyResult struct {
	VerifierIndex    string `json:",omitempty"`
	VerifierResponse string `json:",omitempty"`
}

type VerifierDisabled struct {
	VerifierAddress string `json:",omitempty"`
}
