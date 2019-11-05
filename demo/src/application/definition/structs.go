// Scry Info.  All rights reserved.
// license that can be found in the license file.

package definition

import "encoding/json"

// MessageIn
type MessageIn struct {
	Name    string          `json:"Name"`
	Payload json.RawMessage `json:"Payload"`
}

// MessageOut
type MessageOut struct {
	Name    string      `json:"Name"`
	Payload interface{} `json:"Payload,omitempty"`
}

// PresetFunc
type PresetFunc = func(*MessageIn) (interface{}, error)

// AccInfo
type AccInfo struct {
	Account  string `json:"interface"`
	Password string `json:"password"`
}

// SDKInitData
type SDKInitData struct {
	FromBlock float64 `json:"fromBlock"`
}

// PublishData
type PublishData struct {
	Price         float64 `json:"price"`
	SupportVerify bool    `json:"supportVerify"`
	Password      string  `json:"password"`
	IDs           IDs     `json:"IDs"`
}
// IDs
type IDs struct {
	MetaDataID   string   `json:"metaDataID"`
	ProofDataIDs []string `json:"proofDataIDs"`
	DetailsID    string   `json:"detailsID"`
}

// OnPublish
type OnPublish struct {
	Title               string   `json:"Title"`
	Keys                string   `json:"Keys"`
	Description         string   `json:"Description"`
	MetaDataExtension   string   `json:"MetaDataExtension"`
	ProofDataExtensions []string `json:"ProofDataExtensions"`
	Seller              string   `json:"Seller"`
	Price               string
	PublishID           string
	SupportVerify       bool
	Block               uint64
}

// BuyData
type BuyData struct {
	Password     string       `json:"password"`
	StartVerify  bool         `json:"startVerify"`
	SelectedData SelectedData `json:"pID"`
}
// SelectedData
type SelectedData struct {
	PublishID string `json:"PublishID"`
}

// OnApprove
type OnApprove struct {
	Block uint64
}

// OnVerifiersChosen
type OnVerifiersChosen struct {
	TransactionID  string
	PublishID      string
	ProofFileNames []string
	TxState        string
	Block          uint64
}

// OnTransactionCreate
type OnTransactionCreate struct {
	TransactionID  string
	PublishID      string
	ProofFileNames []string
	Buyer          string
	StartVerify    bool
	TxState        string
	Block          uint64
}

// Prepared
type Prepared struct {
	Extensions []string `json:"extensions"`
}

// PurchaseData
type PurchaseData struct {
	Password   string       `json:"password"`
	SelectedTx SelectedTxPD `json:"tID"`
}
// SelectedTxPD
type SelectedTxPD struct {
	TransactionID string `json:"TransactionID"`
}

// OnPurchase
type OnPurchase struct {
	TransactionID           string
	PublishID               string
	MetaDataIdEncWithSeller []byte
	TxState                 string
	UserIndex               string
	Buyer                   string // temp
	Block                   uint64
}

// ReEncryptData
type ReEncryptData struct {
	Password   string        `json:"password"`
	SelectedTx SelectedTxRED `json:"tID"`
}
// SelectedTxRED
type SelectedTxRED struct {
	TransactionID           string `json:"TransactionID"`
	Buyer                   string `json:"Buyer"`
	Seller                  string `json:"Seller"`
	MetaDataIDEncWithSeller []byte `json:"MetaDataIDEncWithSeller"`
}

// OnReadyForDownload
type OnReadyForDownload struct {
	TransactionID          string
	MetaDataIdEncWithBuyer []byte
	UserIndex              string
	TxState                string
	Block                  uint64
}

// DecryptData
type DecryptData struct {
	Password   string       `json:"password"`
	SelectedTx SelectedTxDD `json:"tID"`
}
// SelectedTxDD
type SelectedTxDD struct {
	MetaDataIDEncrypt []byte `json:"MetaDataIDEncrypt"`
	MetaDataExtension string `json:"MetaDataExtension"`
	User              string `json:"User"`
}

// ConfirmData
type ConfirmData struct {
	Password   string       `json:"password"`
	SelectedTx SelectedTxCD `json:"tID"`
	Truth      bool         `json:"confirmData"`
}
// SelectedTxCD
type SelectedTxCD struct {
	TransactionID string `json:"TransactionID"`
}

// OnClose
type OnClose struct {
	TransactionID string
	UserIndex     string
	TxState       string
	Block         uint64
}

// RegisterVerifierData
type RegisterVerifierData struct {
	Password string `json:"password"`
}
// OnRegisterAsVerifier
type OnRegisterAsVerifier struct {
	Block uint64
}

// VerifyData
type VerifyData struct {
	Password      string `json:"password"`
	TransactionID string `json:"tID"`
	Verify        Verify `json:"verify"`
}
// Verify
type Verify struct {
	Suggestion bool   `json:"suggestion"`
	Comment    string `json:"comment"`
}

// OnVote
type OnVote struct {
	TransactionID    string
	VerifierResponse string
	VerifierIndex    string
	TxState          string
	Block            uint64
}

// CreditData
type CreditData struct {
	Password   string        `json:"password"`
	SelectedTx SelectedTxCrD `json:"tID"`
	Credit     Credit        `json:"credit"`
}
// SelectedTxCrD
type SelectedTxCrD struct {
	TransactionID string `json:"TransactionID"`
}
// Credit
type Credit struct {
	Verifier1Revert bool    `json:"verifier1Revert"`
	Verifier1Credit float64 `json:"verifier1Credit"`
	Verifier2Revert bool    `json:"verifier2Revert"`
	Verifier2Credit float64 `json:"verifier2Credit"`
}

// OnVerifierDisable
type OnVerifierDisable struct {
	Block uint64
}
