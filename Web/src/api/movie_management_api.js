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

// 编辑用户
export function movieEditApi(data) {
    return Service.put("/api/movies", data)
}