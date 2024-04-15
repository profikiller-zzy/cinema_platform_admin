import {Service} from "@/service/request.js";
// async 用于异步请求
export async function adminLogin(data) {
    return Service.post("/api/admin_login", data)
}