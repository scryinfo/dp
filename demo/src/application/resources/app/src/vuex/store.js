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

const mutations = {
    dlAdd: function (state, dl) {
        state.datalist.push({ID: dl.ID, Title: dl.Title, Price: dl.Price, Keys: dl.Keys, Description: dl.Description,Owner: dl.Owner})
    },
    mtAdd: function (state, mt) {
        state.mytransaction.push({Title: mt.Title, TransactionID: mt.TransactionID, Seller:mt.Seller, Buyer:mt.Buyer, State: mt.State})
    },
    accAdd: function (state, acc) {
        state.accounts.push({address: acc.address})
    },
    accNew: function (state, acc) {
        state.account = acc
    }
}

const actions = {
    addDL: function (context, dl) {
        context.commit('dlAdd', dl)
    },
    addMT: function (context, mt) {
        context.commit('mtAdd', mt)
    },
    addAcc: function (context, acc) {
        context.commit('accAdd', acc)
    },
    addAc: function (context, acc) {
        context.commit('accNew', acc)
    }
}

const Store = new Vuex.Store({
    state,
    mutations,
    actions
})

export default Store
