var ScryProtocol = artifacts.require("./ScryProtocol.sol")

var ptl
contract('ScryProtocol', function(accounts) {
    it("Event Publish should be watched", function() {
        return ScryProtocol.deployed().then(function(instance){
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
    it("Event TransactionCreate should be watched", function() {
        return ScryProtocol.deployed().then(function(instance){
            ptl = instance
            return ptl.publishDataInfo("publishId", 1000, "0", ["1","2"], "2", true);
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
    })   
})
