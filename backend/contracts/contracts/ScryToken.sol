// Scry Info.  All rights reserved.
// license that can be found in the license file.

pragma solidity ^0.4.24;

import "zeppelin-solidity/contracts/token/ERC20/StandardToken.sol";


contract ScryToken is StandardToken {
    string public name = "ScryToken";
    string public symbol = "yyy";
    uint8 public decimals = 2;
//    uint256 public INITIAL_SUPPLY = 1000000000;

    constructor () public {
        totalSupply_ = 1000000000;
        balances[msg.sender] = 1000000000;
    }
}