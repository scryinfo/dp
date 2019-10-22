// Scry Info.  All rights reserved.
// license that can be found in the license file.

import Vue  from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

let state = {
    // show data dictionary below each item.

    datalist: [],
    // {Title: "", Price: 0, Keys: "", Description: "", Seller: "", SupportVerify: false, pId: "",
    //     MetaDataExtension: "", ProofDataExtensions: []} = 9
    //  primary key: PublishId.

    transactionsell: [],
    // {Title: "", Price: 0, Keys: "", Description: "", Buyer: "", Seller: "", State: "", SupportVerify: false, StartVerify: false,
    //     MetaDataExtension: "", ProofDataExtensions: [], MetaDataIdEncWithSeller: "", MetaDataIdEncWithBuyer: "", MetaDataIdEncWithArbitrator: "",
    //     Verifier1Response: "", Verifier2Response: "", ArbitrateResult: false, pId: "", tId: ""， Identify: 0}
    //
    // identify match: 0 -> seller, 1 -> buyer, 2 -> verifier, 3 -> arbitrator
    //
    // transaction's data dictionary = [datalist] + {Buyer, State, StartVerify, MetaDataIdEncWithSeller, MetaDataIdEncWithBuyer, MetaDataIdEncWithArbitrator,
    //     Verifier1Response, Verifier2Response, ArbitrateResult, tId, identify} = 9 + 11 = 20
    //  primary key: TransactionId.

    transactionbuy: [],
    // [txs] + {SVDisplay} = 21
    
    transactionverifier: [],
    // the same as txs, all tx arrays use the same item in database.

    transactionarbitrator: [],
    // the same as txs, all tx arrays use the same item in database.
    
    accounts: [],
    // {address: "", nickname: "", fromBlock: 0(uint64), isVerifier: false}
    //  primary key: address.

    account: "",
    nickname: "用户昵称加载中...",

    balance: [
        // balance[0]: eth
        {
            Balance: "-",
            Time: "-"
        },
        // balance[1]: scry token
        {
            Balance: "-",
            Time: "-"
        }
    ]

};

let Store = new Vuex.Store({
    state
});

export default Store;
