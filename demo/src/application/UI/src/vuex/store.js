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
    // {Title: "", Price: 0, Keys: "", Description: "", Buyer: "", Seller: "", State: "", SupportVerify: false, StartVerify: false,
    //     MetaDataExtension: "", ProofDataExtensions: [], MetaDataIDEncWithSeller: "", MetaDataIDEncWithBuyer: "", MetaDataIDEncWithArbitrator: "", 
    //     ArbitrateResult: false, Verifier1: "", Verifier2: "", Verifier1Response: "", Verifier2Response: "", pID: "", tID: ""}
    //
    // transaction's data dictionary = [datalist] + {Buyer, State, StartVerify, MetaDataIDEncWithSeller, MetaDataIDEncWithBuyer,
    //     MetaDataIDEncWithArbitrator, Verifier1, Verifier2, Verifier1Response, Verifier2Response, ArbitrateResult, tID} = 9 + 12 = 21
    //  primary key: TransactionID.

    transactionsell: [],
    // the same as transactionbuy, four arrays use the same database.
    
    transactionverifier: [],
    // the same as transactionbuy, four arrays use the same database.
    
    accounts: [],
    // {address: "", fromBlock: 0(uint64), isVerifier: false}
    //  primary key: address.

    account: ""
}

let Store = new Vuex.Store({
    state
})

export default Store
