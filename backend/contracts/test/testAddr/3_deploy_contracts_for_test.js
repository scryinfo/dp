// Scry Info.  All rights reserved.
// license that can be found in the license file.

let testAddr = artifacts.require("testAddr");

module.exports = function(deployer, network, accounts) {
    deployer.deploy(testAddr);
};