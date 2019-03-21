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

type PubDataIDs struct {
	MetaDataID    string   `json:"metaDataID"`
	ProofDataIDs  []string `json:"proofDataIDs"`
	DetailsID     string   `json:"detailsID"`
	Price         float64  `json:"price"`
	SupportVerify bool     `json:"supportVerify"`
	Password      string   `json:"password"`
}

type DataDetails struct {
	Title               string   `json:"Title"`
	Keys                string   `json:"Keys"`
	Description         string   `json:"Description"`
	MetaDataExtension   string   `json:"MetaDataExtension"`
	ProofDataExtensions []string `json:"ProofDataExtensions"`
	Price               string
	PublishID           string
	SupportVerify       bool // not implement.
	Block               uint64
}

type BuyData struct {
	Password  string `json:"password"`
	PublishID string `json:"pID"`
}

type ApproveDetails struct {
	Block uint64
}

type TransactionDetails struct {
	TransactionID  string
	PublishID      string
	ProofFileNames []string
	Buyer          string
	TxState        string
	Block          uint64
}

type PurchaseData struct {
	Password   string       `json:"password"`
	SelectedTx SelectedTxPD `json:"tID"`
}
type SelectedTxPD struct {
	TransactionID string `json:"TransactionID"`
}

type PurchaseDetails struct {
	TransactionID           string
	MetaDataIdEncWithSeller []byte
	TxState                 string
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

type ReEncryptDetails struct {
	TransactionID          string
	MetaDataIdEncWithBuyer []byte
	TxState                string
	Block                  uint64
}

type DecryptData struct {
	Password   string       `json:"password"`
	SelectedTx SelectedTxDD `json:"tID"`
}
type SelectedTxDD struct {
	MetaDataIDEncWithBuyer []byte `json:"MetaDataIDEncWithBuyer"`
	MetaDataExtension      string `json:"MetaDataExtension"`
	Buyer                  string `json:"Buyer"`
}

type ConfirmData struct {
	Password   string       `json:"password"`
	SelectedTx SelectedTxCD `json:"tID"`
	Arbitrate  bool         `json:"startArbitrate"`
}
type SelectedTxCD struct {
	TransactionID string `json:"TransactionID"`
}

type CloseDetails struct {
	TransactionID string
	TxState       string
	Block         uint64
}
