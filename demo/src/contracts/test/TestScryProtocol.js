let scryProtocol = artifacts.require("./ScryProtocol.sol")
let scryToken = artifacts.require("./ScryToken.sol")

let ptl, ste
let deployer, seller, buyer, verifier1, verifier2, verifier3, chosenVerfiers
let publishId, txId
contract('ScryProtocol', async accounts => {

    before(function() {
        InitUsers()
        InitContracts()
    })

    it("Normal procedure with verifier", async () => {
        await timeout(60000)

        let r = await ptl.publishDataInfo("seqno", "publishId", 1000, "0", ["1", "2"], "2", true, {from: seller})
        assert(checkEvent("DataPublish", r), "failed to watch event DataPublish")

        r = await ste.approve(ptl.address, 10000, {from: verifier1})
        assert(checkEvent("Approval", r), "no Approval event watched")

        //register verifier
        r = await ptl.registerAsVerifier("seqno1", {from: verifier1})
        assert(checkEvent("RegisterVerifier", r), "failed to watch event RegisterVerifier")

        r = await ste.approve(ptl.address, 10000, {from: verifier2})
        assert(checkEvent("Approval", r), "no Approval event watched")

        r = await ptl.registerAsVerifier("seqno1", {from: verifier2})
        assert(checkEvent("RegisterVerifier", r), "failed to watch event RegisterVerifier")

        r = await ste.approve(ptl.address, 10000, {from: verifier3})
        assert(checkEvent("Approval", r), "no Approval event watched")

        r = await ptl.registerAsVerifier("seqno1", {from: verifier3})
        assert(checkEvent("RegisterVerifier", r), "failed to watch event RegisterVerifier")

        r = await ste.approve(ptl.address, 1600, {from: buyer})
        assert(checkEvent("Approval", r), "no Approval event watched")

        r = await ptl.createTransaction("seqno3",  "publishId", true, {from: buyer})
        assert(checkEvent("TransactionCreate", r), "failed to watch event TransactionCreate")

        verifiers = getEventField("TransactionCreate", r, "verifiers")
        txId = getEventField("TransactionCreate", r, "transactionId")

        console.log("verifiers:", verifiers, ", txId:", txId)

        assert(verifiers.length == 2, "Invalid chosen verifiers")

        r = await ptl.vote("seqNo4", txId, true, "comments from verifier1", {from: verifiers[0]})
        assert(checkEvent("Vote", r), "failed to watch event Vote")

        r = await ptl.buyData("seqNo5", txId, {from: buyer})
        assert(checkEvent("Buy", r), "failed to watch event Buy")

        r = await ptl.submitMetaDataIdEncWithBuyer("seqNo6", txId, "0", {from: seller})
        assert(checkEvent("ReadyForDownload", r), "failed to watch event ReadyForDownload")

        r = await ptl.confirmDataTruth("seqNO7", txId, false, {from: buyer})
        assert(checkEvent("ArbitratingBegin", r), "failed to watch event ArbitratingBegin")

        arbitrators = getEventField("ArbitratingBegin", r, "users");
        r = await ptl.arbitrate("seqNO8", txId, true, {from: arbitrators[0]})
        assert(checkEvent("Payed", r), "failed to watch event Payed")
    })

    function InitContracts() {
        let nc = new Promise(function() {
            scryToken.deployed().then(function (instance) {
                ste = instance
                console.log("ste:", ste.address)
                ste.transfer(seller, 10000)
                ste.transfer(buyer, 30000)
                ste.transfer(verifier1, 13000)
                ste.transfer(verifier2, 13000)
                ste.transfer(verifier3, 13000)
            }).then(function() {
                scryProtocol.deployed().then(function (instance) {
                    ptl = instance
                    console.log("ptl:", ptl.address)
                })
            })
        })

        return nc
    }

    function InitUsers() {
        password = "111111"
        deployer = accounts[0]

        seller = web3.personal.newAccount(password)
        web3.personal.unlockAccount(seller, password)
        web3.eth.sendTransaction({
                from: deployer,
                to: seller,
                value: 16721975000000000000
            }, function(err, transactionHash) {
                if (err)
                    console.log(transactionHash, "error", err);
        })

        buyer = web3.personal.newAccount(password)
        web3.personal.unlockAccount(buyer, password)
        web3.eth.sendTransaction({
                from: deployer,
                to: buyer,
                value: 16721975000000000000
            }, function(err, transactionHash) {
                if (err)
                    console.log(transactionHash, "error", err);
        })

        verifier1 = web3.personal.newAccount(password)
        web3.personal.unlockAccount(verifier1, password)
        web3.eth.sendTransaction({
            from: deployer,
            to: verifier1,
            value: 16721975000000000000
        }, function(err, transactionHash) {
            if (err)
                console.log(transactionHash, "error", err);
        })

        verifier2 = web3.personal.newAccount(password)
        web3.personal.unlockAccount(verifier2, password)
        web3.eth.sendTransaction({
            from: deployer,
            to: verifier2,
            value: 16721975000000000000
        }, function(err, transactionHash) {
            if (err)
                console.log(transactionHash, "error", err);
        })

        verifier3 = web3.personal.newAccount(password)
        web3.personal.unlockAccount(verifier3, password)
        web3.eth.sendTransaction({
            from: deployer,
            to: verifier3,
            value: 16721975000000000000
        }, function(err, transactionHash) {
            if (err)
                console.log(transactionHash, "error", err);
        })
    }

    function timeout(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }
})

function checkEvent(eventName, receipt) {
    for (let i = 0; i < receipt.logs.length; i++) {
        let log = receipt.logs[i]

        if (log.event == eventName) {
            console.log("Event " + eventName + " watched")
            return true
        }
    }
}

function getEventField(eventName, receipt, fieldName) {
    for (let i = 0; i < receipt.logs.length; i++) {
        let log = receipt.logs[i]

        if (log.event == eventName) {
            return log.args[fieldName]
        }
    }
}
