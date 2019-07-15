import {grpc} from "@improbable-eng/grpc-web";
import {BinaryService} from "../binary_pb_service";
import {TxParams, ClientInfo, CreateAccountParams, TransferEthParams, TransferTokenParams, PublishParams, SubscribeInfo} from "../binary_pb";

//create account
var deployerAddr = "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8";
var deployerPassword = "111111";

var sellerAddr;
var sellerPassword = "222222";

start();

function start() {
    console.log("start...");

    createAccount();
    console.log("account created");
}

function onCreateAccount() {
    createEventsChannel();
    console.log("streaming channel created");

    subscribeEvent();
    console.log("event subscribed");

    transferEth();
    console.log("eth is transferred");
}

function onTransferEth() {
    transferToken();
    console.log("token is transferred");
}

function onTransferToken() {
    publish();
    console.log("data published");
}

function createAccount() {
    let req = new CreateAccountParams();
    req.setPassword(sellerPassword);
    grpc.unary(BinaryService.CreateAccount, {
        request: req,
        host: "http://localhost:6868",
        onEnd: res => {
            const { status, statusMessage, headers, message, trailers } = res;
            if (status === grpc.Code.OK && message) {
                console.log("all ok: CreateAccount", message.toObject());
                sellerAddr = message.getAccountid();
                onCreateAccount();
            } else {
                console.log("error: CreateAccount", message);
            }

        },
    });
}

function transferEth() {
    let req = new TransferEthParams();
    req.setFrom(deployerAddr);
    req.setPassword(deployerPassword);
    console.log("transfer eth, account:", sellerAddr);
    req.setTo(sellerAddr);
    req.setValue(1000000000000000000);
    grpc.unary(BinaryService.TransferEth, {
        request: req,
        host: "http://localhost:6868",
        onEnd: res => {
            const { status, statusMessage, headers, message, trailers } = res;
            if (status === grpc.Code.OK && message) {
                console.log("all ok.: TransferEth", message.toObject());
                onTransferEth();
            } else {
                console.log("error: TransferEth", message);
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

function transferToken() {
    let req = new TransferTokenParams();
    req.setTxparam(makeTxParams(deployerAddr, deployerPassword));
    req.setTo(sellerAddr);
    req.setValue(10000);
    grpc.unary(BinaryService.TransferTokens, {
        request: req,
        host: "http://localhost:6868",
        onEnd: res => {
            const { status, statusMessage, headers, message, trailers } = res;
            if (status === grpc.Code.OK && message) {
                console.log("all ok.: TransferTokens", message.toObject());
                onTransferToken();
            } else {
                console.log("error: TransferTokens", message);
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
                console.log("all ok.: Publish", message.toObject());
            } else {
                console.log("error: Publish", message);
            }
        },
    });
}

function createEventsChannel() {
    let req = new ClientInfo();
    req.setAddress(sellerAddr);
    req.setPassword(sellerPassword);
    grpc.invoke(BinaryService.RecvEvents, {
        request: req,
        host: "http://localhost:6868",
        onHeaders: ((headers) => {
            console.log("onHeaders", headers);
        }),
        onMessage: ((message) => {
            console.log("onMessage", message);
        }),
        onEnd: ((status, statusMessage, trailers) => {
            console.log("onEnd", status, statusMessage, trailers);
        }),
    });
}

function subscribeEvent() {
    let req = new SubscribeInfo();
    req.setAddress(sellerAddr);
    req.setEvent("DataPublish");

    grpc.unary(BinaryService.SubscribeEvent, {
        request: req,
        host: "http://localhost:6868",
        onEnd: res => {
            const { status, statusMessage, headers, message, trailers } = res;
            if (status === grpc.Code.OK && message) {
                console.log("all ok.: subscribeEvent", message.toObject());
            } else {
                console.log("error: subscribeEvent", message);
            }
        },
    });
}