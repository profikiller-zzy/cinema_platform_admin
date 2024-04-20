import {Service} from "@/service/request.js";

export function userListApi(params) {
    return Service.get("/api/users", {params})
}

export function userCreateApi(data) {
    return Service.post("/api/users", data)
}

// 删除用户 参数是用户的ID列表
export function userRemoveApi(id_list) {
    return Service.delete("/api/users", {data: {id_list}})
}

// 编辑用户
export function userEditApi(data) {
    return Service.put("/api/users", data)
}