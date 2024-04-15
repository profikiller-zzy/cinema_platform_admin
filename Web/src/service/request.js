import axios from "axios";

// 初始化一个axios实例
export const Service = axios.create({
    timeout: 7000, // 响应时间为七秒
    baseURL: "",
    headers: {
        "Content-Type": "application/json"
    }
})

// Service.interceptors.request.use(request => {
//     // 中间件一般用于存储管理员的token
//     const store = UserInfoStore()
//     request.headers["token"] = store.userInfo.token
//     return request
// })

Service.interceptors.response.use(response => {
    // 一般只用于返回的数据
    return response.data
})