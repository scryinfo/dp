// Scry Info.  All rights reserved.
// license that can be found in the license file.

/**
    Author: Scry R&D
    Date  : 2019.4.28
    Notes :
    1. notice: deployer_keyjson and deployer_password need to be changed to yours before testing.
    2. for test, it is necessary to persistence some files in your IPFS node.
       cd [dir] , ipfs add persistence_data -r , and you can use IDS below directly.

       meta data ID    : QmRJCQ6uYcWppEoz7oogP9mivwvoCnEKFJn7j4fWCZ5aer
        proof data IDs  : QmRhPUCtgFLzxbewyGB2Zzg25kxdowJHrFYodD6Q1AmMKy,
                          QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew
        describe data ID: QmQihohHMHGfwkzaaYEizzEAXy4mehJh12Sr63LdHAt3Ho
    3. an unexpected situation: run this file from start to end, contract will emit a event several times.
       It still make me puzzled. (╯︵╰)
*/

package sdkinterface

import (
    "fmt"
    "github.com/scryinfo/dp/demo/src/application/definition"
    "github.com/scryinfo/dp/demo/src/application/sdkinterface/settings"
    "github.com/scryinfo/dp/demo/src/sdk"
    "github.com/scryinfo/dp/demo/src/sdk/core/ethereum/events"
    "math/big"
    "testing"
    "time"
)

var (
    publishTest            = false
    approveTest            = false
    verifiersChosenTest    = false
    txCreateTest           = false
    purchaseTest           = false
    readyForDownloadTest   = false
    closeTest              = false
    registerAsVerifierTest = false
    voteTest               = false
    verifierDisableTest    = false

    eventName = []string{"DataPublish", "Approval", "VerifiersChosen", "TransactionCreate", "Buy", "ReadyForDownload",
        "TransactionClose", "RegisterVerifier", "Vote", "VerifierDisable"}
    transactionID           string
    metaDataIDEncWithSeller []byte
    metaDataIDEncWithBuyer  []byte

    supportVerify    = false
    needVerify       = false
    confirmDataTruth = true
)

func initialize() error {
    scryinfo, err := settings.LoadSettings()
    err = sdk.Init(scryinfo.Chain.Ethereum.EthNode,
        scryinfo.Services.Keystore,
        scryinfo.Chain.Contracts.ProtocolAddr,
        scryinfo.Chain.Contracts.TokenAddr,
        scryinfo.Services.Ipfs,
        "D:/EnglishRoad/workspace/Go/src/github.com/scryinfo/dp/demo/src/log/scry_sdk.log",
        "SDK interface unit test",
    )
    SetScryInfo(scryinfo)
    SetFromBlock(0)
    return err
}

func TestInitialize(t *testing.T) {
    if err := initialize(); err != nil {
        fmt.Println(err)
        t.Fail()
    }

}

func TestCreateUserWithLogin(t *testing.T) {
    if _, err := CreateUserWithLogin("111"); err != nil {
        fmt.Println(err)
        t.Fail()
    }
}

func TestTransferTokenFromDeployer(t *testing.T) {
    _ = initialize()
    _, _ = CreateUserWithLogin("111")
    err := TransferTokenFromDeployer(big.NewInt(1))
    if err != nil {
        fmt.Println(err)
        t.Fail()
    }
}

func TestUserLogin(t *testing.T) {
    _ = initialize()
    _, _ = CreateUserWithLogin("111")
    login, err := UserLogin(curUser.Account.Address, "111")
    if !login || err != nil {
        fmt.Println(err)
        t.Fail()
    }
}

func TestSubscribeEvents(t *testing.T) {
    _ = initialize()
    _, _ = CreateUserWithLogin("111")
    if err = SubscribeEvents(eventName, onPublish, onApprove, onVerifiersChosen, onTransactionCreate, onPurchase,
        onReadyForDownload, onClose, onRegisterAsVerifier, onVote, onVerifierDisable); err != nil {
        fmt.Println(err)
        t.Fail()
    }
}

func TestUnsubscribeEvents(t *testing.T) {
    _ = initialize()
    _, _ = CreateUserWithLogin("111")
    _ = SubscribeEvents(eventName, onPublish, onApprove, onVerifiersChosen, onTransactionCreate, onPurchase, onReadyForDownload,
        onClose, onRegisterAsVerifier, onVote, onVerifierDisable)
    if err = UnsubscribeEvents(eventName); err != nil {
        fmt.Println(err)
        t.Fail()
    }
}

func TestPublishData(t *testing.T) {
    _ = initialize()
    _, _ = CreateUserWithLogin("111")

    _ = SubscribeEvents([]string{"DataPublish"}, onPublish)

    _, err := PublishData(&definition.PublishData{
        IDs: definition.IDs{
            MetaDataID: "QmRJCQ6uYcWppEoz7oogP9mivwvoCnEKFJn7j4fWCZ5aer",
            ProofDataIDs: []string{"QmRhPUCtgFLzxbewyGB2Zzg25kxdowJHrFYodD6Q1AmMKy",
                "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew"},
            DetailsID: "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew",
        },
        Price:         float64(1000),
        SupportVerify: supportVerify,
        Password:      "111",
    })

    time.Sleep(time.Second * 5)

    if err != nil || !publishTest {
        fmt.Println(err, "|", publishTest)
        t.Fail()
    }
}

func TestApproveTransferForRegisterAsVerifier(t *testing.T) {
    _ = initialize()
    _, _ = CreateUserWithLogin("111")
    _ = TransferTokenFromDeployer(big.NewInt(1000000))

    _ = SubscribeEvents([]string{"Approval"}, onApprove)
    err := ApproveTransferForRegisterAsVerifier("111")

    time.Sleep(time.Second * 5)

    if err != nil || !approveTest {
        fmt.Println(err, "|", approveTest)
        t.Fail()
    }
}

func TestApproveTransferForBuying(t *testing.T) {
    _ = initialize()
    _, _ = CreateUserWithLogin("111")
    _ = TransferTokenFromDeployer(big.NewInt(1000000))

    _ = SubscribeEvents([]string{"Approval"}, onApprove)
    err := ApproveTransferForBuying("111")

    time.Sleep(time.Second * 5)

    if err != nil || !approveTest {
        fmt.Println(err, "|", approveTest)
        t.Fail()
    }
}

func TestCreateTransaction(t *testing.T) {
    _ = initialize()

    _, _ = CreateUserWithLogin("111")
    _ = SubscribeEvents([]string{"DataPublish"}, onPublish)
    publishID, _ := PublishData(&definition.PublishData{
        IDs: definition.IDs{
            MetaDataID: "QmRJCQ6uYcWppEoz7oogP9mivwvoCnEKFJn7j4fWCZ5aer",
            ProofDataIDs: []string{"QmRhPUCtgFLzxbewyGB2Zzg25kxdowJHrFYodD6Q1AmMKy",
                "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew"},
            DetailsID: "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew",
        },
        Price:         float64(1000),
        SupportVerify: supportVerify,
        Password:      "111",
    })

    time.Sleep(time.Second * 5)

    _, _ = CreateUserWithLogin("222")
    _ = SubscribeEvents([]string{"Approval"}, onApprove)
    _ = SubscribeEvents([]string{"TransactionCreate"}, onTransactionCreate)
    _ = TransferTokenFromDeployer(big.NewInt(1600))
    _ = ApproveTransferForBuying("222")

    err := CreateTransaction(publishID, "222", needVerify)

    time.Sleep(time.Second * 5)

    if err != nil || !txCreateTest || (needVerify != verifiersChosenTest) {
        fmt.Println(err, "|", txCreateTest, "|", needVerify != verifiersChosenTest)
        t.Fail()
    }
}

func TestBuy(t *testing.T) {
    _ = initialize()

    _, _ = CreateUserWithLogin("111")
    _ = SubscribeEvents([]string{"DataPublish"}, onPublish)
    publishID, _ := PublishData(&definition.PublishData{
        IDs: definition.IDs{
            MetaDataID: "QmRJCQ6uYcWppEoz7oogP9mivwvoCnEKFJn7j4fWCZ5aer",
            ProofDataIDs: []string{"QmRhPUCtgFLzxbewyGB2Zzg25kxdowJHrFYodD6Q1AmMKy",
                "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew"},
            DetailsID: "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew",
        },
        Price:         float64(1000),
        SupportVerify: supportVerify,
        Password:      "111",
    })

    time.Sleep(time.Second * 5)

    _, _ = CreateUserWithLogin("222")
    _ = SubscribeEvents([]string{"Approval"}, onApprove)
    _ = SubscribeEvents([]string{"TransactionCreate"}, onTransactionCreate)
    _ = SubscribeEvents([]string{"Buy"}, onPurchase)
    _ = TransferTokenFromDeployer(big.NewInt(1600))
    _ = ApproveTransferForBuying("222")

    _ = CreateTransaction(publishID, "222", needVerify)

    time.Sleep(time.Second * 5)

    err := Buy(transactionID, "222")

    time.Sleep(time.Second * 5)

    if err != nil || !purchaseTest {
        fmt.Println(err, "|", purchaseTest)
        t.Fail()
    }
}

func TestSubmitMetaDataIdEncWithBuyer(t *testing.T) {
    _ = initialize()

    seller, _ := CreateUserWithLogin("111")
    _ = SubscribeEvents([]string{"DataPublish"}, onPublish)
    publishID, _ := PublishData(&definition.PublishData{
        IDs: definition.IDs{
            MetaDataID: "QmRJCQ6uYcWppEoz7oogP9mivwvoCnEKFJn7j4fWCZ5aer",
            ProofDataIDs: []string{"QmRhPUCtgFLzxbewyGB2Zzg25kxdowJHrFYodD6Q1AmMKy",
                "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew"},
            DetailsID: "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew",
        },
        Price:         float64(1000),
        SupportVerify: supportVerify,
        Password:      "111",
    })

    time.Sleep(time.Second * 5)

    buyer, _ := CreateUserWithLogin("222")
    _ = SubscribeEvents([]string{"Approval"}, onApprove)
    _ = SubscribeEvents([]string{"TransactionCreate"}, onTransactionCreate)
    _ = SubscribeEvents([]string{"Buy"}, onPurchase)
    _ = TransferTokenFromDeployer(big.NewInt(1600))
    _ = ApproveTransferForBuying("222")

    _ = CreateTransaction(publishID, "222", needVerify)

    time.Sleep(time.Second * 5)

    _ = Buy(transactionID, "222")

    time.Sleep(time.Second * 5)

    _, _ = UserLogin(seller, "111")
    _ = SubscribeEvents([]string{"ReadyForDownload"}, onReadyForDownload)

    err := SubmitMetaDataIdEncWithBuyer(transactionID, "111", seller, buyer, metaDataIDEncWithSeller)

    time.Sleep(time.Second * 5)

    if err != nil || !readyForDownloadTest {
        fmt.Println(err, "|", readyForDownloadTest)
        t.Fail()
    }
}

func TestDecryptAndGetMetaDataFromIPFS(t *testing.T) {
    _ = initialize()

    seller, _ := CreateUserWithLogin("111")
    _ = SubscribeEvents([]string{"DataPublish"}, onPublish)
    publishID, _ := PublishData(&definition.PublishData{
        IDs: definition.IDs{
            MetaDataID: "QmRJCQ6uYcWppEoz7oogP9mivwvoCnEKFJn7j4fWCZ5aer",
            ProofDataIDs: []string{"QmRhPUCtgFLzxbewyGB2Zzg25kxdowJHrFYodD6Q1AmMKy",
                "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew"},
            DetailsID: "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew",
        },
        Price:         float64(1000),
        SupportVerify: supportVerify,
        Password:      "111",
    })

    time.Sleep(time.Second * 5)

    buyer, _ := CreateUserWithLogin("222")
    _ = SubscribeEvents([]string{"Approval"}, onApprove)
    _ = SubscribeEvents([]string{"TransactionCreate"}, onTransactionCreate)
    _ = SubscribeEvents([]string{"Buy"}, onPurchase)
    _ = TransferTokenFromDeployer(big.NewInt(1600))
    _ = ApproveTransferForBuying("222")

    _ = CreateTransaction(publishID, "222", needVerify)

    time.Sleep(time.Second * 5)

    _ = Buy(transactionID, "222")

    time.Sleep(time.Second * 5)

    _, _ = UserLogin(seller, "111")
    _ = SubscribeEvents([]string{"ReadyForDownload"}, onReadyForDownload)

    _ = SubmitMetaDataIdEncWithBuyer(transactionID, "111", seller, buyer, metaDataIDEncWithSeller)

    time.Sleep(time.Second * 5)

    _, _ = UserLogin(buyer, "222")
    _, err := DecryptAndGetMetaDataFromIPFS("222", metaDataIDEncWithBuyer, buyer, ".txt")

    time.Sleep(time.Second * 5)

    if err != nil {
        fmt.Println(err)
        t.Fail()
    }
}

func TestConfirmDataTruth(t *testing.T) {
    _ = initialize()

    seller, _ := CreateUserWithLogin("111")
    _ = SubscribeEvents([]string{"DataPublish"}, onPublish)
    publishID, _ := PublishData(&definition.PublishData{
        IDs: definition.IDs{
            MetaDataID: "QmRJCQ6uYcWppEoz7oogP9mivwvoCnEKFJn7j4fWCZ5aer",
            ProofDataIDs: []string{"QmRhPUCtgFLzxbewyGB2Zzg25kxdowJHrFYodD6Q1AmMKy",
                "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew"},
            DetailsID: "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew",
        },
        Price:         float64(1000),
        SupportVerify: supportVerify,
        Password:      "111",
    })

    time.Sleep(time.Second * 5)

    buyer, _ := CreateUserWithLogin("222")
    _ = SubscribeEvents([]string{"Approval"}, onApprove)
    _ = SubscribeEvents([]string{"TransactionCreate"}, onTransactionCreate)
    _ = SubscribeEvents([]string{"Buy"}, onPurchase)
    _ = TransferTokenFromDeployer(big.NewInt(1600))
    _ = ApproveTransferForBuying("222")

    _ = CreateTransaction(publishID, "222", needVerify)

    time.Sleep(time.Second * 5)

    _ = Buy(transactionID, "222")

    time.Sleep(time.Second * 5)

    _, _ = UserLogin(seller, "111")
    _ = SubscribeEvents([]string{"ReadyForDownload"}, onReadyForDownload)

    _ = SubmitMetaDataIdEncWithBuyer(transactionID, "111", seller, buyer, metaDataIDEncWithSeller)

    time.Sleep(time.Second * 5)

    _, _ = UserLogin(buyer, "222")
    time.Sleep(time.Second * 5)

    err := ConfirmDataTruth(transactionID, "222", confirmDataTruth)

    time.Sleep(time.Second * 5)

    if err != nil {
        fmt.Println(err)
        t.Fail()
    }
}

func TestCancelTransaction(t *testing.T) {
    _ = initialize()
    {
        _, _ = CreateUserWithLogin("111")
        _ = SubscribeEvents([]string{"DataPublish"}, onPublish)
        publishID, _ := PublishData(&definition.PublishData{
            IDs: definition.IDs{
                MetaDataID: "QmRJCQ6uYcWppEoz7oogP9mivwvoCnEKFJn7j4fWCZ5aer",
                ProofDataIDs: []string{"QmRhPUCtgFLzxbewyGB2Zzg25kxdowJHrFYodD6Q1AmMKy",
                    "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew"},
                DetailsID: "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew",
            },
            Price:         float64(1000),
            SupportVerify: supportVerify,
            Password:      "111",
        })

        time.Sleep(time.Second * 5)

        _, _ = CreateUserWithLogin("222")
        _ = SubscribeEvents([]string{"Approval"}, onApprove)
        _ = SubscribeEvents([]string{"TransactionCreate"}, onTransactionCreate)
        _ = SubscribeEvents([]string{"Buy"}, onPurchase)
        _ = TransferTokenFromDeployer(big.NewInt(1600))
        _ = ApproveTransferForBuying("222")

        _ = CreateTransaction(publishID, "222", needVerify)

        time.Sleep(time.Second * 5)
        _ = SubscribeEvents([]string{"TransactionClose"}, onClose)
        err := CancelTransaction(transactionID, "222")
        time.Sleep(time.Second * 5)
        if err != nil || !closeTest {
            fmt.Println(err, "|", closeTest)
            t.Fail()
        } else {
            fmt.Println("Cancel transaction test success, time node: after create. ")
        }
    }
    {
        closeTest = false
        _, _ = CreateUserWithLogin("111")
        _ = SubscribeEvents([]string{"DataPublish"}, onPublish)
        publishID, _ := PublishData(&definition.PublishData{
            IDs: definition.IDs{
                MetaDataID: "QmRJCQ6uYcWppEoz7oogP9mivwvoCnEKFJn7j4fWCZ5aer",
                ProofDataIDs: []string{"QmRhPUCtgFLzxbewyGB2Zzg25kxdowJHrFYodD6Q1AmMKy",
                    "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew"},
                DetailsID: "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew",
            },
            Price:         float64(1000),
            SupportVerify: supportVerify,
            Password:      "111",
        })

        time.Sleep(time.Second * 5)

        _, _ = CreateUserWithLogin("222")
        _ = SubscribeEvents([]string{"Approval"}, onApprove)
        _ = SubscribeEvents([]string{"TransactionCreate"}, onTransactionCreate)
        _ = SubscribeEvents([]string{"Buy"}, onPurchase)
        _ = TransferTokenFromDeployer(big.NewInt(1600))
        _ = ApproveTransferForBuying("222")

        _ = CreateTransaction(publishID, "222", needVerify)

        time.Sleep(time.Second * 5)

        _ = Buy(transactionID, "222")

        time.Sleep(time.Second * 5)
        _ = SubscribeEvents([]string{"TransactionClose"}, onClose)
        err := CancelTransaction(transactionID, "222")
        time.Sleep(time.Second * 5)
        if err != nil || !closeTest {
            fmt.Println(err, "|", closeTest)
            t.Fail()
        } else {
            fmt.Println("Cancel transaction test success, time node: after purchase. ")
        }
    }
    {
        closeTest = false
        seller, _ := CreateUserWithLogin("111")
        _ = SubscribeEvents([]string{"DataPublish"}, onPublish)
        publishID, _ := PublishData(&definition.PublishData{
            IDs: definition.IDs{
                MetaDataID: "QmRJCQ6uYcWppEoz7oogP9mivwvoCnEKFJn7j4fWCZ5aer",
                ProofDataIDs: []string{"QmRhPUCtgFLzxbewyGB2Zzg25kxdowJHrFYodD6Q1AmMKy",
                    "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew"},
                DetailsID: "QmZchYGj9xaoJuuXCdw7DYF7mLN8rDHe7kKQgpS1JAtDew",
            },
            Price:         float64(1000),
            SupportVerify: supportVerify,
            Password:      "111",
        })

        time.Sleep(time.Second * 5)

        buyer, _ := CreateUserWithLogin("222")
        _ = SubscribeEvents([]string{"Approval"}, onApprove)
        _ = SubscribeEvents([]string{"TransactionCreate"}, onTransactionCreate)
        _ = SubscribeEvents([]string{"Buy"}, onPurchase)
        _ = TransferTokenFromDeployer(big.NewInt(1600))
        _ = ApproveTransferForBuying("222")

        _ = CreateTransaction(publishID, "222", needVerify)

        time.Sleep(time.Second * 5)

        _ = Buy(transactionID, "222")

        time.Sleep(time.Second * 5)

        _, _ = UserLogin(seller, "111")
        _ = SubscribeEvents([]string{"ReadyForDownload"}, onReadyForDownload)

        _ = SubmitMetaDataIdEncWithBuyer(transactionID, "111", seller, buyer, metaDataIDEncWithSeller)

        time.Sleep(time.Second * 5)
        _, _ = UserLogin(buyer, "222")
        _ = SubscribeEvents([]string{"TransactionClose"}, onClose)
        err := CancelTransaction(transactionID, "222")
        time.Sleep(time.Second * 5)
        if err == nil || closeTest {
            fmt.Println(err, "|", closeTest)
            t.Fail()
        } else {
            fmt.Println("Cancel transaction test success, time node: after re-encrypt. ")
        }
    }
}

func onPublish(event events.Event) bool {
    publishTest = true
    return true
}

func onApprove(event events.Event) bool {
    approveTest = true
    return true
}

func onVerifiersChosen(event events.Event) bool {
    verifiersChosenTest = true
    return true
}

func onTransactionCreate(event events.Event) bool {
    transactionID = event.Data.Get("transactionId").(*big.Int).String()
    txCreateTest = true
    return true
}

func onPurchase(event events.Event) bool {
    metaDataIDEncWithSeller = event.Data.Get("metaDataIdEncSeller").([]byte)
    purchaseTest = true
    return true
}

func onReadyForDownload(event events.Event) bool {
    metaDataIDEncWithBuyer = event.Data.Get("metaDataIdEncBuyer").([]byte)
    readyForDownloadTest = true
    return true
}

func onClose(event events.Event) bool {
    closeTest = true
    return true
}

func onRegisterAsVerifier(event events.Event) bool {
    registerAsVerifierTest = true
    return true
}

func onVote(event events.Event) bool {
    voteTest = true
    return true
}

func onVerifierDisable(event events.Event) bool {
    verifierDisableTest = true
    return true
}
