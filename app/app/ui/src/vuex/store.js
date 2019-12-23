// Scry Info.  All rights reserved.
// license that can be found in the license file.

import Vue  from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

let state = {
    accounts: [],
    // {address: ""} = 1

    datalist: [],
    // {Title: "", Price: "", Keys: "", Description: "", Seller: "", SupportVerify: false, PublishId: ""} = 7 + {SVDisplay} = 8

    transactionsell: [],
    // {PublishId: "", TransactionId: "", Title: "", Price: "", Keys: "", Description: "", State: "", ArbitrateResult: false} = 8 + {SVDisplay, NVDisplay} = 10

    transactionbuy: [],
    // {PublishId: "", TransactionId: "", Title: "", Price: "", Keys: "", Description: "", State: "", ArbitrateResult: false,
    //     Verifier1Response: "", Verifier2Response: "", SupportVerify: false} = 11 + {SVDisplay, NVDisplay} = 13
    
    transactionverifier: [],
    // {PublishId: "", TransactionId: "", Title: "", Price: "", Keys: "", Description: ""} = 6

    transactionarbitrator: [],
    // {PublishId: "", TransactionId: "", Title: "", Price: "", Keys: "", Description: ""} = 6

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
