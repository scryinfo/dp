pragma solidity ^0.4.24;

contract Demo {
//    enum TransactionState {Begin, Created, Voted, Buying, ReadyForDownload, Arbitrating, Payed, Closed}
    
    /*struct DataInfoPublished {
        uint256 price;
        bytes metaDataIdEncSeller;
        bytes32[] proofDataIds;
        uint256 numberOfProof;
        string despDataId;
        address seller;
        bool supportVerify;
        bool used;
    } 

    struct TransactionItem {
        TransactionState state;
        address buyer;
        address seller;
        address[] verifiers;
        bool[] creditGived; 
        string publishId;
        bytes meteDataIdEncBuyer;
        bytes metaDataIdEncSeller;
        uint256 buyerDeposit;
        bool used;
    }*/
    
    struct Arbitrator {
        address addr;
    }

    struct ArbitrateResult {
        bool judge;
    }

    Arbitrator[] arbitrators;
//    mapping (string => DataInfoPublished) mapPublishedData;
//    mapping (uint256 => TransactionItem) mapTransaction;
    mapping (uint256 => mapping(address => ArbitrateResult)) mapArbitrate;
    mapping (uint256 => bool[]) mapCount;

//    event TransactionCreate(string seqNo, uint256 transactionId, string publishId, bytes32[] proofIds, bool supportVerify, address[] users);
//    event TransactionClose(string seqNo, uint256 transactionId, address[] users);
    
    event ArbitratorsChosen(string seqNo,uint256 txId,address[] selectedArbitrators);
    event Arbitrate(string seqNo,uint256 txId,address[] users);
   
    uint256 transactionSeq = 0;
    uint256 encryptedIdLen = 32;

    address token_address = 0x0;
    address owner         = 0x0;
    ERC20   token;

    constructor (address _token) public {
//        require(_token != 0x0);
//
//        owner = msg.sender;
//        token_address = _token;
//        token = ERC20(_token);

        //the first element used for empty usage
        // verifiers[verifiers.length++] = Verifier(0x00, 0, 0, 0, false);
        arbitrators[arbitrators.length++] = Arbitrator(0x00);
    }

    function inCreateTransaction(string seqNo, string publishId) external {
//        //published data info
//        DataInfoPublished memory data = mapPublishedData[publishId];
//        require(data.used, "Publish data does not exist");
//        require((token.balanceOf(msg.sender)) >= data.price, "No enough balance");
//        require(token.transferFrom(msg.sender, address(this), data.price), "Failed to transfer token from caller");
//
//        //create transaction
//        uint txId = getTransactionId();
//        bytes memory metaDataIdEncBuyer = new bytes(encryptedIdLen);

        /*address[] memory selectedVerifiers; 
        bool[] memory creditGived;
        if (data.supportVerify) {
            //choose verifiers randomly
            selectedVerifiers = chooseVerifiers(verifierNum);
            creditGived = new bool[](verifierNum);
            emit VerifiersChosen(seqNo, txId,data.proofDataIds, selectedVerifiers);
        }*/
        address[] memory selectedArbitrators;
        selectedArbitrators = chooseArbitrators();
        emit ArbitratorsChosen(seqNo, txId, selectedArbitrators);

//        address[] memory users = new address[](1);
//        users[0] = msg.sender;
//        emit TransactionCreate(seqNo, txId, publishId, data.proofDataIds, data.supportVerify, users);
//
//        mapTransaction[txId] = TransactionItem(TransactionState.Created, msg.sender, data.seller, selectedVerifiers, creditGived,
//                                                publishId, metaDataIdEncBuyer, data.metaDataIdEncSeller, data.price, true);
    }
    
    function chooseArbitrators() internal view returns (address[] memory) {
        address[] memory chosenArbitrators = new address[](2);
        
        for (uint8 i = 0;i < 2;i++) {
            uint index = getRandomNumber(arbitrators.length);
            chosenArbitrators[i] = arbitrators[index].addr;
        }
        
        return chosenArbitrators;
    }

//    function getTransactionId() internal returns(uint) {
//        return transactionSeq++;
//    }

//    function getRandomNumber(uint mod) internal view returns (uint) {
//        return uint(keccak256(now, msg.sender)) % mod;
//    }
    
    function arbitrate(string seqNo,uint txId,bool judge) external {
        TransactionItem storage txItem = mapTransaction[txId];
        DataInfoPublished storage data = mapPublishedData[txItem.publishId];
        require(txItem.used, "Transaction does not exist");
        
        mapArbitrate[txId][msg.sender] = ArbitrateResult(judge);
        mapCount[txId]++;

        txItem.state = TransactionState.Arbitrating;
        
        if (mapCount[txId].length == 3) {
            uint memory truth;
            for (uint i = 0;i < 3;i++) {
                if (mapArbitrate[txId][i]) {
                    truth++;
                }
            }
            delete mapCount[txId];
            if (truth > 1) {
                payToSeller(txItem, data);
            }else {
                // reward arbitrators.
            }
            address[] memory users = new address[](1);
            users[0] = txItem.buyer;
            emit Arbitrate(seqNo, txId, users);
        }
    }

    /*function confirmDataTruth(string seqNo, uint256 txId, bool truth) external {
        //validate
        TransactionItem storage txItem = mapTransaction[txId];
        require(txItem.used, "Transaction does not exist");
        require(txItem.buyer == msg.sender, "Invalid buyer");

        DataInfoPublished storage data = mapPublishedData[txItem.publishId];
        require(data.used, "Publish data does not exist");

        require(txItem.state == TransactionState.ReadyForDownload);
        if (!data.supportVerify) {
            if(truth) {
                payToSeller(txItem, data);
            }
            
            closeTransaction(txItem, seqNo, txId);
        } else {
            if (truth) {
                payToSeller(txItem, data);
                closeTransaction(txItem, seqNo, txId);
            } else {
                
            }
        }
    }*/
}