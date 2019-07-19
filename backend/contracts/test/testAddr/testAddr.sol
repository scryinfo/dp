pragma solidity ^0.4.24;

import "../../contracts/ScryToken.sol";

contract testAddr {
    function () public payable {}

    function getEthBalance() public view returns(uint256) {
        return address(this).balance;
    }

    function approveTransfer(address testContract, uint count, ERC20 token) public returns(bool) {
        return token.approve(testContract, count);
    }
}
