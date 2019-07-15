var ScryToken = artifacts.require("ScryToken");
var ScryProtocol = artifacts.require("ScryProtocol");

var tokenContract;
module.exports = function(deployer, network, accounts) {
    deployer.deploy(ScryToken).then(function(instance) {
        tokenContract = instance;
        console.log(tokenContract.address);

        return deployer.deploy(ScryProtocol, tokenContract.address);
    }).then (function(ptl) {
        console.log(ptl.address);
        console.log("interface:", accounts.length);
    });
};
