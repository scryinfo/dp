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
            {path: "/dl", component: DataList, name: "Data list"},
            {path: "/tb", component: TransactionBuy, name: "Transaction-buy"},
            {path: "/ts", component: TransactionSell, name: "Transaction-sell"},
            {path: "/pd", component: PublishNewData, name: "Publish new data"},
            {path: "/vf", component: Verify, name: "Verify"},
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
