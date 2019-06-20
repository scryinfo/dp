// Scry Info.  All rights reserved.
// license that can be found in the license file.

let ScryToken = artifacts.require("./ScryToken.sol");
let ScryProtocol = artifacts.require("./ScryProtocol.sol");

let Common = artifacts.require("./lib/common.sol");
let Transaction = artifacts.require("./lib/transaction.sol");
let Verification = artifacts.require("./lib/verification.sol");

let tokenContract;
module.exports = function(deployer, network, accounts) {
    deployer.deploy(Common);
    deployer.link(Common, [Transaction, Verification, ScryProtocol]);

    deployer.deploy(Transaction);
    deployer.link(Transaction, ScryProtocol);
    deployer.deploy(Verification);
    deployer.link(Verification, ScryProtocol);

    deployer.deploy(ScryToken).then(function(instance) {
        tokenContract = instance;
        console.log(tokenContract.address);

        return deployer.deploy(ScryProtocol, tokenContract.address);
    }).then (function(ptl) {
        console.log(ptl.address);
        console.log("interface:", accounts.length);
    });
};
