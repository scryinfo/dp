var scryProtocol = artifacts.require("./ScryProtocol.sol")
var scryToken = artifacts.require("./ScryToken.sol")

var ptl
var ste
contract('ScryProtocol', function(accounts) {
    it("Event Publish should be watched", function() {
        return scryProtocol.deployed().then(function(instance){
            ptl = instance
            return ptl.publishDataInfo("publishId", 1000, "0", ["1","2"], "2", true);
        }).then(function(result){
            for (var i = 0; i < result.logs.length; i++) {
                var log = result.logs[i];
            
                if (log.event == "Publish") {
                  console.log("Event Published watched")
                  break;
                }
              }
        }).catch(function(err){
            assert.fail()
            console.log("catched error:", err)
        })
    })
    /*it("Event TransactionCreate should be watched", function() {
        return scryToken.deployed().then(function(instance){
            ste = instance;
        }).then(function(){
            return scryProtocol.deployed(ste.address).then(function(instance){
                ptl = instance;
            })
        }).then(function(){
            return ste.transfer(ptl.address, 1000);
        }).then(function(result){
            console.log("result:", result)
            return ptl.publishDataInfo("publishId", 1000, "0", ["1","2"], "2", true);            
        }).then(() => {
            return ste.balanceOf.call(ste.address).then(bal => {
              console.info(`balance: ${bal}`);
            });
        }).then(function(result){
            return ptl.prepareToBuy("publishId");
        }).then(function(result){
            for (var i = 0; i < result.logs.length; i++) {
                var log = result.logs[i];
            
                if (log.event == "TransactionCreate") {
                  console.log("Event TransactionCreate watched")
                  break;
                }
              }
        }).catch(function(err){
            assert.fail()
            console.log("catched error:", err)
        })
    
    })*/   
})
