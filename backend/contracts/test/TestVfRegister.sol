pragma solidity ^0.4.24;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/lib/verification.sol";
import "../contracts/lib/common.sol";
import "../contracts/ScryToken.sol";

contract TestVfRegister is ScryToken {
    common.DataSet ds;

    ERC20 token = this;

    function testRegister() public {
        prepareParamsForRegister();

        verification.register(ds, "seqNo", token);

        checkParamsForRegister();

        bool r = address(this).call(abi.encodePacked(this.reRegister.selector));
        Assert.isFalse(r, "register twice should not be allowed");
    }

    function reRegister() public {
        verification.register(ds, "seqNo", token);
    }

    function prepareParamsForRegister() internal {
        approve(address(this), 10000);
        require(allowance(msg.sender, address(this)) == 10000, "extra address approve test contract transfer token failed. ");

        ds.verifiers.list[ds.verifiers.list.length++] = common.Verifier(0x00, 0, 0, 0, false);
        ds.conf = common.Configuration(2, 10000, 300,   1, 0, 500,   0, 5, 2,   0, 32);
    }

    function checkParamsForRegister() internal view {
        require(balanceOf(address(this)) == 10000, "balance of verifier1 is wrong. ");
        require(allowance(msg.sender, address(this)) == 0, "allowance of extra address - test contract is wrong. ");

        common.Verifier memory v1 = verification.getVerifier(ds.verifiers, msg.sender);

        require(v1.addr == msg.sender, "verifier1 not found. ");
        require(v1.enable, "verifier1 is invalid. ");
        require(ds.verifiers.validVerifierCount == 1, "valid verifier count is wrong. ");
    }
}
