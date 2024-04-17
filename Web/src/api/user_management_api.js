import {Service} from "@/service/request.js";

export function userListApi(params) {
    return Service.get("/api/user_list", {params})
}

export function userCreateApi(data) {
    return Service.post("/api/user_create", data)
}

// 删除用户 参数是用户的ID列表
export function userRemoveApi(id_list) {
    return Service.delete("/api/user_remove", {data: {id_list}})
}

// 编辑用户
export function userEditApi(data) {
    return Service.put("/api/user_edit", data)
}