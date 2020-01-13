// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Vuex from "vuex";
import App from './App.vue'
import router from './router'
import store from "./vuex/store";
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

Vue.use(Vuex);
Vue.use(ElementUI);
// Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
    router,
    store,
    render: h => h(App)
}).$mount("#app");
