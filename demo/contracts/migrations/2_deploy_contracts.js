var ScryProtocol = artifacts.require("./ScryProtocol.sol");

module.exports = function(deployer) {
  deployer.deploy(ScryProtocol);
};
