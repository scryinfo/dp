// Scry Info.  All rights reserved.
// license that can be found in the license file.

import NotFound        from "./views/404.vue";
import Login           from "./views/login.vue";
import Home            from "./views/home.vue";
import DataList        from "./views/options/datalist.vue";
import TransactionBuy  from "./views/options/transactionBuy.vue";
import TransactionSell from "./views/options/transactionSell.vue";
import Publish         from "./views/options/publish.vue";
import Verify          from "./views/options/verify.vue";
// import Arbitrate       from "./views/options/t_arbitrate.vue";     // contract not implement
import Balance         from "./views/options/balance.vue";
import NickName        from "./views/options/nickname.vue";
import Message         from "./views/options/message.vue";
// import test            from "./views/options/t_test.vue";          // for test
// import test2           from "./views/options/t_test_two.vue";      // for test
// import administrator   from "./views/options/t_administrator.vue"; // for test

let routes = [
    {
        path: "/",
        component: Login,
        name: "login",
        hidden: true
    },
    {
        path: "/home",
        component: Home,
        name: "home",
        children: [
            {path: "/dl",  component: DataList,        name: "数据列表"},
            {path: "/tb",  component: TransactionBuy,  name: "我购买的数据"},
            {path: "/ts",  component: TransactionSell, name: "我出售的数据"},
            {path: "/pd",  component: Publish,  name: "发布新数据"},
            {path: "/vf",  component: Verify,          name: "我验证的数据"},
            // {path: "/at",  component: Arbitrate,       name: "我仲裁的数据",  hidden: true}, // contract not implement
            {path: "/blc", component: Balance,         name: "Balance",       hidden: true},
            {path: "/ncn", component: NickName,        name: "NickName",      hidden: true},
            {path: "/msg", component: Message,         name: "Short Message", hidden: true},
            // {path: "/test_page",      component: test,          name: "Test",                    hidden: true}, // for test
            // {path: "/test_page2",     component: test2,         name: "Test2",                   hidden: true}, // for test
            // {path: "/administrator",  component: administrator, name: "Administrator Functions", hidden: true}  // for test
        ]
    },
    {
        path: "/404",
        component: NotFound,
        name: "not found",
        hidden: true
    },
    {
        path: "*",
        redirect: { path: "/404" },
        hidden: true
    }
];

export default routes;
