// Code generated by MockGen. DO NOT EDIT.
// Source: chain_wrapper.go

// Package mock_scry is a generated GoMock package.
package mock_scry

import (
	common "github.com/ethereum/go-ethereum/common"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	gomock "github.com/golang/mock/gomock"
	transaction "github.com/scryinfo/dp/dots/eth/transaction"
	big "math/big"
	reflect "reflect"
)

// MockChainWrapper is a mock of ChainWrapper interface
type MockChainWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockChainWrapperMockRecorder
}

// MockChainWrapperMockRecorder is the mock recorder for MockChainWrapper
type MockChainWrapperMockRecorder struct {
	mock *MockChainWrapper
}

// NewMockChainWrapper creates a new mock instance
func NewMockChainWrapper(ctrl *gomock.Controller) *MockChainWrapper {
	mock := &MockChainWrapper{ctrl: ctrl}
	mock.recorder = &MockChainWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChainWrapper) EXPECT() *MockChainWrapperMockRecorder {
	return m.recorder
}

// Conn mocks base method
func (m *MockChainWrapper) Conn() *ethclient.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Conn")
	ret0, _ := ret[0].(*ethclient.Client)
	return ret0
}

// Conn indicates an expected call of Conn
func (mr *MockChainWrapperMockRecorder) Conn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Conn", reflect.TypeOf((*MockChainWrapper)(nil).Conn))
}

// Publish mocks base method
func (m *MockChainWrapper) Publish(txParams *transaction.TxParams, price *big.Int, metaDataID []byte, proofDataIDs []string, proofNum int32, detailsID string, supportVerify bool) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", txParams, price, metaDataID, proofDataIDs, proofNum, detailsID, supportVerify)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish
func (mr *MockChainWrapperMockRecorder) Publish(txParams, price, metaDataID, proofDataIDs, proofNum, detailsID, supportVerify interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockChainWrapper)(nil).Publish), txParams, price, metaDataID, proofDataIDs, proofNum, detailsID, supportVerify)
}

// AdvancePurchase mocks base method
func (m *MockChainWrapper) AdvancePurchase(txParams *transaction.TxParams, publishId string, startVerify bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdvancePurchase", txParams, publishId, startVerify)
	ret0, _ := ret[0].(error)
	return ret0
}

// AdvancePurchase indicates an expected call of AdvancePurchase
func (mr *MockChainWrapperMockRecorder) AdvancePurchase(txParams, publishId, startVerify interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdvancePurchase", reflect.TypeOf((*MockChainWrapper)(nil).AdvancePurchase), txParams, publishId, startVerify)
}

// ConfirmPurchase mocks base method
func (m *MockChainWrapper) ConfirmPurchase(txParams *transaction.TxParams, txId *big.Int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfirmPurchase", txParams, txId)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfirmPurchase indicates an expected call of ConfirmPurchase
func (mr *MockChainWrapperMockRecorder) ConfirmPurchase(txParams, txId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmPurchase", reflect.TypeOf((*MockChainWrapper)(nil).ConfirmPurchase), txParams, txId)
}

// CancelPurchase mocks base method
func (m *MockChainWrapper) CancelPurchase(txParams *transaction.TxParams, txId *big.Int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelPurchase", txParams, txId)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelPurchase indicates an expected call of CancelPurchase
func (mr *MockChainWrapperMockRecorder) CancelPurchase(txParams, txId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelPurchase", reflect.TypeOf((*MockChainWrapper)(nil).CancelPurchase), txParams, txId)
}

// ReEncrypt mocks base method
func (m *MockChainWrapper) ReEncrypt(txParams *transaction.TxParams, txId *big.Int, encodedData []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReEncrypt", txParams, txId, encodedData)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReEncrypt indicates an expected call of ReEncrypt
func (mr *MockChainWrapperMockRecorder) ReEncrypt(txParams, txId, encodedData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReEncrypt", reflect.TypeOf((*MockChainWrapper)(nil).ReEncrypt), txParams, txId, encodedData)
}

// ConfirmData mocks base method
func (m *MockChainWrapper) ConfirmData(txParams *transaction.TxParams, txId *big.Int, truth bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfirmData", txParams, txId, truth)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfirmData indicates an expected call of ConfirmData
func (mr *MockChainWrapperMockRecorder) ConfirmData(txParams, txId, truth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmData", reflect.TypeOf((*MockChainWrapper)(nil).ConfirmData), txParams, txId, truth)
}

// ApproveTransfer mocks base method
func (m *MockChainWrapper) ApproveTransfer(txParams *transaction.TxParams, spender common.Address, value *big.Int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApproveTransfer", txParams, spender, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApproveTransfer indicates an expected call of ApproveTransfer
func (mr *MockChainWrapperMockRecorder) ApproveTransfer(txParams, spender, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApproveTransfer", reflect.TypeOf((*MockChainWrapper)(nil).ApproveTransfer), txParams, spender, value)
}

// Vote mocks base method
func (m *MockChainWrapper) Vote(txParams *transaction.TxParams, txId *big.Int, judge bool, comments string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Vote", txParams, txId, judge, comments)
	ret0, _ := ret[0].(error)
	return ret0
}

// Vote indicates an expected call of Vote
func (mr *MockChainWrapperMockRecorder) Vote(txParams, txId, judge, comments interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vote", reflect.TypeOf((*MockChainWrapper)(nil).Vote), txParams, txId, judge, comments)
}

// RegisterAsVerifier mocks base method
func (m *MockChainWrapper) RegisterAsVerifier(txParams *transaction.TxParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAsVerifier", txParams)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterAsVerifier indicates an expected call of RegisterAsVerifier
func (mr *MockChainWrapperMockRecorder) RegisterAsVerifier(txParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAsVerifier", reflect.TypeOf((*MockChainWrapper)(nil).RegisterAsVerifier), txParams)
}

// GradeToVerifier mocks base method
func (m *MockChainWrapper) GradeToVerifier(txParams *transaction.TxParams, txId *big.Int, index, credit uint8) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GradeToVerifier", txParams, txId, index, credit)
	ret0, _ := ret[0].(error)
	return ret0
}

// GradeToVerifier indicates an expected call of GradeToVerifier
func (mr *MockChainWrapperMockRecorder) GradeToVerifier(txParams, txId, index, credit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GradeToVerifier", reflect.TypeOf((*MockChainWrapper)(nil).GradeToVerifier), txParams, txId, index, credit)
}

// Arbitrate mocks base method
func (m *MockChainWrapper) Arbitrate(txParams *transaction.TxParams, txId *big.Int, judge bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Arbitrate", txParams, txId, judge)
	ret0, _ := ret[0].(error)
	return ret0
}

// Arbitrate indicates an expected call of Arbitrate
func (mr *MockChainWrapperMockRecorder) Arbitrate(txParams, txId, judge interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Arbitrate", reflect.TypeOf((*MockChainWrapper)(nil).Arbitrate), txParams, txId, judge)
}

// GetBuyer mocks base method
func (m *MockChainWrapper) GetBuyer(txParams *transaction.TxParams, txId *big.Int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBuyer", txParams, txId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBuyer indicates an expected call of GetBuyer
func (mr *MockChainWrapperMockRecorder) GetBuyer(txParams, txId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBuyer", reflect.TypeOf((*MockChainWrapper)(nil).GetBuyer), txParams, txId)
}

// GetArbitrators mocks base method
func (m *MockChainWrapper) GetArbitrators(txParams *transaction.TxParams, txId *big.Int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArbitrators", txParams, txId)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArbitrators indicates an expected call of GetArbitrators
func (mr *MockChainWrapperMockRecorder) GetArbitrators(txParams, txId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArbitrators", reflect.TypeOf((*MockChainWrapper)(nil).GetArbitrators), txParams, txId)
}

// TransferTokens mocks base method
func (m *MockChainWrapper) TransferTokens(txParams *transaction.TxParams, to common.Address, value *big.Int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTokens", txParams, to, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransferTokens indicates an expected call of TransferTokens
func (mr *MockChainWrapperMockRecorder) TransferTokens(txParams, to, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTokens", reflect.TypeOf((*MockChainWrapper)(nil).TransferTokens), txParams, to, value)
}

// GetTokenBalance mocks base method
func (m *MockChainWrapper) GetTokenBalance(txParams *transaction.TxParams, owner common.Address) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenBalance", txParams, owner)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTokenBalance indicates an expected call of GetTokenBalance
func (mr *MockChainWrapperMockRecorder) GetTokenBalance(txParams, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenBalance", reflect.TypeOf((*MockChainWrapper)(nil).GetTokenBalance), txParams, owner)
}

// ModifyContractParam mocks base method
func (m *MockChainWrapper) ModifyContractParam(txParams *transaction.TxParams, paramName, paramValue string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyContractParam", txParams, paramName, paramValue)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyContractParam indicates an expected call of ModifyContractParam
func (mr *MockChainWrapperMockRecorder) ModifyContractParam(txParams, paramName, paramValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyContractParam", reflect.TypeOf((*MockChainWrapper)(nil).ModifyContractParam), txParams, paramName, paramValue)
}
