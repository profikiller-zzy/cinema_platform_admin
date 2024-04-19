import {Service} from "@/service/request.js";
export function baseApiList(url, params) {
    return Service.get(url, {params})
}