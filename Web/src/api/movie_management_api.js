import {Service} from "@/service/request.js";

export function movieCoverUploadApi(formData) {
  return Service.post('/api/movie_cover/', formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  });
}

export function movieCreateApi(data) {
    return Service.post("/api/movies", data)
}

// 编辑电影
export function movieEditApi(data) {
    return Service.put("/api/movies", data)
}

// 获取电影票房排行榜
export function movieRankApi() {
    return Service.get("/api/movies_rank")
}

// 获取电影票房
export function movieBoxOfficeApi(data) {
    return Service.post("/api/movie_box_office", data)
}