var ScryProtocol = artifacts.require("./ScryProtocol.sol")

var ptl
contract('ScryProtocol', function(accounts) {
    it("Event Published should be watched", function() {
        return ScryProtocol.deployed().then(function(instance){
            ptl = instance
            return ptl.publishDataInfo("publishId", "0", "1", "2", true);
        }).then(function(result){
            for (var i = 0; i < result.logs.length; i++) {
                var log = result.logs[i];
            
                if (log.event == "Published") {
                  console.log("Event Published watched")
                  break;
                }
              }
        }).catch(function(err){
            assert.fail()
            console.log("catched error:", err)
        })
    })
    it("Published data should be saved into map", function() {
        return ScryProtocol.deployed().then(function(instance){
            ptl = instance
            return ptl.publishDataInfo("publishId", "0", "1", "2", true);
        }).then(function(result){
            return ptl.isPublishedDataExisted.call("publishId")
        }).then(function(result){
            assert.equal(result, true)
        }).catch(function(err){
            assert.fail()
            console.log("catched error:", err)
        })
    })   
})
