import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const state = {
    datalist: [],
    mytransaction: [],
    accounts: [
        {address: ""},
        {address: "0x5124852365789564128564723598621475354895"}
    ],
    account: ""
}

const Store = new Vuex.Store({
    state
})

export default Store
