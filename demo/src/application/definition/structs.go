package definition

type AccInfo struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type SDKInitData struct {
	FromBlock float64 `json:"fromBlock"`
}

//type PubData struct {
//	MetaData      string   `json:"Data"`
//	ProofData     []string `json:"Proofs"`
//	DespData      string   `json:"Description"`
//	Price         float64  `json:"Price"`
//	Seller        string   `json:"Owner"`
//	Password      string   `json:"Password"`
//	SupportVerify bool     `json:"SupportVerify"`
//}

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

type BuyData struct {
	Password     string       `json:"password"`
	StartVerify  bool         `json:"startVerify"`
	SelectedData SelectedData `json:"pID"`
}
type SelectedData struct {
	PublishID string `json:"PublishID"`
}

type OnApprove struct {
	Block uint64
}

type OnVerifiersChosen struct {
	TransactionID  string
	PublishID      string
	ProofFileNames []string
	TxState        string
	Block          uint64
}

type OnTransactionCreate struct {
	TransactionID  string
	PublishID      string
	ProofFileNames []string
	Buyer          string
	StartVerify    bool
	TxState        string
	Block          uint64
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

type OnPurchase struct {
	TransactionID           string
	PublishID               string
	MetaDataIdEncWithSeller []byte
	TxState                 string
	UserIndex               string
	Block                   uint64
}

type ReEncryptData struct {
	Password   string        `json:"password"`
	SelectedTx SelectedTxRED `json:"tID"`
}
type SelectedTxRED struct {
	TransactionID           string `json:"TransactionID"`
	Buyer                   string `json:"Buyer"`
	Seller                  string `json:"Seller"`
	MetaDataIDEncWithSeller []byte `json:"MetaDataIDEncWithSeller"`
}

type OnReadyForDownload struct {
	TransactionID          string
	MetaDataIdEncWithBuyer []byte
	UserIndex              string
	TxState                string
	Block                  uint64
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

type OnClose struct {
	TransactionID string
	UserIndex     string
	TxState       string
	Block         uint64
}

type RegisterVerifierData struct {
	Password string `json:"password"`
}
type OnRegisterAsVerifier struct {
	Block uint64
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

type OnVote struct {
	TransactionID    string
	VerifierResponse string
	VerifierIndex    string
	TxState          string
	Block            uint64
}

type CreditData struct {
	Password      string        `json:"password"`
	TransactionID string `json:"tID"`
	Credit        Credit        `json:"credit"`
}
type Credit struct {
	Verifier1Revert bool    `json:"verifier1Revert"`
	Verifier1Credit float64 `json:"verifier1Credit"`
	Verifier2Revert bool    `json:"verifier2Revert"`
	Verifier2Credit float64 `json:"verifier2Credit"`
}

type OnVerifierDisable struct {
	Block uint64
}
