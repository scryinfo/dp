import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const state = {
    // show data dictionary.

    datalist: [],
    // {Title: "", Price: 0, Keys: "", Description: "", Owner: "", pID: ""}  // primary key: publishID.

    mytransaction: [
    {Buyer: "0x0000", Seller: "0x0001", State: "Created", Title: "test title", ArbitrateResult: false, tID: "test transactionID",
        Verifier1Response: "1,v1r", Verifier2Response: "1,v2r", Verifier3Response: "1,v3r"} // primary key: transactionID.
    ],

    
    accounts: [],
    // {address: ""} // primary key: address.

    account: ""
}

const Store = new Vuex.Store({
    state
})

export default Store
