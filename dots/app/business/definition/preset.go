// Scry Info.  All rights reserved.
// license that can be found in the license file.

package definition

type AccInfo struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type SDKInitData struct {
	FromBlock float64 `json:"fromBlock"`
}

type PublishData struct {
	Price         float64 `json:"price"`
	SupportVerify bool    `json:"supportVerify"`
	Password      string  `json:"password"`
	IDs           IDs     `json:"IDs"`
}
type IDs struct {
	MetaDataID   string   `json:"metaDataID"`
	ProofDataIDs []string `json:"proofDataIDs"`
	DetailsID    string   `json:"detailsID"`
}

type BuyData struct {
	Password     string       `json:"password"`
	StartVerify  bool         `json:"startVerify"`
	SelectedData SelectedData `json:"pID"`
}
type SelectedData struct {
	PublishID string  `json:"PublishID"`
	Price     float64 `json:"Price"`
}

type Prepared struct {
	Extensions []string `json:"extensions"`
}

type PurchaseData struct {
	Password   string       `json:"password"`
	SelectedTx SelectedTxPD `json:"tID"`
}
type SelectedTxPD struct {
	TransactionID string `json:"TransactionID"`
}

type ReEncryptData struct {
	Password   string        `json:"password"`
	SelectedTx SelectedTxRED `json:"tID"`
}
type SelectedTxRED struct {
	TransactionID           string `json:"TransactionID"`
	Seller                  string `json:"Seller"`
	MetaDataIDEncWithSeller []byte `json:"MetaDataIDEncWithSeller"`
}

type DecryptData struct {
	Password   string       `json:"password"`
	SelectedTx SelectedTxDD `json:"tID"`
}
type SelectedTxDD struct {
	MetaDataIDEncrypt []byte `json:"MetaDataIDEncrypt"`
	MetaDataExtension string `json:"MetaDataExtension"`
	User              string `json:"User"`
}

type ConfirmData struct {
	Password   string       `json:"password"`
	SelectedTx SelectedTxCD `json:"tID"`
	Truth      bool         `json:"confirmData"`
}
type SelectedTxCD struct {
	TransactionID string `json:"TransactionID"`
}

type RegisterVerifierData struct {
	Password string `json:"password"`
}

type VerifyData struct {
	Password      string `json:"password"`
	TransactionID string `json:"tID"`
	Verify        Verify `json:"verify"`
}
type Verify struct {
	Suggestion bool   `json:"suggestion"`
	Comment    string `json:"comment"`
}

type CreditData struct {
	Password   string        `json:"password"`
	SelectedTx SelectedTxCrD `json:"tID"`
	Credit     Credit        `json:"credit"`
}
type SelectedTxCrD struct {
	TransactionID string `json:"TransactionID"`
}
type Credit struct {
	Verifier1Revert bool    `json:"verifier1Revert"`
	Verifier1Credit float64 `json:"verifier1Credit"`
	Verifier2Revert bool    `json:"verifier2Revert"`
	Verifier2Credit float64 `json:"verifier2Credit"`
}

type ArbitrateData struct {
	Password        string       `json:"password"`
	SelectedTx      SelectedTxAD `json:"tID"`
	ArbitrateResult bool         `json:"arbitrateResult"`
}
type SelectedTxAD struct {
	TransactionId string `json:"TransactionID"`
}

type Balance struct {
	Balance   string
	TimeStamp string
}
