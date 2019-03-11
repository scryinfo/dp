import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const state = {
    // show data dictionary.

    datalist: [],
    // {Title: "", Price: 0, Keys: "", Description: "", SupportVerify: false, pID: ""}  // primary key: publishID.

    transactionbuy: [],
    // {Buyer: "", Seller: "", State: "", Title: "", Price: 0, pID: "", tID: "",
    //     Verifier1Response: "", Verifier2Response: "", ArbitrateResult: false} // primary key: transactionID.

    transactionsell: [],
    // the same as transactionbuy.
    
    accounts: [],
    // {address: ""} // primary key: address.

    account: ""
}

const Store = new Vuex.Store({
    state
})

export default Store
