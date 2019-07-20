/**
 * The testing has no GUI.
 * You can observe the testing procedure in browser console (Press f12 if your browser is Chrome/Firefox).
 * prerequisite:
 * a. set the protocol contract address
 * b. server started at local host
 * c. geth started
 */

import {grpc} from "@improbable-eng/grpc-web";
import {BinaryService} from "../binary_pb_service";
import {
    TxParams,
    ClientInfo,
    CreateAccountParams,
    TransferEthParams,
    TransferTokenParams,
    PublishParams,
    SubscribeInfo,
    PrepareParams,
    ApproveTransferParams,
    BuyParams,
    ReEncryptDataParams,
    DataConfirmParams
} from "../binary_pb";

//create account
var deployerAddr = "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8";
var deployerPassword = "111111";

var sellerAddr;
var sellerPassword = "222222";

var buyerAddr;
var buyerPassword = "333333";

//please set the protocol contract address before running this test
var protocolContractAddr = "0x3420c44090c6a2c444ce85cb914087760ac0a78b";

var publishId;

start();

function start() {
    console.log("start...");

    createSeller();
    console.log("account created");
}

function onCreateSeller() {
    authenticateSeller();
    createBuyer();
}

function onCreateBuyer() {
    createEventsChannel(sellerAddr, sellerPassword, dispatchSellerEvent);
    console.log("streaming channel for seller created");

    createEventsChannel(buyerAddr, buyerPassword, dispatchBuyerEvent);
    console.log("streaming channel for buyer created");
}

function onTransferEth() {
    transferToken(buyerAddr);
    console.log("token is transferred");
}

function onTransferToken() {
    publish();
    console.log("data published");
}

function createSeller() {
    let req = new CreateAccountParams();
    req.setPassword(sellerPassword);
    grpc.unary(BinaryService.CreateAccount, {
        request: req,
        host: "http://localhost:6868",
        onEnd: res => {
            const { status, statusMessage, headers, message, trailers } = res;
            if (status === grpc.Code.OK && message) {
                console.log("ok: CreateAccount", message.toObject());
                sellerAddr = message.getAccountid();
                onCreateSeller();
            } else {
                console.log("error: CreateAccount", status, statusMessage, headers, trailers);
            }

        },
    });
}


function createBuyer() {
    let req = new CreateAccountParams();
    req.setPassword(buyerPassword);
    grpc.unary(BinaryService.CreateAccount, {
        request: req,
        host: "http://localhost:6868",
        onEnd: res => {
            const { status, statusMessage, headers, message, trailers } = res;
            if (status === grpc.Code.OK && message) {
                console.log("ok: CreateAccount", message.toObject());
                buyerAddr = message.getAccountid();
                onCreateBuyer();
            } else {
                console.log("error: CreateAccount", status, statusMessage, headers, trailers);
            }
        },
    });
}


function authenticateSeller() {
    let req = new ClientInfo();
    req.setAddress(sellerAddr);
    req.setPassword(sellerPassword);
    grpc.unary(BinaryService.Authenticate, {
        request: req,
        host: "http://localhost:6868",
        onEnd: res => {
            const { status, statusMessage, headers, message, trailers } = res;
            if (status === grpc.Code.OK && message) {
                console.log("ok: Authenticate", message.toObject());
            } else {
                console.log("error: Authenticate", status, statusMessage, headers, trailers);
            }
        },
    });
}


function transferEth(to) {
    let req = new TransferEthParams();
    req.setFrom(deployerAddr);
    req.setPassword(deployerPassword);
    console.log("transfer eth, account:", to);
    req.setTo(to);
    req.setValue(1000000000000000000);
    grpc.unary(BinaryService.TransferEth, {
        request: req,
        host: "http://localhost:6868",
        onEnd: res => {
            const { status, statusMessage, headers, message, trailers } = res;
            if (status === grpc.Code.OK && message) {
                console.log("ok: TransferEth", message.toObject());

                if (to === sellerAddr) {
                    transferEth(buyerAddr);
                }

                if (to === buyerAddr) {
                    onTransferEth();
                }

            } else {
                console.log("error: TransferEth", status, statusMessage, headers, trailers);
            }
        },
    });
}

function makeTxParams(from, password) {
    let p = new TxParams();
    p.setFrom(from);
    p.setPassword(password);
    return p;
}

function transferToken(to) {
    let req = new TransferTokenParams();
    req.setTxparam(makeTxParams(deployerAddr, deployerPassword));
    req.setTo(to);
    req.setValue(10000);
    grpc.unary(BinaryService.TransferTokens, {
        request: req,
        host: "http://localhost:6868",
        onEnd: res => {
            const { status, statusMessage, headers, message, trailers } = res;
            if (status === grpc.Code.OK && message) {
                console.log("ok: TransferTokens", message.toObject());

                if (to === buyerAddr) {
                    onTransferToken();
                }

            } else {
                console.log("error: TransferTokens", status, statusMessage, headers, trailers);
            }
        },
    });
}

function publish() {
    let req = new PublishParams();
    req.setTxparam(makeTxParams(sellerAddr, sellerPassword));
    req.setDetailsid("QmSsw6EcnwEiTT9c4rnAGeSENvsJMepNHmbrgi2S9bXNJr");
    req.setMetadataid("QmSsw6EcnwEiTT9c4rnAGeSENvsJMepNHmbrgi2S9bXNJ1");
    req.setPrice(2000);
    req.setProofdataidsList(["QmSsw6EcnwEiTT9c4rnAGeSENvsJMepNHmbrgi2S9bXNJ3", "QmSsw6EcnwEiTT9c4rnAGeSENvsJMepNHmbrgi2S9bXNJ4"]);
    req.setProofnum(2);
    req.setSupportverify(true);
    grpc.unary(BinaryService.Publish, {
        request: req,
        host: "http://localhost:6868",
        onEnd: res => {
            const { status, statusMessage, headers, message, trailers } = res;
            if (status === grpc.Code.OK && message) {
                console.log("ok: Publish", message.toObject());
            } else {
                console.log("error: Publish", status, statusMessage, headers, trailers);
            }
        },
    });
}

function approveTransferToken() {
    let req = new ApproveTransferParams();
    req.setTxparam(makeTxParams(buyerAddr, buyerPassword));
    req.setSpenderaddr(protocolContractAddr);
    req.setValue(10000);

    grpc.unary(BinaryService.ApproveTransfer, {
        request: req,
        host: "http://localhost:6868",
        onEnd: res => {
            const { status, statusMessage, headers, message, trailers } = res;
            if (status === grpc.Code.OK && message) {
                console.log("ok: ApproveTransferToken", message.toObject());
            } else {
                console.log("error: ApproveTransferToken", status, statusMessage, headers, trailers );
            }
        },
    });
}

function prepareToBuy(publishId) {
    let req = new PrepareParams();
    req.setTxparam(makeTxParams(buyerAddr, buyerPassword));
    req.setPublishid(publishId);
    req.setStartverify(false);
    grpc.unary(BinaryService.PrepareToBuy, {
        request: req,
        host: "http://localhost:6868",
        onEnd: res => {
            const { status, statusMessage, headers, message, trailers } = res;
            if (status === grpc.Code.OK && message) {
                console.log("ok: PrepareToBuy", message.toObject());
            } else {
                console.log("error: PrepareToBuy", status, statusMessage, headers, trailers );
            }
        },
    });
}

function buy(txId) {
    let req = new BuyParams();
    req.setTxparam(makeTxParams(buyerAddr, buyerPassword));
    req.setTxid(txId);
    grpc.unary(BinaryService.BuyData, {
        request: req,
        host: "http://localhost:6868",
        onEnd: res => {
            const { status, statusMessage, headers, message, trailers } = res;
            if (status === grpc.Code.OK && message) {
                console.log("ok: buy", message.toObject());
            } else {
                console.log("error: buy", status, statusMessage, headers, trailers );
            }
        },
    });
}

function submitEncryptedId(txId, encryptedData) {
    let req = new ReEncryptDataParams();
    req.setTxparam(makeTxParams(sellerAddr, sellerPassword));
    req.setTxid(txId);
    req.setEncodeddatawithseller(encryptedData);
    grpc.unary(BinaryService.ReEncryptMetaDataId, {
        request: req,
        host: "http://localhost:6868",
        onEnd: res => {
            const { status, statusMessage, headers, message, trailers } = res;
            if (status === grpc.Code.OK && message) {
                console.log("ok: submitEncryptedId", message.toObject());
            } else {
                console.log("error: submitEncryptedId", status, statusMessage, headers, trailers );
            }
        },
    });
}

function createEventsChannel(addr, password, eventDispatcher) {
    let req = new ClientInfo();
    req.setAddress(addr);
    req.setPassword(password);
    grpc.invoke(BinaryService.RecvEvents, {
        request: req,
        host: "http://localhost:6868",
        onHeaders: ((headers) => {
            console.log("createEventsChannel: onHeaders", addr, headers);
        }),
        onMessage: ((message) => {
            //console.log("onMessage", message);
            eventDispatcher(message);
        }),
        onEnd: ((status, statusMessage, trailers) => {
            console.log("createEventsChannel: onEnd", addr, status, statusMessage, trailers);
        }),
    });
}

function subscribeEvent(addr, events) {
    let req = new SubscribeInfo();
    req.setAddress(addr);
    req.setEventList(events);

    grpc.unary(BinaryService.SubscribeEvent, {
        request: req,
        host: "http://localhost:6868",
        onEnd: res => {
            const { status, statusMessage, headers, message, trailers } = res;
            if (status === grpc.Code.OK && message) {
                console.log("ok: subscribeEvent", message.toObject());

                if (addr === buyerAddr) {
                    transferEth(sellerAddr);
                }

                console.log("eth is transferred");
            } else {
                console.log("error: subscribeEvent ", status, statusMessage, headers, trailers);
            }
        },
    });
}

function unSubscribeEvent(addr, events) {
    let req = new SubscribeInfo();
    req.setAddress(addr);
    req.setEventList(events);

    grpc.unary(BinaryService.UnSubscribeEvent, {
        request: req,
        host: "http://localhost:6868",
        onEnd: res => {
            const { status, statusMessage, headers, message, trailers } = res;
            if (status === grpc.Code.OK && message) {
                console.log("ok: UnSubscribeEvent", message.toObject());

                if (addr === sellerAddr) {
                    unSubscribeEvent(buyerAddr, ["DataPublish", "Approval", "TransactionCreate", "Buy", "ReadyForDownload", "TransactionClose"]);
                }
            } else {
                console.log("error: subscribeEvent ", status, statusMessage, headers, trailers);
            }
        },
    });
}

function confirmDataTruth(txId) {
    let req = new DataConfirmParams();
    req.setTxparam(makeTxParams(buyerAddr, buyerPassword));
    req.setTxid(txId);
    req.setTruth(true);

    grpc.unary(BinaryService.ConfirmDataTruth, {
        request: req,
        host: "http://localhost:6868",
        onEnd: res => {
            const { status, statusMessage, headers, message, trailers } = res;
            if (status === grpc.Code.OK && message) {
                console.log("ok: confirmDataTruth", message.toObject());
            } else {
                console.log("error: confirmDataTruth ", status, statusMessage, headers, trailers);
            }
        },
    });
}

function dispatchSellerEvent(message) {
    let evt = parseEvent(message);
    console.log("event to seller:", evt);

    switch (evt.EventName) {
        case "ChannelCreated":
            subscribeEvent(sellerAddr, ["DataPublish", "Approval", "TransactionCreate", "Buy", "ReadyForDownload", "TransactionClose"]);
            console.log("seller event subscribed");
            break;

        case "DataPublish":
            break;
        case "Buy":
            submitEncryptedId(evt.EventData.transactionId, evt.EventData.metaDataIdEncSeller);
            break;
        case "ReadyForDownload":
            break;
        case "TransactionClose":
            console.log("Seller: Transaction Closed.");
            break;
    }
}

function dispatchBuyerEvent(message) {
    let evt = parseEvent(message);
    console.log("event to buyer:", evt);

    switch (evt.EventName) {
        case "ChannelCreated":
            subscribeEvent(buyerAddr, ["DataPublish", "Approval", "TransactionCreate", "Buy", "ReadyForDownload", "TransactionClose"]);
            console.log("buyer event subscribed");
            break;
        case "DataPublish":
            publishId = evt.EventData.publishId;
            console.log("publishID:", publishId);
            approveTransferToken();
            break;
        case "Approval":
            prepareToBuy(publishId);
            break;
        case "TransactionCreate":
            buy(evt.EventData.transactionId);
            break;
        case "Buy":
            break;
        case "ReadyForDownload":
            confirmDataTruth(evt.EventData.transactionId);
            break;
        case "TransactionClose":
            console.log("Buyer: Transaction Closed.");
            unSubscribeEvent(sellerAddr, ["DataPublish", "Approval", "TransactionCreate", "Buy", "ReadyForDownload", "TransactionClose"]);
            break;
    }
}

function parseEvent(message) {
    let obj = JSON.parse(message.array[1]);
    obj.EventData = JSON.parse(obj.EventData);

    return obj;
}