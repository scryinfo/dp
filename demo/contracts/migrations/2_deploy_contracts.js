var ScryToken = artifacts.require("./ScryToken.sol")
var ScryProtocol = artifacts.require("./ScryProtocol.sol");

var tokenContract
module.exports = function(deployer, network, accounts) {
  deployer.deploy(ScryToken).then(function(instance) {
    tokenContract = instance
    return deployer.deploy(ScryProtocol, ScryToken.address);
  }).then (function() {
    tokenContract.transfer(accounts[1], 10000000)
    tokenContract.transfer(accounts[2], 10000000)
    tokenContract.transfer(accounts[3], 10000000)
    return tokenContract.transfer(accounts[4], 10000000)
  }).then(function(result) {
    return tokenContract.balanceOf.call(accounts[0])
  }).then(function(balance) {
    console.log("balance of account2", accounts[0], balance) 
  }) 
  
};
