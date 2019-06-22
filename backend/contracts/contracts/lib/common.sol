pragma solidity ^0.4.24;

library common {
    struct DataSet {
        PublishedData pubData;
        TransactionData txData;
        VoteData voteData;
        ArbitratorData arbitratorData;
        Verifiers verifiers;
        Configuration conf;
    }

    enum TransactionState {Begin, Created, Voted, Buying, ReadyForDownload, Closed}

    struct DataInfoPublished {
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
        bool[] creditGiven;
        address[] arbitrators;
        string publishId;
        bytes meteDataIdEncBuyer;
        bytes metaDataIdEncSeller;
        bytes metaDataIdEncArbitrator;
        uint256 buyerDeposit;
        uint256 verifierBonus;
        uint256 arbitratorBonus;
        bool needVerify;
        bool used;
    }

    struct Verifier {
        address addr;
        uint256 deposit;
        uint8 credits;  //credit: 0-5
        uint256 creditTimes;
        bool enable;
    }

    struct VoteResult {
        bool judge;
        string comments;
        bool used;
    }

    struct ArbitratorResult {
        address addr;
        bool judge;
        bool used;
    }

    struct PublishedData {
        mapping (string => DataInfoPublished) map;
    }

    struct TransactionData {
        mapping (uint256 => TransactionItem) map;
    }

    struct VoteData {
        mapping (uint256 => mapping(address => VoteResult)) map;
    }

    struct ArbitratorData {
        mapping (uint256 => mapping(uint8 => ArbitratorResult)) map;
    }

    struct Verifiers {
        Verifier[] list;
        uint8 validVerifierCount;
    }

    struct Configuration {
        uint8 verifierNum;
        uint256 verifierDepositToken;
        uint256 verifierBonus;

        uint8 arbitratorNum;
        uint8 arbitrateCredit;
        uint256 arbitratorBonus;

        uint8 creditLow;
        uint8 creditHigh;
        uint8 creditThreshold;

        uint256 transactionSeq;
        uint256 encryptedIdLen;
    }
}
