import {defineStore} from 'pinia'
import { message } from 'ant-design-vue'

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
            localStorage.setItem("adminInfo", JSON.stringify(info))
        },
        loadAdminInfo(){
            let adminInfo = JSON.parse(localStorage.getItem("adminInfo"))
            if (adminInfo === null){
                return
            }
            let exp = adminInfo.exp
            let nowTime = new Date().getTime()
            if (exp *1000 - nowTime < 0){
                message.warn("当前登录已失效")
            }
            this.setAdminInfo(adminInfo)
        }
    }
})