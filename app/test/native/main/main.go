// Scry Info.  All rights reserved.
// license that can be found in the license file.

package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/line"
	"github.com/scryinfo/dp/dots/binary"
	"github.com/scryinfo/dp/dots/binary/scry"
	"github.com/scryinfo/dp/dots/eth/event"
	"github.com/scryinfo/dp/dots/eth/transaction"
	"github.com/scryinfo/scryg/sutils/ssignal"
	"go.uber.org/zap"
	"math/big"
	"os"
	"time"
)

var (
	seller    scry.Client
	buyer     scry.Client
	verifier1 scry.Client
	verifier2 scry.Client
	verifier3 scry.Client

	protocolContractAddr = "0x3420c44090c6a2c444ce85cb914087760ac0a78b"
	clientPassword       = "888888"
	suAddress            = "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8"
	suPassword           = "111111"

	bin   *binary.Binary
	chain scry.ChainWrapper

	publishId                        = ""
	txId                    *big.Int = big.NewInt(0)
	metaDataIdEncWithSeller []byte
)

func main() {
	logger := dot.Logger()
	l, err := line.BuildAndStart(func(l dot.Line) error {
		l.PreAdd(binary.BinTypeLiveWithGrpc()...)
		return nil
	})

	if err != nil {
		dot.Logger().Debugln("Line init failed. ", zap.NamedError("", err))
		return
	}

	d, err := l.ToInjecter().GetByLiveId(dot.LiveId(binary.BinLiveId))
	if err != nil {
		logger.Errorln("load Binary component failed.")
	}

	if bin, ok := d.(*binary.Binary); !ok {
		logger.Errorln("load Binary component failed.", zap.Any("d", d))
	} else {
		chain = bin.ChainWrapper()
	}

	defer line.StopAndDestroy(l, true)

	go Start()

	ssignal.WaitCtrlC(func(s os.Signal) bool {
		return false //quit
	})
}

// Start start test console
func Start() {
	InitUsers()

	time.Sleep(time.Second * 5)

	StartTestingWithoutVerify()
}

// InitUsers init users with eth and token
func InitUsers() {
	var err error
	seller, err = CreateClientWithToken(big.NewInt(10000000), big.NewInt(10000000))
	if err != nil {
		panic(err)
	}

	buyer, err = CreateClientWithToken(big.NewInt(10000000), big.NewInt(10000000))
	if err != nil {
		panic(err)
	}

	//verifier1, err = CreateClientWithToken(big.NewInt(10000000), big.NewInt(10000000))
	//if err != nil {
	//    panic(err)
	//}
	//
	//verifier2, err = CreateClientWithToken(big.NewInt(10000000), big.NewInt(10000000))
	//if err != nil {
	//    panic(err)
	//}
	//
	//verifier3, err = CreateClientWithToken(big.NewInt(10000000), big.NewInt(10000000))
	//if err != nil {
	//    panic(err)
	//}
}

// StartTestingWithoutVerify start test without verify
func StartTestingWithoutVerify() {
	subscribeAllEvents()
	SellerPublishData()
}

// CreateClientWithToken create a scry Client with eth and token
func CreateClientWithToken(token *big.Int, eth *big.Int) (scry.Client, error) {
	client, err := scry.CreateScryClient(clientPassword, chain)
	if err != nil {
		return nil, err
	}

	dot.Logger().Debugln("client:" + client.Account().Addr)

	err = client.TransferEthFrom(
		common.HexToAddress(suAddress),
		suPassword,
		eth,
		chain.Conn(),
	)
	if err != nil {
		return nil, err
	}

	txParam := transaction.TxParams{
		From:     common.HexToAddress(suAddress),
		Password: suPassword,
		Value:    big.NewInt(0),
		Pending:  false,
	}
	err = chain.TransferTokens(
		&txParam,
		common.HexToAddress(client.Account().Addr),
		token,
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// SellerPublishData simulate seller publish data
func SellerPublishData() {
	//publish data
	metaData := []byte("QmcHXkMXwgvZP56tsUJNtcfedojHkqrDsgkC4fbsBM1zre")
	proofData := []string{"QmNrTHX545s7hGfbEVrJuCiMqKVwUWJ4cPwXrAPv3GW5pm", "Qmb7csVP7wGco16XHVNqqRXUE7vQMjuA24QRypnZdkeQMD"}
	despData := "QmQXqZdEwXnWgKpfJmUVaACuVar9R7vpBxtZgddQMTa2UN"

	txParam := transaction.TxParams{
		From:     common.HexToAddress(seller.Account().Addr),
		Password: clientPassword,
		Value:    big.NewInt(0),
		Pending:  false,
	}

	var err error
	publishId, err = chain.Publish(
		&txParam,
		big.NewInt(1000),
		metaData,
		proofData,
		2,
		despData,
		false,
	)

	if err != nil {
		fmt.Println("SellerPublishData:", err)
	}

}

func subscribeAllEvents() {
	_ = seller.SubscribeEvent("DataPublish", onPublish)

	_ = buyer.SubscribeEvent("Approval", onApprovalBuyerTransfer)
	_ = buyer.SubscribeEvent("TransactionCreate", onTransactionCreate)
}

// BuyerApproveTransfer simulate buyer approve contract transfer token from him
func BuyerApproveTransfer() {
	txParam := transaction.TxParams{
		From:     common.HexToAddress(buyer.Account().Addr),
		Password: clientPassword,
		Value:    big.NewInt(0),
		Pending:  false,
	}

	err := chain.ApproveTransfer(&txParam, common.HexToAddress(protocolContractAddr), big.NewInt(1600))
	if err != nil {
		fmt.Println("BuyerApproveTransfer:", err)
	}
}

// PrepareToBuy simulate buyer create a tx
func PrepareToBuy(publishId string) {
	txParam := transaction.TxParams{
		From:     common.HexToAddress(buyer.Account().Addr),
		Password: clientPassword,
		Value:    big.NewInt(0),
		Pending:  false,
	}
	err := chain.AdvancePurchase(&txParam, publishId, false)
	if err != nil {
		fmt.Println("failed to prepareToBuy, error:", err)
	}
}

// Buy simulate buyer confirm to purchase specific data
func Buy(txId *big.Int) {
	txParam := transaction.TxParams{
		From:     common.HexToAddress(buyer.Account().Addr),
		Password: clientPassword,
		Value:    big.NewInt(0),
		Pending:  false,
	}

	err := chain.ConfirmPurchase(&txParam, txId)
	if err != nil {
		fmt.Println("failed to buyData, error:", err)
	}
}

func onApprovalBuyerTransfer(event event.Event) bool {
	fmt.Println("onApprovalBuyerTransfer:", event)

	PrepareToBuy(publishId)
	return true
}

func onTransactionCreate(event event.Event) bool {
	fmt.Println("onTransactionCreated:", event)
	txId = event.Data.Get("transactionId").(*big.Int)
	Buy(txId)

	return true
}

func onPublish(event event.Event) bool {
	fmt.Println("onPublish: ", event)

	publishId = event.Data.Get("publishId").(string)

	BuyerApproveTransfer()

	return true
}
