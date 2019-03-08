import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const state = {
    // show data dictionary.

    datalist: [],
    // {Title: "", Price: 0, Keys: "", Description: "", SupportVerify: false, pID: ""}  // primary key: publishID.

    mytransaction: [],
    // {Buyer: "", Seller: "", State: "", Title: "", ArbitrateResult: false, Price: 0, tID: "",
    //     Verifier1Response: "", Verifier2Response: "", Verifier3Response: ""} // primary key: transactionID.
    
    accounts: [],
    // {address: ""} // primary key: address.

    account: ""
}

const Store = new Vuex.Store({
    state
})

export default Store
