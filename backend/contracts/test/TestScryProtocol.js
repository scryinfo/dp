// Scry Info.  All rights reserved.
// license that can be found in the license file.

let scryProtocol = artifacts.require("ScryProtocol");
let scryToken = artifacts.require("ScryToken");

let ptl, ste;
let deployer, seller, buyer, verifier1, verifier2, verifier3, verifier4, verifierSelected, arbitratorSelected;
let publishId, txId, password;
contract('ScryProtocol', async accounts => {

    before(function() {
        InitUsers();
        InitContracts();
    });

    it("Normal procedure with verifier", async () => {
        await timeout(1000);
        let r = await ptl.publishDataInfo("seqno", "publishId", 1000, "0", ["1", "2"], "2", true, {from: seller});
        assert(checkEvent("DataPublish", r), "failed to watch event DataPublish");

        r = await ste.approve(ptl.address, 10000, {from: verifier1});
        assert(checkEvent("Approval", r), "no Approval event watched");

        //register verifier
        r = await ptl.registerAsVerifier("seqno1", {from: verifier1});
        assert(checkEvent("RegisterVerifier", r), "failed to watch event RegisterVerifier");

        r = await ste.approve(ptl.address, 10000, {from: verifier2});
        assert(checkEvent("Approval", r), "no Approval event watched");

        r = await ptl.registerAsVerifier("seqno1", {from: verifier2});
        assert(checkEvent("RegisterVerifier", r), "failed to watch event RegisterVerifier");

        r = await ste.approve(ptl.address, 10000, {from: verifier3});
        assert(checkEvent("Approval", r), "no Approval event watched");

        r = await ptl.registerAsVerifier("seqno1", {from: verifier3});
        assert(checkEvent("RegisterVerifier", r), "failed to watch event RegisterVerifier");

        r = await ste.approve(ptl.address, 10000, {from: verifier4});
        assert(checkEvent("Approval", r), "no Approval event watched");

        r = await ptl.registerAsVerifier("seqno1", {from: verifier4});
        assert(checkEvent("RegisterVerifier", r), "failed to watch event RegisterVerifier");

        r = await ste.approve(ptl.address, 2100, {from: buyer});
        assert(checkEvent("Approval", r), "no Approval event watched");

        r = await ptl.createTransaction("seqno3",  "publishId", true, {from: buyer});
        assert(checkEvent("VerifiersChosen", r), "failed to watch event VerifiersChosen");
        assert(checkEvent("TransactionCreate", r), "failed to watch event TransactionCreate");

        verifierSelected = getEventField("VerifiersChosen", r, "users");
        console.log("> verifiers:", verifierSelected);

        txId = getEventField("TransactionCreate", r, "transactionId");
        console.log("> txId:", txId);

        r = await ptl.vote("seqNo4", txId, true, "comments from verifier1", {from: verifierSelected[0]});
        assert(checkEvent("Vote", r), "failed to watch event Vote");

        r = await ptl.buyData("seqNo5", txId, {from: buyer});
        assert(checkEvent("Buy", r), "failed to watch event Buy");

        r = await ptl.reEncryptMetaDataIdBySeller("seqNo6", txId, "0", "0", {from: seller});
        assert(checkEvent("ReadyForDownload", r), "failed to watch event ReadyForDownload");

        r = await ptl.confirmDataTruth("seqNO7", txId, false, {from: buyer});
        assert(checkEvent("ArbitrationBegin", r), "failed to watch event ArbitrationBegin");

        arbitratorSelected = getEventField("ArbitrationBegin", r, "users");
        console.log("> arbitrators:", arbitratorSelected);

        r = await ptl.arbitrate("seqN08", txId, true, {from: arbitratorSelected[0]});
        assert(checkEvent("ArbitrationResult", r), "failed to watch event ArbitrationResult");

        assert(checkEvent("TransactionClose", r), "failed to watch event TransactionClose");

        r = await ptl.creditsToVerifier("seqNO9", txId, 0, 1, {from: buyer});
        assert(checkEvent("VerifierDisable", r), "failed to watch event VerifierDisable");
    });

    function InitContracts() {
        return new Promise(function() {
            scryToken.deployed().then(function (instance) {
                ste = instance;
                console.log("> ste:", ste.address);
                ste.transfer(seller, 10000);
                ste.transfer(buyer, 40000);
                ste.transfer(verifier1, 13000);
                ste.transfer(verifier2, 13000);
                ste.transfer(verifier3, 13000);
                ste.transfer(verifier4, 13000);
            }).then(function() {
                scryProtocol.deployed().then(function (instance) {
                    ptl = instance;
                    console.log("> ptl:", ptl.address);
                })
            })
        })
    }

    function InitUsers() {
        password = "111111";
        deployer = accounts[0];

        seller = web3.personal.newAccount(password);
        web3.personal.unlockAccount(seller, password);
        web3.eth.sendTransaction({
                from: deployer,
                to: seller,
                value: 1672197500000000000
            }, function(err, transactionHash) {
                if (err) {
                    console.log(transactionHash, "error", err);
                }
        });

        buyer = web3.personal.newAccount(password);
        web3.personal.unlockAccount(buyer, password);
        web3.eth.sendTransaction({
                from: deployer,
                to: buyer,
                value: 1672197500000000000
            }, function(err, transactionHash) {
                if (err) {
                    console.log(transactionHash, "error", err);
                }
        });

        verifier1 = web3.personal.newAccount(password);
        web3.personal.unlockAccount(verifier1, password);
        web3.eth.sendTransaction({
            from: deployer,
            to: verifier1,
            value: 1672197500000000000
        }, function(err, transactionHash) {
            if (err) {
                console.log(transactionHash, "error", err);
            }
        });

        verifier2 = web3.personal.newAccount(password);
        web3.personal.unlockAccount(verifier2, password);
        web3.eth.sendTransaction({
            from: deployer,
            to: verifier2,
            value: 1672197500000000000
        }, function(err, transactionHash) {
            if (err) {
                console.log(transactionHash, "error", err);
            }
        });

        verifier3 = web3.personal.newAccount(password);
        web3.personal.unlockAccount(verifier3, password);
        web3.eth.sendTransaction({
            from: deployer,
            to: verifier3,
            value: 1672197500000000000
        }, function(err, transactionHash) {
            if (err) {
                console.log(transactionHash, "error", err);
            }
        });

        verifier4 = web3.personal.newAccount(password);
        web3.personal.unlockAccount(verifier4, password);
        web3.eth.sendTransaction({
            from: deployer,
            to: verifier4,
            value: 1672197500000000000
        }, function(err, transactionHash) {
            if (err) {
                console.log(transactionHash, "error", err);
            }
        });
    }

    function timeout(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }
});

function checkEvent(eventName, receipt) {
    //console.log("event:", eventName, " receipt:", receipt);
    for (let i = 0; i < receipt.logs.length; i++) {
        let log = receipt.logs[i];
        if (log.event === eventName) {
            console.log("Event " + eventName + " watched");
            return true
        }
    }
}

function getEventField(eventName, receipt, fieldName) {
    for (let i = 0; i < receipt.logs.length; i++) {
        let log = receipt.logs[i];

        if (log.event === eventName) {
            return log.args[fieldName];
        }
    }
}
