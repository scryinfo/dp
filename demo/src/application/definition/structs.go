package definition

type AccInfo struct {
	Account  string `json:"account"`
	Password string `json:"password"`
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
	MetaDataExtension   string   `json:"metaDataExtension"`
	ProofDataExtensions []string `json:"proofDataExtensions"`
	Price               string
	PublishID           string
	SupportVerify       bool // not implement.
}

type BuyData struct {
	Password  string `json:"password"`
	PublishID string `json:"pID"`
}

type TransactionDetails struct {
	TransactionID  string
	PublishID      string
	ProofFileNames []string
	TxState        string
}

type PurchaseData struct {
	Password      string `json:"password"`
	TransactionID string `json:"tID"`
}

type PurchaseDetails struct {
	TransactionID           string
	MetaDataIdEncWithSeller string
	TxState                 string
}

type ReEncryptData struct {
	Password   string     `json:"password"`
	SelectedTx SelectedTx `json:"tID"`
}
type SelectedTx struct {
	TransactionID           string `json:"TransactionID"`
	Buyer                   string `json:"Buyer"`
	Seller                  string `json:"Seller"`
	MetaDataIDEncWithSeller string `json:"MetaDataIDEncWithSeller"`
}

type ReEncryptDetails struct {
	TransactionID          string
	MetaDataIdEncWithBuyer string
	TxState                string
}

type DecryptData struct {
	Password   string      `json:"password"`
	SelectedTx SelectedTxD `json:"tID"`
}
type SelectedTxD struct {
	MetaDataIDEncWithBuyer string `json:"MetaDataIDEncWithBuyer"`
	MetaDataExtension      string `json:"MetaDataExtension"`
	Buyer                  string `json:"Buyer"`
}

type ConfirmData struct {
	Password      string `json:"password"`
	TransactionID string `json:"tID"`
	Arbitrate     bool   `json:"startArbitrate"`
}
