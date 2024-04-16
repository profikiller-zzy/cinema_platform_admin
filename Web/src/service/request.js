import axios from "axios";
import {AdminInfoStore} from "@/stores/admin_info.js";

// 初始化一个axios实例
export const Service = axios.create({
    timeout: 7000, // 响应时间为七秒
    baseURL: "",
    headers: {
        "Content-Type": "application/json"
    }
})

// 请求拦截器
Service.interceptors.request.use(request => {
    // 中间件用于存储管理员的token
    const store = AdminInfoStore();
    const token = store.adminInfo.token;
    console.log("Token in AdminInfoStore:", token); // 添加日志，输出 token 的值
    request.headers["token"] = token;
    console.log("Request headers after adding token:", request.headers); // 添加日志，输出添加 token 后的 headers
    return request;
});

Service.interceptors.response.use(response => {
    // 一般只用于返回的数据
    return response.data
})