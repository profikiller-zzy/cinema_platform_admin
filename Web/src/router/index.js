import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
        routes: [
            {
                path :"/login",
                name :"login",
                component : ()=>import("../views/admin_login.vue")
            },
            {
                path :"/admin",
                name :"admin_home",
                component : ()=>import("../views/admin/admin.vue"),
                children:[]
            }
        ]
}
)

export default router
