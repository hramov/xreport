import {createRouter, createWebHashHistory} from "vue-router";
import Home from "./components/Home.vue";

const routes = [
    { path: '/', component: Home },
    { path: '/create-service', component: () => import('./components/create-service/CreateService.vue') },
    { path: '/create-report', component: () => import('./components/create-report/CreateReport.vue') },
    { path: '/create-data-source', component: () => import('./components/create-data-source/CreateDataSource.vue') },
    { path: '/report/:id', component: () => import('./components/report/Report.vue') },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

export default router