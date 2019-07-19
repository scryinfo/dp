pragma solidity ^0.4.24;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";

import "./testAddr/testAddr.sol";

import "../contracts/ScryToken.sol";

// a sample for testing function including token transfer.
// application scenarios: write the whole process test with solidity, this sample will give out some help about addresses.
// some examples:
//   1.transfer eth to a contract address, an anonymous function with key word 'payable' is necessary in target contract.
//   2.extends ScryToken, solidity copy codes in ScryToken to TestSample contract, (especially functions)
//     so 'transfer', 'transferFrom' and so on below, just likes function call within the same contract,
//     which means that param 'msg.sender' in 'transfer' is also extra address rather than contract address.
//   3.require reason string should be less than 32 bytes(chars), but in test contracts, whatever works.
//   4.sample has passed test with truffle + ganache.
contract TestSample is ScryToken {
    uint public initialBalance = 10 ether; // contract will get 10 eth when deployed but we do not use it unfortunately.
    uint public ethBalance = 1 ether;
    uint public tokenBalance = 10000;
    uint public tokenTotal = 1000000000;

    address testAddress;

    testAddr testContract;

    ERC20 token = this;
    address owner;

    constructor() public {
        testAddress = DeployedAddresses.testAddr();

        testContract = testAddr(testAddress);

        owner = msg.sender;
    }

    function testSimulateTransfer() public {
        transferETH(testAddress, ethBalance);
        require(testContract.getEthBalance() == ethBalance, "Error Node: seller.ethBalance. ");

        transfer(testAddress, tokenBalance);
        require(balanceOf(testAddress) == tokenBalance, "Error Node: buyer balance, at start. ");

        testContract.approveTransfer(owner, 2000, token);
        require(allowance(testAddress, owner) == 2000, "Error Node: buyer approve. ");

        transferFrom(testAddress, owner, 1000);
        require(balanceOf(testAddress) == 9000 &&
            balanceOf(owner) == tokenTotal - 9000 &&
            allowance(testAddress, owner) == 1000,
            "Error Node: buyer balance, after first transfer. ");
    }

    function transferETH(address _to, uint _amount) public payable {
        _to.transfer(_amount);
    }
}
