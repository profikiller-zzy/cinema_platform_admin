import {defineStore} from 'pinia'
import { message } from 'ant-design-vue'
import { useRouter } from "vue-router"

export const AdminInfoStore = defineStore('adminInfo', {
    state: () => {
        return {
            adminInfo : {
                token: '',
                user_name: '',
                user_id: 0,
                avatar: '',
                exp: 0,
            }
        }
    },
    getters: {},
    actions: {
        setAdminInfo(info){
            this.$patch({
                adminInfo: info
            })
            // 持久化
            localStorage.setItem("adminInfo", JSON.stringify(info))
        },
        loadAdminInfo(){
            let adminInfo = JSON.parse(localStorage.getItem("adminInfo"))
            if (adminInfo === null){
                return
            }
            // 判断时间是否过期
            let exp = adminInfo.exp
            let nowTime = new Date().getTime()
            if (exp *1000 - nowTime < 0){
                message.warn("当前登录已失效")
                // 将当前页面跳转至登录页面
                const router = useRouter()
                router.push({ name: 'login' }) // 假设登录页面的名称为 'login'
                return
            }
            this.setAdminInfo(adminInfo)
        }
    }
})