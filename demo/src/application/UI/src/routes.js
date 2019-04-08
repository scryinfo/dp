import NotFound from "./views/404.vue"
import Login from "./views/Login.vue"
import Home from "./views/Home.vue"
import DataList from "./views/options/DataList.vue"
import TransactionBuy from "./views/options/TransactionBuy.vue"
import TransactionSell from "./views/options/TransactionSell.vue"
import PublishNewData from "./views/options/PublishNewData.vue"
import Verify from "./views/options/Verify.vue"
import Message from "./views/options/Message.vue"
import Test from "./views/options/test.vue"  // for test

let routes = [
    {
        path: "/",
        component: Login,
        name: "login",
        hidden: true
    },
    {
        path: "/404",
        component: NotFound,
        name: "not found",
        hidden: true
    },
    {
        path: "/home",
        component: Home,
        name: "home",
        children: [
            {path: "/dl", component: DataList, name: "数据列表"},
            {path: "/tb", component: TransactionBuy, name: "我购买的数据"},
            {path: "/ts", component: TransactionSell, name: "我出售的数据"},
            {path: "/vf", component: Verify, name: "我验证的数据"},
            {path: "/pd", component: PublishNewData, name: "发布新数据"},
            {path: "/msg", component: Message, name: "Short Message", hidden: true},
            {path: "/test", component: Test, name: "Test", hidden: true} // for test
        ]
    },
    {
        path: "*",
        redirect: { path: "/404" },
        hidden: true
    }
]

export default routes;
