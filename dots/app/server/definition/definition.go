// Scry Info.  All rights reserved.
// license that can be found in the license file.

package definition

import "encoding/json"

// MessageIn unified structure deserialize msg from client
type MessageIn struct {
	Name    string          `json:"Name"`
	Payload json.RawMessage `json:"Payload"`
}

// MessageOut unified structure serialize msg send to client
type MessageOut struct {
	Name    string      `json:"Name"`
	Payload interface{} `json:"Payload,omitempty"`
}

// PresetFunc preset system functions' handler
type PresetFunc = func(*MessageInPayload) (interface{}, error)

// MessageInPayload pre-define structure for deserialize msg from client
type MessageInPayload struct {
	PublishId     string `json:"PublishId"`
	TransactionId string `json:"TransactionId"`

	Address       string `json:"address"`
	Password      string `json:"password"`
	Price         string `json:"price"`
	SupportVerify bool   `json:"supportVerify"`
	StartVerify   bool   `json:"startVerify"`

	Salt string

	Ids       Ids       `json:"Ids"`
	Confirm   Confirm   `json:"confirm"`
	Verify    Verify    `json:"verify"`
	Grade     Grade     `json:"grade"`
	Arbitrate Arbitrate `json:"arbitrate"`

	Nickname ModifyNickname      `json:"modifyNickname"`
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
	Verifier1Revert bool   `json:"verifier1Revert"`
	Verifier1Grade  string `json:"verifier1Grade"`
	Verifier2Revert bool   `json:"verifier2Revert"`
	Verifier2Grade  string `json:"verifier2Grade"`
}

// Arbitrate specific structure for arbitrate func
type Arbitrate struct {
	ArbitrateResult bool `json:"arbitrateResult,omitempty"`
}

// ModifyNickname specific structure for MN func
type ModifyNickname struct {
	Nickname string `json:"nickname"`
}

// ModifyContractParam specific structure for MCP func
type ModifyContractParam struct {
	ParamName  string `json:"paramName"`
	ParamValue string `json:"paramValue"`
}
