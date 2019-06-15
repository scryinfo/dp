// Scry Info.  All rights reserved.
// license that can be found in the license file.

import Vue  from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

let state = {
    // show data dictionary below each item.

    datalist: [],
    // {Title: "", Price: 0, Keys: "", Description: "", Seller: "", SupportVerify: false, pID: "", 
    //     MetaDataExtension: "", ProofDataExtensions: []}  
    //  primary key: PublishID.

    transactionbuy: [],
    // {Title: "", Price: 0, Keys: "", Description: "", Buyer: "", Seller: "", State: "", SupportVerify: false, StartVerify: false,
    //     MetaDataExtension: "", ProofDataExtensions: [], MetaDataIDEncWithSeller: "", MetaDataIDEncWithBuyer: "",
    //     Verifier1Response: "", Verifier2Response: "", pID: "", tID: ""}
    //
    // transaction's data dictionary = [datalist] + {Buyer, State, StartVerify, MetaDataIDEncWithSeller, MetaDataIDEncWithBuyer,
    //     Verifier1Response, Verifier2Response, tID} = 9 + 8 = 17
    //  primary key: TransactionID.

    transactionsell: [],
    // the same as txb, all tx arrays use the same database.
    
    transactionverifier: [],
    // the same as txb, all tx arrays use the same database.
    
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
