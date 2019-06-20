pragma solidity ^0.4.24;

import "./common.sol";

library verification {
    function register(common.Verifiers storage self, common.Configuration storage conf, address applicant) internal {
        self.list[self.list.length] = common.Verifier(applicant, conf.verifierDepositToken, 0, 0, true);
    }

    function chooseVerifiers(common.Verifiers storage self, common.Configuration storage conf) public returns (address[] memory) {
        uint256 len = self.list.length;
        address[] memory chosenVerifiers = new address[](conf.verifierNum);

        for (uint8 i = 0; i < conf.verifierNum; i++) {
            uint256 index = uint256(keccak256(abi.encodePacked(now, msg.sender))) % len;
            common.Verifier memory v = verifiers[index];

            //loop if invalid verifier was chosen until get valid verifier
            address vb = v.addr;
            while (!v.enable || addressExist(chosenVerifiers, v.addr)) {
                v = verifiers[(++index) % len];
                require(v.addr != vb, "Disordered verifiers");
            }

            chosenVerifiers[i] = v.addr;
        }

        return chosenVerifiers;
    }

    function inVerifier(common.Verifiers storage self, address addr) internal view returns (bool exist, uint256 index) {
        for (uint256 i = 0; i < self.list.length; i++) {
            if (self.list[i].addr == addr) {
                exist = true;
                index = i;
                break;
            }
        }

        return ;
    }

    function addressExist(address[] addrArray, address addr) internal view returns (bool exist, uint8 index) {
        for (uint8 i = 0; i < addrArray.length; i++) {
            if (addr == addrArray[i]) {
                exist = true;
                index = i;
                break;
            }
        }

        return ;
    }

    function validVerifier(common.Verifiers storage self, address[] addrArray, address applicant) internal view returns (bool, uint8) {
        uint256 index256;
        {
            bool exist;
            (exist, index256) = inVerifier(self, applicant);
            if (!(exist && self.list[index256].enable)) {
                return (exist, 0);
            }
        }

        bool exist;
        uint8 index8;
        (exist, index8) = addressExist(addrArray, applicant);

        return (exist, index8);
    }
}
