import NotFound from './views/404.vue'
import Login from './views/Login.vue'
import Home from './views/Home.vue'
import DataList from './views/options/DataList.vue'
import MyTransaction from './views/options/MyTransaction.vue'
import PublishNewData from './views/options/PublishNewData.vue'
import Message from './views/Message.vue'

let routes = [
    {
        path: '/',
        component: Login,
        name: 'login',
        hidden: true
    },
    {
        path: '/404',
        component: NotFound,
        name: 'not found',
        hidden: true
    },
    {
        path: '/home',
        component: Home,
        name: 'home',
        children: [
            {path: '/dl', component: DataList, name: 'Data list'},
            {path: '/mt', component: MyTransaction, name: 'My transaction'},
            {path: '/pd', component: PublishNewData, name: 'Publish new data'},
            {path: '/msg', component: Message, name: 'Short Message', hidden: true}
        ]
    },
    {
        path: '*',
        redirect: { path: '/404' },
        hidden: true
    }
]

export default routes;
