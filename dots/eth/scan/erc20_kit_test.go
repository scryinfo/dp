package scan

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestErc20Kit_PackInputForTransfer(t *testing.T) {
	erc, err := NewErc20Kit()
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, erc.Name)
	assert.NotEqual(t, nil, erc.Abi)
	assert.NotEqual(t, nil, erc.Name)
	assert.NotEqual(t, nil, erc.Symbol)
	assert.NotEqual(t, nil, erc.Decimals)
	assert.NotEqual(t, nil, erc.TotalSupply)
	assert.NotEqual(t, nil, erc.BalanceOf)
	assert.NotEqual(t, nil, erc.Transfer)
	assert.NotEqual(t, nil, erc.TransferFrom)
	assert.NotEqual(t, nil, erc.Approve)
	assert.NotEqual(t, nil, erc.Allowance)

	to := common.HexToAddress("0x0123456789012345678901234567890123456789")
	value := big.NewInt(100)
	exData := []byte{1, 2}
	bs, err := erc.PackInputsForTransfer(&to, value, exData)
	assert.Equal(t, nil, err)
	to2, value2, exData2, err := erc.UnPackInputsForTransfer(bs)
	assert.Equal(t, nil, err)
	assert.Equal(t, to, *to2)
	assert.Equal(t, value, value2)
	assert.Equal(t, exData, exData2)

	exData = nil
	bs, err = erc.PackInputsForTransfer(&to, value, exData)
	assert.Equal(t, nil, err)
	to2, value2, exData2, err = erc.UnPackInputsForTransfer(bs)
	assert.Equal(t, nil, err)
	assert.Equal(t, to, *to2)
	assert.Equal(t, value, value2)
	assert.Equal(t, exData, exData2)

}

func TestErc20Kit_PackInputForTransferFrom(t *testing.T) {
	erc, err := NewErc20Kit()
	assert.Equal(t, nil, err)
	from := common.HexToAddress("0x0123456789012345678901234567890123456780")
	to := common.HexToAddress("0x0123456789012345678901234567890123456789")
	value := big.NewInt(100)
	exData := []byte{1, 2}
	bs, err := erc.PackInputsForTransferFrom(&from, &to, value, exData)
	assert.Equal(t, nil, err)
	from2, to2, value2, exData2, err := erc.UnPackInputsForTransferFrom(bs)
	assert.Equal(t, nil, err)
	assert.Equal(t, from, *from2)
	assert.Equal(t, to, *to2)
	assert.Equal(t, value, value2)
	assert.Equal(t, exData, exData2)

	exData = nil
	bs, err = erc.PackInputsForTransferFrom(&from, &to, value, exData)
	assert.Equal(t, nil, err)
	from2, to2, value2, exData2, err = erc.UnPackInputsForTransferFrom(bs)
	assert.Equal(t, nil, err)
	assert.Equal(t, from, *from2)
	assert.Equal(t, to, *to2)
	assert.Equal(t, value, value2)
	assert.Equal(t, exData, exData2)
}
