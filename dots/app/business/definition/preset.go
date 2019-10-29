// Scry Info.  All rights reserved.
// license that can be found in the license file.

package definition

type Preset struct {
	PublishId     string `json:"PublishId"`
	TransactionId string `json:"TransactionId"`

	Address             string   `json:"address"`
	Password            string   `json:"password"`
	Price               string  `json:"price"`
	SupportVerify       bool     `json:"supportVerify"`
	StartVerify         bool     `json:"startVerify"`
	MetaDataExtension   string   `json:"MetaDataExtension"`
	ProofDataExtensions []string `json:"ProofDataExtensions"`

	Ids         Ids         `json:"Ids"`
	EncryptedId EncryptedId `json:"encryptedId"`
	Confirm     Confirm     `json:"confirm"`
	Verify      Verify      `json:"verify"`
	Grade       Grade       `json:"grade"`
	Arbitrate   Arbitrate   `json:"arbitrate"`

	Contract ModifyContractParam `json:"modifyContractParam"`
	Balance  Balance

	CurrentUser User
}

type Ids struct {
	MetaDataId   string   `json:"metaDataId"`
	ProofDataIds []string `json:"proofDataIds"`
	DetailsId    string   `json:"detailsId"`
}

type Verify struct {
	Suggestion bool   `json:"suggestion"`
	Comment    string `json:"comment"`
}

type Grade struct {
	Verifier1Revert bool    `json:"verifier1Revert"`
	Verifier1Grade  float64 `json:"verifier1Grade"`
	Verifier2Revert bool    `json:"verifier2Revert"`
	Verifier2Grade  float64 `json:"verifier2Grade"`
}

type Confirm struct {
	Truth bool `json:"confirmResult"`
}

type Balance struct {
	Balance   string
	TimeStamp string
}

type ModifyContractParam struct {
	ParamName  string `json:"paramName"`
	ParamValue string `json:"paramValue"`
}
