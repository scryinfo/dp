package scan

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/scryinfo/dot/dot"
	"go.uber.org/zap"
	"math/big"
	"strings"
)

const ScryTokenABI = "[{\"constant\":true,\"inputs\":[],\"Name\":\"Name\",\"outputs\":[{\"Name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"Name\":\"_spender\",\"type\":\"address\"},{\"Name\":\"_value\",\"type\":\"uint256\"}],\"Name\":\"approve\",\"outputs\":[{\"Name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"Name\":\"totalSupply\",\"outputs\":[{\"Name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"Name\":\"_from\",\"type\":\"address\"},{\"Name\":\"_to\",\"type\":\"address\"},{\"Name\":\"_value\",\"type\":\"uint256\"}],\"Name\":\"transferFrom\",\"outputs\":[{\"Name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"Name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"Name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"Name\":\"decimals\",\"outputs\":[{\"Name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"Name\":\"_spender\",\"type\":\"address\"},{\"Name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"Name\":\"decreaseApproval\",\"outputs\":[{\"Name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"Name\":\"_owner\",\"type\":\"address\"}],\"Name\":\"balanceOf\",\"outputs\":[{\"Name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"Name\":\"Symbol\",\"outputs\":[{\"Name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"Name\":\"_to\",\"type\":\"address\"},{\"Name\":\"_value\",\"type\":\"uint256\"}],\"Name\":\"transfer\",\"outputs\":[{\"Name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"Name\":\"_spender\",\"type\":\"address\"},{\"Name\":\"_addedValue\",\"type\":\"uint256\"}],\"Name\":\"increaseApproval\",\"outputs\":[{\"Name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"Name\":\"_owner\",\"type\":\"address\"},{\"Name\":\"_spender\",\"type\":\"address\"}],\"Name\":\"allowance\",\"outputs\":[{\"Name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"Name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"Name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"Name\":\"value\",\"type\":\"uint256\"}],\"Name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"Name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"Name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"Name\":\"value\",\"type\":\"uint256\"}],\"Name\":\"Transfer\",\"type\":\"event\"}]"

//Erc20Kit some kits for erc20
type Erc20Kit struct {
	Abi          *abi.ABI
	Name         *abi.Method //name of erc20
	Symbol       *abi.Method //symbol of erc20
	Decimals     *abi.Method //decimals of erc20
	TotalSupply  *abi.Method //totalSupply of erc20
	BalanceOf    *abi.Method //balance of erc20
	Transfer     *abi.Method //transfer of erc20
	TransferFrom *abi.Method //transfer from of erc20
	Approve      *abi.Method //approve of erc20
	Allowance    *abi.Method //allowance of erc20
}

//NewErc20Kit new Erc20Kit
func NewErc20Kit() (*Erc20Kit, error) {
	erc := &Erc20Kit{}
	temp, err := abi.JSON(strings.NewReader(ScryTokenABI))
	if err != nil {
		erc.Abi = nil
		return nil, err
	}
	erc.Abi = &temp

	erc.Name = &abi.Method{}
	*erc.Name = erc.Abi.Methods["Name"]
	erc.Symbol = &abi.Method{}
	*erc.Symbol = erc.Abi.Methods["Symbol"]
	erc.Decimals = &abi.Method{}
	*erc.Decimals = erc.Abi.Methods["decimals"]
	erc.TotalSupply = &abi.Method{}
	*erc.TotalSupply = erc.Abi.Methods["totalSupply"]
	erc.BalanceOf = &abi.Method{}
	*erc.BalanceOf = erc.Abi.Methods["balanceOf"]
	erc.Transfer = &abi.Method{}
	*erc.Transfer = erc.Abi.Methods["transfer"]
	erc.TransferFrom = &abi.Method{}
	*erc.TransferFrom = erc.Abi.Methods["transferFrom"]
	erc.Approve = &abi.Method{}
	*erc.Approve = erc.Abi.Methods["approve"]
	erc.Allowance = &abi.Method{}
	*erc.Allowance = erc.Abi.Methods["allowance"]

	return erc, nil
}

//PackInputs the exData will be appended to end of the transaction
func PackInputs(m *abi.Method, exData []byte, param ...interface{}) ([]byte, error) {
	data, err := m.Inputs.Pack(param...)
	if err != nil {
		return nil, err
	}
	data = append(m.ID(), data...)
	if len(exData) > 0 {
		data = append(data, exData...)
	}
	return data, nil
}

//UnPackInputs
func UnPackInputs(m *abi.Method, inputData []byte, inputs interface{}) (exData []byte, err error) {
	encodedData := inputData[4:] //去掉方法名
	err = m.Inputs.Unpack(inputs, encodedData)
	if err != nil {
		dot.Logger().Errorln("Erc20Kit", zap.Error(err))
		return nil, err
	} else {
		exData = encodedData[m.Inputs.LengthNonIndexed()*32:]
		if len(exData) < 1 {
			exData = nil
		}
		return exData, nil
	}
}

//GenerateTransfer
func (c *Erc20Kit) GenerateTransfer(to *common.Address, value *big.Int, exData []byte, nonce uint64, price *big.Int, limit uint64) (*types.Transaction, error) {
	data, err := c.PackInputsForTransfer(to, value, exData)
	if err != nil {
		return nil, err
	}
	return types.NewTransaction(nonce, *to, nil, limit, price, data), nil
}

//PackInputsForTransfer
func (c *Erc20Kit) PackInputsForTransfer(to *common.Address, value *big.Int, exData []byte) ([]byte, error) {
	return PackInputs(c.Transfer, exData, to, value)
}

//UnPackInputsForTransfer
func (c *Erc20Kit) UnPackInputsForTransfer(inputData []byte) (to *common.Address, value *big.Int, exData []byte, err error) {
	type Inputs struct {
		To    common.Address
		Value *big.Int
	}
	inputs := &Inputs{}
	exData, err = UnPackInputs(c.Transfer, inputData, inputs)
	if err != nil {
		return nil, nil, nil, err
	} else {
		return &inputs.To, inputs.Value, exData, nil
	}
}

//GenerateTransfer
func (c *Erc20Kit) GenerateTransferFrom(from *common.Address, to *common.Address, value *big.Int, exData []byte, nonce uint64, price *big.Int, limit uint64) (*types.Transaction, error) {
	data, err := c.PackInputsForTransferFrom(from, to, value, exData)
	if err != nil {
		return nil, err
	}
	return types.NewTransaction(nonce, *to, nil, limit, price, data), nil
}

//PackInputsForTransferFrom
func (c *Erc20Kit) PackInputsForTransferFrom(from *common.Address, to *common.Address, value *big.Int, exData []byte) ([]byte, error) {
	return PackInputs(c.TransferFrom, exData, from, to, value)
}

//UnPackInputsForTransferFrom
func (c *Erc20Kit) UnPackInputsForTransferFrom(inputData []byte) (from *common.Address, to *common.Address, value *big.Int, exData []byte, err error) {
	type Inputs struct {
		From  common.Address
		To    common.Address
		Value *big.Int
	}
	inputs := &Inputs{}
	exData, err = UnPackInputs(c.TransferFrom, inputData, inputs)
	if err != nil {
		return nil, nil, nil, nil, err
	} else {
		return &inputs.From, &inputs.To, inputs.Value, exData, nil
	}
}
