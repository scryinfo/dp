import Vue from "vue";
import Vuex from "vuex";
import store from "./vuex/store";
import VueRouter from "vue-router";
import routes from "./routes";
import ElementUI from "element-ui";
import "element-ui/lib/theme-chalk/index.css";
import App from "./App";

Vue.use(Vuex);
Vue.use(VueRouter);
Vue.use(ElementUI);

const router = new VueRouter({
  routes
});

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");

