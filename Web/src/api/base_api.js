import {Service} from "@/service/request.js";
export function baseApiList(url, params) {
    return Service.get(url, {params})
}

export function baseRemoveApi(url, id_list) {
    return Service.delete(url, {data: {id_list}})
}