package definition

type AccInfo struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type Datalist struct {
	ID          string
	Title       string
	Price       int
	Keys        string
	Description string
	Owner       string
}

type Transaction struct {
	Title         string
	TransactionID int
	Seller        string
	Buyer         string
	State         byte
	Verifier1Response string
	Verifier2Response string
	Verifier3Response string
	ArbitrateResult bool
}

type PubData struct {
	MetaData  string   `json:"Data"`
	ProofData []string `json:"Proofs"`
	DespData  string   `json:"Description"`
	Price     float64  `json:"Price"`
	Seller    string   `json:"Owner"`
	Password  string   `json:"Password"`
	SupportVerify bool `json:"supportVerify"`
}

type BuyData struct {
	Buyer string   `json:"buyer"`
	IDs   []string `json:"ids"`
}
