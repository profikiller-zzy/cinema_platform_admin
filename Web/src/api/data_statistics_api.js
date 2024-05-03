import {Service} from "@/service/request.js";

export function dataStatisticsApi() {
    return Service.get("/api/data_statistics/",{})
}

export function dataStatisticsInSevenDaysApi() {
    return Service.get("/api/statistics_in_seven_days/",{})
}