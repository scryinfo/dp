var scryProtocol = artifacts.require("./ScryProtocol.sol")
var scryToken = artifacts.require("./ScryToken.sol")

var ptl
var ste
contract('ScryProtocol', function(accounts) {
    it("Event Publish should be watched", function() {
        return scryProtocol.deployed().then(function(instance){
            ptl = instance
            return ptl.publishDataInfo("seqno", "publishId", 1000, "0", ["1","2"], "2", true);
        }).then(function(result){
            console.log(result)

            for (var i = 0; i < result.logs.length; i++) {
                var log = result.logs[i];
        
                if (log.event == "DataPublish") {
                    console.log("Event DataPublish watched");
                    break;
                }
            }
        }).catch(function(err){
            assert.fail()
            console.log("catched error:", err)
        })
    })/*
    it("Event TransactionCreate should be watched", function() {
        return scryToken.deployed().then(function(instance){
            ste = instance;
        }).then(function(){
            return scryProtocol.deployed(ste.address).then(function(instance){
                ptl = instance;
            })
        }).then(function() {
            return ptl.publishDataInfo("seqno1", "publishId", 1000, "0", ["1","2"], "2", true);
        }).then(function(){
            return ste.approve(ptl.address, 1000);
        }).then(function() {
            console.log("Start creating transaction...");
            return ptl.createTransaction("seqno2", "publishId");
        }).then(function(result) {
            console.log("result", result);
            for (var i = 0; i < result.logs.length; i++) {
                var log = result.logs[i];
        
                if (log.event == "TransactionCreate") {
                    console.log("Event TransactionCreate watched");
                    break;
                }
            }
        }).catch(function(err){
            assert.fail()
            console.log("catched error:", err)
        })
    
    })*/
})

/*
.then(function(){
    return ste.approve(ptl.address, 1000);
}).then(function(result){
    console.log(result)

    for (var i = 0; i < result.logs.length; i++) {
        var log = result.logs[i];

        if (log.event == "Approval") {
            console.log("Event Approval watched");
            break;
        }
    }
})*/