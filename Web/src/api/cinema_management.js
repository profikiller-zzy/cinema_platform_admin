import {Service} from "@/service/request.js";

// 获取影院票房排行榜
export function cinemaRankApi() {
    return Service.get("/api/cinemas_rank")
}