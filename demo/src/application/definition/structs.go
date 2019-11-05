// Scry Info.  All rights reserved.
// license that can be found in the license file.

package definition

import "encoding/json"

// MessageIn deserialize msg received from client
type MessageIn struct {
	Name    string          `json:"Name"`
	Payload json.RawMessage `json:"Payload"`
}

// MessageOut serialize msg send to client
type MessageOut struct {
	Name    string      `json:"Name"`
	Payload interface{} `json:"Payload,omitempty"`
}

// PresetFunc type define
type PresetFunc = func(*MessageIn) (interface{}, error)

// AccInfo acc
type AccInfo struct {
	Account  string `json:"interface"`
	Password string `json:"password"`
}

// SDKInitData from block
type SDKInitData struct {
	FromBlock float64 `json:"fromBlock"`
}

// PublishData publish data
type PublishData struct {
	Price         float64 `json:"price"`
	SupportVerify bool    `json:"supportVerify"`
	Password      string  `json:"password"`
	IDs           IDs     `json:"IDs"`
}

// IDs publish ids
type IDs struct {
	MetaDataID   string   `json:"metaDataID"`
	ProofDataIDs []string `json:"proofDataIDs"`
	DetailsID    string   `json:"detailsID"`
}

// OnPublish publish event
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

// BuyData advance purchase
type BuyData struct {
	Password     string       `json:"password"`
	StartVerify  bool         `json:"startVerify"`
	SelectedData SelectedData `json:"pID"`
}

// SelectedData selected data
type SelectedData struct {
	PublishID string `json:"PublishID"`
}

// OnApprove on approve
type OnApprove struct {
	Block uint64
}

// OnVerifiersChosen on verifiers chosen
type OnVerifiersChosen struct {
	TransactionID  string
	PublishID      string
	ProofFileNames []string
	TxState        string
	Block          uint64
}

// OnTransactionCreate on tc
type OnTransactionCreate struct {
	TransactionID  string
	PublishID      string
	ProofFileNames []string
	Buyer          string
	StartVerify    bool
	TxState        string
	Block          uint64
}

// Prepared deserialize proof file extensions
type Prepared struct {
	Extensions []string `json:"extensions"`
}

// PurchaseData confirm purchase
type PurchaseData struct {
	Password   string       `json:"password"`
	SelectedTx SelectedTxPD `json:"tID"`
}

// SelectedTxPD selected tx
type SelectedTxPD struct {
	TransactionID string `json:"TransactionID"`
}

// OnPurchase on purchase
type OnPurchase struct {
	TransactionID           string
	PublishID               string
	MetaDataIdEncWithSeller []byte
	TxState                 string
	UserIndex               string
	Buyer                   string // temp
	Block                   uint64
}

// ReEncryptData re-encrypt data
type ReEncryptData struct {
	Password   string        `json:"password"`
	SelectedTx SelectedTxRED `json:"tID"`
}

// SelectedTxRED selected tx
type SelectedTxRED struct {
	TransactionID           string `json:"TransactionID"`
	Buyer                   string `json:"Buyer"`
	Seller                  string `json:"Seller"`
	MetaDataIDEncWithSeller []byte `json:"MetaDataIDEncWithSeller"`
}

// OnReadyForDownload on RFD
type OnReadyForDownload struct {
	TransactionID          string
	MetaDataIdEncWithBuyer []byte
	UserIndex              string
	TxState                string
	Block                  uint64
}

// DecryptData decrypt data
type DecryptData struct {
	Password   string       `json:"password"`
	SelectedTx SelectedTxDD `json:"tID"`
}

// SelectedTxDD selected tx
type SelectedTxDD struct {
	MetaDataIDEncrypt []byte `json:"MetaDataIDEncrypt"`
	MetaDataExtension string `json:"MetaDataExtension"`
	User              string `json:"User"`
}

// ConfirmData confirm data
type ConfirmData struct {
	Password   string       `json:"password"`
	SelectedTx SelectedTxCD `json:"tID"`
	Truth      bool         `json:"confirmData"`
}

// SelectedTxCD selected tx
type SelectedTxCD struct {
	TransactionID string `json:"TransactionID"`
}

// OnClose on close
type OnClose struct {
	TransactionID string
	UserIndex     string
	TxState       string
	Block         uint64
}

// RegisterVerifierData register
type RegisterVerifierData struct {
	Password string `json:"password"`
}

// OnRegisterAsVerifier on register
type OnRegisterAsVerifier struct {
	Block uint64
}

// VerifyData verify
type VerifyData struct {
	Password      string `json:"password"`
	TransactionID string `json:"tID"`
	Verify        Verify `json:"verify"`
}

// Verify verify
type Verify struct {
	Suggestion bool   `json:"suggestion"`
	Comment    string `json:"comment"`
}

// OnVote on vote
type OnVote struct {
	TransactionID    string
	VerifierResponse string
	VerifierIndex    string
	TxState          string
	Block            uint64
}

// CreditData credit
type CreditData struct {
	Password   string        `json:"password"`
	SelectedTx SelectedTxCrD `json:"tID"`
	Credit     Credit        `json:"credit"`
}

// SelectedTxCrD selected tx
type SelectedTxCrD struct {
	TransactionID string `json:"TransactionID"`
}

// Credit credit
type Credit struct {
	Verifier1Revert bool    `json:"verifier1Revert"`
	Verifier1Credit float64 `json:"verifier1Credit"`
	Verifier2Revert bool    `json:"verifier2Revert"`
	Verifier2Credit float64 `json:"verifier2Credit"`
}

// OnVerifierDisable on verifier disabled
type OnVerifierDisable struct {
	Block uint64
}
