// Scry Info.  All rights reserved.
// license that can be found in the license file.

let Migrations = artifacts.require("./Migrations.sol");

module.exports = function(deployer) {
  deployer.deploy(Migrations);
};
