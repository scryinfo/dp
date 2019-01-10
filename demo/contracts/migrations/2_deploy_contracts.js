var ScryToken = artifacts.require("./ScryToken.sol")
var ScryProtocol = artifacts.require("./ScryProtocol.sol");

module.exports = function(deployer) {
  deployer.deploy(ScryToken).then(function() {
    return deployer.deploy(ScryProtocol, ScryToken.address);
  })
  
};
