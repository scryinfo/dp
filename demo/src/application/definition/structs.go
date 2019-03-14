package definition

const (
	Created = byte(iota)
	Voted
	Buying
	ReadyForDownload
	Arbitrating
	Payed
	Closed
)

type AccInfo struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type Datalist struct {
	PublishID     string
	Title         string
	Price         float64
	Keys          string
	Description   string
	SupportVerify bool
}

type Transaction struct {
	Title             string
	TransactionID     float64
	PublishID         string
	Price             float64
	Seller            string
	Buyer             string
	State             byte
	Verifier1Response string
	Verifier2Response string
	ArbitrateResult   bool
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
	Price               float64
	SupportVerify       bool	// not implement.
}

type BuyData struct {
	Password string `json:"password"`
	IDs      string `json:"ids"`
}

type PurchaseData struct {
	Password      string  `json:"password"`
	TransactionID float64 `json:"ids"`
}

type ReEncryptData struct {
	Password   string     `json:"password"`
	SelectedTx SelectedTx `json:"ids"`
}
type SelectedTx struct {
	TransactionID           float64 `json:"ID"`
	Buyer                   string  `json:"Buyer"`
	Seller                  string  `json:"Seller"`
	MetaDataIDEncWithSeller string  `json:"MetaDataIDEncWithSeller"`
}

type DecryptData struct {
	Password               string `json:"password"`
	MetaDataIDEncWithBuyer string `json:"metaDataIDEncWithBuyer"`
	Buyer                  string `json:"buyer"`
}

type ConfirmData struct {
	Password      string  `json:"password"`
	TransactionID float64 `json:"ids"`
	Arbitrate     bool    `json:"startArbitrate"`
}
