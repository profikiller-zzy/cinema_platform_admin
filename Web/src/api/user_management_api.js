import {Service} from "@/service/request.js";

export function userListApi(params) {
    return Service.get("/api/user_list", {params})
}

export function userCreateApi(data) {
    return Service.post("/api/user_create", data)
}
