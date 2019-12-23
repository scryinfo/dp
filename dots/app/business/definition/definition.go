// Scry Info.  All rights reserved.
// license that can be found in the license file.

package definition

// Preset pre-define structure for deserialize msg from client
type Preset struct {
	PublishId     string `json:"PublishId"`
	TransactionId string `json:"TransactionId"`

	Address       string `json:"address"`
	Password      string `json:"password"`
	Price         string `json:"price"`
	SupportVerify bool   `json:"supportVerify"`
	StartVerify   bool   `json:"startVerify"`

	Ids       Ids       `json:"Ids"`
	Confirm   Confirm   `json:"confirm"`
	Verify    Verify    `json:"verify"`
	Grade     Grade     `json:"grade"`
	Arbitrate Arbitrate `json:"arbitrate"`

	Contract ModifyContractParam `json:"modifyContractParam"`
}

// Ids specific structure for publish func
type Ids struct {
	MetaDataId   string   `json:"metaDataId"`
	ProofDataIds []string `json:"proofDataIds"`
	DetailsId    string   `json:"detailsId"`
}

// Confirm specific structure for confirm func
type Confirm struct {
	Truth bool `json:"confirmResult"`
}

// Verify specific structure for vote func
type Verify struct {
	Suggestion bool   `json:"suggestion"`
	Comment    string `json:"comment"`
}

// Grade specific structure for grade func
type Grade struct {
	Verifier1Revert bool    `json:"verifier1Revert"`
	Verifier1Grade  float64 `json:"verifier1Grade"`
	Verifier2Revert bool    `json:"verifier2Revert"`
	Verifier2Grade  float64 `json:"verifier2Grade"`
}

// Arbitrate specific structure for arbitrate func
type Arbitrate struct {
	ArbitrateResult bool `json:"arbitrateResult,omitempty"`
}

// ModifyContractParam specific structure for MCP func
type ModifyContractParam struct {
	ParamName  string `json:"paramName"`
	ParamValue string `json:"paramValue"`
}
