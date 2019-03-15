import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

let state = {
    // show data dictionary below each item.

    datalist: [],
    // {Title: "", Price: 0, Keys: "", Description: "", Seller: "", SupportVerify: false, pID: "", 
    //     MetaDataExtension: "", ProofDataExtensions: []}  
    //  primary key: PublishID.

    transactionbuy: [],
    // {Title: "", Price: 0, Keys: "", Description: "", Buyer: "", Seller: "", State: "", SupportVerify: false,
    //     MetaDataExtension: "", ProofDataExtensions: [], MetaDataIDEncWithSeller: "", MetaDataIDEncWithBuyer: "",
    //     Verifier1Response: "", Verifier2Response: "", ArbitrateResult: false, pID: "", tID: "",}
    //
    // transaction's data dictionary = [datalist] + {Buyer, State, MetaDataIDEncWithSeller, MetaDataIDEncWithBuyer,
    //     Verifier1Response, Verifier2Response, ArbitrateResult, tID} = 9 + 8 = 17
    //  primary key: TransactionID.

    transactionsell: [],
    // the same as transactionbuy, the two arrays use the same database.
    
    accounts: [],
    // {address: ""} 
    //  primary key: address.

    account: ""
}

let Store = new Vuex.Store({
    state
})

export default Store
