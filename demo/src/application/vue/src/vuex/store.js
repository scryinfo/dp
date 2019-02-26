import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const state = {
    datalist: [],
    mytransaction: [],
    accounts: [
        {address: ""}
    ],
    account: ""
}

const Store = new Vuex.Store({
    state
})

export default Store
