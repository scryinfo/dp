import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const state = {
    datalist: [
        {Title: "", Price: 0, Keys: "", Description: "", Owner: "", pID: ""} // primary key: publishID.
    ],
    mytransaction: [
        {Buyer: "", Seller: "", State: "", Title: "", tID: ""} // primary key: transactionID.
    ],
    accounts: [
        {address: ""} // primary key: address.
    ],
    account: ""
}

const Store = new Vuex.Store({
    state
})

export default Store
