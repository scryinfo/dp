package definition

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

type OnPurchase struct {
    TransactionID           string
    PublishID               string
    MetaDataIdEncWithSeller []byte
    TxState                 string
    Block                   uint64
}

type OnReadyForDownload struct {
    TransactionID          string
    MetaDataIdEncWithBuyer []byte
    TxState                string
    Block                  uint64
}

type OnClose struct {
    TransactionID string
    TxState       string
    Block         uint64
}

type OnRegisterAsVerifier struct {
    Block uint64
}

type OnVote struct {
    TransactionID    string
    VerifierResponse string
    VerifierIndex    string
    TxState          string
    Block            uint64
}

type OnVerifierDisable struct {
    Verifier string
    Block    uint64
}

type OnArbitrationBegin struct {
    TransactionId               string
    PublishId                   string
    ProofFileNames              []string
    MetaDataIdEncWithArbitrator []byte
    Block                       uint64
}

type OnArbitrationResult struct {
    TransactionId   string
    ArbitrateResult string
    Block           uint64
}
