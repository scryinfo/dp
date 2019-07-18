// Scry Info.  All rights reserved.
// license that can be found in the license file.

let ScryToken = artifacts.require("ScryToken");
let ScryProtocol = artifacts.require("ScryProtocol");

let Common = artifacts.require("common");
let Transaction = artifacts.require("transaction");
let Verification = artifacts.require("verification");

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
        console.log("> token   : ", tokenContract.address);

        return deployer.deploy(ScryProtocol, tokenContract.address);
    }).then (function(ptl) {
        console.log("> protocol: ", ptl.address);
        console.log("> accounts: ", accounts.length);
    });
};
