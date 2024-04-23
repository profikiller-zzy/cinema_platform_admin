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
                name :"admin",
                component : ()=>import("../views/admin/admin.vue"),
                children:[
                    {
                        path :"",
                        name :"adminHome",
                        redirect:"admin/home"
                    },
                    {
                        path :"home",
                        name :"admin_home",
                        component : ()=>import("@/views/admin/home/admin_home.vue")
                    },
                    {
                        path :"user_list",
                        name :"user_list",
                        component : ()=>import("@/views/admin/user_management/user_list.vue")
                    },
                    {
                        path :"movie_list",
                        name :"movie_list",
                        component : ()=>import("@/views/admin/movie_management/movie_list.vue")
                    },
                    {
                        path :"cinema_list",
                        name :"cinema_list",
                        component : ()=>import("@/views/admin/cinema_management/cinema_list.vue")
                    },
                ]
            }
        ]
}
)

export default router
